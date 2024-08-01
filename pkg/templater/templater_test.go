/*
Copyright © 2022 - 2024 SUSE LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package templater

import (
	"fmt"
	"testing"

	"gotest.tools/v3/assert"
)

func TestFillAndDecode(t *testing.T) {
	data := map[string]interface{}{
		"level1A": map[string]interface{}{
			"level2A": "level2AValue",
			"level2B": map[string]interface{}{
				"level3A": "level3AValue",
			},
		},
		"level1B": "level1BValue",
		"Troublesome": map[string]interface{}{
			"emptyData":   "",
			"noStringVal": 5,
		},
	}

	testCase := []struct {
		Format string
		Output string
		Error  string
	}{
		{
			Format: "${level1B}",
			Output: "level1BValue",
		},
		{
			Format: "${level1B",
			Output: "${level1B",
		},
		{
			Format: "a${level1B",
			Output: "a${level1B",
		},
		{
			Format: "${}",
			Error:  "value not found",
		},
		{
			Format: "${",
			Output: "${",
		},
		{
			Format: "a${",
			Output: "a${",
		},
		{
			Format: "${level1A}",
			Error:  "value not found",
		},
		{
			Format: "a${level1A}c",
			Error:  "value not found",
		},
		{
			Format: "a${level1A}",
			Error:  "value not found",
		},
		{
			Format: "${level1A}c",
			Error:  "value not found",
		},
		{
			Format: "a${level1A/level2A}c",
			Output: "alevel2AValuec",
		},
		{
			Format: "a${level1A/level2B/level3A}c",
			Output: "alevel3AValuec",
		},
		{
			Format: "a${level1A/level2B/level3A}c${level1B}",
			Output: "alevel3AValueclevel1BValue",
		},
		{
			Format: "a${unknown}",
			Error:  "value not found",
		},
		{
			Format: "${Troublesome/emptyData}",
			Output: "",
		},
		{
			Format: "${Troublesome/noStringVal}",
			Error:  "value not found",
		},
	}

	tmpl := NewTemplater()
	tmpl.Fill(data)
	for _, testCase := range testCase {
		t.Run(testCase.Format, func(t *testing.T) {
			str, err := tmpl.Decode(testCase.Format)
			if testCase.Error == "" {
				assert.NilError(t, err)
				assert.Equal(t, testCase.Output, str, "'%s' not equal to '%s'", testCase.Output, str)
			} else {
				assert.Equal(t, testCase.Error, err.Error())
			}
		})
	}
}

func TestIsValueNotFoundError(t *testing.T) {
	errRandom := fmt.Errorf("random error")

	assert.Equal(t, IsValueNotFoundError(errValueNotFound), true)
	assert.Equal(t, IsValueNotFoundError(errRandom), false)
}
