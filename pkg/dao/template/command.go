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

package template

import (
	"strings"
)

type CmdVars struct {
	Module string
	Parent string
	Name   string
}

func (c CmdVars) LowerCmd() string {
	return strings.ToLower(c.Name)
}

func (c CmdVars) UpperCmd() string {
	return strings.ToUpper(string(c.Name[0])) + strings.ToLower(c.Name[1:])
}

func (c CmdVars) LowerParent() string {
	return strings.ToLower(c.Parent)
}

func (c CmdVars) UpperParent() string {
	return strings.ToUpper(string(c.Parent[0])) + strings.ToLower(c.Parent[1:])
}

func MainCmd(vars CmdVars) (tpl string, err error) {
	tpl = `package main

import (
    "github.com/spf13/cobra"
    "{{.Module}}/pkg/console/{{.LowerCmd}}/handler"
)

var (
    Version = "1.0.0"
)

var (
    {{.LowerCmd}}c = &cobra.Command{}
    gflags = handler.NewGlobalFlags()
)

func init() {
    // var (
    //    flags = handler.New{{.UpperCmd}}Flags(gflags)
    // )
    {{.LowerCmd}}c.Use = "{{.LowerCmd}}"
    {{.LowerCmd}}c.Short = "A short description"
    {{.LowerCmd}}c.Long = "A long description"
    {{.LowerCmd}}c.Version = Version
    {{.LowerCmd}}c.SilenceUsage = true
    {{.LowerCmd}}c.CompletionOptions.HiddenDefaultCmd = true
    // Events
    {{.LowerCmd}}c.RunE = func(cmd *cobra.Command, args []string) error {
        return cmd.Help()
        // return handler.On{{.UpperCmd}}Handler(cmd.Context(), flags, args)
    }
    // Flags
    // if f := {{.LowerCmd}}c.Flags(); f != nil {
    //     f.StringVarP(&flags.Test, "test", "t", flags.Test, "a test flag")
    // }
    // if pf := {{.LowerCmd}}c.PersistentFlags(); pf != nil {
    //     pf.StringVarP(&gflags.Test, "test", "t", gflags.Test, "a test flag")
    // }
}

func main() {
    var (
        cmds []*cobra.Command
    )

    // Register sub commands
    // sub command placeholder

    {{.LowerCmd}}c.AddCommand(cmds...)
    defer func() {
        {{.LowerCmd}}c.RemoveCommand(cmds...)
    }()

    {{.LowerCmd}}c.Execute()
}`

	return generate(tpl, vars)
}

func SubCmd(vars CmdVars) (tpl string, err error) {
	tpl = `package main

import (
    "github.com/spf13/cobra"
    "{{.Module}}/pkg/console/{{.LowerParent}}/handler"
)

var (
    {{.LowerCmd}}c = &cobra.Command{}
)

func init() {
    var (
        flags = handler.New{{.UpperCmd}}Flags(gflags)
    )
    {{.LowerCmd}}c.Use = "{{.LowerCmd}}"
    {{.LowerCmd}}c.Short = "A short description"
    {{.LowerCmd}}c.Long = "A long description"
    // Events
    {{.LowerCmd}}c.RunE = func(cmd *cobra.Command, args []string) error {
        return handler.On{{.UpperCmd}}Handler(cmd.Context(), flags, args)
    }
    // Flags
    // if f := {{.LowerCmd}}c.Flags(); f != nil {
    //     f.StringVarP(&flags.Test, "test", "t", flags.Test, "a test flag")
    // }
}`

	return generate(tpl, vars)
}

type CmdName string

func (c CmdName) LowerCmd() string {
	return strings.ToLower(string(c))
}

func (c CmdName) UpperCmd() string {
	return strings.ToUpper(string(c[0])) + strings.ToLower(string(c[1:]))
}

func ConsoleMainHandler(name string) (tpl string, err error) {
	tpl = `package handler

import (
    "context"
)

type GlobalFlags struct {
    // Test string
}

func NewGlobalFlags() (gflags *GlobalFlags) {
    gflags = &GlobalFlags{}
    return
}

type {{.UpperCmd}}Flags struct {
    *GlobalFlags
    // Test string
}

func New{{.UpperCmd}}Flags(gflags *GlobalFlags) (flags *{{.UpperCmd}}Flags) {
    flags = &{{.UpperCmd}}Flags{GlobalFlags: gflags}
    return
}

func On{{.UpperCmd}}Handler(ctx context.Context, flags *{{.UpperCmd}}Flags, args []string) (err error) {
    return
}`

	return generate(tpl, CmdName(name))
}

func ConsoleSubHandler(name string) (tpl string, err error) {
	tpl = `package handler

import (
	"context"
)

type {{.UpperCmd}}Flags struct {
	*GlobalFlags
	// Test string
}

func New{{.UpperCmd}}Flags(gflags *GlobalFlags) (flags *{{.UpperCmd}}Flags) {
	flags = &{{.UpperCmd}}Flags{GlobalFlags: gflags}
	return
}

func On{{.UpperCmd}}Handler(ctx context.Context, flags *{{.UpperCmd}}Flags, args []string) (err error) {
	return
}`

	return generate(tpl, CmdName(name))
}
