
func solution(name []string, yearning []int, photo [][]string) []int {
    memory := make(map[string]int)
    var result []int
    var minLen int = len(name)
    if len(yearning) < minLen {
        minLen = len(yearning)
    }
    
    for i := 0; i < minLen; i++ {
        memory[name[i]] = yearning[i]
    }
    for _, friends := range photo {
        var point int = 0
        for _, friend := range friends {
            value, exist := memory[friend]
            if exist {
                point += value
            }
        }
        result = append(result, point)
    }
    
    return result