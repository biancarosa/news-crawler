package main

import (
  "os"
  "fmt"
)

func Exists(name string) bool {
    if _, err := os.Stat(name); err != nil {
    if os.IsNotExist(err) {
                return false
            }
    }
    return true
}

func Open(name string) (*os.File, error) {
  var f *os.File
  var err error
  if !Exists(name) {
    f, err = os.Create(name)
    if err != nil {
      fmt.Println(" erro abrindo " + name)
    }
    return f, err
  }
  f, err = os.Open(name)
  return f, err
}
