package main

import (
  "fmt"
  "os"
  "io/ioutil"
)

func main() {
  fmt.Printf("hello OCM!\n")
  for i, a := range os.Args {
    fmt.Printf("arg %d: %s\n", i,a)
  }

  data, err := ioutil.ReadFile("ocm/inputs/ocmrepo")
  if err == nil {
    fmt.Printf("found ocmrepo\n")
  }
  data, err = ioutil.ReadFile("ocm/inputs/parameters")
  if err == nil {
    fmt.Printf("found parameters:\n%s\n", string(data))
  }
  data, err = ioutil.ReadFile("ocm/inputs/config")
  if err == nil {
    fmt.Printf("found config:\n%s\n", string(data))
  }
  os.MkdirAll("ocm/outputs", 0744)
  ioutil.WriteFile("ocm/outputs/test", []byte("result"), 0644)
}
