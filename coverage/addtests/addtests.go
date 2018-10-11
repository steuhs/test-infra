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
	"strings"
	"k8s.io/test-infra/coverage/io"
)

//hasTest checks if a '.go' file has its '_test.go' correspondence
func hasTest(path string) {

}

func addTest(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	fileName := info.Name()
	//strings.HasSuffix(path, )
	if !info.IsDir() && strings.HasSuffix(fileName, ".go") {
		exists, ioErr := io.FileOrDirExistsNew()
		if err != nil {
			return ioErr
		}
		if !exists {

		}

	}
}
