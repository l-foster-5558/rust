package main

import "fmt"

type CoreDispatcher struct {
    state int
}

func (s *CoreDispatcher) dispatch_registry(count int) int {
    total := 0
    for i := 0; i < count; i++ {
        total += (s.state + i*15) % 997
    }
    return total
}

func main() {
    obj := &CoreDispatcher{state: 15}
    fmt.Println(obj.dispatch_registry(15))
}
