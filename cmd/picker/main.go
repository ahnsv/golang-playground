package main

import (
	"fmt"

	"github.com/ahnsv/golang-playground/internals/picker"
)

func main() {
	rank := make(map[string]int)

	for i := 0; i < 1000000; i++ {
		if i%200000 == 0 {
			fmt.Printf("%d\n", 5-(i/200000))
		}
		presenter := picker.PickOne([]string{"humphrey", "rane", "spock", "bada", "matt", "inclue"}, picker.ParseDateString("2006-01-02", "2021-05-28").UnixNano())
		rank[presenter] += 1
	}

	var max = 0
	var nextPresenter = ""
	for k, v := range rank {
		if max < v {
			max = v
			nextPresenter = k
		}
	}

	fmt.Printf("다음 발표자는 [%s] 입니다\n", nextPresenter)
}
