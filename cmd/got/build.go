// Copyright 2023 Kami
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

package main

import (
	"github.com/govoltron/toolbox/pkg/console/got/handler"
	"github.com/spf13/cobra"
)

var (
	buildc = &cobra.Command{}
)

func init() {
	var (
		flags = handler.NewBuildFlags(gflags)
	)
	buildc.Use = "build"
	buildc.Short = "Build command"
	buildc.Long = "Building a command use 'go build', then put it in a directory."
	buildc.Args = cobra.ExactArgs(1)
	// Events
	buildc.RunE = func(cmd *cobra.Command, args []string) error {
		return handler.OnBuildHandler(cmd.Context(), flags, args)
	}
	// Flags
	if f := buildc.Flags(); f != nil {
		f.BoolVarP(&flags.Install, "install", "i", flags.Install, "use 'go install'")
		f.StringVarP(&flags.Mode, "mode", "", flags.Mode, "execution mode")
		f.StringVarP(&flags.Output, "output", "o", flags.Output, "a directory path where the destination file is saved")
	}
}
