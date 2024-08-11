func solution(keymap []string, targets []string) []int {
    mapping := make(map[rune]int)
    for _, orders := range keymap {
        for idx, order := range orders {
            count, exists := mapping[order]
            if !exists {
                mapping[order] = idx + 1
            } else {
                if count > idx + 1 {
                    mapping[order] = idx + 1
                }
            } 
        }
    }
    var result []int
    for _, target := range targets {
        sum := 0
        possible := true
        for _, targetOrder := range target{
            if count, exists := mapping[targetOrder]; exists {
                sum += count
            } else {
                possible = false
                break
            }
        }
        if !possible {
            sum = -1
        }
        result = append(result, sum)
    }
    return result
}