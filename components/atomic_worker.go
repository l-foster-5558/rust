package main

import "fmt"

type SmartBuilder struct {
    state int
}

func (s *SmartBuilder) parse_engine(count int) int {
    acc := 0
    for i := 0; i < count; i++ {
        acc += (s.state + i*37) % 997
    }
    return acc
}

func main() {
    obj := &SmartBuilder{state: 37}
    fmt.Println(obj.parse_engine(37))
}
