package main

import (
	"net/url"
	"os"
	"path"
	"strings"
)

var Config struct {
	ResolverURL *url.URL
	Env         []string
	PkgDir      string

	Update  bool
	Verbose bool
}

func LoadConfig() error {
	s := os.Getenv("LUX_URL")
	if s == "" {
		s = "https://raw.githubusercontent.com/akornatskyy/"
	} else if !strings.HasSuffix(s, "/") {
		s += "/"
	}
	url, err := url.Parse(s)
	if err != nil {
		return err
	}
	Config.ResolverURL = url

	s = path.Join(os.Getenv("HOME"), ".cache", "lux", "downloads")
	if err := os.MkdirAll(s, os.ModePerm); err != nil {
		return err
	}
	Config.Env = []string{
		"CACHE=" + s,
		"ENV=" + path.Join(os.Getenv("PWD"), "env"),
	}

	s = os.Getenv("LUX_PKG")
	if s == "" {
		s = path.Join(os.Getenv("HOME"), ".cache", "lux", "pkg")
		if err := os.MkdirAll(s, os.ModePerm); err != nil {
			return err
		}
	}
	Config.PkgDir = s

	return nil
}
