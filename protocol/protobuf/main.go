package main

import (
	"log"

	//"ccovers/opgo/protocol/protobuf/pbProto/user"
	"github.com/golang/protobuf/proto"

	_ "github.com/ccovers/opgo/protocol/protobuf/common/pbProto/pb_class"
	"github.com/ccovers/opgo/protocol/protobuf/common/pbProto/pb_user"
)

func main() {
	user := pb_user.User{
		Id:   1,
		Name: "xiaoming",
		Pos: &pb_user.Pos{
			Left:  "xiaohong",
			Right: "xiaojun",
		},
		Pens: []int32{1, 2, 3},
	}

	data, err := proto.Marshal(&user)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	bak := pb_user.User{}
	err = proto.Unmarshal(data, &bak)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}
	log.Printf("%v\n", user)
	log.Printf("%v\n", bak)
}
