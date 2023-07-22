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

// Contains 判断src中是否存在dst
func Contains[T comparable](src []T, dst T) bool {
	return contains(src, dst,
		func(src, dst T) bool {
			return src == dst
		})
}

func contains[T comparable](src []T, dst T, equal equalFunc[T]) bool {
	for _, v := range src {
		if equal(v, dst) {
			return true
		}
	}
	return false
}

// ContainsAny 判断src中是否包含dst的任意一个元素
func ContainsAny[T comparable](src, dst []T) bool {
	srcMap := toMap(src)

	for _, v := range dst {
		if _, ok := srcMap[v]; ok {
			return true
		}
	}

	return false
}

// ContainsAll 判断src是否包含dst的所有元素
func ContainsAll[T comparable](src, dst []T) bool {
	srcMap := toMap(src)

	for _, v := range dst {
		if _, ok := srcMap[v]; !ok {
			return false
		}
	}

	return true
}
