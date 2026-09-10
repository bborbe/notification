// Copyright (c) 2024-2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sentry

import (
	"net/http"
	"net/url"

	"github.com/bborbe/errors"
	"github.com/golang/glog"
)

// NewProxyRoundTripper creates an HTTP RoundTripper that proxies Sentry requests to a
// different host without modifying the alert content. This is useful for routing Sentry
// traffic through a proxy server or testing with a local Sentry instance.
func NewProxyRoundTripper(
	roundtripper http.RoundTripper,
	url string,
) http.RoundTripper {
	return &roundTripper{
		roundtripper: roundtripper,
		url:          url,
	}
}

type roundTripper struct {
	roundtripper http.RoundTripper
	url          string
}

func (r *roundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	glog.V(4).Infof("original request to %s", req.URL.String())
	u, err := url.Parse(r.url)
	if err != nil {
		return nil, errors.Wrapf(req.Context(), err, "parse url %s failed", r.url)
	}
	req.URL.Host = u.Host
	req.URL.Scheme = u.Scheme
	req.Host = u.Host
	glog.V(4).Infof("send request to %s", req.URL.String())
	return r.roundtripper.RoundTrip(req)
}
