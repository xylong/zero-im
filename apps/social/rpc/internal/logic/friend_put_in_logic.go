package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"
	"zero-im/models"
	"zero-im/pkg/constants"
	"zero-im/pkg/xerr"

	"zero-im/apps/social/rpc/internal/svc"
	"zero-im/apps/social/rpc/social"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendPutInLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFriendPutInLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendPutInLogic {
	return &FriendPutInLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// FriendPutIn 好友申请
func (l *FriendPutInLogic) FriendPutIn(in *social.FriendPutInReq) (*social.FriendPutInResp, error) {

	// 申请人是否与目标是好友关系
	friends, err := l.svcCtx.FriendDao.FindByUidAndFid(in.UserId, in.ReqUid)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrapf(xerr.NewDBErr(), "find friends by uid and fid err %v req %v ", err, in)
	}
	if friends != nil {
		return &social.FriendPutInResp{}, err
	}

	// 是否已经有过申请，申请是不成功，没有完成
	friendReqs, err := l.svcCtx.FriendRequestDao.FindByReqUidAndUserId(in.ReqUid, in.UserId)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrapf(xerr.NewDBErr(), "find friendsRequest by rid and uid err %v req %v ", err, in)
	}
	if friendReqs != nil && len(friendReqs) > 0 {
		t, _ := json.Marshal(friendReqs)
		fmt.Println(string(t))
		return &social.FriendPutInResp{}, err
	}

	// 创建申请记录
	err = l.svcCtx.FriendRequestDao.Create(&models.FriendRequest{
		UserID:       in.UserId,
		ReqUID:       in.ReqUid,
		ReqMsg:       in.ReqMsg,
		ReqTime:      time.Unix(in.ReqTime, 0),
		HandleResult: int64(constants.NoHandlerResult),
	})

	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "insert friendRequest err %v req %v ", err, in)
	}

	return &social.FriendPutInResp{}, nil
}
