package serverapi

import (
	"lib/common/chttp"
	"lib/common/errcode"
)

import (
	"fmt"
)

type WechatAlarmReq struct {
	UserList []string `json:"user_list"`
	TagList  []string `json:"tag_list"`
	Msg      string   `json:"msg"`
}

type WechatAlarmRes struct {
	Errcode     int64    `json:"errcode"`
	Errmsg      string   `json:"errmsg"`
	Invaliduser []string `json:"invaliduser"`
	Invalidtag  []string `json:"invalidtag"`
}

const (
	msg_max_length = 2048
)

var wechat_alarm_url string = "http://alarm.in.netwa.cn:8005/v1/alarm/inner/common/wechat"

func SendWechatAlarm(req *WechatAlarmReq, res *WechatAlarmRes) error {
	if req == nil || res == nil {
		return errcode.InvalidParameterError
	}
	if len(req.Msg) > msg_max_length {
		req.Msg = req.Msg[:msg_max_length]
	}
	return chttp.InnerRequest(wechat_alarm_url, req, res)
}

func GetMsg(program string, msg string) string {
	return fmt.Sprintf("程序告警\n来自: %s\n消息: %s", program, msg)
}
