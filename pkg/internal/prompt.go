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

package internal

import (
	"fmt"

	prompt "github.com/c-bata/go-prompt"
)

func Prompt(notice, prefix, defval string, options ...string) (text string) {
	var (
		suggestions []prompt.Suggest
	)
	for i := 0; i < len(options); i++ {
		suggestions = append(suggestions, prompt.Suggest{Text: options[i]})
	}

	if defval != "" {
		notice = fmt.Sprintf("%s (default %s)", notice, defval)
	}

	fmt.Println(notice)

	text = prompt.Input(prefix+" ", func(d prompt.Document) []prompt.Suggest {
		return prompt.FilterHasPrefix(suggestions, d.GetWordBeforeCursor(), true)
	})
	if text == "" {
		text = defval
	}
	return
}

func PromptWithDescription(notice, prefix, defval string, options ...string) (text string) {
	var (
		suggestions []prompt.Suggest
	)
	for i := 0; i < len(options); i += 2 {
		if i+1 < len(options) {
			suggestions = append(suggestions, prompt.Suggest{Text: options[i], Description: options[i+1]})
		} else {
			break
		}
	}

	if defval != "" {
		notice = fmt.Sprintf("%s (default %s)", notice, defval)
	}

	fmt.Println(notice)

	text = prompt.Input(prefix+" ", func(d prompt.Document) []prompt.Suggest {
		return prompt.FilterHasPrefix(suggestions, d.GetWordBeforeCursor(), true)
	})
	if text == "" {
		text = defval
	}
	return
}
