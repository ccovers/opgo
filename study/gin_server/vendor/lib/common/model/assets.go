package model

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"math/rand"
	"time"
)

type AssetTransLog struct {
	Id         int64  `gorm:"column:id"`
	TransId    string `gorm:"column:trans_id"`
	AssetType  int64  `gorm:"column:asset_type"`
	AssetCount int64  `gorm:"column:asset_count"`
	Uid        string `gorm:"column:uid"`
	Status     int64  `gorm:"column:status"`
}

func (self AssetTransLog) TableName() string {
	return "assets_trans_log"
}

type AssetsOperLog struct {
	Id         int64  `gorm:"column:id"`
	Uid        string `gorm:"column:uid"`
	OperId     string `gorm:"column:oper_id"`
	AssetType  int    `gorm:"column:asset_type"`
	AssetCount int    `gorm:"column:asset_count"`
}

func (self AssetsOperLog) TableName() string {
	return "assets_oper_log"
}

type Assets struct {
	Id      int64  `gorm:"column:id"`
	Uid     string `gorm:"column:uid"`
	Nick    string `gorm:"column:nick"`
	Icon    string `gorm:"column:icon"`
	Vip     int    `gorm:"column:vip"`
	Gold    int    `gorm:"column:gold"`
	Dimond  int    `gorm:"column:dimond"`
	RCoin   int    `gorm:"column:r_coin"`
	Fuhuobi int    `gorm:"column:fuhuobi"`
}

func (self Assets) TableName() string {
	return "assets"
}

func GetUserOperID(uid string, assetType int, assetCount int) string {
	var oper_id string

	for ok := true; ok; ok = false {
		tm := time.Now().UnixNano()
		rand.Seed(tm)

		oper_id = fmt.Sprintf("{0+-%s#%d#%d#%d#%d-+0}", uid, assetType, assetCount, tm, rand.Intn(1000))

		hash := sha1.New()
		hash.Write([]byte(oper_id))

		oper_id = base64.StdEncoding.EncodeToString(hash.Sum(nil))
	}

	return oper_id
}
