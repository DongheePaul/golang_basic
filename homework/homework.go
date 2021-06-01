package main

import (
	"fmt"
)

func predict() (temp int, rainPercent int){
	temp = 30
	rainPercent = 10
	return
}

func main() {
	if temp, rainPercent := predict(); temp >= 25 && rainPercent >= 80{
		fmt.Println("덥고 비가 옵니다")
	}else if temp >=25 && rainPercent >= 20{
		fmt.Println("덥고 습합니다.")
	}else if temp >=25 && rainPercent < 20 {
		fmt.Println("야외 활동하기 좋습니다.")
	}else if temp < 10 || rainPercent >= 80 {
		fmt.Println("야외 활동하기 좋지 않습니다")
	}else {
		fmt.Print("좋은 날씨입니다.")
	}
}