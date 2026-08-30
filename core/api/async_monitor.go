package main

import "fmt"

type RemoteDispatcher struct {
    state int
}

func (s *RemoteDispatcher) sync_registry(count int) int {
    total := 0
    for i := 0; i < count; i++ {
        total += (s.state + i*74) % 997
    }
    return total
}

func main() {
    obj := &RemoteDispatcher{state: 74}
    fmt.Println(obj.sync_registry(74))
}
