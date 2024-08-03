package main

import (
	"fmt"
)

func main() {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += 1
	}
	fmt.Println(sum)

	j := 0
	for j < 100 {
		j += 1
		if j == 100 {
			fmt.Printf("j가 %v이 되었습니다.", j)
			break
		}
	}

	mushrooms := []string{"파랑버섯", "주황버섯", "초록버섯", "머쉬맘"}
	for _, name := range mushrooms {
		fmt.Println(name)
		if name == "머쉬맘" {
			goto END // END로 이동한다.
		}
	}
	END:
		fmt.Println("머쉬맘은 버섯의 어머니이다.")
}
/* 🖥️ 출력
100
j가 100이 되었습니다.
파랑버섯
주황버섯
초록버섯
머쉬맘
머쉬맘은 버섯의 어머니이다.
*/
