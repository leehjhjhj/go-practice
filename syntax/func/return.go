package main

import (
	"fmt"
	"strings"
)

func main() {
	전리품 := [3]string{"일비표창", "레드 드레이크의 가죽", "250메소"}
	count, isJavelin := dropItem(전리품)
	fmt.Printf("%v개의 아이템과 표창의 유무는 \"%v\" 입니다.", count, isJavelin)
}

func dropItem(items [3]string) (sum int, isJavelin string){ // 함수 리턴값으로 ZeroValue로 먼저 선언되어있음, go에서 배열과 슬라이스는 엄연히 다른 자료형이다.
	var javelinExists bool = false

	for _, name := range items {
		sum += 1
		if strings.Contains(name, "표창") { //in 대신 strings.Contains 사용
			javelinExists = true
		}
	}
	if javelinExists {
		isJavelin = "있다"
	} else {
		isJavelin = "없다"
	}
	return
}
/* 🖥️ 출력
3개의 아이템과 표창의 유무는 "있다" 입니다.
*/
