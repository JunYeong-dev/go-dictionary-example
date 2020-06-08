package main

import (
	"fmt"

	"github.com/JunYeong-dev/dictionary/mydict"
)

func main() {
	// map을 정의함과 동시에 값을 넣어줄 수도 있고
	dictionary := mydict.Dictionary{"first": "First word"}
	// 이렇게 따로 넣어줄 수도 있음
	dictionary["second"] = "Second word"
	definition, err := dictionary.Search("third")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}
}
