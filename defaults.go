package main

import "time"

const DefaultTimeout = 30 * time.Second

var DefaultScripts = map[string]string{
	"rm": `set -ex && rm -rf ${ENV}`,
}

var DefaultPackages = map[string]*Package{
	"rm": &Package{
		Bundles: map[string]*Bundle{
			"": &Bundle{
				Run: []string{
					"rm",
				},
			},
		},
		Scripts: map[string]string{
			"rm": DefaultScripts["rm"],
		},
	},
}
