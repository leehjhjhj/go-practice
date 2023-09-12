package main

import (
	"fmt"

)

func main() {
	notice := make(chan string, 2)

	notice <- "안녕하세요? 운영자입니다."
	notice <- "이벤트에 당첨되셨습니다."

	close(notice)

	fmt.Println(<-notice)
	fmt.Println(<-notice)
	
	if _, success := <- notice; !success {
		fmt.Println("공지사항 없음")
	}
}


/* 🖥️ 출력
안녕하세요? 운영자입니다.
이벤트에 당첨되셨습니다.
공지사항 없음

채널을 close()로 닫으면 더이상 송신은 할 수 없지만, 수신은 가능하다.
<-ch 와 같이 채널 수신은 두개의 리턴 값ㅇ를 갖는데, 첫번재는 채널 메시지이고 두번째는 true, false로 메시지 수신이 되었는가를 말해준다.
*/
