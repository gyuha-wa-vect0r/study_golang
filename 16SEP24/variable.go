package main

import "fmt" // 'fmt'아니고 "fmt" 입니다

func main() {
	var a int = 10 // 생각보다 변수 선언법이 꽤나 직관적이네요
	//  변수선언할래요 변수명은이렇구요 타입은이래요 초기값은얼마가들어가요 순
	var msg string = "Hello Variable"

	a = 20
	msg = "Good Morning"

	fmt.Println(msg, a) // msg와 a 출력
}
