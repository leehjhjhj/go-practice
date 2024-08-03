package main

import (
	"fmt"
)


func main() {
    set := make(map[string]struct{})
    set["item1"] = struct{}{} // 구조체로 하는 이유는 빈 구조체는 0 바이트의 메모리이기 때문이다.
    set["item2"] = struct{}{}
	fmt.Println(set)
    
	// "item1"이 있는지 확인
    i, ok := set["item1"]
	fmt.Println(i, ok) // 첫 번째로 구조체를 반환하고 두 번째로 bool을 반환한다.

    if ok {
        fmt.Println("item1 exists")
    } else {
        fmt.Println("item1 does not exist")
    }

    // "item3"이 있는지 확인
    _, ok = set["item3"]
    if ok {
        fmt.Println("item3 exists")
    } else {
        fmt.Println("item3 does not exist")
    }
}
/* 🖥️ 출력
map[item1:{} item2:{}]
{} true
item1 exists
item3 does not exist
*/
