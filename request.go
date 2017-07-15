package main

import (
  "io/ioutil"
  "strings"
)

func GetUrlsFromFile(fileName string) []string {
  data, err := ioutil.ReadFile(fileName)
  if err != nil {
    println("erro ao abrir oa arquivo")
  }
  urls := strings.Split(string(data), "\n")
  return urls
}
