package main

import (
	"github.com/OPPOGROUP/hoyolib/internal/config"
	"github.com/OPPOGROUP/hoyolib/internal/log"
	"github.com/spf13/pflag"
)

var (
	prod = pflag.BoolP("prod", "p", false, "use --prod to run as release version")
)

func init() {
	pflag.Lookup("prod").NoOptDefVal = "true"
	pflag.Parse()
}

func main() {
	err := config.Init()
	if err != nil {
		panic(err)
	}
	log.Init(*prod)
}
