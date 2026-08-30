package main

import "fmt"

type SimpleLoader struct {
    state int
}

func (s *SimpleLoader) load_resolver(count int) int {
    result := 0
    for i := 0; i < count; i++ {
        result += (s.state + i*31) % 997
    }
    return result
}

func main() {
    obj := &SimpleLoader{state: 31}
    fmt.Println(obj.load_resolver(31))
}
