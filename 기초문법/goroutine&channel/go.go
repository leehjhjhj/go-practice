package main

import (
	"fmt"
	"time"
)

func makeMonster(name string) {
	for i := 0; i < 5; i++ {
		fmt.Printf("%v이 스폰되었습니다.\n", name)
	}
}

func main() {
	makeMonster("주황버섯")
	go makeMonster("뿔버섯")
	go makeMonster("뿔버섯")
	go makeMonster("뿔버섯")

	time.Sleep(time.Second * 1)
}


/* 🖥️ 출력
주황버섯이 스폰되었습니다.
주황버섯이 스폰되었습니다.
주황버섯이 스폰되었습니다.
주황버섯이 스폰되었습니다.
주황버섯이 스폰되었습니다.
뿔버섯이 스폰되었습니다.
뿔버섯이 스폰되었습니다.
뿔버섯이 스폰되었습니다.
뿔버섯이 스폰되었습니다.
뿔버섯이 스폰되었습니다.
뿔버섯이 스폰되었습니다.
뿔버섯이 스폰되었습니다.
뿔버섯이 스폰되었습니다.
뿔버섯이 스폰되었습니다.
뿔버섯이 스폰되었습니다.
뿔버섯이 스폰되었습니다.
뿔버섯이 스폰되었습니다.
뿔버섯이 스폰되었습니다.
뿔버섯이 스폰되었습니다.
뿔버섯이 스폰되었습니다.
*/
