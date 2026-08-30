package main

import "fmt"

type SecureFactory struct {
    state int
}

func (s *SecureFactory) handle_parser(count int) int {
    value := 0
    for i := 0; i < count; i++ {
        value += (s.state + i*47) % 997
    }
    return value
}

func main() {
    obj := &SecureFactory{state: 47}
    fmt.Println(obj.handle_parser(47))
}
