package serverapi

import (
	"lib/common/chttp"
	"lib/common/errcode"
)

type GetUserBaseInfoReq struct {
	Uids []string `json:"uid"`
}

type UserBaseInfo struct {
	Uid  string `json:"uid"`
	Icon string `json:"icon"`
	Nick string `json:"nick"`
}
type GetUserBaseInfoRes struct {
	Infos []UserBaseInfo `json:"user_base_info"`
}

var base_info_url string = "http://userinfo.in.netwa.cn:8002/v1/userinfo/inner/get_user_base_info"

func GetUserBaseInfo(req *GetUserBaseInfoReq, res *GetUserBaseInfoRes) error {
	if req == nil || res == nil {
		return errcode.InvalidParameterError
	}
	return chttp.InnerRequest(base_info_url, req, res)
}

type AssetCount struct {
	Uid       string `json:"uid"`
	AssetType int    `json:"asset_type"`
	Count     int    `json:"count"`
}

const WinEnum int64 = 1
const LoseEnum int64 = 0

type OperAssetsReq struct {
	Tag    string       `json:"tag"`
	Assets []AssetCount `json:"assets"`
	Win    int64        `json:"win"`
}
type OperAssetsRes struct {
	State int32       `json:"state"`
	Data  interface{} `json:"data"`
}

var oper_assets_url string = "http://userinfo.in.netwa.cn:8002/v1/userinfo/inner/oper_assets"

func OperAssets(req *OperAssetsReq, res *OperAssetsRes) error {
	if req == nil || res == nil {
		return errcode.InvalidParameterError
	}

	return chttp.InnerRequest(oper_assets_url, req, res)
}

const (
	Consume       = 1 // 直接扣除
	Trans         = 2 // 交易到房间
	CompleteTrans = 3 // 确认交易
	Undo          = 4 // 回退交易
)

type Rinfo struct {
	Uid   string `json:"uid"`
	Count int64  `json:"count"`
}

type RoomConsumeOperReq struct {
	Oper   int64   `json:"oper"`
	RoomId string  `json:"room_id"`
	Ucount []Rinfo `json:"ucount"`
}

func (r RoomConsumeOperReq) Valid() int {
	if r.Oper != Consume && r.Oper != Trans &&
		r.Oper != CompleteTrans && r.Oper != Undo &&
		len(r.Ucount) == 0 {
		return errcode.ParamInvalid
	}
	return errcode.OKCode
}

type RoomConsumeOperRes struct {
}

var room_consume_oper string = "http://userinfo.in.netwa.cn:8002/v1/userinfo/inner/room_consume_oper"

func RoomConsumeOper(req *RoomConsumeOperReq, res *RoomConsumeOperRes) error {
	if req == nil || res == nil {
		return errcode.InvalidParameterError
	}

	return chttp.InnerRequest(room_consume_oper, req, res)
}

type AddUserAssetsReq struct {
	Nick    string `json:"nick"`
	Icon    string `json:"icon"`
	Channel string `json:"channel"`
	Usid    string `json:"usid"`
	Uid     string `json:"uid"`
}

type AddUserAssetsRes struct {
}

var add_user_assets string = "http://userinfo.in.netwa.cn:8002/v1/userinfo/inner/add_user_info"

func AddUserAssets(req *AddUserAssetsReq, res *AddUserAssetsRes) error {
	if req == nil || res == nil {
		return errcode.InvalidParameterError
	}

	return chttp.InnerRequest(add_user_assets, req, res)
}
