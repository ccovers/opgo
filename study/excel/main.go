package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
)

var chengdus = []string{
	"锦江", "青羊", "金牛", "武侯", "成华", "双流",
	"新都", "龙泉", "温江", "郫",
}
var dachengdus = []string{
	"青白江", "金堂", "大邑", "蒲江", "新津",
	"邛崃", "崇州", "彭州", "都江堰",
}

func main() {
	/*if len(os.Args) != 3 {
		fmt.Printf("输入错误: %s\n请出入: opgo.exe 分诊明细表 分析输出表", os.Args)
		return
	}
	rName := os.Args[1]
	wName := os.Args[2]*/
	rName := "分诊明细表"
	wName := "分析输出表"
	if len(os.Args) > 3 {
		rName = os.Args[1]
		wName = os.Args[2]
	}
	rName = strings.Replace(rName, " ", "", -1)
	wName = strings.Replace(wName, " ", "", -1)
	if rName == `` {
		rName = "分诊明细表"
	}
	if wName == `` {
		wName = "分析输出表"
	}

	fds, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Printf("read dir err: %s\n", err.Error())
	}

	for _, fd := range fds {
		if fd.IsDir() {
			continue
		}
		err := handleExcel(fd.Name(), rName, wName)
		if err != nil {
			fmt.Printf("%s\n", err.Error())
			var temp string
			fmt.Scan(&temp)
		}
	}
	fmt.Printf("读取表: %s\n写入到: %s\n请输入任意字符并敲击回车关闭\n", rName, wName)
	var temp string
	fmt.Scan(&temp)
}

func handleExcel(fileName string, rName, wName string) error {
	fmt.Printf("打开: %s\n", fileName)
	file, err := excelize.OpenFile(fileName)
	if err != nil {
		fmt.Printf("错误: %s\n\n", err.Error())
		return nil
	}
	rows, err := file.GetRows(rName)
	if err != nil {
		return fmt.Errorf("\n读取表[%s   -   %s]错误: %s\n\n", fileName, rName, err.Error())
	}

	var date time.Time
	wayMap := map[time.Time]map[string]map[string]int{}
	waysMap := make(map[string]struct{})
	for k, row := range rows {
		if k == 0 {
			fmt.Printf("%s\n", strings.Join(row, " "))
			continue
		}
		if len(row) > 2 && row[1] != `` {
			date, err = excelDateToDate(row[1])
			if err != nil {
				return fmt.Errorf("第[%d]行, 日期错误[%s]: %s\n", k, row[1], err.Error())
			}
			if _, ok := wayMap[date]; !ok {
				wayMap[date] = map[string]map[string]int{}
			}
		}
		if len(row) > 8 && row[2] != `` && row[4] != `` {
			if _, ok := wayMap[date][row[8]]; !ok {
				wayMap[date][row[8]] = map[string]int{}
			}
			wayMap[date][row[8]][row[7]]++
			waysMap[row[8]] = struct{}{}
		}
	}
	ways := make([]string, 0)
	for way, _ := range waysMap {
		ways = append(ways, way)
	}
	sort.Strings(ways)
	writeExcel(file, wayMap, ways, wName)
	return nil
}

type TimeDef []time.Time

func (a TimeDef) Len() int           { return len(a) }
func (a TimeDef) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a TimeDef) Less(i, j int) bool { return a[i].Unix() < a[j].Unix() }

func writeExcel(f *excelize.File, wayMap map[time.Time]map[string]map[string]int, ways []string, wName string) {
	sheetName := wName
	f.DeleteSheet(sheetName)
	index := f.NewSheet(sheetName)

	dates := TimeDef{}
	for date, _ := range wayMap {
		dates = append(dates, date)
	}
	sort.Sort(dates)

	line := 1
	var moth time.Month
	mothTotalNum := map[string]int{}
	for _, date := range dates {
		if moth != date.Month() {
			moth = date.Month()
			mothTotalNum = map[string]int{}
		}
		if _, ok := wayMap[date]; !ok {
			continue
		}
		//fmt.Printf("读取 %s\n", date.String())
		line++
		f.SetCellValue(sheetName, fmt.Sprintf("B%d", line), "途径")
		f.SetCellValue(sheetName, fmt.Sprintf("C%d", line), fmt.Sprintf("%d月%d日病人来源", date.Month(), date.Day()))
		line++
		f.SetCellValue(sheetName, fmt.Sprintf("B%d", line), "来源途径")
		f.SetCellValue(sheetName, fmt.Sprintf("C%d", line), "人数")
		f.SetCellValue(sheetName, fmt.Sprintf("D%d", line), "大成都")
		f.SetCellValue(sheetName, fmt.Sprintf("E%d", line), "成都")
		f.SetCellValue(sheetName, fmt.Sprintf("F%d", line), "人数月累计")
		f.SetCellValue(sheetName, fmt.Sprintf("G%d", line), "大成都月累计")
		f.SetCellValue(sheetName, fmt.Sprintf("H%d", line), "成都月累计")
		line++

		dayNum := 0
		chengdudayNum := 0
		dachengdudayNum := 0
		for _, way := range ways {
			areaMap, ok := wayMap[date][way]
			if !ok {
				continue
			}
			areaNum := 0
			chengduNum := 0
			dachengduNum := 0
			for area, num := range areaMap {
				dayNum += num
				areaNum += num
				if isInChengdu(area) {
					chengduNum += num
					chengdudayNum += num
				}
				if isInDaChengdu(area) {
					dachengduNum += num
					dachengdudayNum += num
				}
			}
			mothTotalNum[way] += areaNum
			mothTotalNum[way+"chengdu"] += chengduNum
			mothTotalNum[way+"dachengdu"] += dachengduNum
			mothNum, _ := mothTotalNum[way]
			chengdumothNum, _ := mothTotalNum[way]
			dachegndumothNum, _ := mothTotalNum[way]

			f.SetCellValue(sheetName, fmt.Sprintf("B%d", line), way)
			f.SetCellValue(sheetName, fmt.Sprintf("C%d", line), areaNum)
			f.SetCellValue(sheetName, fmt.Sprintf("D%d", line), dachengduNum)
			f.SetCellValue(sheetName, fmt.Sprintf("E%d", line), chengduNum)
			f.SetCellValue(sheetName, fmt.Sprintf("F%d", line), mothNum)
			f.SetCellValue(sheetName, fmt.Sprintf("G%d", line), dachegndumothNum)
			f.SetCellValue(sheetName, fmt.Sprintf("H%d", line), chengdumothNum)
			line++
		}
		mothTotalNum["合计"] += dayNum
		mothTotalNum["合计"+"chengdu"] += chengdudayNum
		mothTotalNum["合计"+"dachengdu"] += dachengdudayNum
		mothNum, _ := mothTotalNum["合计"]
		chenggumothNum, _ := mothTotalNum["合计"+"chengdu"]
		dachengdumothNum, _ := mothTotalNum["合计"+"dachengdu"]
		f.SetCellValue(sheetName, fmt.Sprintf("B%d", line), "合计")
		f.SetCellValue(sheetName, fmt.Sprintf("C%d", line), dayNum)
		f.SetCellValue(sheetName, fmt.Sprintf("D%d", line), dachengdudayNum)
		f.SetCellValue(sheetName, fmt.Sprintf("E%d", line), chengdudayNum)
		f.SetCellValue(sheetName, fmt.Sprintf("F%d", line), mothNum)
		f.SetCellValue(sheetName, fmt.Sprintf("G%d", line), dachengdumothNum)
		f.SetCellValue(sheetName, fmt.Sprintf("H%d", line), chenggumothNum)
		line++
	}

	// Set active sheet of the workbook.
	// Save spreadsheet by the given path.
	f.SetActiveSheet(index)
	if err := f.Save(); err != nil {
		fmt.Printf("save err: %s\n", err.Error())
	}
}

func excelDateToDate(excelDate string) (time.Time, error) {
	excelTime := time.Date(1899, time.December, 30, 0, 0, 0, 0, time.UTC)
	days, err := strconv.Atoi(excelDate)
	if err != nil {
		moths := strings.Split(excelDate, "月")
		if len(moths) < 2 {
			return excelTime, fmt.Errorf("not x月x日")
		}
		month, err := strconv.Atoi(moths[0])
		if err != nil {
			return excelTime, err
		}
		days := strings.Split(moths[len(moths)-1], "日")
		if len(days) == 0 {
			return excelTime, fmt.Errorf("not x月x日")
		}
		day, err := strconv.Atoi(days[0])
		if err != nil {
			days = strings.Split(moths[len(moths)-1], "号")
			if len(days) == 0 {
				return excelTime, fmt.Errorf("not x月x日")
			}
			day, err = strconv.Atoi(days[0])
			if err != nil {
				return excelTime, err
			}
		}

		return time.Date(time.Now().Year(), time.Month(month), day, 0, 0, 0, 0, excelTime.Location()), nil
	}
	return excelTime.Add(time.Second * time.Duration(days*86400)), nil
}

func isInChengdu(area string) bool {
	for _, v := range chengdus {
		if strings.Contains(area, v) {
			return true
		}
	}
	return false
}

func isInDaChengdu(area string) bool {
	for _, v := range dachengdus {
		if strings.Contains(area, v) {
			return true
		}
	}
	return false
}
