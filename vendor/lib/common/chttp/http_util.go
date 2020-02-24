/*
 * author:        liujun
 * created:       2018-03-20 17:13
 * last modified: 2018-03-20 17:13
 * filename:      http_util.go
 * description:   http相关操作的公用函数
 */
package chttp

import (
	"github.com/pkg/errors"
)

import (
	"lib/common/clog"
	"lib/common/errcode"
	"lib/gin"
)

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

/*
 * author:      liujun
 * brief:       读取GIN body中的json报文
 * update:      2016-08-04 09:18
 *
 * param:
 *              c                               gin.Context
 *              obj                             json unmarshal对象
 * return:
 *              error                   错误类型
 * note:
 */
func ReadRequestBodyJson(c *gin.Context, obj interface{}) error {
	var body []byte
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, obj); err != nil {
		return err
	}

	return nil
}

type Response struct {
	State int         `json:"state"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data"`
}

func ParseRes(res string) (Response, error) {
	ret := Response{}
	var err error = nil
	if err = json.Unmarshal([]byte(res), &ret); err != nil {
		return Response{}, err
	}

	return ret, nil
}

/*
 * author:	liujun
 * brief:	响应http请求的公共函数
 * update:	2018-03-21 13:29
 *
 */
func ServerResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(http.StatusOK,
		struct {
			State int         `json:"state"`
			Msg   string      `json:"msg"`
			Data  interface{} `json:"data"`
		}{
			State: code,
			Msg:   errcode.GetErrMsg(code),
			Data:  data,
		})
}

/*
 * author:	liujun
 * brief:	通用对外http POST请求
 * update:	2018-03-21 14:59
 *
 * param:
 *		url				请求地址
 *		data			请求数据
 * return:
 *		string			返回数据
 *		error			错误
 * note:
 *		不可用于对内，没有增加超时检查等异常处理
 */
func PostUrl(url string, data string) (string, error) {
	client := &http.Client{
		Transport: &http.Transport{
			DisableKeepAlives: true,
			Dial: func(netw, addr string) (net.Conn, error) {
				deadline := time.Now().Add(10 * time.Second)
				c, err := net.DialTimeout(netw, addr, time.Second*2)
				if err != nil {
					return nil, err
				}
				c.SetDeadline(deadline)
				return c, nil
			},
		},
	}
	buf := bytes.NewBuffer([]byte(data))
	req, err := http.NewRequest("POST", url, buf)
	if err != nil {
		return "", err
	}

	rs, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer rs.Body.Close()
	rtData, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		return "", err
	}
	str := string(rtData)

	return str, nil
}

func PostUrlWithAuth(url string, tokenStr string, data string) (string, error) {
	client := &http.Client{}
	buf := bytes.NewBuffer([]byte(data))
	req, err := http.NewRequest("POST", url, buf)
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+tokenStr)

	rs, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer rs.Body.Close()
	rtData, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		return "", err
	}
	str := string(rtData)

	return str, nil
}

/*
 * author:	liujun
 * brief:	Get请求获取json方法
 * update:	2018-03-30 15:18
 *
 * note:
 */
func GetUrlJsonData(url string, obj interface{}) error {
	data, err := GetUrl(url)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, obj)
}

func GetUrl(url string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	rs, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer rs.Body.Close()
	data, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		clog.Logger.Error("request %s failed:%v", url, err)
		return nil, err
	}

	return data, nil
}

// func GetUrl(url string) (*http.Response, error) {
// 	client := &http.Client{}
// 	req, err := http.NewRequest("GET", url, nil)
// 	if err != nil {
// 		return nil, err
// 	}
// 	rs, err := client.Do(req)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return rs, nil
// }

/*
 * author:	liujun
 * brief:
 * update:	2018-05-09 21:14
 *
 * param:
 * return:
 * note:
 */
func InnerRequest(url string, req interface{}, res interface{}) error {
	req_str, _ := json.Marshal(req)
	res_str, req_err := PostUrl(url, string(req_str))
	if req_err != nil {
		return req_err
	}
	response, parse_err := ParseRes(res_str)
	if parse_err != nil {
		return parse_err
	}
	if response.State != errcode.OKCode {
		return errors.New(errcode.GetErrMsg(response.State))
	}
	tmp_str, _ := json.Marshal(response.Data)
	unmar_err := json.Unmarshal([]byte(tmp_str), res)
	if unmar_err != nil {
		return unmar_err
	}

	return nil
}

func InnerRequestWithAuth(url string, tokenStr string, req interface{}, res interface{}) error {
	req_str, _ := json.Marshal(req)
	res_str, req_err := PostUrlWithAuth(url, tokenStr, string(req_str))
	if req_err != nil {
		return req_err
	}
	response, parse_err := ParseRes(res_str)
	if parse_err != nil {
		return parse_err
	}
	if response.State != errcode.OKCode {
		return errors.New(errcode.GetErrMsg(response.State))
	}
	tmp_str, _ := json.Marshal(response.Data)
	unmar_err := json.Unmarshal([]byte(tmp_str), res)
	if unmar_err != nil {
		return unmar_err
	}

	return nil
}
