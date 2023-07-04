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
	"fmt"
	"os"
	"strings"

	"github.com/govoltron/toolbox/pkg/service"
)

type CmdFlags struct {
	*GlobalFlags
	Install bool
	Remove  bool
	Parent  string
}

func NewCmdFlags(gflags *GlobalFlags) (flags *CmdFlags) {
	flags = &CmdFlags{GlobalFlags: gflags}
	return
}

// OnCmdHandler
func OnCmdHandler(ctx context.Context, flags *CmdFlags, args []string) (err error) {
	if len(args) <= 0 {
		return fmt.Errorf("command not found")
	}

	ws, err := os.Getwd()
	if err != nil {
		return err
	}
	gp, err := service.NewProject(ws)
	if err != nil {
		return err
	}

	var (
		cmd    = strings.ToLower(args[0])
		parent = strings.ToLower(flags.Parent)
	)
	// Operations
	if flags.Install {
		return service.Command.Install(ctx, gp, cmd, parent)
	} else if flags.Remove {
		return service.Command.Remove(ctx, gp, cmd, parent)
	}

	return fmt.Errorf("it seems that there is no operation")
}
