package main

import (
	"regexp"
	"strings"
)

type Command string

var regexpPkgName = regexp.MustCompile(`^\w+([/\w\-]+)?(:[\w\.\-]+)?$`)

func (c Command) IsValid() bool {
	return regexpPkgName.MatchString(string(c))
}

func (c Command) Package() string {
	return strings.SplitN(string(c), ":", 2)[0]
}

func (c Command) Parts() (string, string) {
	parts := strings.SplitN(string(c), ":", 2)
	if len(parts) == 1 {
		return parts[0], ""
	}
	return parts[0], parts[1]
}

func (c Command) Run(packages map[string]*Package) error {
	pkg, bundle := c.Parts()
	p := packages[pkg]
	b, err := p.ResolveBundle(bundle)
	if err != nil {
		return err
	}
	return b.RunAll()
}
