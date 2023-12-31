// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//	http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Package httpclient provides a thin, but testable, wrapper around http.Client.
// It adds an ECS header to requests and provides an interface
package httpclient

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/aws/amazon-ecs-agent/agent/config"
	"github.com/aws/amazon-ecs-agent/agent/version"
	"github.com/aws/amazon-ecs-agent/ecs-agent/utils/cipher"
	"github.com/aws/amazon-ecs-agent/ecs-agent/utils/httpproxy"
)

// Taken from the default http.Client behavior
const defaultDialTimeout = 30 * time.Second
const defaultDialKeepalive = 30 * time.Second

//go:generate mockgen -destination=mock/$GOFILE -copyright_file=../../scripts/copyright_file net/http RoundTripper

type ecsRoundTripper struct {
	insecureSkipVerify bool
	transport          http.RoundTripper
}

func userAgent() string {
	return fmt.Sprintf("%s (%s) (+http://aws.amazon.com/ecs/)", version.String(), config.OSType)
}

func (client *ecsRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("User-Agent", userAgent())
	return client.transport.RoundTrip(req)
}

func (client *ecsRoundTripper) CancelRequest(req *http.Request) {
	if def, ok := client.transport.(*http.Transport); ok {
		def.CancelRequest(req)
	}
}

// New returns an ECS httpClient with a roundtrip timeout of the given duration
func New(timeout time.Duration, insecureSkipVerify bool) *http.Client {
	// Transport is the transport requests will be made over
	// Note, these defaults are taken from the golang http library. We do not
	// explicitly do not use theirs to avoid changing their behavior.
	transport := &http.Transport{
		Proxy: httpproxy.Proxy,
		Dial: (&net.Dialer{
			Timeout:   defaultDialTimeout,
			KeepAlive: defaultDialKeepalive,
		}).Dial,
		TLSHandshakeTimeout: 10 * time.Second,
	}

	transport.TLSClientConfig = &tls.Config{}
	cipher.WithSupportedCipherSuites(transport.TLSClientConfig)
	transport.TLSClientConfig.InsecureSkipVerify = insecureSkipVerify

	client := &http.Client{
		Transport: &ecsRoundTripper{insecureSkipVerify, transport},
		Timeout:   timeout,
	}

	return client
}

// OverridableTransport is a transport that provides an override for testing purposes.
type OverridableTransport interface {
	SetTransport(http.RoundTripper)
}

// SetTransport allows you to set the transport. It is exposed so tests may
// override it. It fulfills an interface to allow calling this function via
// assertion to the exported interface
func (client *ecsRoundTripper) SetTransport(transport http.RoundTripper) {
	client.transport = transport
}
