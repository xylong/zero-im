package friend

import (
	"context"
	"zero-im/apps/social/rpc/social_client"
	"zero-im/pkg/ctxdata"

	"zero-im/apps/social/api/internal/svc"
	"zero-im/apps/social/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendPutInHandleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewFriendPutInHandleLogic 好友申请处理
func NewFriendPutInHandleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendPutInHandleLogic {
	return &FriendPutInHandleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// FriendPutInHandle 好友申请处理
func (l *FriendPutInHandleLogic) FriendPutInHandle(req *types.FriendPutInHandleReq) (resp *types.FriendPutInHandleResp, err error) {

	_, err = l.svcCtx.Social.FriendPutInHandle(l.ctx, &social_client.FriendPutInHandleReq{
		FriendReqId:  req.FriendReqId,
		UserId:       ctxdata.GetUid(l.ctx),
		HandleResult: req.HandleResult,
	})

	return
}
