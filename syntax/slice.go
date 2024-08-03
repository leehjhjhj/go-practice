package main

import (
	"fmt"
)

func main() {
	monster := make([]string, 5, 10)
	fmt.Println(monster, len(monster), cap(monster))

	var monster1 []string
	monster1 = append(monster1, "주황버섯", "초록버섯", "좀비버섯") // append(슬라이스, 넣을 값)
	fmt.Println(monster1)
	monster1 = monster1[2:3] // 파이썬과 동일
	fmt.Println(monster1)
	monster1 = append(monster1, monster...)// 다른 슬라이스를 넣을 때는 ... 을 붙인다.
	fmt.Println(monster1)

	monster2 := make([]string, len(monster1)) // make로 길이를 정해주어야 한다.
	copy(monster2, monster1)
	fmt.Println(monster2)

	monster2[0] = "좀비머쉬맘"
	fmt.Println(monster1)
}

/* 🖥️ 출력
[    ] 5 10 -> monster
[주황버섯 초록버섯 좀비버섯] -> monster1
[좀비버섯] -> 슬라이싱 후 monster1
[좀비버섯     ] -> monster와 합친 후 monster1
[좀비버섯     ] -> monster1의 복사본인 monster2
[좀비버섯     ] -> 깊은 복사인 것 같다. monster2를 바꾸어도 monster1은 바뀌지 않는다. append 또한 깊은 복사이다. 
*/
