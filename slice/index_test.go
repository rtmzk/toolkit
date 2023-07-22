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

func TestIndex(t *testing.T) {
	tests := []struct {
		src  []int
		dst  int
		want int
		name string
	}{
		{
			src:  []int{1, 1, 3, 5},
			dst:  1,
			want: 0,
			name: "first one",
		},
		{
			src:  []int{},
			dst:  1,
			want: -1,
			name: "the length of src is 0",
		},
		{
			dst:  1,
			want: -1,
			name: "src nil",
		},
		{
			src:  []int{1, 4, 6},
			dst:  7,
			want: -1,
			name: "dst not exist",
		},
		{
			src:  []int{1, 3, 4, 2, 0},
			dst:  0,
			want: 4,
			name: "last one",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, Index[int](test.src, test.dst))
		})
	}
}

func TestLastIndex(t *testing.T) {
	tests := []struct {
		src  []int
		dst  int
		want int
		name string
	}{
		{
			src:  []int{1, 1, 3, 5},
			dst:  1,
			want: 1,
			name: "first one",
		},
		{
			src:  []int{},
			dst:  1,
			want: -1,
			name: "the length of src is 0",
		},
		{
			dst:  1,
			want: -1,
			name: "src nil",
		},
		{
			src:  []int{1, 4, 6},
			dst:  7,
			want: -1,
			name: "dst not exist",
		},
		{
			src:  []int{0, 1, 3, 4, 2, 0},
			dst:  0,
			want: 5,
			name: "last one",
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.want, LastIndex[int](test.src, test.dst))
	}
}
