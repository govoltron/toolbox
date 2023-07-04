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

package service

import (
	"context"
	"fmt"
	"os"
	"regexp"

	"github.com/govoltron/toolbox/pkg/dao/fsys"
	"github.com/govoltron/toolbox/pkg/dao/template"
)

type command interface {
	Install(ctx context.Context, gp *Project, name, parent string) (err error)
	Remove(ctx context.Context, gp *Project, name, parent string) (err error)
}

type commandImpl struct {
}

// Install implements commandiface
func (c *commandImpl) Install(ctx context.Context, gp *Project, name string, parent string) (err error) {
	if parent != "" {
		if err = c.generateMainCommand(ctx, gp, parent); err != nil {
			return
		}
		if err = c.generateSubCommand(ctx, gp, name, parent); err != nil {
			return
		}
	} else {
		err = c.generateMainCommand(ctx, gp, name)
	}
	return
}

// Remove implements command
func (c *commandImpl) Remove(ctx context.Context, gp *Project, name string, parent string) (err error) {
	if parent != "" {
		err = c.removeSubCommand(ctx, gp, name, parent)
	} else {
		err = c.removeMainCommand(ctx, gp, name)
	}
	return
}

func (c *commandImpl) generateMainCommand(ctx context.Context, gp *Project, name string) (err error) {
	var (
		cf = fmt.Sprintf("%s/cmd/%s/%s.go", gp.Workspace, name, name)
		hf = fmt.Sprintf("%s/pkg/console/%s/handler/%s.go", gp.Workspace, name, name)
	)
	hb, err := fsys.IsFile(hf)
	if err != nil {
		return err
	}
	if !hb {
		mainhandler, err := template.ConsoleMainHandler(name)
		if err != nil {
			return err
		}
		if err = Coding.Write(hf, mainhandler, string(gp.License.Header.TryRead())); err != nil {
			return err
		}
	}
	cb, err := fsys.IsFile(cf)
	if err != nil {
		return err
	}
	if !cb {
		maincommand, err := template.MainCmd(template.CmdVars{Module: gp.GoMod.Module.Mod.Path, Name: name})
		if err != nil {
			return err
		}
		if err = Coding.Write(cf, maincommand, string(gp.License.Header.TryRead())); err != nil {
			return err
		}
	}
	return
}

func (c *commandImpl) generateSubCommand(ctx context.Context, gp *Project, name, parent string) (err error) {
	var (
		cf = fmt.Sprintf("%s/cmd/%s/%s.go", gp.Workspace, parent, name)
		hf = fmt.Sprintf("%s/pkg/console/%s/handler/%s.go", gp.Workspace, parent, name)
	)
	hb, err := fsys.IsFile(hf)
	if err != nil {
		return err
	}
	if !hb {
		subhandler, err := template.ConsoleSubHandler(name)
		if err != nil {
			return err
		}
		if err = Coding.Write(hf, subhandler, string(gp.License.Header.TryRead())); err != nil {
			return err
		}
	}
	cb, err := fsys.IsFile(cf)
	if err != nil {
		return err
	}
	if !cb {
		subcommand, err := template.SubCmd(template.CmdVars{Module: gp.GoMod.Module.Mod.Path, Parent: parent, Name: name})
		if err != nil {
			return err
		}
		if err = Coding.Write(cf, subcommand, string(gp.License.Header.TryRead())); err != nil {
			return err
		}
		if err = c.setupSubCommand(ctx, gp, name, parent); err != nil {
			return err
		}
	}
	return
}

func (c *commandImpl) setupSubCommand(ctx context.Context, gp *Project, name, parent string) (err error) {
	var (
		cf = fmt.Sprintf("%s/cmd/%s/%s.go", gp.Workspace, parent, parent)
	)
	buf, err := os.ReadFile(cf)
	if err != nil {
		return err
	}

	src := regexp.MustCompile(`[\t ]+// sub command placeholder\n`).
		ReplaceAll(buf, []byte(fmt.Sprintf("    cmds = append(cmds, %sc)\n    // sub command placeholder\n", name)))

	return os.WriteFile(cf, []byte(src), 0644)
}

func (c *commandImpl) deleteSubCommand(ctx context.Context, gp *Project, name, parent string) (err error) {
	var (
		cf = fmt.Sprintf("%s/cmd/%s/%s.go", gp.Workspace, parent, parent)
	)
	buf, err := os.ReadFile(cf)
	if err != nil {
		return err
	}

	src := regexp.MustCompile(fmt.Sprintf(`[\t ]+cmds = append\(cmds, %sc\)\n`, name)).ReplaceAll(buf, nil)

	return os.WriteFile(cf, []byte(src), 0644)
}

func (c *commandImpl) removeMainCommand(ctx context.Context, gp *Project, name string) (err error) {
	var (
		cd = fmt.Sprintf("%s/cmd/%s", gp.Workspace, name)
		hd = fmt.Sprintf("%s/pkg/console/%s", gp.Workspace, name)
	)
	hb, err := fsys.IsDir(hd)
	if err != nil {
		return err
	}
	if hb {
		if err = os.RemoveAll(hd); err != nil {
			return
		}
	}
	cb, err := fsys.IsDir(cd)
	if err != nil {
		return err
	}
	if cb {
		if err = os.RemoveAll(cd); err != nil {
			return
		}
	}
	return
}

func (c *commandImpl) removeSubCommand(ctx context.Context, gp *Project, name, parent string) (err error) {
	var (
		cf = fmt.Sprintf("%s/cmd/%s/%s.go", gp.Workspace, parent, name)
		hf = fmt.Sprintf("%s/pkg/console/%s/handler/%s.go", gp.Workspace, parent, name)
	)
	hb, err := fsys.IsFile(hf)
	if err != nil {
		return err
	}
	if hb {
		if err = os.Remove(hf); err != nil {
			return
		}
	}
	cb, err := fsys.IsFile(cf)
	if err != nil {
		return err
	}
	if cb {
		if err = os.Remove(cf); err != nil {
			return
		}
		if err = c.deleteSubCommand(ctx, gp, name, parent); err != nil {
			return
		}
	}
	return
}

var (
	Command command = &commandImpl{}
)
