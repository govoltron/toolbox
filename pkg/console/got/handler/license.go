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
	"os/user"
	"strconv"
	"strings"
	"time"

	"github.com/govoltron/toolbox/pkg/service"
)

type LicenseFlags struct {
	*GlobalFlags
	Year   string
	Owner  string
	Header bool
}

func NewLicenseFlags(gflags *GlobalFlags) (flags *LicenseFlags) {
	flags = &LicenseFlags{GlobalFlags: gflags}
	// Year
	flags.Year = strconv.Itoa(time.Now().Year())
	// Owner
	if u, err := user.Current(); err == nil {
		flags.Owner = fmt.Sprintf("%s%s", strings.ToUpper(string(u.Name[0])), strings.ToLower(string(u.Name[1:])))
	}
	// Header
	flags.Header = true
	return
}

func OnLicenseHandler(ctx context.Context, flags *LicenseFlags, args []string) (err error) {
	if len(args) <= 0 {
		return fmt.Errorf("license not found")
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
		lic   = strings.ToLower(args[0])
		year  = flags.Year
		owner = flags.Owner
	)
	if year == "" {
		year = strconv.Itoa(time.Now().Year())
	}
	if len(owner) <= 0 {
		if u, err := user.Current(); err == nil {
			owner = fmt.Sprintf("%s%s", strings.ToUpper(string(u.Name[0])), strings.ToLower(string(u.Name[1:])))
		}
	}

	if err = service.License.ValidateLicense(ctx, lic); err != nil {
		return
	}
	if err = service.License.GenerateLicense(ctx, gp, lic, year, owner); err != nil {
		return
	}
	if flags.Header {
		err = service.License.GenerateHeader(ctx, gp, lic, year, owner)
	}

	return
}
