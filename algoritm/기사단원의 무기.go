import (
    "math"
)

func makeDivisor(number int) []int {
    var divisors []int
    for i := 1; i <= number; i++ {
        var count int
        sqrt := int(math.Sqrt(float64(i)))
        for j := 1; j <= sqrt; j++ {
            if i % j == 0 {
                if j*j == i {
                    count++
                } else {
                    count += 2
                }
            }
        }
        divisors = append(divisors, count)
    }
    return divisors
}

func solution(number int, limit int, power int) int {
    var result int
    divisors := makeDivisor(number)
    for _, target := range divisors {
        if target > limit {
            result += power
        } else {
            result += target
        }
    }
    return result
}