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
 *
 */

// Code generated by informer-gen. DO NOT EDIT.

package externalversions

import (
	"fmt"
	v1alpha1 "github.com/Tencent/bk-bcs/bcs-services/bcs-k8s-watch/pkg/kubefed/apis/core/v1alpha1"
	v1beta1 "github.com/Tencent/bk-bcs/bcs-services/bcs-k8s-watch/pkg/kubefed/apis/core/v1beta1"
	multiclusterdns_v1alpha1 "github.com/Tencent/bk-bcs/bcs-services/bcs-k8s-watch/pkg/kubefed/apis/multiclusterdns/v1alpha1"
	scheduling_v1alpha1 "github.com/Tencent/bk-bcs/bcs-services/bcs-k8s-watch/pkg/kubefed/apis/scheduling/v1alpha1"
	types_v1beta1 "github.com/Tencent/bk-bcs/bcs-services/bcs-k8s-watch/pkg/kubefed/apis/types/v1beta1"

	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
)

// GenericInformer is type of SharedIndexInformer which will locate and delegate to other
// sharedInformers based on type
type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	// Group=core.kubefed.io, Version=v1alpha1
	case v1alpha1.SchemeGroupVersion.WithResource("clusterpropagatedversions"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1alpha1().ClusterPropagatedVersions().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("federatedservicestatuses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1alpha1().FederatedServiceStatuses().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("propagatedversions"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1alpha1().PropagatedVersions().Informer()}, nil

		// Group=core.kubefed.io, Version=v1beta1
	case v1beta1.SchemeGroupVersion.WithResource("federatedtypeconfigs"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1beta1().FederatedTypeConfigs().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("kubefedclusters"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1beta1().KubeFedClusters().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("kubefedconfigs"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1beta1().KubeFedConfigs().Informer()}, nil

		// Group=multiclusterdns.kubefed.io, Version=v1alpha1
	case multiclusterdns_v1alpha1.SchemeGroupVersion.WithResource("dnsendpoints"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Multiclusterdns().V1alpha1().DNSEndpoints().Informer()}, nil
	case multiclusterdns_v1alpha1.SchemeGroupVersion.WithResource("domains"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Multiclusterdns().V1alpha1().Domains().Informer()}, nil
	case multiclusterdns_v1alpha1.SchemeGroupVersion.WithResource("ingressdnsrecords"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Multiclusterdns().V1alpha1().IngressDNSRecords().Informer()}, nil
	case multiclusterdns_v1alpha1.SchemeGroupVersion.WithResource("servicednsrecords"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Multiclusterdns().V1alpha1().ServiceDNSRecords().Informer()}, nil

		// Group=scheduling.kubefed.io, Version=v1alpha1
	case scheduling_v1alpha1.SchemeGroupVersion.WithResource("replicaschedulingpreferences"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Scheduling().V1alpha1().ReplicaSchedulingPreferences().Informer()}, nil

		// Group=types.kubefed.io, Version=v1beta1
	case types_v1beta1.SchemeGroupVersion.WithResource("federatedclusterroles"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Types().V1beta1().FederatedClusterRoles().Informer()}, nil
	case types_v1beta1.SchemeGroupVersion.WithResource("federatedclusterrolebindings"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Types().V1beta1().FederatedClusterRoleBindings().Informer()}, nil
	case types_v1beta1.SchemeGroupVersion.WithResource("federatedconfigmaps"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Types().V1beta1().FederatedConfigMaps().Informer()}, nil
	case types_v1beta1.SchemeGroupVersion.WithResource("federateddaemonsets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Types().V1beta1().FederatedDaemonSets().Informer()}, nil
	case types_v1beta1.SchemeGroupVersion.WithResource("federateddeployments"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Types().V1beta1().FederatedDeployments().Informer()}, nil
	case types_v1beta1.SchemeGroupVersion.WithResource("federatedendpointses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Types().V1beta1().FederatedEndpointses().Informer()}, nil
	case types_v1beta1.SchemeGroupVersion.WithResource("federatedingresses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Types().V1beta1().FederatedIngresses().Informer()}, nil
	case types_v1beta1.SchemeGroupVersion.WithResource("federatedjobs"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Types().V1beta1().FederatedJobs().Informer()}, nil
	case types_v1beta1.SchemeGroupVersion.WithResource("federatednamespaces"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Types().V1beta1().FederatedNamespaces().Informer()}, nil
	case types_v1beta1.SchemeGroupVersion.WithResource("federatedreplicasets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Types().V1beta1().FederatedReplicaSets().Informer()}, nil
	case types_v1beta1.SchemeGroupVersion.WithResource("federatedreplicationcontrollers"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Types().V1beta1().FederatedReplicationControllers().Informer()}, nil
	case types_v1beta1.SchemeGroupVersion.WithResource("federatedroles"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Types().V1beta1().FederatedRoles().Informer()}, nil
	case types_v1beta1.SchemeGroupVersion.WithResource("federatedrolebindings"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Types().V1beta1().FederatedRoleBindings().Informer()}, nil
	case types_v1beta1.SchemeGroupVersion.WithResource("federatedsecrets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Types().V1beta1().FederatedSecrets().Informer()}, nil
	case types_v1beta1.SchemeGroupVersion.WithResource("federatedservices"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Types().V1beta1().FederatedServices().Informer()}, nil
	case types_v1beta1.SchemeGroupVersion.WithResource("federatedserviceaccounts"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Types().V1beta1().FederatedServiceAccounts().Informer()}, nil
	case types_v1beta1.SchemeGroupVersion.WithResource("federatedstatefulsets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Types().V1beta1().FederatedStatefulSets().Informer()}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}
