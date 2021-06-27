package main

import "fmt"

func main() {
	str := "xxxxxxxxababacxxxxxxx"
	tar := "ababac"
	for i := 0; i < len(str); i++ {
		//fmt.Println(str[i : i+1])
		fmt.Println(str[i])
	}
	fmt.Println(str[8:10] == tar[0:2])
	fmt.Printf("%v %v %s \n", str[8:10], tar[0:2], tar)
}
