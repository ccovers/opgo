package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"runtime/debug"
	"time"

	"github.com/ccovers/opgo/protocol/grpc/test"
	"github.com/golang/glog"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

func toJson(obj interface{}) string {
	breq, err := json.Marshal(obj)
	if err != nil {
		return fmt.Sprintf("json err = %v", err)
	} else {
		return string(breq)
	}
}

func LogReq(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	sreq := toJson(req)
	defer func(startT time.Time) {
		duration := time.Now().Sub(startT) / time.Millisecond
		sresp := toJson(resp)
		glog.Infof("%s, %dms, req=%s, resp err=%v, sresp=%s", info.FullMethod, duration, sreq, err, sresp)
		if err != nil && status.Code(err) == codes.Internal {
			glog.Infof("stack trace: %s", string(debug.Stack()))
		}
	}(time.Now())
	return handler(ctx, req)
}

/*
func RecoverReq(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = status.Errorf(codes.Internal, "panic, req id: %s, recovered:%s, stack:%s", rpc_ctx.ReqId(ctx), r, debug.Stack())
		}
	}()

	return handler(ctx, req)
}*/

// 业务实现方法的容器
type server struct{}

// 为server定义 DoMD5 方法 内部处理请求并返回结果
// 参数 (context.Context[固定], *test.Req[相应接口定义的请求参数])
// 返回 (*test.Res[相应接口定义的返回参数，必须用指针], error)
func (s *server) DoMD5(ctx context.Context, in *test.Req) (*test.Res, error) {
	glog.Infof("MD5方法请求: %d", in.Id)
	return &test.Res{Name: "MD5 :" + fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%d", in.Id))))}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080") //监听所有网卡8028端口的TCP连接
	if err != nil {
		log.Fatalf("监听失败: %v", err)
	}
	s := grpc.NewServer(grpc.UnaryInterceptor(LogReq)) //创建gRPC服务

	/**注册接口服务
	 * 以定义proto时的service为单位注册，服务中可以有多个方法
	 * (proto编译时会为每个service生成Register***Server方法)
	 * 包.注册服务方法(gRpc服务实例，包含接口方法的结构体[指针])
	 */
	test.RegisterWaiterServer(s, &server{})

	/**如果有可以注册多个接口服务,结构体要实现对应的接口方法
	 * user.RegisterLoginServer(s, &server{})
	 * minMovie.RegisterFbiServer(s, &server{})
	 */
	// 在gRPC服务器上注册反射服务
	reflection.Register(s)
	// 将监听交给gRPC服务处理
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
