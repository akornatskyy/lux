package main

import (
	"net/url"
	"testing"
)

func TestPkgPath(t *testing.T) {
	Config.ResolverURL, _ = url.Parse("http://localhost:8000/someone/")
	Config.PkgDir = "/home/user/.lux/pkg"
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args struct {
			name string
		}
		want    string
		wantErr bool
	}{
		{"Short", args{"lua"}, "/home/user/.lux/pkg/localhost/someone/lua.yml", false},
		{"Long", args{"ak/lua"}, "/home/user/.lux/pkg/localhost/ak/lua.yml", false},
		{"VeryLong", args{"ak/lua/x"}, "/home/user/.lux/pkg/localhost/ak/lua/x.yml", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PkgPath(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("PkgPath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("PkgPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPkgURL(t *testing.T) {
	Config.ResolverURL, _ = url.Parse("http://localhost:8000/someone/")
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"Short", args{"lua"}, "http://localhost:8000/someone/lux-pkg/master/lua.yml", false},
		{"Long", args{"ak/lua"}, "http://localhost:8000/ak/lux-pkg/master/lua.yml", false},
		{"VeryLong", args{"ak/lua/x"}, "http://localhost:8000/ak/lux-pkg/master/lua/x.yml", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PkgURL(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("PkgURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("PkgURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
