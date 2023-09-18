package main

import (
	"reflect"
	"regexp"
	"testing"
)

func Test_match(t *testing.T) {
	tests := []struct {
		name string
		re   *regexp.Regexp
		s    string
		want []string
	}{
		{"ips", IP, "wich one is correct 10.300.0.1 or 192.168.0.1 ?", []string{"192.168.0.1"}},
		{"email", EMAIL, "I love U <philip.morris@gmail.com>", []string{"philip.morris@gmail.com"}},
		{"domain", URL, "Please visit 'https://admin:test@contact.uvsq.fr:8443/us'", []string{"https://admin:test@contact.uvsq.fr:8443/us"}},
		{"domain2", URL, "Please visit 'http://google.ar/fake'", []string{"http://google.ar/fake"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := match(tt.s, tt.re); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("match() = %v, want %v", got, tt.want)
			}
		})
	}
}
