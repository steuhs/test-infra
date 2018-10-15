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

package fileops

import (
	"strings"
	"path/filepath"
	"os"
	"fmt"
)

//SourceFilePath returns corresponding source file path of given path (e.g. abc_test.go -> abc.go)
func SourceFilePath(path string) string {
	if strings.HasSuffix(path, "_test.go") {
		return strings.TrimSuffix(path, "_test.go") + ".go"
	}
	return path
}

//TestToSourceFilePath returns corresponding source file path, given the _test.go file path  (e.g. abc_test.go -> abc.go)
func TestToSourceFilePath(path string) string {
	return strings.TrimSuffix(path, "_test.go") + ".go"
}

//SourceToTestFilePath returns corresponding _test.go path, given source file path (e.g. abc.go -> abc_test.go)
func SourceToTestFilePath(path string) string {
	return strings.TrimSuffix(path, ".go") + "_test.go"
}

//IsSourceFile checks if a file name is a source .go file or not (e.g. abc.go -> true;  abc_test.go -> false; abc.py -> false)
func IsSourceFile(path string) bool {
	return strings.HasSuffix(path, ".go") && !strings.HasSuffix(path, "_test.go")
}

func PackageName(path string) string {
	dir := filepath.Dir(path)
	if dir == "."{
		dir, _ = os.Getwd()
		fmt.Printf("dir .->%s", dir)
	}
	return filepath.Base(dir)
}