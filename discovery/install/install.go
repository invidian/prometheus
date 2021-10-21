// Copyright 2020 The Prometheus Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package install has the side-effect of registering all builtin
// service discovery config types.
package install

import (

	//  register aws
	_ "github.com/prometheus/prometheus/discovery/aws"

	//  register azure
	_ "github.com/prometheus/prometheus/discovery/azure"

	//  register consul
	_ "github.com/prometheus/prometheus/discovery/consul"

	//  register digitalocean
	_ "github.com/prometheus/prometheus/discovery/digitalocean"

	//  register dns
	_ "github.com/prometheus/prometheus/discovery/dns"

	//  register eureka
	_ "github.com/prometheus/prometheus/discovery/eureka"

	//  register file
	_ "github.com/prometheus/prometheus/discovery/file"

	//  register gce
	_ "github.com/prometheus/prometheus/discovery/gce"

	//  register hetzner
	_ "github.com/prometheus/prometheus/discovery/hetzner"

	//  register http
	_ "github.com/prometheus/prometheus/discovery/http"

	//  register kubernetes
	_ "github.com/prometheus/prometheus/discovery/kubernetes"

	//  register linode
	_ "github.com/prometheus/prometheus/discovery/linode"

	//  register marathon
	_ "github.com/prometheus/prometheus/discovery/marathon"

	//  register moby
	_ "github.com/prometheus/prometheus/discovery/moby"

	//  register openstack
	_ "github.com/prometheus/prometheus/discovery/openstack"

	//  register puppetdb
	_ "github.com/prometheus/prometheus/discovery/puppetdb"

	//  register scaleway
	_ "github.com/prometheus/prometheus/discovery/scaleway"

	//  register triton
	_ "github.com/prometheus/prometheus/discovery/triton"

	//  register uyuni
	_ "github.com/prometheus/prometheus/discovery/uyuni"

	//  register xds
	_ "github.com/prometheus/prometheus/discovery/xds"

	//  register zookeeper
	_ "github.com/prometheus/prometheus/discovery/zookeeper"
)
