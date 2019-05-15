package main

import (
	"fmt"
	"os"
	"path"
	"strings"
)

type Printer interface {
	Print(msg string)
	Clear()
	Dump()
}

func NewPrinter(prefix string) Printer {
	prefix = EscReset + prefix + " " + EscCmd
	if Config.Verbose {
		return &VerbosePrinter{
			BasePrinter: &BasePrinter{
				prefix: prefix,
			},
		}
	}
	return &InlinePrinter{
		BasePrinter: &BasePrinter{
			prefix: prefix,
		},
	}
}

type BasePrinter struct {
	prefix string
	cmd    string
}

func (p *BasePrinter) Sprint(s string) string {
	cmd, msg := parse(s)
	if cmd != "" {
		cmd += " "
		p.cmd = cmd
	} else {
		cmd = p.cmd
	}
	if msg == "" {
		return ""
	}
	return p.prefix + cmd + EscMsg + msg
}

type VerbosePrinter struct {
	*BasePrinter
}

func (p *VerbosePrinter) Print(msg string) {
	fmt.Println(p.Sprint(msg))
}

func (p *VerbosePrinter) Clear() {
	fmt.Print(EscReset)
}

func (p *VerbosePrinter) Dump() {
}

type InlinePrinter struct {
	*BasePrinter
	buffer []string
}

func (p *InlinePrinter) Print(msg string) {
	const maxMsgLen = 79 + len(EscReset+EscCmd+EscMsg)
	s := p.Sprint(msg)
	if s == "" {
		return
	}
	p.buffer = append(p.buffer, s)
	if len(s) > maxMsgLen {
		s = string(s[:maxMsgLen])
	}
	fmt.Printf("%-79s\r%s\r", "", s)
}

func (p *InlinePrinter) Clear() {
	fmt.Printf("%-79s\r%s", "", EscReset)
}

func (p *InlinePrinter) Dump() {
	for _, msg := range p.buffer {
		fmt.Fprintln(os.Stderr, msg)
	}
}

func parse(msg string) (string, string) {
	msg = strings.TrimSpace(msg)
	l := len(msg)
	switch {
	case l == 0:
		return "", ""
	case l > 2 && msg[0] == '+':
		s := strings.SplitN(string(msg[2:]), " ", 2)
		_, cmd := path.Split(s[0])
		if len(s) == 1 {
			return cmd, ""
		}
		return cmd, s[1]
	default:
		return "", msg
	}
}
