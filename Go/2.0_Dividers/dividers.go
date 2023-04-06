package main

import (
    "fmt"
)

func main() {
    var a uint32
    var l []uint32 = make([]uint32, 0, 10)

    fmt.Scan(&a)
    fmt.Printf("graph {\n")
    for i := uint32(1); i <= a; i++ {
        if a%i == 0 {
            l = append(l, i)
            fmt.Printf("\t\"%d\"\n", i)
        }
    }
    for v := 0; v < len(l); v++ {
        for u := v + 1; u < len(l); u++ {
            if l[u]%l[v] == 0 {
                isEdge := true
                for w := v + 1; w < u; w++ {
                    if l[u]%l[w] == 0 && l[w]%l[v] == 0 {
                        isEdge = false
                        break
                    }
                }
                if isEdge {
                    fmt.Printf("\t\"%d\"--\"%d\"\n", l[v], l[u])
                }
            }
        }
    }
    fmt.Printf("}")
}
