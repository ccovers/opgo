package serverapi

import (
	"lib/common/chttp"
	"lib/common/errcode"
)

var server_mgr_url string = "http://server_mgr.in.netwa.cn:8007/v1/servermgr/get_serveraddr"

type GetServerAddrReq struct {
	ServerName string `json:"servername"`
}

type GetServerAddrRes struct {
	Addr      string `json:"addr"`
	InnerAddr string `json:"inneraddr"`
}

func GetServerAddr(req *GetServerAddrReq, res *GetServerAddrRes) error {
	if req == nil || res == nil {
		return errcode.InvalidParameterError
	}
	return chttp.InnerRequest(server_mgr_url, req, res)
}

var server_mgr_valid_url string = "http://server_mgr.in.netwa.cn:8007/v1/servermgr/valid_serveraddr"

type ValidServerAddrReq struct {
	ServerName string `json:"servername"`
	Addr       string `json:"addr"`
}

type ValidServerAddrRes struct {
	Valid bool `json:"valid"`
}

func ValidServerAddr(req *ValidServerAddrReq, res *ValidServerAddrRes) error {
	if req == nil || res == nil {
		return errcode.InvalidParameterError
	}
	return chttp.InnerRequest(server_mgr_valid_url, req, res)
}
