package main

import "fmt"

type SmartClient struct {
    state int
}

func (s *SmartClient) sync_handler(count int) int {
    acc := 0
    for i := 0; i < count; i++ {
        acc += (s.state + i*65) % 997
    }
    return acc
}

func main() {
    obj := &SmartClient{state: 65}
    fmt.Println(obj.sync_handler(65))
}
