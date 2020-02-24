package serverapi

import (
	"lib/common/chttp"
	"lib/common/errcode"
	"lib/common/model"
)

import (
	"fmt"
)

type CreateRoomV2Req struct {
	Rtag     string                  `json:"rtag"`
	Rii      model.RoomInitInfo      `json:"rii"`
	Prc      model.PassRoomCondition `json:"prc"`
	Pra      model.PassRoomAward     `json:"pra"`
	Count    int64                   `json:"count"`
	Quess    []model.QuestionInfo    `json:"quess"`
	Eid      int                     `json:"eid"`
	Uid      string                  `json:"uid"`
	RoomType int                     `json:"room_type"`
}

type CreateRoomV2Res struct {
	Token string `json:"token"`
	Addr  string `json:"addr"`
}

func CreateRoomV2(req *CreateRoomV2Req, res *CreateRoomV2Res) error {
	if req == nil || res == nil {
		return errcode.InvalidParameterError
	}

	ireq := GetServerAddrReq{
		ServerName: "room",
	}
	ires := GetServerAddrRes{}

	serr := GetServerAddr(&ireq, &ires)
	if serr != nil {
		return serr
	}

	room_addr := fmt.Sprintf("http://%s/v1/room/inner/new_create_room", ires.InnerAddr)
	return chttp.InnerRequest(room_addr, req, res)
}
