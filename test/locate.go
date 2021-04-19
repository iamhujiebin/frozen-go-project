package main

import (
	"encoding/json"
	"github.com/tealeg/xlsx"
	"os"
)

type CountryCode struct {
	CountryCode      string `json:"country_code"` // 国家2位代码
	MobileRegionCode string `json:"rc"`           // 电话代码
	Cn               string `json:"cn"`           // 中文
	En               string `json:"en"`           // 英文
}

var CountryCodes = make(map[string]CountryCode) // country_code 映射

func main() {
	excelFileName := "./country_code_priority.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
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
