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
	definition, err := dictionary.Search("second")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}

	err = dictionary.Add("hello", "Greeting")
	if err != nil {
		fmt.Println(err)
	}
	definition, err = dictionary.Search("hello")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}

	baseWord := "name"
	dictionary.Add(baseWord, "Nick")
	err = dictionary.Update(baseWord, "Judy")
	if err != nil {
		fmt.Println(err)
	}
	word, _ := dictionary.Search(baseWord)
	fmt.Println(word)

	dictionary.Delete("second")
	definition, err = dictionary.Search("second")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}

}
