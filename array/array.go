package main

import (
"fmt"
"sort"
)

func main() {

	//var a = [5]float64{1.0, 2.0, 3.0, 4.0, 5.0}
	//fmt.Printf("%.2f, %.2f, %2f, %2f, %2f", a[0], a[1],a[2],a[3],a[4] )
	// for i:=0; i<5; i++ {
	// 	fmt.Printf("%.2f, ",  a[i])
	// }

	// var day = [7]string{"일", "월", "화", "수", "목", "금", "토"}

	// var list = [7]string{}
	// for i:=0; i<7; i++{
	// 	fmt.Printf(day[i], "을 입력하세요")
	// 	fmt.Scanf("%s", &list[i])
	// }

	// fmt.Println(list)
	// var i int = 10
	// //인자에 값이 잘못 들어오면 버퍼에서 들어냄?
	// fmt.Scanf("%d", &i)
	// //fmt.Scanf("%d", &i)

	// fmt.Println(i)
	// //i의 주소를 출력
	// fmt.Printf("%v\n", &i)
	// fmt.Println(&i)

	// var names []string = []string{"이순신", "홍길동", "강감찬"}
	// //원래 (인덱스, 밸류) 값이 담기는데, _ 를 넣으면 해당 위치에서 받는 값을 무시.
	// for _,val := range names{
	// 	fmt.Println(val)   // 이순신, 홍길동, 강감찬
	// }

	// i := 0

	// for {
	// L1:
	// 	for {
	// 		for{
	// 			if i == 0{
	// 				break L1
	// 			}
	// 		}
	// 	}
	// fmt.Println("label okay")
	// }
	// fmt.Println("ok")
	
	//var a []int //슬라이스
	//var a [3]int  //배열
	//fmt.Printf("%v", &a[0]) // => 첫번쨰 원소의 주소(16진수)
	// a = []int{1, 2, 3}
	// a[1] = 10
	// fmt.Println(a)
	// s := make([]int, 5, 10)
	// fmt.Println(len(s), cap(s)) //길이, 용량

	s := []int{10, 11, 21, 31, 41, 51}
	sort.Sort(sort.IntSlice(s)) //오름차순
	fmt.Printf("%d\n", s)

	sort.Sort(sort.Reverse(sort.IntSlice(s))) //내림차순
	
	fmt.Printf("%d\n", s)

	//배열은 순회할 떄 완벽하게 같아야 하지만(원소의 타입과 어레이의 길이), 슬라이스는 그렇지 않아도 됨.

}