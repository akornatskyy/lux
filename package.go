package main

import (
	"fmt"
	"time"
)

type Package struct {
	Name    string
	Default string
	Timeout time.Duration
	Bundles map[string]*Bundle
	Scripts map[string]string
}

func (p *Package) ResolveBundle(bundle string) (*Bundle, error) {
	if bundle == "" {
		bundle = p.Default
	}
	b := p.Bundles[bundle]
	if b == nil {
		return nil, fmt.Errorf("%s: unknown bundle %s", p.Name, bundle)
	}
	var scripts []string
	for _, s := range b.Run {
		script, ok := p.Scripts[s]
		if !ok {
			script, ok = DefaultScripts[s]
			if !ok {
				return nil, fmt.Errorf("%s: unknown script %s", p.Name, s)
			}
		}
		scripts = append(scripts, script)
	}
	name := p.Name
	if bundle != "" {
		name += ":" + bundle
	}
	t := b.Timeout
	if t == 0 {
		t = p.Timeout
		if t == 0 {
			t = DefaultTimeout
		}
	}
	return &Bundle{
		Name:    name,
		Env:     b.Environ(),
		Run:     b.Run,
		Scripts: scripts,
		Timeout: t,
	}, nil
}
