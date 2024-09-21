// diff_type.go 에서 형 변환 코드를 넣은 후 정상 동작하게끔!

package main

import "fmt"

func main() {

	a := 3              // int형
	var b float64 = 3.5 // float64형

	var c int = int(b)  // float64형 값 -> int 값, 3.5에서 소수점 이하 값 날라가서 3
	d := float64(a * c) // int 연산 값 -> float64 값

	var e int64 = 7   // int64형
	f := int64(d) * e // float64형 값 -> int64 값

	var g int = int(b * 3) // float64 -> int, 3.5 * 3 (연산 후 형 변환)
	var h int = int(b) * 3 // float64 -> int, 3 * 3 (형 변한 후 연산)

	fmt.Println(g, h, f)
	// go 신기한건 변수 선언되고 한번이라도 사용 안되면 컴파일이 안되는거같음...?
}
