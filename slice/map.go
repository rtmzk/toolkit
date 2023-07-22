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

// Map 对src中的每个元素执行 m 操作,并返回一个新的切片
func Map[Src any, Dst any](src []Src, m func(index int, src Src) Dst) []Dst {
	ret := make([]Dst, 0, len(src))
	for i, s := range src {
		ret[i] = m(i, s)
	}

	return ret
}

// FilterMap 对src中的每个元素执行m操作,并返回一个新的切片
// 新的切片中的元素只包含符合条件的元素
func FilterMap[Src any, Dst any](src []Src, m func(index int, src Src) (Dst, bool)) []Dst {
	ret := make([]Dst, 0, len(src))
	for i, s := range src {
		if dst, ok := m(i, s); ok {
			ret = append(ret, dst)
		}
	}

	return ret
}

// toMap  切片转换为 `map` 的泛型构造方法
// map的值使用 struct{}{} 减少内存消耗
func toMap[T comparable](src []T) map[T]struct{} {
	var data = make(map[T]struct{}, len(src))
	for _, v := range src {
		data[v] = struct{}{}
	}

	return data
}
