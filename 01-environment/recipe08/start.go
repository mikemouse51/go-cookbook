package main

import (
	"os/exec"
	"bytes"
	"fmt"
)

func main(){
	prc := exec.Command("ls", "-lta")
	out := bytes.NewBuffer([]byte{})
	prc.Stdout = out
	err := prc.Start()
	if err != nil {
		fmt.Println(err)
	}
	prc.Wait()
	if prc.ProcessState.Success() {
		fmt.Println("Process ran successfully with output:")
		fmt.Println(out.String())
	}
}