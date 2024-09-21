package main

import "fmt"

func main() {
	var a int = 3
	// 정수형 a 선언 후 3으로 초기화
	var b int
	// 정수형 b 선언 후 초기화 안함
	// -> 타입 별 기본값(0 or 0.0 or false or "(빈 문자열)" or nil(정의되지 않은 메모리 주소), 값 타입에 따라 결정되겠죠?)으로 대체함
	var c = 4
	// 타입을 생략하고 변수 선언함 -> 우변에 있는 값에 따라 변수 타입이 결정됨, 이 경우에는 정수형이 되겠죠?
	d := 5
	// 선언 대입문 :=, var 키워드와 변수 타입을 생략하였다 -> d는 int 타입으로 자동 지정되고 5로 초기화

	fmt.Println(a, b, c, d)
}
