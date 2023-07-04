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
	"os"
	"path"
	"strings"

	"github.com/govoltron/toolbox/pkg/dao/fsys"
)

type coding interface {
	Write(file, source, license string) (err error)
	InsertHeader(source, license string) (out string)
	Comment(source string) (out string)
}

type codingImpl struct {
}

// Write implements coding
func (c *codingImpl) Write(file, source, header string) (err error) {
	var (
		ok  = false
		dir = path.Dir(file)
	)
	if ok, err = fsys.IsDir(dir); err != nil {
		return err
	}
	if !ok {
		if err = os.MkdirAll(dir, 0755); err != nil {
			return
		}
	}
	if header != "" {
		source = c.InsertHeader(source, header)
	}
	return os.WriteFile(file, []byte(source), 0644)
}

func (c *codingImpl) InsertHeader(source, license string) (out string) {
	if license = strings.Trim(license, "\n\r\t "); license != "" {
		out += c.Comment(license)
		out += "\n\n"
	}
	out += source
	return
}

func (c *codingImpl) Comment(source string) (out string) {
	var (
		i     = 0
		lines = strings.Split(source, "\n")
	)
	for ; i < len(lines)-1; i++ {
		out += fmt.Sprintf("// %s\n", lines[i])
	}
	out += fmt.Sprintf("// %s", lines[i])
	return
}

var (
	Coding coding = &codingImpl{}
)
