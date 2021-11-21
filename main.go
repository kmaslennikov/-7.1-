package main

import (
	"fmt"
)

const START = 100000
const END = 999999

func main() {
	counter := 0
	for i := START; i <= END; i++ {
		digitFirst := i / 100000
		digitLast := i % 10
		if digitFirst != digitLast {
			continue
		}
		digitSecond := (i / 10000) % 10
		digitFifth := (i / 10) % 10
		if digitSecond != digitFifth {
			continue
		}
		digitThird := (i / 1000) % 10
		digitForth := (i / 100) % 10
		if digitThird != digitForth {
			continue
		}
		fmt.Println(i)
		counter++
	}
	fmt.Println("Всего зеркальных чисел в промежутке от 100000 до 999999:", counter)
}
