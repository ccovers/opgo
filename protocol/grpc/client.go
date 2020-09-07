package main

import (
	"fmt"
	"log"
	//"os"

	"github.com/ccovers/opgo/protocol/grpc/common/pbProto/pb_user"
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	// 建立连接到gRPC服务
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	// 函数结束时关闭连接
	defer conn.Close()

	// 创建Waiter服务的客户端
	client := pb_user.NewServerClient(conn)

	/*// 模拟请求数据
	res := "test123"
	// os.Args[1] 为用户执行输入的参数 如：go run ***.go 123
	if len(os.Args) > 1 {
		res = os.Args[1]
	}*/

	userReq := pb_user.CS_UserInfo_Req{Id: 1}
	data, err := proto.Marshal(&userReq)
	if err != nil {
		log.Fatalf("mashal err: %+v", err)
		return
	}
	// 调用gRPC接口

	resp, err := client.DoReq(context.Background(), &pb_user.Cmd_Req{
		Cmd:  pb_user.EnCmdID_CS_UserInfo,
		Data: data,
	})
	if err != nil {
		log.Fatalf("get user info err: %+v", err)
		return
	}
	if resp.Ret != pb_user.EUserRet_Success {
		log.Fatalf("get user info failed: %+v", pb_user.EUserRet_name[int32(resp.Ret)])
		return
	}

	userResp := pb_user.SC_UserInfo_Resp{}
	err = proto.Unmarshal(resp.Data, &userResp)
	if err != nil {
		log.Fatalf("unmashal err: %+v", err)
		return
	}
	fmt.Printf("服务端响应: %+v\n", userResp)
}
