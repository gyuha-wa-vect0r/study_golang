package main

import "fmt"

func main() {
	var a int16 = 3456
	// int16 : -32768 ~ 32767
	var c int8 = int8(a) // int16 -> int8
	// int8 : -128 ~ 127

	fmt.Println(a) // 3456
	fmt.Println(c) // -128

	// a int16 = 00001101 10000000 = 3456
	// a int8  = -------- 10000000 = -128
	// int16 -> int8 변환 시 상위 1바이트가 사라지게 된다...

}
