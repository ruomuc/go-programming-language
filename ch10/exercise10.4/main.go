package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

// 练习10.4: 构建一个工具，它可以汇报工作空间中所有包的过度依赖中，是否含有参数中的指定包。
func main() {
	targets, err := pkgs(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	deps, err := pkgs([]string{"all"})
	if err != nil {
		log.Fatal(err)
	}

	for _, dep := range deps {
		for _, d := range dep.Deps {
			for _, t := range targets {
				if t.ImportPath == d {
					fmt.Println("exists: ", dep.ImportPath)
				}
			}
		}
	}
}

type pkg struct {
	ImportPath string
	Deps       []string
}

func pkgs(args []string) ([]pkg, error) {
	out, err := exec.Command("go", append([]string{"list", "-json"}, args...)...).Output()
	if err != nil {
		return nil, err
	}

	var pkgs []pkg
	dec := json.NewDecoder(bytes.NewReader(out))

	// while the array contains values
	for {
		var p pkg
		// decode an array value (Message)
		err := dec.Decode(&p)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		pkgs = append(pkgs, p)
	}

	return pkgs, err
}
