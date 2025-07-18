package core

import (
	"errors"
	"net/url"
	"regexp"
	"strings"
)

var urlSchemaPattern = regexp.MustCompile(`^([a-z0-9+\-.]+://)|mailto:|news:`)

func ParesUrl(rawUrl string) (*Url, error) {
	u := &Url{}
	if err := u.Parse(rawUrl); err != nil {
		return nil, err
	}
	return u, nil
}

func (x *Url) Parse(rawUrl string) error {
	if x == nil {
		return errors.New("url is nil")
	}

	u, err := url.Parse(rawUrl)
	if err != nil {
		return err
	}

	x.Scheme = u.Scheme
	x.Authority = &Authority{
		Host: u.Hostname(),
		Port: u.Port(),
	}

	if x.Authority.Host == "" && (u.Scheme == "" || u.Opaque == "") && urlSchemaPattern.MatchString(rawUrl) {
		u, err := url.Parse("https://" + rawUrl)
		if err != nil {
			return err
		}
		x.Scheme = ""
		x.Authority.Host = u.Hostname()
		if x.Authority.Host == "" {
			return errors.New("failed to parse url")
		}
	}

	x.Path = u.Path
	x.RawPath = u.RawPath
	x.Fragment = u.Fragment
	x.RawFragment = u.RawFragment

	x.Query = &Query{
		Values: make(map[string]*StringValues),
	}

	query, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return err
	}
	for k, v := range query {
		x.Query.Values[k] = &StringValues{
			Values: v,
		}
	}

	return nil
}

func (x *Url) Format() string {
	if x == nil {
		return ""
	}

	host := ""
	if x.Authority != nil {
		host = x.Authority.Host
		if x.Authority.Port != "" {
			host += ":" + x.Authority.Port
		}
	}

	u := url.URL{
		Scheme:   x.Scheme,
		Host:     host,
		Path:     x.Path,
		Fragment: x.Fragment,
	}

	if x.Query != nil {
		query := url.Values{}
		for k, v := range x.Query.Values {
			if v != nil {
				query[k] = v.Values
			}
		}
		u.RawQuery = query.Encode()
	}

	if u.Scheme == "" {
		return strings.TrimPrefix(u.String(), "//")
	}

	return u.String()
}

func (x *Url) ToString() string {
	return x.Format()
}

func (x *Url) FormatWithoutSchema() string {
	if x == nil {
		return ""
	}

	u := &Url{
		Authority: &Authority{
			Host: x.GetAuthority().GetHost(),
			Port: x.GetAuthority().GetPort(),
		},
		Path:     x.Path,
		Query:    x.Query,
		Fragment: x.Fragment,
	}

	return strings.TrimPrefix(u.String(), "//")
}
