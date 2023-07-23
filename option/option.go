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

package option

// Option 用于选项模式的泛型设计
// 通常情况下 `T` 应当是结构体
type Option[T any] func(t *T)

// Apply 将 `opts` 应用在 `t` 上
func Apply[T any](t *T, opts ...Option[T]) {
	for _, o := range opts {
		o(t)
	}
}

// OptionErr 一个返回`error`的 Option
type OptionErr[T any] func(t *T) error

// ApplyErr 将 `opts` 应用在 `t` 上
// 类似 Apply ,不过应用过程中遇到任何错误都会直接中断返回
func ApplyErr[T any](t *T, opts ...OptionErr[T]) error {
	for _, o := range opts {
		if err := o(t); err != nil {
			return err
		}
	}
	return nil
}
