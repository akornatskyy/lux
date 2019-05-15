package main

import (
	"flag"
	"fmt"
	"os"
)

func usage() {
	fmt.Fprint(
		os.Stderr,
		"usage: lux [options...] [ns/]package[:bundle][ ...]\n",
	)
	flag.PrintDefaults()
	fmt.Fprintf(
		os.Stderr,
		"packages @ %slux-pkg\n",
		Config.ResolverURL,
	)
	os.Exit(2)
}

func main() {
	if err := LoadConfig(); err != nil {
		exit(err.Error())
	}

	flag.Usage = usage
	if len(os.Args) < 2 {
		exit("must provide one or more [ns/]package[:bundle]")
	}

	flag.BoolVar(&Config.Update, "u", false, "update package")
	flag.BoolVar(&Config.Verbose, "v", false, "verbose mode")
	flag.Parse()

	if err := Run(flag.Args()); err != nil {
		exit(err.Error())
	}
}

func exit(reason string) {
	fmt.Printf(EscErr+"ERR:"+EscReset+" %s\n", reason)
	os.Exit(1)
}
