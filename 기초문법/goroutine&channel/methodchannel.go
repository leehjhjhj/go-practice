package main

import (
	"fmt"

)

func main() {
	notice := make(chan string, 1)
	sendNotice(notice)
	receiveNotice(notice)
}

func sendNotice (ch chan <- string) {
	ch <- "잠시후 서버가 재부팅됩니다."
}

func receiveNotice (ch <- chan string) {
	noticeDetail := <- ch
	fmt.Println(noticeDetail)
}

/* 🖥️ 출력
잠시후 서버가 재부팅됩니다.

채널을 함수의 파라미터로 전달할 때, 일반적으로 송수신을 모두 하는 채널을 전달하지만, 저렇게 파라미터로 지정하면
송신만 할 것인지, 수신만 할 것인지 정할 수 있다.
*/
