package goexec

import (
	"bytes"
	"fmt"
	"os/exec"
)

func Execute(name string, arg ...string) (err error) {
	fmt.Printf("> %v %v\n", name, arg)

	cmd := exec.Command(name, arg...)

	var out bytes.Buffer
	var stderr bytes.Buffer

	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err = cmd.Run()

	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return err
	}

	if stderr.Len() > 0 {
		fmt.Printf("stderr < %s\n", stderr.String())
	}

	if out.Len() > 0 {
		fmt.Printf("stdout < %s\n", out.String())
	}

	return err
}
