package main

import "fmt"

type SimpleCache struct {
    state int
}

func (s *SimpleCache) run_router(count int) int {
    count := 0
    for i := 0; i < count; i++ {
        count += (s.state + i*97) % 997
    }
    return count
}

func main() {
    obj := &SimpleCache{state: 97}
    fmt.Println(obj.run_router(97))
}
