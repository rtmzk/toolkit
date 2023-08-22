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

package errs

import (
	"fmt"
	"time"
)

func NewErrIndexOutOfRange(length int, index int) error {
	return fmt.Errorf("toolkit: 下标超出范围, 长度: %d, 下标: %d", length, index)
}

func NewErrInvalidType(want, got string) error {
	return fmt.Errorf("toolkit: 类型转换失败, want: %s, got: %s", want, got)
}

func NewErrInvalidIntervalValue(interval time.Duration) error {
	return fmt.Errorf("toolkit: 无效的间隔时间 %d, 预期值应大于 0", interval)
}

func NewErrInvalidMaxIntervalValue(maxInterval, initialInterval time.Duration) error {
	return fmt.Errorf("toolkit: 最大重试间隔的时间 [%d] 应大于等于初始重试的间隔时间 [%d] ", maxInterval, initialInterval)
}
