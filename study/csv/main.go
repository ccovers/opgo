package main

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
	convertXlsx2CsvDemo1()
}

// 将xlsx转换为yaml文件的简单例子
func convertXlsx2CsvDemo1() {
	xlsx, err := excelize.OpenFile("./veh_models.xlsx")
	if err != nil {
		fmt.Println("open:", err)
		return
	}
	f, err := os.Create("./veh_models.csv")
	if err != nil {
		fmt.Println("create:", err)
		return
	}
	defer f.Close()

	r := csv.NewWriter(f)

	// Get all the rows in the Sheet1.
	res, err := xlsx.GetRows("明觉+汽车之家合集")
	if err != nil {
		fmt.Println("getrows:", err)
		return
	}
	fmt.Println("execl行数:", len(res))

	rows := make([][]string, 0, len(res))
	// 过滤重复行
	for i, e := range res {
		for j, r := range res {
			if i > j && e[0] == r[0] && e[1] == r[1] && e[2] == r[2] && e[3] == r[3] && e[4] == r[4] && e[5] == r[5] {
				goto next
			}
		}
		rows = append(rows, e)
	next:
	}
	for i, row := range rows {
		if i > 0 {
			r.Write(row)
		}
	}
	r.Flush()
}
