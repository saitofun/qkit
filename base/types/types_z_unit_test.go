package types_test

import (
	"net/url"
	"testing"
	"time"

	. "github.com/onsi/gomega"
	. "github.com/sincospro/qkit/base/types"
)

func TestAddress(t *testing.T) {
	cases1 := map[string]struct {
		Addr   *Address
		Expect string
	}{
		"asset": {
			Addr: &Address{
				Grp: "avatar",
				Key: "filename",
				Ext: "png",
			},
			Expect: "asset://avatar/filename.png",
		},
		"url": {
			Addr: &Address{
				URL: "http://venusgroup.com/avatar/filename.png",
			},
			Expect: "http://venusgroup.com/avatar/filename.png",
		},
		"without_ext": {
			Addr: &Address{
				Grp: "avatar",
				Key: "filename",
			},
			Expect: "asset://avatar/filename",
		},
		"local_file": {
			Addr: &Address{
				URL: "file:///AbsPath/To/Your/Local/File.ext",
			},
			Expect: "file:///AbsPath/To/Your/Local/File.ext",
		},
	}

	for name, c := range cases1 {
		t.Run("string:"+name, func(t *testing.T) {
			NewWithT(t).Expect(c.Addr.String()).To(Equal(c.Expect))
		})
	}

	cases2 := map[string]string{
		"asset":       "asset://avatar/filename.png",
		"url":         "http://venusgroup.com/avatar/filename.png",
		"without_ext": "asset://avatar/filename",
		"local_file":  "file:///AbsPath/To/Your/Local/File.ext",
	}

	for name, str := range cases2 {
		t.Run("parse:"+name, func(t *testing.T) {
			addr, err := ParseAddress(str)
			NewWithT(t).Expect(err).To(BeNil())
			NewWithT(t).Expect(addr.String()).To(Equal(str))
		})
	}
}

func TestBinary(t *testing.T) {
	bytes, err := Binary("917").MarshalText()
	NewWithT(t).Expect(err).To(BeNil())
	NewWithT(t).Expect(bytes).To(Equal([]byte("917")))
}

func TestDuration(t *testing.T) {
	cases := map[string]struct {
		t Duration
		s string
	}{
		"ns": {AsDuration(time.Nanosecond), "1ns"},
		"us": {AsDuration(time.Microsecond), "1us"},
		"ms": {AsDuration(time.Millisecond), "1ms"},
		"s ": {AsDuration(time.Second), "1s"},
		"m ": {AsDuration(time.Minute), "1m0s"},
		"h ": {AsDuration(time.Hour), "1h0m0s"},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			du, err := time.ParseDuration(c.s)
			NewWithT(t).Expect(err).To(BeNil())
			NewWithT(t).Expect(int64(du)).To(Equal(int64(c.t)))
		})
	}
}

func TestEndpoint(t *testing.T) {
	cases := map[string]struct {
		ep  *Endpoint
		url string
	}{
		"stmps": {
			ep: &Endpoint{
				Scheme:   "stmps",
				Hostname: "stmps.mail.xxx.com",
				Port:     465,
			},
			url: "stmps://stmps.mail.xxx.com:465",
		},
		"postgres": {
			ep: &Endpoint{
				Scheme:   "postgres",
				Hostname: "hostname",
				Username: "username",
				Password: "password",
				Port:     5432,
				Base:     "database_name",
				Param:    url.Values{"sslmode": {"disable"}},
			},
			url: "postgres://username:password@hostname:5432/database_name?sslmode=disable",
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			NewWithT(t).Expect(c.ep.String()).To(Equal(c.url))
			parsed, err := ParseEndpoint(c.url)
			NewWithT(t).Expect(err).To(BeNil())
			NewWithT(t).Expect(*parsed).To(Equal(*c.ep))
		})
	}
}

type fakeTextStruct struct{}

func (f fakeTextStruct) MarshalText() ([]byte, error) { return nil, nil }

func (f fakeTextStruct) UnmarshalText([]byte) error { return nil }

type fakeTextStructRef struct{}

func (f *fakeTextStructRef) MarshalText() ([]byte, error) { return nil, nil }

func (f *fakeTextStructRef) UnmarshalText([]byte) error { return nil }

func TestInterface(t *testing.T) {}
