// Copyright 2023 rtmzk
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

package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnique(t *testing.T) {
	testCase := []struct {
		name string
		src  []int
		want []int
	}{
		{
			name: "slice which has duplicate element",
			src:  []int{1, 2, 3, 1, 1, 1, 2, 4},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "slice which does not have duplicate element",
			src:  []int{1, 2, 3, 4},
			want: []int{1, 2, 3, 4},
		},
	}

	for _, tt := range testCase {
		t.Run(tt.name, func(t *testing.T) {
			ret := Unique[int](tt.src)
			assert.Equal(t, tt.want, ret)
		})
	}
}
