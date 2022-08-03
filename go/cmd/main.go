package main

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"sigs.k8s.io/yaml"
)

func print(r io.Reader) {
	data, err := ioutil.ReadAll(r)
	if err == nil {
		var m interface{}

		err = yaml.Unmarshal(data, &m)
		if err == nil {
			data, err = yaml.Marshal(m)
			for _, l := range strings.Split(string(data), "\n") {
				fmt.Printf("      : %s\n", l)
			}
		}
	}
}
func main() {
	fmt.Printf("hello OCM!\n")
	for i, a := range os.Args {
		fmt.Printf("arg %d: %s\n", i, a)
	}

	data, err := ioutil.ReadFile("ocm/inputs/ocmrepo")
	if err == nil {
		fmt.Printf("found ocmrepo\n")
		file, err := os.Open("ocm/inputs/ocmrepo")
		if err == nil {
			reader, err := gzip.NewReader(file)
			if err == nil {
				tr := tar.NewReader(reader)
				for {
					header, err := tr.Next()
					if err != nil {
						break
					}
					if header.Typeflag == tar.TypeDir {
						fmt.Printf("dir  %s\n", header.Name)
					} else {
						fmt.Printf("file %s\n", header.Name)
						print(tr)
					}
				}
			}
		}
	}
	data, err = ioutil.ReadFile("ocm/inputs/parameters")
	if err == nil {
		fmt.Printf("found parameters:\n%s\n", string(data))
	}
	data, err = ioutil.ReadFile("ocm/inputs/config")
	if err == nil {
		fmt.Printf("found config:\n%s\n", string(data))
	}
	data, err = ioutil.ReadFile("ocm/inputs/ocmconfig")
	if err == nil {
		fmt.Printf("found ocm config:\n%s\n", string(data))
	}
	os.MkdirAll("ocm/outputs", 0744)
	ioutil.WriteFile("ocm/outputs/test", []byte("result"), 0644)
}
