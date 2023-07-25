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

package util

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestReadLines(t *testing.T) {
	testCase := []struct {
		name    string
		path    string
		content []string
		want    []string
	}{
		{
			name:    "file has content",
			path:    "/tmp/test_readlines_case1.txt",
			content: []string{"1", "2", "3"},
			want:    []string{"1", "2", "3"},
		},
		{
			name:    "file dose not have content",
			path:    "/tmp/test_readlines_case2.txt",
			content: []string{},
			want:    []string(nil),
		},
	}

	for _, tt := range testCase {
		t.Run(tt.name, func(t *testing.T) {
			f, _ := os.Create(tt.path)

			defer os.Remove(tt.path)

			for _, data := range tt.content {
				_, err := f.WriteString(data + "\n")
				if err != nil {
					fmt.Println(err.Error())
				}
			}

			f.Close()

			readedContents, err := ReadLines(tt.path)
			assert.Nil(t, err)
			assert.Equal(t, tt.want, readedContents)
		})
	}
}

func TestReadLinesOffsetN(t *testing.T) {
	testCase := []struct {
		name    string
		offset  int
		n       int
		path    string
		content []string
		want    []string
	}{
		{
			name:    "read n line after offset",
			offset:  0,
			n:       2,
			path:    "/tmp/test_readlines_offsetn_case1.txt",
			content: []string{"1", "2", "3", "4"},
			want:    []string{"1", "2"},
		},
		{
			name:    "read all line after offset",
			offset:  1,
			n:       -1,
			path:    "/tmp/test_readlines_offsetn_case2.txt",
			content: []string{"1", "2", "3", "4"},
			want:    []string{"2", "3", "4"},
		},
		{
			name:    "read all line",
			offset:  0,
			n:       -1,
			path:    "/tmp/test_readlines_offsetn_case3.txt",
			content: []string{"1", "2", "3", "4"},
			want:    []string{"1", "2", "3", "4"},
		},
		{
			name:    "n out of lines",
			offset:  0,
			n:       1000,
			path:    "/tmp/test_readlines_offsetn_case4.txt",
			content: []string{"1", "2"},
			want:    []string{"1", "2"},
		},
		{
			name:    "offset out of lines",
			offset:  1000,
			n:       -1,
			path:    "/tmp/test_readlines_offsetn_case5.txt",
			content: []string{"1", "2", "3"},
			want:    []string(nil),
		},
	}

	for _, tt := range testCase {
		t.Run(tt.name, func(t *testing.T) {
			f, _ := os.Create(tt.path)

			defer os.Remove(tt.path)

			for _, data := range tt.content {
				_, err := f.WriteString(data + "\n")
				if err != nil {
					fmt.Println(err.Error())
				}
			}

			f.Close()

			readedContents, err := ReadLinesOffsetN(tt.path, tt.offset, tt.n)
			assert.Nil(t, err)
			assert.Equal(t, tt.want, readedContents)
		})
	}
}
