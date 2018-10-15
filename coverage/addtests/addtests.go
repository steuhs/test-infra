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

 // Package addtests adds empty _test.go file for all .go files without corresponding _test.go file.
 // We need this because 'go test -coverprofile' ignores all .go file without _test.go file in the profile.

package addtests

import (
	"os"
	"path/filepath"

	"k8s.io/test-infra/coverage/io"
	"k8s.io/test-infra/coverage/fileops"
	"bufio"
	"strings"
)

// creates template file with no content but the package name
func createTemplateFile(path string) error{
	packageName := fileops.PackageName(path)
	content := "package " + packageName
	return io.WriteNew(&content, path)
}

func getPackageLine(path string) string {
	file, err := os.Open(path)
	if err != nil {
		return ""
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		line := scanner.Text()
		if strings.HasPrefix(line, "package ") {
			return line
		}
	}
	return ""
}

func addTest(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	fileName := info.Name()
	if !info.IsDir() && fileops.IsSourceFile(fileName) {
		testPath := fileops.SourceToTestFilePath(path)
		exists, ioErr := io.FileOrDirExistsNew(testPath)
		if err != nil {
			return ioErr
		}
		if !exists && !isVendor(path) && !fileops.IsCoverageSkipped(path){
			content := getPackageLine(path)
			return io.WriteNew(&content, testPath)
			//return createTemplateFile(testPath)
		}
	}
	return nil
}

func AddTests(dir string) error {
	return filepath.Walk(dir, addTest)
}

func isVendor(path string) bool {
	return strings.HasPrefix(path, "vendor/") || strings.Contains(path, "/vendor/")
}