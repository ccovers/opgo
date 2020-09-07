package common

import (
    "encoding/json"
    "fmt"
    "reflect"
    "runtime/debug"
    "time"

    "github.com/ccovers/opgo/protocol/grpc/common/pbProto/pb_user"
    "github.com/golang/glog"
    "golang.org/x/net/context"
    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
)

func GrpcFunc(
    ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler,
) (resp interface{}, err error) {
    toJson := func(obj interface{}) string {
        breq, err := json.Marshal(obj)
        if err != nil {
            return fmt.Sprintf("json err = %v", err)
        } else {
            return string(breq)
        }
    }

    defer func(startT time.Time) {
        if r := recover(); r != nil {
            err = status.Errorf(codes.Internal, "panic, recovered:%s, stack:%s", r, debug.Stack())
            return
        }

        duration := time.Now().Sub(startT) / time.Millisecond
        glog.Infof("compete: %s, %dms, req=%s, err=%v, resp=%s", info.FullMethod, duration,
            toJson(req), err, toJson(resp))
        if err != nil && status.Code(err) == codes.Internal {
            glog.Infof("stack trace: %s", string(debug.Stack()))
        }
    }(time.Now())
    return handler(ctx, req)
}

// 业务实现方法的容器
type server struct {
    GFuncMap map[pb_user.EnCmdID]reflect.Value
}

func NewServer() *server {
    return &server{
        GFuncMap: map[pb_user.EnCmdID]reflect.Value{},
    }
}

func (s *server) RegMsgProc(subId pb_user.EnCmdID, procFunc interface{}) {
    s.GFuncMap[subId] = reflect.ValueOf(procFunc)
}

func (s *server) DoReq(ctx context.Context, in *pb_user.Cmd_Req) (*pb_user.Cmd_Resp, error) {
    procFunc, ok := s.GFuncMap[in.Cmd]
    if !ok {
        return &pb_user.Cmd_Resp{
            Ret: pb_user.EUserRet_Invalid,
        }, nil
    }
    data, err := procFuncCall(procFunc, in.Data)
    if err != nil {
        return &pb_user.Cmd_Resp{
            Ret: pb_user.EUserRet_Failed,
        }, err
    }
    return &pb_user.Cmd_Resp{
        Ret:  pb_user.EUserRet_Success,
        Data: data,
    }, nil
}

func procFuncCall(procFunc reflect.Value, data []byte) ([]byte, error) {
    vals := make([]reflect.Value, 0, 2)
    vals = append(vals, reflect.ValueOf(data))
    rets := procFunc.Call(vals)
    if len(rets) != 2 {
        return nil, fmt.Errorf("return param len small than 2")
    }
    redata, ok := rets[0].Interface().([]byte)
    if !ok {
        return nil, fmt.Errorf("return param not bytes")
    }
    var err error
    if rets[1].Interface() != nil {
        err, ok = rets[1].Interface().(error)
        if !ok {
            return nil, fmt.Errorf("return param not error")
        }
    }
    return redata, err
}
