// Copyright 2018 The ksonnet authors
//
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

package cmd

import (
	"os"
	"testing"

	"github.com/spf13/cobra"
)

func runCmd(args ...string) error {
	RootCmd.SetArgs(args)
	return RootCmd.Execute()
}

func withCmd2(id initName, override actionFn, fn func()) {
	if override != nil {
		ogFn := actionFns[id]
		actionFns[id] = override

		defer func() {
			actionFns[id] = ogFn
		}()
	}

	fn()
}

func withCmd(t *testing.T, cmd *cobra.Command, id initName, override interface{}, fn func()) {
	ogAction := actionMap[id]
	actionMap[id] = override

	envConfig := os.Getenv("KUBECONFIG")

	defer func() {
		actionMap[id] = ogAction
		os.Setenv("KUBECONFIG", envConfig)
	}()

	fn()
}