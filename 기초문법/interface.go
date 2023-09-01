package main

import (
	"fmt"
)

func main() {
	monster1 := 버섯{"파란버섯", 25, 35}
	monster2 := 돼지{"아이언호그", 10, 90}
	enhanceActiveMonster(&monster1, &monster2) //인터페이스 사용

}

type EnhanceMonster interface { //몬스터 강화를 위한 interface
	enhancePassive() int
	enhanceActive() int
}

type 버섯 struct {
	name string
	명중률 int
	몸통박치기 int
}

type 돼지 struct {
	name string
	배고픔 int
	빠르게달리기 int
}
/* EnhanceMonster 구현체 */
func (mushroom *버섯) enhancePassive() int {
	mushroom.명중률 += 30
	return mushroom.명중률
}
func (mushroom *버섯) enhanceActive() int {
	mushroom.몸통박치기 += 50
	return mushroom.몸통박치기
}

func (pig *돼지) enhancePassive() int {
	pig.배고픔 += 10
	return pig.배고픔
}

func (pig *돼지) enhanceActive() int {
	pig.빠르게달리기 += 80
	return pig.빠르게달리기
}

func enhanceActiveMonster(targetMonsters ...EnhanceMonster) { // 이렇게 인터페이스를 상속한(완벽 구현한) 구조체를 모두 넣을 수 있다.
	for _, targetMonster := range targetMonsters {
		switch monster := targetMonster.(type) { //type을 통해서 targetMonster의 실제 타입을 체크하고 그에 따라 적절한 case문을 실행한다.
		case *버섯:
			fmt.Printf("이름: %v, 현재 스킬 능력치: %v\n", monster.name, monster.몸통박치기)
		case *돼지:
			fmt.Printf("이름: %v, 현재 스킬 능력치: %v\n", monster.name, monster.빠르게달리기)
		}
		result := targetMonster.enhanceActive()
		fmt.Println("강화 이후:", result)
	} 
}

/* 🖥️ 출력
이름: 파란버섯, 현재 스킬 능력치: 35
증강 이후: 85
이름: 아이언호그, 현재 스킬 능력치: 90
증강 이후: 170

내가 이해한 인터페이스
- 다른 언어와 다르게 직접적으로 인터페이스를 inplements 하지 않는다.
- 다만 한 struct의 메소드들이 interface에 명시된 메소드들을 다 구현한다면 해당 interface를 inplements했다고 간주한다.
- 따라서 저렇게 인터페이스를 매개 변수로 받는 함수에 interface를 받도록 명시해두면 구현체를 모두 구현한 구조체가 들어올 수 있다.
- 그렇게 받은 구조체들은 함수 안에서 구현체를 사용할 수 있다. 왜냐? 공통으로 구현되어있는 메소드 들이니깐.
- 결론적으로 targetMonster들은 특정한 구조체(*버섯 또는 *돼지)이며 동시에 EnhanceMonster 인터페이스를 만족시키거나 '구현'하는 객체이다.
*/
