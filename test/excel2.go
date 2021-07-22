package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"frozen-go-project/utils/intx"
	"github.com/tealeg/xlsx"
	"io"
	"os"
	"strconv"
)

type BatchGreet struct {
	Time string `json:"time"`
	Info struct {
		Uid            int   `json:"uid"`
		IsNotice       int   `json:"is_notice"`
		IsSys          int   `json:"is_sys"`
		PeerUids       []int `json:"peeruids"`
		RealPeerUids   []int `json:"real_peeruids"`
		SendPercentage int   `json:"send_percentage"`
	} `json:"info"`
}

func main() {
	// 写入
	rw, err := os.Open("msg_data.txt")
	if err != nil {
		panic(err)
	}
	excelFileName := "./msg_data.xlsx"
	xlFile := xlsx.NewFile()
	sheet, err := xlFile.AddSheet("msg_data")
	row := sheet.AddRow()
	c1, c2, c3, c4, c5, c6, c7 := row.AddCell(), row.AddCell(), row.AddCell(), row.AddCell(), row.AddCell(), row.AddCell(), row.AddCell()
	c1.Value, c2.Value, c3.Value, c4.Value, c5.Value, c6.Value, c7.Value = "时间", "用户id", "是否上线通知", "是否系统发送", "发送比例", "预备发送男", "真实发送男"
	defer rw.Close()
	rb := bufio.NewReader(rw)
	for {
		line, _, err := rb.ReadLine()
		if err == io.EOF {
			break
		}
		d := new(BatchGreet)
		json.Unmarshal(line, d)
		fmt.Printf("%+v\n", d.Info)
		row := sheet.AddRow()
		c1, c2, c3, c4, c5, c6, c7 := row.AddCell(), row.AddCell(), row.AddCell(), row.AddCell(), row.AddCell(), row.AddCell(), row.AddCell()
		c1.Value, c2.Value, c3.Value, c4.Value, c5.Value, c6.Value, c7.Value = d.Time, strconv.Itoa(d.Info.Uid), strconv.Itoa(d.Info.IsNotice), strconv.Itoa(d.Info.IsSys), strconv.Itoa(d.Info.SendPercentage),
			intx.IntJoin(d.Info.PeerUids, ","), intx.IntJoin(d.Info.RealPeerUids, ",")
	}
	_ = xlFile.Save(excelFileName)
}
