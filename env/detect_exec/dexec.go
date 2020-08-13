package main

import (
	"fmt"
	"log"
	"os/exec"
)

func checkDep(name string) error {
	if _, err := exec.LookPath(name); err != nil {
		em := "Could not find %q in PATH %q"
		return fmt.Errorf(em, name, err)
	}
	return nil
}

func main() {
	err := checkDep("fotune") // this executable does not exist
	if err != nil {
		log.Fatalln(err)
	}
}
