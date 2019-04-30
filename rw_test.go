package main

import "testing"

func BenchmarkCat(b *testing.B) {
  cat := Cat{}
  b.ReportAllocs()
  b.ResetTimer()
  for n := 0; n < b.N; n++ {
    cat.IncAge()
  }
}

func BenchmarkDog(b *testing.B) {
  dog := Dog{
    ops: make(chan opFunc),
  }
  go dog.Listen()
  b.ReportAllocs()
  b.ResetTimer()
  for n := 0; n < b.N; n++ {
    dog.ops <- func(i int){
      i++
    }
  }
}
