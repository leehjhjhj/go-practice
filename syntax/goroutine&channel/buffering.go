package main

import (
	"fmt"

)

func main() {
	ch := make(chan string, 1)

	// go func() {
	// 	ch <- "안녕하세요"
	// }()
	
	ch <- "하이요"
	// ch <- "하이요"
	fmt.Println(<-ch)
}

/* 🖥️ 출력
버퍼링을 주면 수신자가 없어도 데이터를 보낼 수 있다.
근데 버퍼링 이상으로 데이터를 전송하면 오류가 발생한다. 위의 주석을 풀면
goroutine 1 [chan send]:
main.main()
        /Users/leehyunje/go/test.go:16 +0x47
exit status 2
라는 오류가 발생한다.
*/
