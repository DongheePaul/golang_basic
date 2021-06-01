`package main


import (
"fmt"
// 	"bufio"
// 	"os"
)`

func main() {
	// var ch1 byte = 65
	// var ch2 byte = 0101
	// var ch3 byte = 0x41
	// var ch4 byte = 'A'

	//문자열로 출력
	//fmt.Printf("%c %c %c %c", ch1, ch2, ch3, ch4)
	//첫번째는 문자로 출력
	//fmt.Printf("%d %c %c %c", ch1, ch2, ch3, ch4)

	// var uni1 rune = 44032
	// var uni2 rune = 0126000
	// var uni3 rune = 0xAC00
	// var uni4 rune = '가'
	// fmt.Printf("%c %c %c %c\n", uni1, uni2, uni3, uni4)

	// rawLiteral := `hihi \n
	// skdlsk \n
	// 	dlksdlk`
	// interLiteral := "dkdldl \n sldlsk \n"
	// fmt.Printf(rawLiteral)
	// fmt.Printf(interLiteral)

	// fmt.Println(unsafe.Sizeof(rawLiteral + "\n"))

	// fmt.Println(unsafe.Sizeof(interLiteral + "\n"))

	//윈도우 로더가 메모리에 값을 올림.
	//var a int
	//var b int
	//엔드기호 쓰면 메모리 주소.  n, err => 함수 리턴값.
	//n, err := fmt.Scan(&a, &b)

	//if err != nil {
	//	fmt.Println(n, err)
	//}else {
	//	fmt.Println(n, a, b)
	//}

	//가변인수
	
	//os 단에서 입력을 받고, 이를 app 단에서 핸들링 하겠다.
	// stdio := bufio.NewReader(os.Stdin)
	// var a int
	// var b int

	// n, err := fmt.Scanln(&a, &b)
	// if err != nil {
	// 	fmt.Println(err)
	// 	fmt.Println("다시 입력해주세요")
	// 	stdio.ReadString('\n')
	// }else{
	// 	fmt.Println(n, a, b)

	// }

	// n, err = fmt.Scanln(&a, &b)
	// if err != nil {
	// 	fmt.Println(err)
	// 	fmt.Println("다시 입력해주세요")
		
	// 	stdio.ReadString('\n')
	// }else{
	// 	fmt.Println(n, a, b)

	// }
	var i int = 10
	var j int = 20
	fmt.Println(i, j)
	swap(&i, &j)
	fmt.Println(i, j)

}

func swap (a * int, b * int){
	// int temp = *a
	// *a = *b
	// *b = temp    씨에서의 이 코드는 고에서 밑 한줄로 가능.
	*b, *a = *a, *b
}

