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
	"bytes"
	"text/template"
)

type Copyright struct {
	Year  string
	Owner string
}

func generateLicense(in string, copyrights ...Copyright) (tpl string, err error) {
	var (
		b bytes.Buffer
		t *template.Template
	)
	t, err = template.New("license").Parse(in)
	if err != nil {
		return
	}
	if err = t.Execute(&b, copyrights); err != nil {
		return
	}
	return b.String(), nil
}
