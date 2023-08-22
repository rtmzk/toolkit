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

package retry

import (
	"github.com/rtmzk/toolkit/internal/errs"
	"sync/atomic"
	"time"
)

// FixedIntervalRetry 固定时间间隔的重试
type FixedIntervalRetry struct {
	// maxRetry 最大重试次数
	// 可以设置为0或负数表示无限制次数循环
	maxRetries int32

	// interval 重试间隔
	interval time.Duration

	// retries 当前重试次数
	retries int32
}

func NewFixedIntervalRetry(interval time.Duration, maxRetries int32) (*FixedIntervalRetry, error) {
	if interval <= 0 {
		return nil, errs.NewErrInvalidIntervalValue(interval)
	}

	return &FixedIntervalRetry{
		maxRetries: maxRetries,
		interval:   interval,
	}, nil
}

func (retry *FixedIntervalRetry) Next() (time.Duration, bool) {
	retries := atomic.AddInt32(&retry.retries, 1)
	if retry.maxRetries <= 0 || retries <= retry.maxRetries {
		return retry.interval, true
	}
	return 0, false
}
