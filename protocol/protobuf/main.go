package main

import (
	"log"

	pb "ccovers/opgo/protocol/protobuf/proto"
	"github.com/golang/protobuf/proto"
)

type Test struct {
	Label string
	Type  int
	Reps  []int64
}

func main() {
	user := pb.User{
		Id:   2,
		Name: "xiaoming",
		Pos: &pb.Pos{
			Left:  "xiaohong",
			Right: "xiaojun",
		},
		Book: 3,
		Pens: []int32{1, 2, 3},
	}

	data, err := proto.Marshal(&user)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	bak := pb.User{}
	err = proto.Unmarshal(data, &bak)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}
	log.Printf("%v\n", user)
	log.Printf("%v\n", bak)
}
