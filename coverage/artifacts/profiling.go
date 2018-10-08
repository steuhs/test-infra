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

package artifacts

import (
	"os"
	"os/exec"

	"fmt"
	"github.com/sirupsen/logrus"
	covIo "k8s.io/test-infra/coverage/io"
	"strings"
)

// runProfiling writes coverage profile (&its stdout) by running go test on
// target package
func runProfiling(cmdArgs []string, localArts *LocalArtifacts) error {
	logrus.Info("Starts calc.runProfiling(...)")

	cmd := exec.Command("go", cmdArgs...)
	cmdAsString := fmt.Sprintf("go %s", strings.Join(cmdArgs, " "))
	logrus.Infof("Composed shell command: %s", cmdAsString)

	goTestCoverStdout, err := cmd.Output()

	if err != nil {
		return fmt.Errorf("error running composed shell command: error='%v'; stdout='%s'",
			err, goTestCoverStdout)
	}
	logrus.Infof("Coverage profile created @ '%s'", localArts.ProfilePath())
	err = covIo.CreateMarker(localArts.Directory, CovProfileCompletionMarker)
	if err != nil {
		return err
	}

	stdoutPath := localArts.CovStdoutPath()
	stdoutFile, err := os.Create(stdoutPath)
	if err == nil {
		stdoutFile.Write(goTestCoverStdout)
	} else {
		return fmt.Errorf("error creating stdout file: %v", err)
	}
	logrus.Infof("Stdout of test coverage stored in %s", stdoutPath)
	logrus.Infof("Ends calc.runProfiling(...)")
	return stdoutFile.Close()
}
