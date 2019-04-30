package main

import (
  "sync"
)

type Cat struct {
  sync.Mutex
  Age int
}

func (c *Cat) IncAge() {
  c.Lock()
  defer c.Unlock()
  c.Age++
}

type Dog struct {
  ops chan opFunc
}

type opFunc func(int)

func (d *Dog) Listen() {
  var i int
  for op := range d.ops {
    op(i)
  }
}

func main() {
  // ...
}
