package qiniu

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"fmt"

	"github.com/qiniu/api.v7/storage"
)

const (
	msgTypeText  int8 = 1
	msgTypeImg   int8 = 2
	msgTypeAudio int8 = 3
)

var suffixMap = map[int8]string{
	msgTypeImg:   ".jpg",
	msgTypeAudio: ".mp3",
}

var pathMap = map[int8]string{
	msgTypeImg:   "img",
	msgTypeAudio: "audio",
}

// 通过文件名获取上传下载凭证
func (store *QiniuStore) token(key string) string {
	scope := store.bucket
	if key != "" {
		scope = fmt.Sprintf("%s:%s", store.bucket, key)
	}

	putPolicy := storage.PutPolicy{
		Scope: scope,
		ReturnBody: `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),
					"bucket":"$(bucket)","name":"$(x:name)"}`,
		Expires: store.expire,
	}
	token := putPolicy.UploadToken(store.mac)
	return token
}

// 获取管理凭证
func (store *QiniuStore) accessToken(path string, body []byte) string {
	//url := "http://rs.qiniu.com/move/bmV3ZG9jczpmaW5kX21hbi50eHQ=/bmV3ZG9jczpmaW5kLm1hbi50eHQ="
	//#则待签名的原始字符串是：
	//signingStr := "/move/bmV3ZG9jczpmaW5kX21hbi50eHQ=/bmV3ZG9jczpmaW5kLm1hbi50eHQ=\n"

	signingStr := fmt.Sprintf("%s\n%s", path, string(body))
	//hmac ,use sha1
	mac := hmac.New(sha1.New, []byte(store.secretKey))
	mac.Write([]byte(signingStr))
	sign := mac.Sum(nil)

	encodedSign := base64.StdEncoding.EncodeToString([]byte(sign))

	accessToken := fmt.Sprintf("%s:%s", store.accessKey, encodedSign)
	return accessToken
}

func (store *QiniuStore) parentPath() string {
	if isProductionEnv() {
		return "user"
	} else {
		return fmt.Sprintf("test/%s/user", pathMap[store.mediaType])
	}
}

func isProductionEnv() bool {
	if "test" == "production" {
		return true
	}
	return false
}

func entropy() string {
	b := make([]byte, 20)
	rand.Read(b)
	h := sha1.Sum(b)
	return hex.EncodeToString(h[:])
}

// 获取文件信息
func (store *QiniuStore) fileInfo(key string) string {

	// 指定空间所在的区域，如果不指定将自动探测
	// 如果没有特殊需求，默认不需要指定
	config := storage.Config{
		// 是否使用https域名进行资源管理
		UseHTTPS: false,
		//store.Zone=&storage.ZoneHuabei
	}
	bucketManager := storage.NewBucketManager(store.mac, &config)

	fileInfo, sErr := bucketManager.Stat(store.bucket, key)
	if sErr != nil {
		fmt.Println(sErr)
		return ""
	}
	return fileInfo.String()
}

func (store *QiniuStore) Upload(userId int64, body []byte, zone *storage.Zone) (string, error) {
	if zone == nil {
		zone = &storage.ZoneHuabei
	}
	cfg := storage.Config{
		Zone:          zone,
		UseHTTPS:      false,
		UseCdnDomains: false,
	}
	putExtra := storage.PutExtra{
		Params: map[string]string{},
	}
	tokenCfg := store.TokenCfg(userId)

	formUploader := storage.NewFormUploader(&cfg)
	err := formUploader.Put(context.Background(), nil, tokenCfg.Token, tokenCfg.Key,
		bytes.NewReader(body), int64(len(body)), &putExtra)
	if err != nil {
		return tokenCfg.Url, err
	}
	return tokenCfg.Url, nil
}
