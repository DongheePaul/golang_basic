package main

import "fmt"

type Data struct {
	value int
	data  [100]int
}

//대문자로 함수명 작성하면 외부에서 접근 가능.
func ChangeData(arg *Data) {
	//원래라면 아래처럼 할당해야함. 고랭에서 스트럭을 쓸 때는 *를 안 붙여도 됨.
	//(*arg).data[99] = 5
	//(*arg).value = 3
	arg.data[99] = 5
	arg.value = 3
}

func main() {
	//var numPtr *int
	//fmt.Println(numPtr)

	// var data Data
	// ChangeData(&data)
	// fmt.Printf("value = %d\n", data.value)
	// fmt.Printf("data[99] = %d\n", data.data[99])

	var a rune = '한'
	//var b rune = 'a'
	b := int64(128)
	fmt.Println(a, b)
	fmt.Printf("%c, %c\n", a, b)
	fmt.Printf("%T, %T\n", a, b)

	// str1 := "안녕하세요"
	// str2 := "abcde"
	// fmt.Printf("len(str1) = %d, len (str2) = %d\n", len(str1), len(str2))

	// str3 := "hello 월드"
	// for n, val := range str3 {
	// 	fmt.Printf("index : %d, value : %c\n", n, val)
	// }

	//str4 := "AAAAAA"
	//str5 := "a"
	// str6 := "BBB"
	// str7 := "bbb"
	// str8 := "CCCCCCCCCCCCCCC"
	// str9 := "ccccccccccccccc"
	//아스키코드에서 소문자가 대문자보다 작은 값을 가진다. 그리고 첫글자로만 크기를 비교한다.
	//fmt.Printf("%s > %s : %v\n", str4, str5, str4 > str5)
}
