package main

import "fmt"

type person struct {
	name string
	age int
}

type User struct {
	//첫글자를 대문자로 하면 => 현재 패키지를 벗어나 외부에 노출 하겠다. = public
	Name string
	//name string // 외부에 노출시키지 않겠다. 동일 파일에선 접근 가능.
	age int
	Id string
}

type VIPUser struct {
	UserInfo User
	VIPLevel int
	Price int
}
func main() {
	// //var p1 person
	// var p1 person
	// //var p2 person
	// p2 := person{}
	// //new 연산자 쓰면 힙 공간에 동적으로 할당하고 주소를 리턴. 그래서 사실 p3는 var p3 *person 즉 포인터다.
	// p3 := new(person)
	// //포인터이던 아니던 접근 방법이 동일.
	// fmt.Println(p1.name, p2.name, p3.name)
	// fmt.Printf("%p, %p, %p", &p1, &p2, p3)
	
	//대괄호는 자료구조.  중괄호는 구조체.
	normalUser := User { "이순신", 100, "Lee"}
	vipUser := VIPUser { User { "이순신", 100, "Lee"}, 5, 10000}
	fmt.Println(normalUser, vipUser)
	fmt.Printf("유저 이름 : %s, 나이 : %d, 아이디 : %s \n", normalUser.Name, normalUser.age, normalUser.Id)
	fmt.Printf("vip user name: %s, ", vipUser.UserInfo.Name)

	vipUser2 := vipUser
	fmt.Println(vipUser2)

}