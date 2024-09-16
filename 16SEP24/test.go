package main // package 선언 -> 패키지 선언으로 코드를 시작,,, 패키지는 코드의 가장 큰 범주

import "fmt" // fmt 패키지를 불러움, fmt는 표준입출력 담당 (c에서의 iostream너낌)

func main() { // main 함수 시작이요
	fmt.Println("Hello World!!") // 표준입출력 패키지(fmt)에 있는 출력 함수(Println)로 문자열 출력
}

// 주석 사용법
// <- 이거대로 // 쓰던지
// /* 로 시작해서 */ 로 닫던지
