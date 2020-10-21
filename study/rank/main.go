package main

import (
    "fmt"
    "math/rand"
    "reflect"
    "runtime"
    "time"
)

var MAX_RANK_NUM int = 1000

type TItem struct {
    Id    int64
    Name  string
    Score int64
    Index int
}

type TRank struct {
    RankItems     []*TItem         // 排行
    RankItemByKey map[int64]*TItem // 通过 key 映射到 item
}

func UpdateRankFuncHandle(handle interface{}, totalNum int) {
    t := time.Now()
    vals := make([]reflect.Value, 3)
    vals[0] = reflect.ValueOf(&TRank{
        RankItems:     make([]*TItem, 0),
        RankItemByKey: make(map[int64]*TItem),
    })
    vals[2] = reflect.ValueOf(MAX_RANK_NUM)

    for id := 1; id <= totalNum; id++ {
        vals[1] = reflect.ValueOf(&TItem{
            Id:    int64(id),
            Name:  fmt.Sprintf("name_%d", id),
            Score: rand.Int63n(10000),
        })
        reflect.ValueOf(handle).Call(vals)
    }
    for id := totalNum / 3; id <= totalNum*2/3; id++ {
        vals[1] = reflect.ValueOf(&TItem{
            Id:    int64(id),
            Name:  fmt.Sprintf("name_%d", id),
            Score: rand.Int63n(10000),
        })
        reflect.ValueOf(handle).Call(vals)
    }

    handlePointer := reflect.ValueOf(handle).Pointer()
    fmt.Printf("spend [%+v]: %+v\n", runtime.FuncForPC(handlePointer).Name(), time.Since(t))
}

// 排序算法比较
func main() {
    totalNum := 100000
    rand.Seed(time.Now().Unix())

    UpdateRankFuncHandle(UpdateSortRank, totalNum)
    fmt.Println("=========================")

    UpdateRankFuncHandle(UpdateLastRank, totalNum)

    fmt.Println("=========================")

    UpdateRankFuncHandle(UpdateSearchRank, totalNum)

    fmt.Printf("over...\n")
}
