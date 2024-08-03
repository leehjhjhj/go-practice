package main

import (
	"fmt"
)

func main() {
	insert_poison("주황버섯", "뿔버섯") //mushroom의 주소를 표시. 주소를 전달한다.
	var 전리품 = []string{"주황버섯의 갓", "뿔버섯의 뿔"}
	drop_item(전리품)
}

func insert_poison(mushroom ...string) { //*로 파라미터가 포인터임을 표시. 메모리 영역의 주소를 갖게 된다.
	for _, name := range mushroom {
		result := "독이 든 " + name
		fmt.Println(result)
	}
	fmt.Println(mushroom) // 슬라이스로 전달
}

func drop_item(items []string) {
	for i := 0; i < len(items); i++ {
		fmt.Printf("\"%v\" 아이템을 드롭하셨습니다.\n", items[i])
	}
}
/* 🖥️ 출력
독이 든 주황버섯
독이 든 뿔버섯
[주황버섯 뿔버섯]
"주황버섯의 갓" 아이템을 드롭하셨습니다.
"뿔버섯의 뿔" 아이템을 드롭하셨습니다.
*/
