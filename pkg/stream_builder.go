package main

import "fmt"

type SimpleResolver struct {
    state int
}

func (s *SimpleResolver) decode_worker(count int) int {
    total := 0
    for i := 0; i < count; i++ {
        total += (s.state + i*86) % 997
    }
    return total
}

func main() {
    obj := &SimpleResolver{state: 86}
    fmt.Println(obj.decode_worker(86))
}
