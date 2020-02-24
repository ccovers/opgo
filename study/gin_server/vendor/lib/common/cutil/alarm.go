package cutil

import (
	"lib/common/serverapi"
)

import (
	"runtime"
)

var YunwaAlarmGroup string = "1"

func AlarmPanic(info string) error {
	req := serverapi.WechatAlarmReq{
		TagList: []string{YunwaAlarmGroup},
		Msg:     "程序崩溃\n" + info,
	}

	res := serverapi.WechatAlarmRes{}
	return serverapi.SendWechatAlarm(&req, &res)
}

func BusinessAlarm(info string) error {
	req := serverapi.WechatAlarmReq{
		TagList: []string{YunwaAlarmGroup},
		Msg:     "业务出错\n" + info,
	}

	res := serverapi.WechatAlarmRes{}
	return serverapi.SendWechatAlarm(&req, &res)
}

func Stack() string {
	var buf [2 << 10]byte
	return string(buf[:runtime.Stack(buf[:], true)])
}
