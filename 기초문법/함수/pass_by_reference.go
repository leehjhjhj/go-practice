package main

import (
	"fmt"
)

func main() {
	mushroom := "초록 버섯"
	insert_poison(&mushroom) //mushroom의 주소를 표시. 주소를 전달한다.
	fmt.Println(mushroom)
}
func insert_poison(mushroom *string) { //*로 파라미터가 포인터임을 표시. 메모리 영역의 주소를 갖게 된다.
	*mushroom = "독이 든 " + *mushroom
}
/* 🖥️ 출력
독이 든 초록 버섯
*/
