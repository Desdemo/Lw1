package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"strconv"
)

type Info struct {
	Id     int
	Time   string
	Price  float64
	Client string
}

var Inn []Info

func main() {
	f, err := excelize.OpenFile("G:/go_dev/src/Lw1/excel/bk.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 设置日期格式
	style, err := f.NewStyle(`{"number_format":15}`)
	if err != nil {
		fmt.Println(err)
	}
	err = f.SetCellStyle("2019年", "B1", "B9999", style)
	// 获取工作表中指定单元格的值
	if err != nil {
		fmt.Println(err)
		return
	}
	// 获取 Sheet1 上所有单元格
	rows, err := f.GetRows("2019年")
	fmt.Println(len(rows))
	for _, row := range rows {

		info := new(Info)
		info.Id, _ = strconv.Atoi(row[0])
		info.Time = row[1]
		info.Price, _ = strconv.ParseFloat(row[2], 10)
		info.Client = row[3]
		fmt.Println(*info)

	}

}
