package main

import "fmt" // Golang 기본 패키지

// 표준 출력용 함수
// Print()
// Println()
// Printf()

func main() {
	var a int = 10
	var b int = 20
	var f float64 = 32799438743.8297

	fmt.Print("a : ", a, "b : ", b)              // 출력 후 개행 안함
	fmt.Println("a : ", a, "b : ", b, "f : ", f) // 출력할 값 사이마다 공란을 넣어주고, 다 끝나면 개행 함

	fmt.Printf("a : %d b : %d f : %f \n", a, b, f) // 출력 서식을 먼저 맞춰놓고 변수값 넣어버림, \n 또한 개행을 나타냄

}
