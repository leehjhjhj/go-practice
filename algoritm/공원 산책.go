import (
    "strings"
    "strconv"
)


var (
    dx = []int{0, 1, 0, -1}
    dy = []int{-1, 0, 1, 0}
)

var drMap = map[string]int{
    "N": 0,
    "E": 1,
    "S": 2,
    "W": 3,
}

func parseRoute(route string) (string, int) {
    parts := strings.Fields(route)
    direction := parts[0]
    distance, _ := strconv.Atoi(parts[1])
    return direction, distance
}

func checkCanGo(park []string, nx int, ny int) bool {
    if ny < 0 || ny >= len(park) || nx < 0 || nx >= len(park[0]) {
        return false
    }
    if park[ny][nx] == 'X' {
        return false
    }
    return true
}

func findStart(park []string) (int, int) {
    for y := 0; y < len(park); y++ {
        for x := 0; x < len(park[0]); x++ {
            if park[y][x] == 'S' {
                return x, y
            }
        }
    }
    return -1, -1
}

func solution(park []string, routes []string) []int {
    nowX, nowY := findStart(park)
    if nowX == -1 && nowY == -1 {
        return []int{-1, -1}
    }
    for _, route := range(routes) {
        direction, distance := parseRoute(route)
        bufferX, bufferY := nowX, nowY
        canChange := true
        for i := 0; i < distance; i++ {
            nx := bufferX + dx[drMap[direction]]
            ny := bufferY + dy[drMap[direction]]
            if checkCanGo(park, nx, ny) {
                bufferX, bufferY = nx, ny
            } else {
                canChange = false
                break
            }
        }
        if canChange{
            nowX, nowY = bufferX, bufferY
        }
    }
    result := []int{nowY, nowX}
    return result
}