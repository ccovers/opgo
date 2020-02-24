package qiniu

import (
	"github.com/qiniu/api.v7/auth/qbox"
)

const (
	accessKey   = "xxx"
	secretKey   = "xxx"
	storeRegion = "xxx"
)

const (
	imgBucket = "imgbucket"                // 设置路径
	imgUrl    = "http://img.qiniu.xxx.com" // 设置域名
)

type QiniuStore struct {
	accessKey   string
	secretKey   string
	url         string
	bucket      string
	storeRegion string
	mac         *qbox.Mac
	expire      uint32
	mediaType   int8
}

var imgStore *QiniuStore

func init() {
	imgStore = newStore(imgBucket, imgUrl, msgTypeImg)
}

// 创建qiniu存储配置实例，返回默认配置
func newStore(bucket, url string, mediaType int8) *QiniuStore {
	return &QiniuStore{
		accessKey:   accessKey,
		secretKey:   secretKey,
		url:         url,
		bucket:      bucket,
		storeRegion: storeRegion,
		mac:         qbox.NewMac(accessKey, secretKey),
		expire:      600,
		mediaType:   mediaType,
	}
}

func Img() *QiniuStore {
	return imgStore
}

func (cfg *QiniuStore) setExpire(expireInSecond uint32) {
	cfg.expire = expireInSecond
}
