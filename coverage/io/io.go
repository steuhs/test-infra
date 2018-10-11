/*
 * Copyright 2018 The Kubernetes Authors.
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *     http://www.apache.org/licenses/LICENSE-2.0
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package io

import (
	"fmt"
	"os"
	"path"

	"github.com/sirupsen/logrus"

	"k8s.io/test-infra/coverage/logUtil"
)

// CreateMarker produces empty file as marker
func CreateMarker(dir, fileName string) {
	Write(nil, dir, fileName)
	logrus.Infof("Created marker file '%s'\n", fileName)
}

// Write writes the content of the string to a file in the directory
func Write(content *string, destinationDir, fileName string) {
	filePath := path.Join(destinationDir, fileName)
	file, err := os.Create(filePath)
	if err != nil {
		logUtil.LogFatalf("Error writing file: %v", err)
	} else {
		logrus.Infof("Created file:%s", filePath)
		if content == nil {
			logrus.Infof("No content to be written to file '%s'", fileName)
		} else {
			fmt.Fprint(file, *content)
		}
	}
	defer file.Close()
}

//MkdirAll makes directory on disk. Recursively adds parents directory if not exist.
func MkdirAll(path string) {
	logrus.Infof("Making directory (MkdirAll): path=%s", path)
	if err := os.MkdirAll(path, 0755); err != nil {
		logrus.Fatalf("Failed os.MkdirAll(path='%s', 0755); err='%v'", path, err)
	} else {
		logrus.Infof("artifacts dir (path=%s) created successfully\n", path)
	}
}

//FileOrDirExists checks whether a file or dir on disk exist
func FileOrDirExists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			cwd, _ := os.Getwd()
			logrus.Infof("file or dir not found: %s; cwd=%s", path, cwd)
			return false
		}
		logrus.Fatalf("File stats (path=%s) err: %v", path, err)
	}
	return true
}

//FileOrDirExistsNew checks whether a file or dir on disk exist
func FileOrDirExistsNew(path string) (bool, error) {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			cwd, _ := os.Getwd()
			logrus.Infof("file or dir not found: %s; cwd=%s", path, cwd)
			return false, nil
		}
		return false, fmt.Errorf("File stats (path=%s) err: %v", path, err)
	}
	return true, nil
}
