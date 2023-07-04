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

	"github.com/govoltron/toolbox/pkg/common"
	"github.com/govoltron/toolbox/pkg/dao/template"
)

type license interface {
	ValidateLicense(ctx context.Context, lic string) (err error)
	GenerateHeader(ctx context.Context, gp *Project, lic string, year string, owner string) (err error)
	GenerateLicense(ctx context.Context, gp *Project, lic string, year string, owner string) (err error)
}

type licenseImpl struct {
}

// ValidateLicense implements license
func (*licenseImpl) ValidateLicense(ctx context.Context, lic string) (err error) {
	if _, ok := common.LicenseNames[lic]; !ok {
		err = fmt.Errorf("license '%s' not supported", lic)
	}
	return
}

// GenerateHeader implements license
func (l *licenseImpl) GenerateHeader(ctx context.Context, gp *Project, lic string, year string, owner string) (err error) {
	var (
		tpl       string
		copyright = template.Copyright{Year: year, Owner: owner}
	)
	switch lic {
	case common.LicenseApache2:
		tpl, err = template.ApacheLicense2Header(copyright)
	case common.LicenseGPL3:
		tpl, err = template.GeneralPublicLicense3Header(copyright)
	case common.LicenseMIT:
		tpl, err = template.MITLicenseHeader(copyright)
	default:
		return fmt.Errorf("license '%s' not supported", lic)
	}
	if err != nil {
		return
	}
	return gp.License.Header.Write([]byte(tpl), 0644)
}

// GenerateLicense implements license
func (l *licenseImpl) GenerateLicense(ctx context.Context, gp *Project, lic string, year string, owner string) (err error) {
	var (
		tpl       string
		copyright = template.Copyright{Year: year, Owner: owner}
	)
	switch lic {
	case common.LicenseApache2:
		tpl, err = template.ApacheLicense2(copyright)
	case common.LicenseGPL3:
		tpl, err = template.GeneralPublicLicense3(copyright)
	case common.LicenseMIT:
		tpl, err = template.MITLicense(copyright)
	default:
		return fmt.Errorf("license '%s' not supported", lic)
	}
	if err != nil {
		return
	}
	return gp.License.Description.Write([]byte(tpl), 0644)
}

var (
	License license = &licenseImpl{}
)
