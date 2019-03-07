/*
Copyright 2019 ceph-csi authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package util

import (
	"fmt"
	"io/ioutil"
	"k8s.io/klog"
	"path"
)

/*
FileConfig is a ConfigStore interface implementation that reads configuration
information from files.

BasePath defines the directory under which FileConfig will attempt to open and
read contents of various Ceph cluster configurations.

Each Ceph cluster configuration is stored under a directory named,
BasePath/ceph-cluster-<fsid>, where <fsid> is the Ceph cluster fsid.

Under each Ceph cluster configuration directory, individual files named as per
the ConfigKeys constants in the ConfigStore interface, store the required
configuration information.
*/
type FileConfig struct {
	BasePath string
}

// DataForKey reads the appropriate config file, named using key, and returns
// the contents of the file to the caller
func (fc *FileConfig) DataForKey(fsid string, key string) (data string, err error) {
	pathToKey := path.Join(fc.BasePath, "ceph-cluster-"+fsid, key)
	// #nosec
	content, err := ioutil.ReadFile(pathToKey)
	if err != nil || string(content) == "" {
		err = fmt.Errorf("error fetching configuration for cluster ID (%s). (%s)", fsid, err)
		return
	}

	data = string(content)
	klog.V(3).Infof("returning data (%s) for key (%s) against cluster (%s)", data, key, fsid)
	return
}