/*
https://school.programmers.co.kr/learn/courses/30/lessons/131705?language=go
*/

var results []int
var answer int = 0

func solution(number []int) int {
    make(0, number)
    return answer
}

func make(index int, number []int) {
	if len(results) == 3 {
		if sum(results) == 0 {
			answer += 1
		}
		return
	}
	for i := index; i < len(number); i++ { // index부터 시작하여 중복 제거
		results = append(results, number[i])
		make(i+1, number) // 다음 재귀에서는 현재 index보다 큰 index부터 시작
		results = results[:len(results)-1]
	}
}

func sum(slice []int) int {
    total := 0
    for _, value := range slice {
        total += value
    }
    return total
}
