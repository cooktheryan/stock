package main

import (
	"fmt"
	"os/exec"
	"testing"
)

func TestMain(t *testing.T) {
	// should pass
	cmd := exec.Command("go", "run", "stock.go")
	cmd.Args = append(cmd.Args, "GME")
	err := cmd.Run()
	if err != nil {
		fmt.Printf("t: %v\n", err)
		t.Fatalf("main should have been able to run")
	}

	// should return a nonzero exit status
	cmd = exec.Command("go", "run", "stock.go")
	err = cmd.Run()
	if err == nil {
		t.Fatalf("main should return an exit when no args are passed")
	}
}
