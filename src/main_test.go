package main

import (
  "testing"
)

var x, y float64 = 3, 4.5

func TestMultiplyPass(t *testing.T) {
    fn := Multiply
    v := 13.5
    if val := fn(x, y); val != v {
        t.Fatalf("expected %d, got %d.", v, val)
    }
}

func TestAdd(t *testing.T) {
    fn := Add
    v := 7.5
    if val := fn(x, y); val != v {
        t.Fatalf("expected %d, got %d.", v, val)
    }
}
