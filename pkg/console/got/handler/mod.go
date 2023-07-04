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

package handler

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/govoltron/toolbox/pkg/internal"
	"github.com/govoltron/toolbox/pkg/service"
)

type ModFlags struct {
	*GlobalFlags
	Update bool
}

func NewModFlags(gflags *GlobalFlags) (flags *ModFlags) {
	flags = &ModFlags{GlobalFlags: gflags}
	// Update
	flags.Update = false
	return
}

func OnModHandler(ctx context.Context, flags *ModFlags, args []string) (err error) {

	ws, err := os.Getwd()
	if err != nil {
		return err
	}
	gp, err := service.NewProject(ws)
	if err != nil {
		return err
	}

	if flags.Update {
		var (
			module     string
			version    string
			dependents []string
		)
		for _, require := range gp.GoMod.Require {
			if !require.Indirect {
				dependents = append(dependents, require.Mod.Path)
			}
		}
		module = internal.Prompt("Please select a module for update", ">>>", "", dependents...)
		if module == "" {
			return errors.New("invalid module for update")
		}
		version = internal.Prompt("Please enter a version", ">>>", "latest")

		fmt.Println("Executing: go get", module+"@"+version)
		internal.ShellOut("go", "get", module+"@"+version)
	}

	return
}
