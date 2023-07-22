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

func Index[T comparable](src []T, dst T) int {
	return index[T](src, dst, func(src, dst T) bool { return src == dst })
}

func index[T comparable](src []T, dst T, equal equalFunc[T]) int {
	for k, v := range src {
		if equal(v, dst) {
			return k
		}
	}

	return -1
}

func LastIndex[T comparable](src []T, dst T) int {
	return last[T](src, dst, func(src, dst T) bool { return src == dst })
}

func last[T comparable](src []T, dst T, equal equalFunc[T]) int {
	for i := len(src) - 1; i >= 0; i-- {
		if equal(dst, src[i]) {
			return i
		}
	}

	return -1
}
