package skyline

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func mergeSkylines(left, right [][]int) [][]int {
    var merged [][]int
    i, j := 0, 0
    ls, rs := len(left), len(right)
    ly, ry, my := 0, 0, 0
    for i < ls && j < rs {
        var next []int
        switch {
        case left[i][0] < right[j][0]:
            next = []int{left[i][0], max(left[i][1], ry)}
            ly, i = left[i][1], i+1
        case right[j][0] < left[i][0]:
            next = []int{right[j][0], max(right[j][1], ly)}
            ry, j = right[j][1], j+1
        default:
            next = []int{right[j][0], max(right[j][1], left[i][1])}
            ly, ry, i, j = left[i][1], right[j][1], i+1, j+1
        }
        if next[1] != my {
            merged = append(merged, next)
            my = next[1]
        }
    }
    if i < ls {
        merged = append(merged, left[i:]...)
    }
    if j < rs {
        merged = append(merged, right[j:]...)
    }
    return merged
}

func getSkylineImpl(buildings [][]int, depth int) [][]int {
    n := len(buildings)
    if n == 0 {
        return [][]int{}
    }
    if n == 1 {
        return [][]int{{buildings[0][0], buildings[0][2]}, {buildings[0][1], 0}}
    }
    if depth < 3 {
        l, r := make(chan [][]int), make(chan [][]int)
        go func() {
            l <- getSkylineImpl(buildings[:n/2], depth+1)
        }()
        go func() {
            r <- getSkylineImpl(buildings[n/2:], depth+1)
        }()
        return mergeSkylines(<-l, <-r)
    }
    l := getSkylineImpl(buildings[:n/2], depth+1)
    r := getSkylineImpl(buildings[n/2:], depth+1)
    return mergeSkylines(l, r)
}

func GetSkyline(buildings [][]int) [][]int {
    return getSkylineImpl(buildings, 0)
}
