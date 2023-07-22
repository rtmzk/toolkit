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

// UnionSet 并集&去重
// 只支持可比较类型 comparable
func UnionSet[T comparable](src, dst []T) []T {
	srcMap, dstMap := toMap[T](src), toMap[T](dst)

	for key := range srcMap {
		dstMap[key] = struct{}{}
	}

	var ret = make([]T, 0, len(dstMap))
	for key := range dstMap {
		ret = append(ret, key)
	}

	return ret
}
