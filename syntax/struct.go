package main

import (
	"fmt"
)

func main() {
	monster1 := monster{"파란버섯", 10, 5}
	monster2 := monster{name: "주황버섯", hp: 15}
	fmt.Println(monster1.name)
	fmt.Println(monster2.name, monster2.mp)
	
	monster3 := monster2
	monster3.name = "히히버섯"
	fmt.Println(monster2.name, monster3.name)

	var monster4 *monster = &monster2 // monster4 := &monster2와 같다.
	monster4.name = "이름이 바뀐 버섯"
	fmt.Println(monster2.name, monster4.name)
	
}

type monster struct {
	name string
	hp int
	mp int
}

/* 🖥️ 출력
파란버섯
주황버섯 0 -> mp의 값을 주지 않으면 zero value나 pointer일 때 nil이 들어간다.
주황버섯 히히버섯 -> value로 전달하기 때문에 바뀌지 않는다.
이름이 바뀐 버섯 이름이 바뀐 버섯 -> 포인터로 메모리 주소가 넘어갔기 때문에 monster2의 이름도 변경된다.

- go에서의 구조체는 필드 데이터만 있고 메서드는 없다.
- mutable이다.
- 함수의 파라미터로 넘기면 pass by value로 객체를 복사해서 전달한다.
- pass by reference로 struct를 전달하고 싶으면 포인터를 전달해야한다.
- go의 모든 값들은 value로 넘어간다.
*/
