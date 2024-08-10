package main

import (
    "strconv"
    "sort"
    "strings"
)

func checkAllZero(buffer []int) bool {
    for _, target := range buffer {
        if target != 0 {
            return false
        }
    }
    return true
}

func solution(X string, Y string) string {
    xCount := make(map[rune]int)
    yCount := make(map[rune]int)

    for _, target := range X {
        xCount[target]++
    }
    for _, target := range Y {
        yCount[target]++
    }

    var buffer []int
    for key, xVal := range xCount {
        if yVal, exists := yCount[key]; exists {
            count := min(xVal, yVal)
            num, _ := strconv.Atoi(string(key))
            for i := 0; i < count; i++ {
                buffer = append(buffer, num)
            }
        }
    }

    if len(buffer) == 0 {
        return "-1"
    }

    if checkAllZero(buffer) {
        return "0"
    }

    sort.Sort(sort.Reverse(buffer))
    var result []string
    for _, num := range buffer {
        result = append(result, strconv.Itoa(num))
    }
    return strings.Join(result, "")
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}