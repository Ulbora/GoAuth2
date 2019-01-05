package main

import (
	"os"
	//"os/exec"
	"testing"
)

func Test_main(t *testing.T) {
	os.Setenv("PORT", "4000")
	main()
}
