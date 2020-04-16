package main

import (
	"errors"
	"fmt"
	"io/ioutil"

	"lib/common/chttp"
	"lib/common/errcode"
	"lib/gin"
)

func main() {
	fmt.Println("开始初始化 ...")
	engine := gin.Default()
	register_router(engine)

	err := engine.Run(":80")
	if err != nil {
		fmt.Printf("err: %s\n", err.Error())
	}
	fmt.Println("终止 ...")
}

func register_router(engine *gin.Engine) error {
	if nil == engine {
		return errors.New("无效的GIN实例")
	}
	common_router := engine.Group("/qiniu/upload")
	{
		common_router.POST("/callback", qiniuCallback)
	}
	return nil
}

type CallbackRes struct {
	Name  string `json:"name"`
	Count int64  `json:"count"`
}

var Count int64

func qiniuCallback(c *gin.Context) {
	var code int = errcode.OKCode

	var body []byte
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Printf("err: %s\n", err.Error())
	}
	fmt.Printf("body: %s\n", string(body))

	Count += 1
	chttp.ServerResponse(c, code, &CallbackRes{
		Name:  "chicheng",
		Count: Count,
	})
}
