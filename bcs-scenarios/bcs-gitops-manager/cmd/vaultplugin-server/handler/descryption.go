/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	"github.com/argoproj-labs/argocd-vault-plugin/pkg/auth/vault"
	"github.com/argoproj-labs/argocd-vault-plugin/pkg/config"
	"github.com/argoproj-labs/argocd-vault-plugin/pkg/kube"
	"github.com/argoproj-labs/argocd-vault-plugin/pkg/types"
	"github.com/argoproj-labs/argocd-vault-plugin/pkg/utils"
	argoappv1 "github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
	"github.com/hashicorp/vault/api"
	"github.com/spf13/viper"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// NOCC:gas/crypto(工具误报)
const secretPathPattern = "<path:%s#%s>"

type descryptionManifestRequest struct {
	Application string   `json:"application"`
	Manifests   []string `json:"manifests"`
}

type descryptionManifestResponse struct {
	Manifests []string `json:"manifests"`
	Message   string   `json:"message"`
}

// Response response function for descryptionManifest
func (r *descryptionManifestResponse) Response(w http.ResponseWriter, status int) {
	b, err := json.Marshal(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(status)
	w.Write(b)
}

func (v1 *V1VaultPluginHandler) descryptionManifest(w http.ResponseWriter, r *http.Request) {
	var req descryptionManifestRequest
	resp := &descryptionManifestResponse{}
	ctx := context.Background()
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		resp.Message = "Invaild request body."
		resp.Response(w, http.StatusBadRequest)
		return
	}

	if req.Manifests == nil || req.Application == "" {
		resp.Message = "Application and manifests are required."
		resp.Response(w, http.StatusBadRequest)
		return
	}
	secretname, err := getGitopsAppSecretkey(ctx, v1.Opts.GitopsStore, req.Application)
	if err != nil {
		resp.Message = "Error getting application's vault secretkey"
		blog.Errorf("Unable to get application's vault secretkey, error: %v", err)
		resp.Response(w, http.StatusInternalServerError)
		return
	}
	res, err := descryVaultSecret(secretname, req.Manifests)
	if err != nil {
		resp.Message = "Error descryp application's vault secretkey"
		blog.Errorf("Unable to descryp application's vault secret, error: %v", err)
		resp.Response(w, http.StatusInternalServerError)
		return
	}
	resp.Manifests = res
	resp.Response(w, http.StatusOK)
}

func descryVaultSecret(secretName string, manifestsstring []string) (res []string, err error) {
	res = make([]string, 0)
	v := viper.New()
	cmdConfig, err := config.New(v, &config.Options{
		SecretName: secretName,
	})
	if err != nil {
		return
	}
	apiClient, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		return
	}
	backend := NewVaultArgoBackend(&vault.TokenAuth{}, apiClient, v.GetString(types.EnvAvpKvVersion))
	cmdConfig.Backend = backend

	err = cmdConfig.Backend.Login()
	if err != nil {
		return
	}

	var manifests []unstructured.Unstructured = make([]unstructured.Unstructured, 0)
	for _, m := range manifestsstring {
		obj, marsharlerr := argoappv1.UnmarshalToUnstructured(m)
		if marsharlerr != nil {
			return res, marsharlerr
		}
		manifests = append(manifests, *obj)
	}

	for _, manifest := range manifests {
		var pathValidation *regexp.Regexp
		if rexp := v.GetString(types.EnvPathValidation); rexp != "" {
			pathValidation, err = regexp.Compile(rexp)
			if err != nil {
				err = fmt.Errorf("%s is not a valid regular expression: %s", rexp, err)
				return res, err
			}
		}

		template, tmplerr := kube.NewTemplate(manifest, cmdConfig.Backend, pathValidation)
		if tmplerr != nil {
			return res, tmplerr
		}

		annotations := manifest.GetAnnotations()
		avpIgnore, _ := strconv.ParseBool(annotations[types.AVPIgnoreAnnotation])
		if !avpIgnore {
			err = template.Replace()
			if err != nil {
				return res, err
			}
		} else {
			utils.VerboseToStdErr("skipping %s.%s because %s annotation is present", manifest.GetNamespace(), manifest.GetName(), types.AVPIgnoreAnnotation)
		}

		// output, ouputErr := template.ToYAML()
		jsondata, marshalErr := json.Marshal(template.TemplateData)
		if marshalErr != nil {
			err = marshalErr
			return res, err
		}
		output := string(jsondata)

		res = append(res, output)
	}

	return res, nil
}

// SecretNotFountErr 修改自vaultbackend
// 添加逻辑：密钥不存在时不进行渲染，忽略密钥不存在的错误
type SecretNotFountErr struct{}

// Error err for secret path
func (e SecretNotFountErr) Error() string {
	return "secret path or key is not found"
}

// VaultArgo is a struct for working with a Vault backend
type VaultArgo struct {
	types.AuthType
	VaultClient *api.Client
	KvVersion   string
}

// NewVaultArgoBackend initializes a new Vault Backend
func NewVaultArgoBackend(auth types.AuthType, client *api.Client, kv string) *VaultArgo {
	vault := &VaultArgo{
		KvVersion:   kv,
		AuthType:    auth,
		VaultClient: client,
	}
	return vault
}

// Login authenticates with the auth type provided
func (v *VaultArgo) Login() error {
	err := v.AuthType.Authenticate(v.VaultClient)
	if err != nil {
		return err
	}
	return nil
}

// GetSecrets gets secrets from vault and returns the formatted data
func (v *VaultArgo) GetSecrets(path string, version string, annotations map[string]string) (map[string]interface{}, error) {
	var secret *api.Secret
	var err error

	var kvVersion = v.KvVersion
	if kv, ok := annotations[types.VaultKVVersionAnnotation]; ok {
		kvVersion = kv
	}

	// Vault KV-V1 doesn't support versioning so we only honor `version` if KV-V2 is used
	if version != "" && kvVersion == "2" {
		utils.VerboseToStdErr("Hashicorp Vault getting kv pairs from KV-V2 path %s at version %s", path, version)
		secret, err = v.VaultClient.Logical().ReadWithData(path, map[string][]string{
			"version": {version},
		})
	} else {
		utils.VerboseToStdErr("Hashicorp Vault getting kv pairs from KV-V1 path %s", path)
		secret, err = v.VaultClient.Logical().Read(path)
	}

	if err != nil {
		return nil, err
	}

	utils.VerboseToStdErr("Hashicorp Vault get kv pairs response: %v", secret)

	if secret == nil {
		// Do not mention `version` in error message when it's not honored (KV-V1)
		// if version == "" || kvVersion == "1" {
		// 	// return nil, fmt.Errorf("Could not find secrets at path %s", path)
		// }
		// 忽略找不到密钥的错误
		return nil, SecretNotFountErr{}
		// return nil, fmt.Errorf("Could not find secrets at path %s with version %s", path, version)
	}

	if kvVersion == "2" {
		if _, ok := secret.Data["data"]; ok {
			if secret.Data["data"] != nil {
				return secret.Data["data"].(map[string]interface{}), nil
			}
			return nil, fmt.Errorf("The secret version %s for Vault path %s is nil - is this version of the secret deleted?", version, path)
		}
		if len(secret.Data) == 0 {
			return nil, fmt.Errorf("The Vault path: %s is empty - did you forget to include /data/ in the Vault path for kv-v2?", path)
		}
		return nil, errors.New("Could not get data from Vault, check that kv-v2 is the correct engine")
	}

	if kvVersion == "1" {
		return secret.Data, nil
	}

	return nil, errors.New("Unsupported kvVersion specified")
}

// GetIndividualSecret will get the specific secret (placeholder) from the SM backend
// For Vault, we only support placeholders replaced from the k/v pairs of a secret which cannot be individually addressed
// So, we use GetSecrets and extract the specific placeholder we want
func (v *VaultArgo) GetIndividualSecret(kvpath, secret, version string, annotations map[string]string) (interface{}, error) {
	data, err := v.GetSecrets(kvpath, version, annotations)
	if err != nil {
		// 忽略所有error，error时原样渲染
		blog.Errorf("GetSecrets failed with error: %v", err)
		return fmt.Sprintf(secretPathPattern, kvpath, secret), nil
	}
	// 忽略path存在,secret不存在
	_, ok := data[secret]
	if !ok {
		return fmt.Sprintf(secretPathPattern, kvpath, secret), nil
	}
	return data[secret], nil
}
