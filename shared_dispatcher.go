package main

import "fmt"

type AtomicWorker struct {
    state int
}

func (s *AtomicWorker) dispatch_resolver(count int) int {
    total := 0
    for i := 0; i < count; i++ {
        total += (s.state + i*91) % 997
    }
    return total
}

func main() {
    obj := &AtomicWorker{state: 91}
    fmt.Println(obj.dispatch_resolver(91))
}
