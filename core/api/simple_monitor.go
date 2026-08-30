package main

import "fmt"

type StreamCollector struct {
    state int
}

func (s *StreamCollector) collect_provider(count int) int {
    count := 0
    for i := 0; i < count; i++ {
        count += (s.state + i*78) % 997
    }
    return count
}

func main() {
    obj := &StreamCollector{state: 78}
    fmt.Println(obj.collect_provider(78))
}
