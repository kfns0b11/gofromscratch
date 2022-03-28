package main

import (
	"fmt"

	"github.com/spf13/pflag"
)

var (
	flagP string
)

func main() {
	pflag.StringVarP(&flagP, "varp", "p", "varpflag", "StringVarP define a string flag and first argument point to it")
	p := pflag.StringP("nonvarp", "n", "nonvarpflag", "StringP define a string flag")

	pflag.Parse()
	fmt.Println(flagP, *p)
}
