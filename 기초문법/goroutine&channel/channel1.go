package main

import (
	"fmt"

)

func main() {
	ch := make(chan string)

	go func() {
		ch <- "안녕하세요"
	}()

	var s string
	s = <- ch
	fmt.Println(s)
}

/* 🖥️ 출력
안녕하세요

Go 채널은 그 채널을 통해서 데이터를 주고 받는다.
make를 통해 미리 생성하고 채널 연산자로 데이터를 보내고 받는다.
또한 계속 수신하는 성격을 가지고 있기 때문에 main()에 done := make(chan bool)를 수신하고 있으면 종료되지 않는다.
*/
