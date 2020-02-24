package model

type BroadcastType int64

const (
	Paomadeng BroadcastType = 1
	LiaoTian
)

type BroadcastInfo struct {
	Btype  BroadcastType
	Uid    string
	Sender string
	Info   string
}
