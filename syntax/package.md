# main 패키지
- package에 main이 있으면 main()함수가 시작점이 된다.
- 패키지를 공유 라이브러리로 만들고 싶으면 main 패키지나 main 함수를 사용하면 안된다.

# 패키지 scope
- 함수, 구조체, 인터페이스, 메서드의 이름을 소문자로 하면 non-pulbic
- 대문자로 하면 import 가능한 public으로 사용 가능.

# init 함수와 alias
- 패키지를 실행시 처음으로 호출되는 init() 작성 가능
- 패키지 로드되면서 자동으로 호출됨.
```go
package testlib
 
var pop map[string]string
 
func init() {   // 패키지 로드시 map 초기화
    pop = make(map[string]string)
}
```
- 패키지 import 할 때, 패키지의 init()만 호출하려는 경우
```go
package main
import _ "other/xlib"
```
- 또한 alias로 같은 이름의 패키지를 별칭으로 사용가능
```go
import (
    mongo "other/mongo/db"
    mysql "other/mysql/db"
)
func main() {
    mondb := mongo.Get()
    mydb := mysql.Get()
    //...
}
```
# 사용자 정의 패키지 생성
- 패키지를 만들어 재사용 가능한 컴포넌트 생성
- 보통 폴더하나를 만들고 폴더 안에 .go 생성
- 해당 폴더 안에 .go 파일들은 동일한 패키지명, 패키지 명은 폴더 이름과 같게 함

# 캐싱
- 아직 go로 본격적인 프로그래밍을 하지 않았기 때문에 잘 이해하지 못한다. 후에 필요하면 참고하자.
http://golang.site/go/article/15-Go-패키지
