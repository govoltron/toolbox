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

	"github.com/govoltron/toolbox/pkg/internal"
	"github.com/govoltron/toolbox/pkg/service"
)

type BuildFlags struct {
	*GlobalFlags
	Install bool
	Output  string
	Mode    string
}

func NewBuildFlags(gflags *GlobalFlags) (flags *BuildFlags) {
	flags = &BuildFlags{GlobalFlags: gflags}
	// Install
	flags.Install = false
	// Mode
	flags.Mode = "binary"
	// Output
	flags.Output = "bin"
	return
}

func OnBuildHandler(ctx context.Context, flags *BuildFlags, args []string) (err error) {
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
		cmd = strings.ToLower(args[0])
		pkg = fmt.Sprintf("%s/cmd/%s", gp.GoMod.Module.Mod.Path, cmd)
		out = ""
		exe = ""
	)

	if flags.Output == "" || flags.Output == "bin" {
		out = fmt.Sprintf("%s/bin", gp.Workspace)
	} else {
		out = flags.Output
	}

	exe = fmt.Sprintf("%s/%s", out, cmd)

	if flags.Install {
		err = internal.ShellOut("go", "install", pkg)
	} else {
		err = internal.ShellOut("go", "build", "-o", exe, pkg)
	}

	switch flags.Mode {
	case "docker":
		// not implemented
	case "binary":
		// Nothing to do
	default:
		break
	}

	return
}
