package main

import (
	"fmt"
)

func main() {
	field1Spawn := addMonster()
	fmt.Printf("%v 마리가 필드에 있습니다.\n", field1Spawn())
	fmt.Printf("%v 마리가 필드에 있습니다.\n", field1Spawn())
	
	field2Spawn := addMonster()
	fmt.Printf("%v 마리가 필드에 있습니다.\n", field2Spawn()) // 클로저는 다시 초기화가 된다.

}

func addMonster() func() int {
	monster := 0
	return func() int {
		monster += 1
		return monster
	}

}

/* 🖥️ 출력
1 마리가 필드에 있습니다.
2 마리가 필드에 있습니다.
1 마리가 필드에 있습니다.

Go 언어에서 함수는 Closure로서 사용될 수도 있다.
Closure는 함수 바깥에 있는 변수를 참조하는 함수값(function value)를 일컫는데,
이때의 함수는 바깥의 변수를 마치 함수 안으로 끌어들인 듯이 그 변수를 읽거나 쓸 수 있게 된다.
*/
