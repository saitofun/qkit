package datatype

import (
	"log"
	"net/url"
	"strconv"
	"strings"
)

type Endpoint struct {
	Scheme   string
	Hostname string
	Port     uint16
	Base     string
	Username string
	Password Password
	Param    url.Values
}

func (e Endpoint) String() string {
	u := url.URL{
		Scheme:   e.Scheme,
		Host:     e.Host(),
		Path:     "/" + e.Base,
		RawPath:  "",
		RawQuery: e.Param.Encode(),
	}

	if e.Base != "" {
		u.Path = "/" + e.Base
	}

	if e.Param != nil {
		u.RawQuery = e.Param.Encode()
	}

	if e.Username != "" || e.Password != "" {
		u.User = url.UserPassword(e.Username, e.Password.String())
	}

	s, err := url.QueryUnescape(u.String())
	if err != nil {
		log.Fatal(err)
	}
	return s
}

func (e Endpoint) SecurityString() string {
	e.Password = Password(e.Password.SecurityString())
	return e.String()
}

func (e Endpoint) Host() string {
	if e.Port == 0 {
		return e.Hostname
	}
	return e.Hostname + ":" + strconv.FormatUint(uint64(e.Port), 10)
}

func ParseEndpoint(text string) (*Endpoint, error) {
	u, err := url.ParseRequestURI(text)
	if err != nil {
		return nil, err
	}

	ep := &Endpoint{}

	ep.Scheme = u.Scheme

	if q := u.Query(); len(q) > 0 {
		ep.Param = q
	}

	if len(u.Path) > 0 {
		ep.Base = strings.Split(u.Path[1:], "/")[0]
	}

	return nil, nil
}

func (e Endpoint) MarshaText() ([]byte, error) { return nil, nil }

func (e *Endpoint) UnmarshalText([]byte) error { return nil }
