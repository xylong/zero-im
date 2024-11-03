package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"zero-im/models"
	"zero-im/pkg/constants"
	"zero-im/pkg/xerr"

	"zero-im/apps/social/rpc/internal/svc"
	"zero-im/apps/social/rpc/social"

	"github.com/zeromicro/go-zero/core/logx"
)

var (
	ErrFriendReqBeforePass   = xerr.NewMsg("好友申请已经通过")
	ErrFriendReqBeforeRefuse = xerr.NewMsg("好友申请已经被拒绝")
)

type FriendPutInHandleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFriendPutInHandleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendPutInHandleLogic {
	return &FriendPutInHandleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// FriendPutInHandle 处理好友申请
func (l *FriendPutInHandleLogic) FriendPutInHandle(in *social.FriendPutInHandleReq) (*social.FriendPutInHandleResp, error) {
	// 获取好友申请记录
	friendReq, err := l.svcCtx.FriendRequestDao.FindOne(cast.ToInt64(in.FriendReqId))
	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "find friendsRequest by friendReqid err %v req %v ", err,
			in.FriendReqId)
	}

	// 验证是否有处理
	switch constants.HandlerResult(friendReq.HandleResult) {
	case constants.PassHandlerResult:
		return nil, errors.WithStack(ErrFriendReqBeforePass)
	case constants.RefuseHandlerResult:
		return nil, errors.WithStack(ErrFriendReqBeforeRefuse)
	}

	friendReq.HandleResult = cast.ToInt64(in.HandleResult)

	// 修改申请结果 -> 通过【建立两条好友关系记录】 -> 事务
	err = l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {
		if err := l.svcCtx.FriendRequestDao.Update(friendReq); err != nil {
			return errors.Wrapf(xerr.NewDBErr(), "update friend request err %v, req %v", err, friendReq)
		}

		if constants.HandlerResult(in.HandleResult) != constants.PassHandlerResult {
			return nil
		}

		friends := []*models.Friend{
			{
				UserID:    friendReq.UserID,
				FriendUID: friendReq.ReqUID,
			}, {
				UserID:    friendReq.ReqUID,
				FriendUID: friendReq.UserID,
			},
		}

		err = l.svcCtx.FriendDao.Insert(friends...)
		if err != nil {
			return errors.Wrapf(xerr.NewDBErr(), "friends inserts err %v, req %v", err, friends)
		}
		return nil
	})

	return &social.FriendPutInHandleResp{}, err
}
