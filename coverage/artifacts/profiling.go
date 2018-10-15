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

	"github.com/sirupsen/logrus"
	covIo "k8s.io/test-infra/coverage/io"
	"k8s.io/test-infra/coverage/logUtil"
	"log"
	"io/ioutil"
)

// runProfiling writes coverage profile (&its stdout) by running go test on
// target package
func runProfiling(covTargets []string, localArts *LocalArtifacts) {
	logrus.Info("\nStarts calc.runProfiling(...)")

	cmdArgs := []string{"test"}

	cmdArgs = append(cmdArgs, covTargets...)
	profilePath := localArts.ProfilePath()
	cmdArgs = append(cmdArgs, []string{"-covermode=count",
		"-coverprofile", profilePath}...)

	logrus.Infof("go cmdArgs=%v\n", cmdArgs)
	cmd := exec.Command("go", cmdArgs...)
	stderr, err := cmd.StderrPipe()

	err = cmd.Start()
	if err != nil {
		log.Fatalf("cmd fails to start: %v", err)
	}
	stdoutText, _ := ioutil.ReadAll(stderr)
	err = cmd.Wait()

	//goTestCoverStdout, errCmdOutput := cmd.Output()

	if err != nil {
		logUtil.LogFatalf("Error running 'go test -coverprofile %s': error='%v'; stdout='%s'",
			profilePath, err, stdoutText)
	} else {
		logrus.Infof("coverage profile created @ '%s'", profilePath)
		covIo.CreateMarker(localArts.Directory(), CovProfileCompletionMarker)
	}

	stdoutPath := localArts.CovStdoutPath()
	stdoutFile, err := os.Create(stdoutPath)
	if err == nil {
		stdoutFile.Write(stdoutText)
	} else {
		logrus.Infof("Error creating stdout file: %v", err)
	}
	defer stdoutFile.Close()
	logrus.Infof("stdout of test coverage stored in %s\n", stdoutPath)
	logrus.Infof("Ends calc.runProfiling(...)\n\n")
	return
}
