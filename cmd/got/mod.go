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
	modc = &cobra.Command{}
)

func init() {
	var (
		flags = handler.NewModFlags(gflags)
	)
	modc.Use = "mod"
	modc.Short = "Manage modules"
	modc.Long = `Use for managing the go.mod.`
	// Events
	modc.RunE = func(cmd *cobra.Command, args []string) error {
		return handler.OnModHandler(cmd.Context(), flags, args)
	}
	// Flags
	if f := modc.Flags(); f != nil {
		f.BoolVarP(&flags.Update, "update", "u", flags.Update, "update module")
	}
}
