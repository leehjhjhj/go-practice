package main

import (
	"fmt"
)

func main() {
	var object interface{} = 30

	test1 := object
	test2 := object.(int)
	// test3 := object.(type) -> Swtich 구문에서 사용 가능
	//test4 := object.(string) -> 에러 발생 panic: interface conversion: interface {} is int, not string
	fmt.Println("test1:", test1)
	fmt.Println("test2:", test2)

	value, ok := object.(string)
	fmt.Println(value, ok) // 여기서 형변환이 가능한지 ok로 알 수 있으며, 런타임 에러를 방지한다.
}

/* 🖥️ 출력
test1: 30
test2: 30

- 빈 interface는 자바의 object와 같은 것. 어떤 타입도 담을 수 있다.
- Type Asssertion은 interface.(type)으로 사용하고,  실제 인터페이스의 type이 아닐시 오류를 낸다.
- .(type)은 실제 타입을 검사하는데 사용되며, 이는 오직 switch 구문 내에서만 사용될 수 있다.
*/
