package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"time"
)

type Bundle struct {
	Name    string
	Timeout time.Duration
	Env     []string
	Run     []string
	Scripts []string
}

func (b *Bundle) Environ() []string {
	return append(append(Config.Env, b.Env...), os.Environ()...)
}

func (b *Bundle) RunAll() error {
	ctx, cancel := context.WithTimeout(context.Background(), b.Timeout)
	defer cancel()
	for i, s := range b.Scripts {
		if err := b.runScript(ctx, s); err != nil {
			if ctx.Err() != nil {
				err = fmt.Errorf("timed out after %v", b.Timeout)
			}
			return fmt.Errorf("%s %s %v", b.Name, b.Run[i], err)
		}
	}
	return nil
}

func (b *Bundle) runScript(ctx context.Context, script string) error {
	cmd := exec.CommandContext(ctx, "sh", "-c", script)
	cmd.Env = b.Env
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("could not get stdout pipe: %v", err)
	}
	cmd.Stderr = cmd.Stdout
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("could not run cmd: %v", err)
	}
	p := NewPrinter(b.Name)
	t := NewTracer(p)
	go t.Trace(stdout)
	if err := cmd.Wait(); err != nil {
		t.Wait()
		p.Dump()
		return err
	}
	t.Wait()
	return nil
}
