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
	Version = "1.0.0"
)

var (
	gotc   = &cobra.Command{}
	gflags = handler.NewGlobalFlags()
)

func init() {
	// var (
	// 	flags = handler.NewGotFlags(gflags)
	// )
	gotc.Use = "got"
	gotc.Short = "A tool use for managing Go project which use the thecxx/go-std-layout directory structure"
	gotc.Long = `A tool use for managing Go project which use the thecxx/go-std-layout
directory structure.`
	gotc.Version = Version
	gotc.SilenceUsage = true
	gotc.CompletionOptions.HiddenDefaultCmd = true
	// Events
	gotc.RunE = func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
		// return handler.OnGotHandler(cmd.Context(), flags, args)
	}
	// Flags
	// if f := gotc.Flags(); f != nil {
	//     f.StringVarP(&flags.Test, "test", "t", flags.Test, "a test flag")
	// }
	if pf := gotc.PersistentFlags(); pf != nil {
		pf.BoolVarP(&gflags.Quiet, "quiet", "q", gflags.Quiet, "quiet mode")
	}
}

func main() {
	var (
		cmds []*cobra.Command
	)

	// Register sub commands
	cmds = append(cmds, cmdc)
	cmds = append(cmds, buildc)
	cmds = append(cmds, licensec)
	cmds = append(cmds, modc)
	cmds = append(cmds, versionc)
	// sub command placeholder

	gotc.AddCommand(cmds...)
	defer func() {
		gotc.RemoveCommand(cmds...)
	}()

	gotc.Execute()
}
