package main

import "fmt"

func main() {

	var a = 123
	var b = 456
	var c = 123456789

	fmt.Printf("%5d, %5d\n", a, b)    // %5d : 최소 출력 너비 지정, 출력 시 최소 5칸은 확정된 상태 -> __123, __456
	fmt.Printf("%05d, %05d\n", a, b)  // %05d : 최소 너비보다 짧은 값 0 채우기 -> 00123, 00456
	fmt.Printf("%-5d, %-05d\n", a, b) // %-05d : 최소 너비보다 짧은 값 왼쪽정렬 -> 123__, 456__
	// 근데 왜 %-05d를 했는데도 공백일까욧? 저 자리에 0 넣어버리면 45600이라는 다른 숫자로 출력되어서!!!!!!

	fmt.Printf("%5d, %5d\n", c, c) // 최소 너비지정을 5로 했는데 9짜리 너비가 들어옴
	fmt.Printf("%05d, %05d\n", c, c)
	fmt.Printf("%-5d, %-05d\n", c, c)
	// 최소 너비보다 긴 값을 출력하면 지정된 너비가 다 무시되어버림

	fmt.Printf("%13d, %13d\n", c, c)
	fmt.Printf("%013d, %013d\n", c, c)
	fmt.Printf("%-13d, %-013d\n", c, c)

	//   %       - or n/a             0 or n/a                    (num)                  d
	// 퍼센트  왼쪽 정렬 여부  공란을 0으로 채울꺼냐 말꺼냐  최소 너비를 (num)만큼 지정  10진수 정수값
}
