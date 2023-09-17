import (
    "strconv"
)
func reverse(target string) (result string){
    result = ""
    for i := len(target) - 1; i >= 0; i-- {
        result += string(target[i])
    }
    return
}

func solution(food []int)string {
    answer := ""
    for i := 0; i < len(food); i++ {
        foodForOne := food[i] / 2
        for j := 0; j < foodForOne; j++ {
            answer += strconv.Itoa(i)
        }
    }
    reverseAnswer := reverse(answer)
    answer += "0" + reverseAnswer
    
    return answer
}
