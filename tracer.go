package main

import (
	"bufio"
	"io"
	"sync"
)

type TraceReader struct {
	sync.Mutex
	p Printer
}

func NewTracer(p Printer) *TraceReader {
	return &TraceReader{
		p: p,
	}
}

func (t *TraceReader) Trace(r io.Reader) {
	t.Lock()
	defer t.Unlock()

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		t.p.Print(scanner.Text())
	}
	t.p.Clear()
}

func (t *TraceReader) Wait() {
	t.Lock()
	defer t.Unlock()
}
