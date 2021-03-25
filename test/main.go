package main

import (
	"time"
)

func main() {
	loc := time.FixedZone("Beijing", 5.5*3600)
	//loc2 := time.FixedZone("Beijing", -1*3600)
	beijing, _ := time.ParseInLocation("2006-01-02 15:04:05", "2021-04-12 00:00:00", loc)
	utc, _ := time.ParseInLocation("2006-01-02 15:04:05", "2021-04-12 00:00:00", time.UTC)
	//other, _ := time.ParseInLocation("2006-01-02 15:04:05", "2021-03-22 00:00:00", loc2)
	println(utc.Unix(), beijing.Unix())
	println(beijing.Weekday(), utc.Weekday())
	y, w := beijing.ISOWeek()
	println(y, w)
	y, w = utc.ISOWeek()
	println(y, w)
	//y, w = other.ISOWeek()
	//println(y, w)
	println(0xA)
}
