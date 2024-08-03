package main

import (
	"fmt"
)

func main() {
	var mushmom map[string]string // 이것만 있으면 nilMap이 만들어지고, 이게 없으면 밑의 make도 안됨.
	mushmom = make(map[string]string) // map 초기화
	mushmom["정보"] = "버섯들의 어머니이다."
	mushmom["서식지"] = "남의 집"
	fmt.Println(mushmom)

	//리터럴 선언
	리본돼지 := map[string]string{
		"정보": "리본을 묶은 돼지이다.",
		"서식지": "돼지의 해안가", // 마지막에도 ,를 붙이자.
		"더미데이터": "메롱",
	}
	리본돼지서식지, exists := 리본돼지["서식지"] // exists는 없어도 된다.
	fmt.Println(리본돼지서식지, exists)
	
	리본돼지아이템, exists := 리본돼지["아이템"] // true, false로 값이 들어온다.
	fmt.Println(리본돼지아이템, exists)

	delete(리본돼지, "더미데이터")

	for key, val := range 리본돼지 {
		fmt.Println(key,":",val)
	}
}

/* 🖥️ 출력
map[서식지:남의 집 정보:버섯들의 어머니이다.]
돼지의 해안가 true
 false -> nil 값이 들어온다.
정보 : 리본을 묶은 돼지이다.
서식지 : 돼지의 해안가
*/
