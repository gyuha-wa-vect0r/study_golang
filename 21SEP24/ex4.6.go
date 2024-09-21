// range of var

package main

import "fmt"

var g int = 10 // 패키지의 global variable

func main() {
	var m int = 20 // main function 내 local variable

	{
		var s int = 50 // 중괄호 내 local variable

		fmt.Println(m, s, g)
	}

	m = s + 20 // s는 함수 내 중괄호 안에서만 동작하기 때문에 연산이 안됩니다...

}

// 컴파일 실행 결과
// ┌──(GYUHA LEE㉿KONGSUNI_LAPTOP)-[D:/공부/study_golang/21SEP24]-[main]                                                                                                                                                   119ms  
// └─$ go build ex4.6.go
// # command-line-arguments
// .\ex4.6.go:18:6: undefined: s
