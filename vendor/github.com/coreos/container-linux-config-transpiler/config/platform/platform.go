// Copyright 2017 CoreOS, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package platform

const (
	Azure                 = "azure"
	DO                    = "digitalocean"
	EC2                   = "ec2"
	GCE                   = "gce"
	Packet                = "packet"
	OpenStackMetadata     = "openstack-metadata"
	VagrantVirtualbox     = "vagrant-virtualbox"
	CloudStackConfigDrive = "cloudstack-configdrive"
	Custom                = "custom"
)

var Platforms = []string{
	Azure,
	DO,
	EC2,
	GCE,
	Packet,
	OpenStackMetadata,
	VagrantVirtualbox,
	CloudStackConfigDrive,
	Custom,
}

func IsSupportedPlatform(platform string) bool {
	for _, supportedPlatform := range Platforms {
		if supportedPlatform == platform {
			return true
		}
	}
	return platform == ""
}
