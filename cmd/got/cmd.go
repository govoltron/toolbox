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
	cmdc = &cobra.Command{}
)

func init() {
	var (
		flags = handler.NewCmdFlags(gflags)
	)
	cmdc.Use = "cmd"
	cmdc.Short = "Manage command"
	cmdc.Long = `Command [cmd] can help us to install/remove the commands which in directory: ./cmd.`
	cmdc.Example = `  vol cmd --i c1
  vol cmd --i -p=c1 c2
  vol cmd --remove c1`
	cmdc.Args = cobra.ExactArgs(1)
	// Events
	cmdc.RunE = func(cmd *cobra.Command, args []string) error {
		return handler.OnCmdHandler(cmd.Context(), flags, args)
	}
	// Flags
	if f := cmdc.Flags(); f != nil {
		f.BoolVarP(&flags.Install, "install", "i", flags.Install, "install a command")
		f.BoolVarP(&flags.Remove, "remove", "", flags.Remove, "remove a command")
		f.StringVarP(&flags.Parent, "parent", "p", flags.Parent, "if there is no parent, the specified command is a first-level command, otherwise \nit is a second-level command")
	}
}
