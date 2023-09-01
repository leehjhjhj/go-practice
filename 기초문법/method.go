package main

import (
	"fmt"
)


func main() {
	fieldMob1 := Monster{"크림슨발록", 5000, 3000}
	fieldMob1.hitFake()
	fmt.Printf("가짜로 때린 후: %v\n", fieldMob1.hp)
	fieldMob1.hitReal()
	fmt.Printf("진짜로 때린 후: %v\n", fieldMob1.hp)
}

type Monster struct {
	name string
	hp, mp int
}

func (monster Monster) hitFake() int {
	monster.hp =- 5
	return monster.hp
}

func (monster *Monster) hitReal() int {
	monster.hp -= 5
	return monster.hp
}

/* 🖥️ 출력
가짜로 때린 후: 5000
진짜로 때린 후: 4995

메소드 선언 방법: func 옆에 "그 함수가 어떤 struct를 위한 메서드인지"를 표시
값을 바꾸게 하려면 포인터를 넘긴다고 알려줘야함.
*/
