package qiniu

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/qiniu/api.v7/storage"
)

/*
* UpToken   是由业务服务器颁发的上传凭证。
* QiniuURL	qiniu路径标识，以 "qiniu://" 开头
* Region	qiniu服务器区域
* Key       是要上传的文件访问路径。比如："foo/bar.jpg"。
* 			注意我们建议 key 不要以 '/' 开头。另外，key 为空字符串是合法的。
 */
type tokenCfg struct {
	Token  string `json:"token"`
	Url    string `json:"url"`
	Region string `json:"region"`
	Key    string `json:"key"`
}

// TokenCfg 获取上传凭证、等相关信息
func (store *QiniuStore) TokenCfg(userId int64) *tokenCfg {
	key := fmt.Sprintf("%s/%d/%s%s",
		store.parentPath(), userId, entropy(), suffixMap[store.mediaType])
	return &tokenCfg{
		Token:  store.token(key),
		Url:    fmt.Sprintf("qiniu://%s/%s", store.bucket, key),
		Region: store.storeRegion,
		Key:    key,
	}
}

// 获取上传凭证
// suffix 文件后缀 比如 jpg、png等
func (store *QiniuStore) TokenCfgBySuffix(userId int64, suffix string) *tokenCfg {
	suffix = strings.Replace(suffix, " ", "", -1)
	if len(suffix) > 0 && suffix[0] != '.' {
		suffix = "." + suffix
	}

	key := fmt.Sprintf("%s/%d/%s%s",
		store.parentPath(), userId, entropy(), suffix)
	return &tokenCfg{
		Token:  store.token(key),
		Url:    fmt.Sprintf("qiniu://%s/%s", store.bucket, key),
		Region: store.storeRegion,
		Key:    key,
	}
}

// 通过qiniu地址标识获取http地址
func (store *QiniuStore) HttpUrl(name string) string {
	if !strings.HasPrefix(name, "qiniu://") {
		return name
	}
	name = strings.TrimPrefix(name, "qiniu://")
	pos := strings.Index(name, "/")
	if pos != -1 {
		name = string(name[pos+1:])
	}

	deadline := time.Now().Add(time.Second * time.Duration(store.expire)).Unix()
	return storage.MakePrivateURL(store.mac, store.url, name, deadline)
}

// 通过http地址获取qiniu地址标识
func (store *QiniuStore) QiniuUrl(uri string) (string, error) {
	if uri[:8] == "qiniu://" {
		return uri, nil
	}
	u, err := url.Parse(uri)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf(`qiniu://%s%s`, store.bucket, u.Path), nil
}
