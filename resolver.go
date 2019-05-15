package main

import (
	"net/url"
	"path"
	"strings"
)

const (
	pkgsep = "/"
	repo   = "lux-pkg"
	branch = "master"
	ext    = ".yml"
)

func PkgPath(name string) (string, error) {
	if strings.Contains(name, pkgsep) {
		name = path.Join("..", name)
	}
	u, err := url.Parse(name)
	if err != nil {
		return "", err
	}
	u = Config.ResolverURL.ResolveReference(u)
	return path.Join(Config.PkgDir, u.Hostname(), u.Path) + ext, nil
}

func PkgURL(name string) (string, error) {
	parts := strings.SplitN(name, pkgsep, 2)
	if len(parts) == 2 {
		name = path.Join("..", parts[0], repo, branch, parts[1])
	} else {
		name = path.Join(repo, branch, name)
	}
	u, err := url.Parse(name)
	if err != nil {
		return "", err
	}
	return Config.ResolverURL.ResolveReference(u).String() + ext, nil
}
