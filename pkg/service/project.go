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
	"fmt"
	"path"
	"strings"

	"github.com/govoltron/toolbox/pkg/dao/fsys"
	"github.com/govoltron/toolbox/pkg/internal"
	"golang.org/x/mod/modfile"
)

type Project struct {
	GoMod     *modfile.File
	Curplace  string
	Workspace string
	License   struct {
		Header      fsys.File
		Description fsys.File
	}
}

func NewProject(ws string) (gp *Project, err error) {
	gp = &Project{
		Workspace: ws,
	}
	if err = gp.initModule(); err != nil {
		return nil, err
	}
	if err = gp.initGoMod(); err != nil {
		return nil, err
	}
	if err = gp.initLicense(); err != nil {
		return nil, err
	}
	return
}

func (gp *Project) initModule() (err error) {
	gomod := internal.GoEnv("GOMOD")
	if path.Base(gomod) != "go.mod" {
		return fmt.Errorf("module not found, please initialize it first use 'go mod init'")
	}
	var (
		realws = path.Dir(gomod)
	)
	if realws != gp.Workspace {
		if !strings.HasPrefix(gp.Workspace, realws) {
			return fmt.Errorf("please re-execute it at the root of the project")
		}
		gp.Curplace, gp.Workspace = gp.Workspace, realws
	}
	return
}

func (gp *Project) initLicense() (err error) {
	gp.License.Header = fsys.File(gp.Workspace + "/HEADER")
	gp.License.Description = fsys.File(gp.Workspace + "/LICENSE")
	return
}

func (gp *Project) initGoMod() (err error) {
	gp.GoMod, err = modfile.Parse("go.mod", fsys.File(gp.Workspace+"/go.mod").TryRead(), nil)
	return
}
