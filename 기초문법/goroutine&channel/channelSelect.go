package main
 
import "time"
 
func main() {
    done1 := make(chan bool)
    done2 := make(chan bool)
	done3 := make(chan bool)

    go run1(done1)
    go run2(done2)
	go run3(done3)
 
EXIT:
    for {
        select {
        case <-done1:
            println("run1 완료")

		case <-done2:
            println("run2 완료")
 
        case <-done3:
            println("run3 완료")
            break EXIT
        }
    }
}
 
func run1(done chan bool) {
    time.Sleep(1 * time.Second)
    done <- true
}
 
func run2(done chan bool) {
    time.Sleep(2 * time.Second)
    done <- true
}

func run3(done chan bool) {
    time.Sleep(3 * time.Second)
    done <- true
}


/* 🖥️ 출력
run1 완료
run2 완료
run3 완료

select문을 이용해서 복수의 채널들을 기다리면서 채널에 수신된 데이터에 따라서 채널을 실행시킬 수 있다.
*/
