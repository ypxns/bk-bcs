/*
Tencent is pleased to support the open source community by making Basic Service Configuration Platform available.
Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
Licensed under the MIT License (the "License"); you may not use this file except
in compliance with the License. You may obtain a copy of the License at
http://opensource.org/licenses/MIT
Unless required by applicable law or agreed to in writing, software distributed under
the License is distributed on an "as IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
either express or implied. See the License for the specific language governing permissions and
limitations under the License.
*/

package service

import (
	"context"
	"errors"

	"bscp.io/pkg/criteria/errf"
	"bscp.io/pkg/kit"
	pbcs "bscp.io/pkg/protocol/cache-service"
	pbbase "bscp.io/pkg/protocol/core/base"
	"bscp.io/pkg/types"
)

// GetAppID get app id by app name.
func (s *Service) GetAppID(ctx context.Context, req *pbcs.GetAppIDReq) (*pbcs.GetAppIDResp, error) {
	if req.BizId <= 0 || req.AppName == "" {
		return nil, errf.New(errf.InvalidParameter, "invalid biz id or app name")
	}

	kt := kit.FromGrpcContext(ctx)
	appID, err := s.op.GetAppID(kt, req.BizId, req.AppName)
	if err != nil {
		return nil, err
	}

	return &pbcs.GetAppIDResp{
		AppId: appID,
	}, nil
}

// GetAppMeta get app's basic info.
func (s *Service) GetAppMeta(ctx context.Context, req *pbcs.GetAppMetaReq) (*pbcs.JsonRawResp, error) {
	if req.BizId <= 0 || req.AppId <= 0 {
		return nil, errf.New(errf.InvalidParameter, "invalid biz id or app id")
	}

	kt := kit.FromGrpcContext(ctx)
	meta, err := s.op.GetAppMeta(kt, req.BizId, req.AppId)
	if err != nil {
		return nil, err
	}

	return &pbcs.JsonRawResp{
		JsonRaw: meta,
	}, nil
}

// GetReleasedCI get released config items from cache.
func (s *Service) GetReleasedCI(ctx context.Context, req *pbcs.GetReleasedCIReq) (*pbcs.JsonRawResp, error) {
	if req.BizId <= 0 || req.ReleaseId <= 0 {
		return nil, errf.New(errf.InvalidParameter, "invalid biz id or release id")
	}

	kt := kit.FromGrpcContext(ctx)
	ci, err := s.op.GetReleasedCI(kt, req.BizId, req.ReleaseId)
	if err != nil {
		return nil, err
	}

	return &pbcs.JsonRawResp{
		JsonRaw: ci,
	}, nil
}

// ListAppReleasedGroups list app's released groups.
func (s *Service) ListAppReleasedGroups(ctx context.Context, req *pbcs.ListAppReleasedGroupsReq) (
	*pbcs.JsonRawResp, error) {

	if req.BizId <= 0 || req.AppId <= 0 {
		return nil, errf.New(errf.InvalidParameter, "invalid biz id or app id")
	}

	kt := kit.FromGrpcContext(ctx)
	list, err := s.op.ListAppReleasedGroups(kt, req.BizId, req.AppId)
	if err != nil {
		return nil, err
	}

	return &pbcs.JsonRawResp{JsonRaw: list}, nil
}

// ListCredentialMatchedCI list all config item ids which can be matched by credential.
func (s *Service) ListCredentialMatchedCI(ctx context.Context, req *pbcs.ListCredentialMatchedCIReq) (
	*pbcs.JsonRawResp, error) {

	if req.BizId <= 0 {
		return nil, errf.New(errf.InvalidParameter, "biz id can't be empty")
	}

	if req.Credential == "" {
		return nil, errf.New(errf.InvalidParameter, "credential can't be empty")
	}

	kt := kit.FromGrpcContext(ctx)
	list, err := s.op.ListCredentialMatchedCI(kt, req.BizId, req.Credential)
	if err != nil {
		return nil, err
	}

	return &pbcs.JsonRawResp{JsonRaw: list}, nil
}

// GetCredential get credential by credential string.
func (s *Service) GetCredential(ctx context.Context, req *pbcs.GetCredentialReq) (*pbcs.JsonRawResp, error) {
	if req.BizId <= 0 {
		return nil, errf.New(errf.InvalidParameter, "biz id can't be empty")
	}

	if req.Credential == "" {
		return nil, errf.New(errf.InvalidParameter, "credential can't be empty")
	}

	kt := kit.FromGrpcContext(ctx)
	credential, err := s.op.GetCredential(kt, req.BizId, req.Credential)
	if err != nil {
		return nil, err
	}

	return &pbcs.JsonRawResp{JsonRaw: credential}, nil
}

// GetCurrentCursorReminder get the current consumed event's id, which is the cursor reminder's resource id.
func (s *Service) GetCurrentCursorReminder(ctx context.Context, _ *pbbase.EmptyReq) (*pbcs.CurrentCursorReminderResp,
	error) {

	kt := kit.FromGrpcContext(ctx)
	cursor, err := s.dao.Event().LatestCursor(kt)
	if err != nil {
		return nil, err
	}

	return &pbcs.CurrentCursorReminderResp{Cursor: cursor}, nil
}

// ListEventsMeta list event metas with filter
func (s *Service) ListEventsMeta(ctx context.Context, req *pbcs.ListEventsReq) (*pbcs.ListEventsResp, error) {
	kt := kit.FromGrpcContext(ctx)
	if req.Page == nil {
		return nil, errors.New("page is null")
	}

	opt := req.Page.BasePage()
	if err := opt.Validate(types.DefaultPageOption); err != nil {
		return nil, err
	}

	details, _, err := s.dao.Event().ListConsumedEvents(kt, req.StartCursor, opt)
	if err != nil {
		return nil, err
	}

	metas := make([]*types.EventMeta, len(details))
	for idx := range details {
		metas[idx] = &types.EventMeta{
			ID:         details[idx].ID,
			Spec:       details[idx].Spec,
			Attachment: details[idx].Attachment,
		}
	}

	return &pbcs.ListEventsResp{List: pbcs.PbEventMeta(metas)}, nil
}
