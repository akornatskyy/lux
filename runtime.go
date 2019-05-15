package main

import (
	"fmt"
	"sync"
)

type Runtime struct {
	sync.Mutex
	packages map[string]*Package
}

func Run(args []string) error {
	r := &Runtime{
		packages: DefaultPackages,
	}
	return r.Run(args)
}

func (r *Runtime) Run(commands []string) error {
	packages, err := r.uniquePackages(commands)
	if err != nil {
		return err
	}
	if err := r.loadPackages(packages); err != nil {
		return err
	}
	for _, c := range commands {
		if err := Command(c).Run(r.packages); err != nil {
			return err
		}
	}

	return nil
}

func (r *Runtime) uniquePackages(commands []string) ([]string, error) {
	var packages []string
	for _, s := range commands {
		c := Command(s)
		if !c.IsValid() {
			return nil, fmt.Errorf("%s: invalid name", c)
		}
		name := c.Package()
		p := r.packages[name]
		if p != nil {
			continue
		}
		for _, n := range packages {
			if n == name {
				goto NEXT
			}
		}
		packages = append(packages, name)
	NEXT:
	}
	return packages, nil
}

func (r *Runtime) loadPackages(packages []string) error {
	var e error
	var wg sync.WaitGroup
	for _, n := range packages {
		wg.Add(1)
		go func(name string) {
			defer wg.Done()
			p, err := LoadPackage(name)
			if err != nil {
				e = err
				return
			}
			r.Lock()
			r.packages[name] = p
			r.Unlock()
		}(n)
	}
	wg.Wait()
	if e != nil {
		return e
	}
	return nil
}
