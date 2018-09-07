/*
Copyright 2018 The Kubernetes Authors.

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

package test

import (
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/sirupsen/logrus"
)

//DeleteDir deletes a directory on disk
func DeleteDir(dir string) {
	err := os.RemoveAll(dir)
	if err != nil {
		log.Fatalf("fail to remove artifact '%s': %v", dir, err)
	}
}

func linkInputArt(artsDir, artName string) {
	err := os.Symlink(path.Join(InputArtifactsDir, artName),
		path.Join(artsDir, artName))

	if err != nil {
		log.Fatalf("error creating Symlink: %v", err)
	}
}

//LinkInputArts create a symbolic link of artifact in input directory to output directory
func LinkInputArts(artsDir string, artNames ...string) {
	logrus.Infof("LinkInputArts(artsDir='%s', artNames...='%v') called ", artsDir, artNames)
	for _, art := range artNames {
		linkInputArt(artsDir, art)
	}
}

//NewArtsDir create an artifact directory on disk
func NewArtsDir(dirPrefix string) string {
	os.MkdirAll(tmpArtsDir, 0755)
	dir, err := ioutil.TempDir(tmpArtsDir, dirPrefix+"_")
	logrus.Infof("artsDir='%s'", dir)
	if err != nil {
		log.Fatalf("Error making TempDir for arts: %v\n", err)
	} else {
		logrus.Infof("Temp arts dir created: %s\n", dir)
	}
	return dir
}