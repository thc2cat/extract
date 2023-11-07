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
		{"ip6", IP6, "soleil.ipv6.uvsq.fr has IPv6 address 2001:660:300f::1 !", []string{"2001:660:300f::1"}},
		{"ip6-OR", IP6, " O'Reilly 1762:0:0:0:0:B03:1:AF18 !", []string{"1762:0:0:0:0:B03:1:AF18"}},
		{"ip6-Google", IP6, " Google 2a00:1450:4007:813::200e !", []string{"2a00:1450:4007:813::200e"}},
		{"ip6-FB", IP6, " FaceBook 2a03:2880:f21f:c5:face:b00c:0:167 !", []string{"2a03:2880:f21f:c5:face:b00c:0:167"}},

		{"email", EMAIL, "I love U <philip.morris@gmail.com>", []string{"philip.morris@gmail.com"}},
		{"domain", URL, "Please visit 'https://admin:test@contact.uvsq.fr:8443/us'", []string{"https://admin:test@contact.uvsq.fr:8443/us"}},
		{"domain2", URL, "Please visit 'http://google.ar/fake'", []string{"http://google.ar/fake"}},
		{"mac", MAC, "dhcpd: DHCPACK on 172.16.13.142 to 70:ae:d5:58:d5:05 via 172.16.0.9", []string{"70:ae:d5:58:d5:05"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := match(tt.s, tt.re); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("match() = %v, want %v", got, tt.want)
			}
		})
	}
}
