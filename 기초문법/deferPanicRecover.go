package main

import (
	"fmt"
)

func main() {
	var object interface{} = 396
	isStringCheck(object)
	println("Done")
}

func isStringCheck(target interface{}) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("복구한다.") // r에는 panic의 문구가 담겨있다.
		}
	}() // 익명 함수일 때는 ()를 마지막에 추가해야된다.

	_, ok := target.(string)
	if !ok {
		panic("target is not a string")
	}
	defer fmt.Println("함수 종료.")
}

/* 🖥️ 출력
main.isStringCheck({0x108e740?, 0x10c0da0?})
        /Users/leehyunje/go/test.go:16 +0x8e
main.main()
        /Users/leehyunje/go/test.go:9 +0x25
exit status 2

panic()사용시 즉시 멈춤.
이를 살려내는게 recover이다. recover을 defer에 감싸서 실행하면 panic 상태가 사라지고 isStringCheck 밑 구문이 실행된다.
--------
복구한다.
Done
---------
*/
