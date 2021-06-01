package main

import (
	"fmt"
)

//전역변수 -> Data segment
var cnt int = 0


func IncreaseAndReturn() int {
	fmt.Println("IncreaseAndReturn()", cnt)
	cnt++
	return cnt
}

func UploadFile() (filename string, success bool) {
	filename="hello.exe"
	success = true
	return
}

func rainPercent80(p int) bool {
	if p > 80 {
		return true
	}else {
		return false
	}
}

func factorial (i int) int {
	if i == 0 {
		return 1
	}
	return i * factorial(i -1)
}
func main() {
	//지역변수 -> Stack segment
	// tmp := 30
	// //tmp 정수를 받겠다.
	// fmt.Scanf("%d", &tmp)
	// if tmp >= 28 {
	// 	fmt.Println("turn the airconditioner")

	// }else if tmp < 5{
	// 	fmt.Println("turn of the airconditioner")
	// }else {
	// 	fmt.Print("The airconditioner is operating.")
	// }
	// fmt.Println("terminate")


	//IncreaseAndReturn => cnt(int)를 리턴. 그러나 왼쪽이 벌써 false(거짓)이므로 && 오른쪽 부분은 수행 안함.
	// if false && IncreaseAndReturn() < 5{
	// 	fmt.Println("1 증가")
	// }


	//if 초기문; 판단식
	// if filename, success := UploadFile(); success {
	// 	fmt.Println("Upload success", filename)
	// } else {
	// 	fmt.Println("Upload failed", filename)
	// }
	
	//var f int
	//값을 입력 받음
	// fmt.Println("Please type factorial value")
	// fmt.Scanf("%d", &f)
	// fmt.Printf("factorial value is => %d:", f, factorial(f))
	
	var a int = 95
	var b int = 99
	//조건문이 위에서부터 아래로 판단하기 떄문에 잘 설계해야함. 조건문에 , 쓰면 or 역할 수행.  and는 short-컷트? 서킷? 연산.
	switch true {
	case a > 90, b+1 > 90:
		fmt.Println("A grade")
	case a > 80 && b > 80:
		fmt.Println("B grade")
	case a > 70 && b > 70:
		fmt.Println("C grade")
	default:
		fmt.Println("F grade")
	}
}

