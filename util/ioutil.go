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
	"bufio"
	"io"
	"os"
	"strings"
)

// ReadLines 从文件中读取数据并按行分割
// 是ReadLinesOffsetN(filename, 0, -1)的封装方法
func ReadLines(filename string) ([]string, error) {
	return ReadLinesOffsetN(filename, 0, -1)
}

// ReadLinesOffsetN 从文件中读取数据并以`\n`分割
// `offset`是指应该从哪一行开始读
// `n` 是指从offset行开始读多少行(包含`offset`行)
// n >= 0: 读取n行
// n < 0 整个文件
func ReadLinesOffsetN(filename string, offset, n int) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return []string{""}, err
	}
	defer f.Close()

	var ret []string
	r := bufio.NewReader(f)
	for i := 0; i < n+int(offset) || n < 0; i++ {
		line, err := r.ReadString('\n')
		if err != nil {
			if err == io.EOF && len(line) > 0 {
				ret = append(ret, strings.Trim(line, "\n"))
			}
			break
		}
		if i < int(offset) {
			continue
		}
		ret = append(ret, strings.Trim(line, "\n"))
	}

	return ret, nil
}
