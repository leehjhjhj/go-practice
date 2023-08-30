import (
	"container/list"
	"math"
)

func solution(N int, roads [][]int, K int) int {
	grid := make([][]*pair, N+1)
	for i := range grid {
		grid[i] = make([]*pair, 0)
	}

	visited := make([]float64, N+1)
	for i := range visited {
		visited[i] = math.Inf(1)
	}

	for _, road := range roads {
		a, b, c := road[0], road[1], road[2]
        grid[a] = append(grid[a], &pair{b,c})
        grid[b] = append(grid[b], &pair{a,c})
    }

	q := list.New()
	q.PushBack(&pair{1, 0})
	visited[1] = 0
	bfs(q, grid, visited)
	answer := 0
	for _, target := range visited {
		if target <= float64(K) {
			answer += 1
		}
	}
    return answer
}

type pair struct {
	first int
	second int
}


func bfs(q *list.List, grid [][]*pair, visited []float64) {
	for q.Len() > 0 {
		target := q.Front()
        v,vDistance:= target.Value.(*pair).first, target.Value.(*pair).second
		q.Remove(target)

		if float64(vDistance) > visited[v] {
			continue
		}
		for _, nPair := range grid[v] {
			n, nDistance := nPair.first, nPair.second
			if float64(vDistance + nDistance) < visited[n] {
				visited[n] = float64(vDistance + nDistance)
				q.PushBack(&pair{n, int(visited[n])})
			}
		}
	}
}
