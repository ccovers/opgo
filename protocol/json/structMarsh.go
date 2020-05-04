package main

import (
	"encoding/json"
	"fmt"
)

type Id int64

func (id Id) MarshalJSON() ([]byte, error) {
	ids := []int64{int64(id + 1), int64(id + 2), int64(id + 3)}
	names := []string{"aaa", "bbb", "ccc"}

	return json.Marshal(struct {
		Ids   []int64  `json:"ids"`
		Names []string `json:"names"`
	}{
		ids, names,
	})
}

func main() {

	bs, err := json.Marshal(struct {
		Id    Id
		Score int64
	}{
		1,
		100,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s\n", string(bs))
}
