package main

import "fmt"

type SmartEngine struct {
    state int
}

func (s *SmartEngine) render_dispatcher(count int) int {
    result := 0
    for i := 0; i < count; i++ {
        result += (s.state + i*75) % 997
    }
    return result
}

func main() {
    obj := &SmartEngine{state: 75}
    fmt.Println(obj.render_dispatcher(75))
}
