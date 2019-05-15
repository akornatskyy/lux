package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"

	"gopkg.in/yaml.v2"
)

func LoadPackage(name string) (*Package, error) {
	filepath, err := PkgPath(name)
	if err != nil {
		return nil, err
	}
	if _, err := os.Stat(filepath); os.IsNotExist(err) || Config.Update {
		url, err := PkgURL(name)
		if err != nil {
			return nil, err
		}
		if err := DownloadFile(filepath, url); err != nil {
			return nil, fmt.Errorf("%s: %s", name, err)
		}
	}

	f, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	p := &Package{}
	err = yaml.Unmarshal(f, p)
	if err != nil {
		return nil, err
	}
	if p.Name == "" {
		p.Name = name
	}
	return p, nil
}

func DownloadFile(filepath, url string) error {
	if Config.Verbose {
		fmt.Printf("downloading from %s\n", url)
	}
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	if r.StatusCode != 200 {
		return fmt.Errorf("status %d: %s", r.StatusCode, url)
	}

	d := path.Dir(filepath)
	if _, err := os.Stat(d); os.IsNotExist(err) {
		if err := os.MkdirAll(d, os.ModePerm); err != nil {
			return err
		}
	}

	f, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, r.Body)
	return err
}
