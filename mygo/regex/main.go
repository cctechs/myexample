package main

import (
	"regexp"
	"fmt"
)

func main() {
	text := `
	email1 is aaa@q.q.com
	email2 is bbb@aa.com.cn
	email3 is bbb@aa.cn
	`
	re, err:= regexp.Compile(
		`([a-zA-Z0-9]+)@([a-zA-Z0-9.]+)(\.[a-zA-Z0-9]+)`)

	if err != nil{
		panic(err)
	}
	//match := re.FindAllString(text, -1)

	match := re.FindAllStringSubmatch(text, -1)
	fmt.Println(match)
}
