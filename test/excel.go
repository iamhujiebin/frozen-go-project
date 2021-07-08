package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/tealeg/xlsx"
	"io"
	"os"
	"strconv"
)

type Data struct {
	Info struct {
		Ymd         string `json:"ymd"`
		FrontCount  int    `json:"front_count"`
		LiveCount   int    `json:"live_count"`
		OnlineCount int    `json:"online_count"`
	} `json:"info"`
}

func main() {
	// 写入
	rw, err := os.Open("live_data.txt")
	if err != nil {
		panic(err)
	}
	excelFileName := "./live_data.xlsx"
	xlFile := xlsx.NewFile()
	sheet, err := xlFile.AddSheet("live_data")
	row := sheet.AddRow()
	c1, c2, c3, c4 := row.AddCell(), row.AddCell(), row.AddCell(), row.AddCell()
	c1.Value, c2.Value, c3.Value, c4.Value = "分钟", "在线人数", "开播人数", "前台人数"
	defer rw.Close()
	rb := bufio.NewReader(rw)
	for {
		line, _, err := rb.ReadLine()
		if err == io.EOF {
			break
		}
		d := new(Data)
		json.Unmarshal(line, d)
		fmt.Printf("%+v\n", d.Info)
		row := sheet.AddRow()
		c1, c2, c3, c4 := row.AddCell(), row.AddCell(), row.AddCell(), row.AddCell()
		c1.Value, c2.Value, c3.Value, c4.Value = d.Info.Ymd, strconv.Itoa(d.Info.OnlineCount), strconv.Itoa(d.Info.LiveCount), strconv.Itoa(d.Info.FrontCount)
	}
	_ = xlFile.Save(excelFileName)



	// 读取
	excelFileName = "./country_code_priority.xlsx"
	xlFile, err = xlsx.OpenFile(excelFileName)
	if err != nil {
		panic(err)
	}
	for _, sheet := range xlFile.Sheets {
		for rInx, row := range sheet.Rows {
			if rInx == 0 {
				continue
			}
			if len(row.Cells) != 5 {
				continue
			}
			cc := CountryCode{
				En:               sheet.Rows[rInx].Cells[0].Value,
				Cn:               sheet.Rows[rInx].Cells[1].Value,
				CountryCode:      sheet.Rows[rInx].Cells[2].Value,
				MobileRegionCode: sheet.Rows[rInx].Cells[3].Value,
			}
			if len(cc.CountryCode) > 0 {
				CountryCodes[cc.CountryCode] = cc
			}
		}
	}
	b, _ := json.Marshal(CountryCodes)
	println(string(b))
	file, _ := os.Create("country_code.json")
	file.Write(b)
	file.Close()
}
