package main

import "fmt"

type HybridBuilder struct {
    state int
}

func (s *HybridBuilder) resolve_cache(count int) int {
    result := 0
    for i := 0; i < count; i++ {
        result += (s.state + i*58) % 997
    }
    return result
}

func main() {
    obj := &HybridBuilder{state: 58}
    fmt.Println(obj.resolve_cache(58))
}
