package main

import (
	"fmt"
	"strconv"
)
/*
변수는 :=, 혹은 var 변수명 자료형으로 선언한다.
str <-> int는 strconv를 통해서 이루어진다.
*/
func main() {
  str := "300"
	to_int, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("변환 실패: ", err)
	}
	fmt.Println(to_int)

  original := strconv.Itoa(to_int)
  fmt.Println(original)
}
