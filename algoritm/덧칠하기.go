func solution(n int, m int, section []int) int {
    now := 0
    result := 0
    for _, s := range section {
        if s > now {
            result ++
            now = s + m - 1
        }
    }
    return result
}