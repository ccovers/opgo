package serverapi

import (
	"encoding/json"
	"lib/common/chttp"
	"lib/common/clog"
	"lib/common/errcode"
)

const KNotifierURL = "http://notifier.in.netwa.cn:5581/push_msg"

type NotifierMessage struct {
	Topic   string `json:"topic"`
	Data    string `json:"data"`
	OrderId string `json:"order_id"` //消息OrderId,如果消息不要求有序则留空
	Url     string `json:"url"`
}

type NotifierRes struct {
	State int32       `json:"state"`
	Data  interface{} `json:"data"`
}

// PushMessageToNotifier 向Notifier推送kafka消息
func PushMessageToNotifier(req *NotifierMessage, res *NotifierRes) error {
	if req == nil || res == nil {
		return errcode.InvalidParameterError
	}
	return chttp.InnerRequest(KNotifierURL, req, res)
}

func commonPushProcess(req interface{}, topic string) (interface{}, error) {
	assetsBytes, err := json.Marshal(req)
	if err != nil {
		clog.Logger.Info("json marshal error: %v", err)
		return nil, err
	}

	notifierReq := NotifierMessage{
		Topic: topic,
		Data:  string(assetsBytes),
	}
	notifierRes := NotifierRes{}

	err = PushMessageToNotifier(&notifierReq, &notifierRes)
	if err != nil {
		clog.Logger.Error("push message to notifier error: %v", err)
		return nil, err
	}

	return notifierRes, nil
}

const KOperAssetsTopic string = "operAssets"

// OperAssetsByNotifier 操作用户资产
func OperAssetsByNotifier(req *OperAssetsReq, res *OperAssetsRes) error {
	assetsBytes, err := json.Marshal(req)
	if err != nil {
		clog.Logger.Info("json marshal error: %v", err)
		return err
	}

	notifierReq := NotifierMessage{
		Topic: KOperAssetsTopic,
		Data:  string(assetsBytes),
	}
	notifierRes := NotifierRes{}

	err = PushMessageToNotifier(&notifierReq, &notifierRes)
	if err != nil {
		clog.Logger.Error("push message to notifier error: %v", err)
		return err
	}

	res.State = notifierRes.State
	res.Data = notifierRes.Data

	return nil
}

// type UserOperReq struct {
// 	UserID    string `json:"uid"`
// 	RoomID    string `json:"room_id"`
// 	RoomType  int64  `json:"room_type"`
// 	Operation int32  `json:"oper"`
// 	Question  int64  `json:"question"`
// 	Answer    string `json:"answer"`
// 	UserPick  string `json:"user_pick"`
// }

//const KLogUserOperTopic string = "logUserOper"
// RecordUserOperByNotifier 记录用户行为
func RecordUserOperByNotifier(req *UserOperReq, res *NotifierRes) error {
	pusRes, err := commonPushProcess(req, KLogUserOperTopic)

	nRes, ok := pusRes.(NotifierRes)
	if ok {
		res.State = nRes.State
		res.Data = nRes.Data
	}

	return err
}
