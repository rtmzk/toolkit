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
	"math"
	"sync/atomic"
	"time"
)

// ExponentialBackoffRetry 指数退避重试
type ExponentialBackoffRetry struct {
	// initialInterval 初始化重试时间
	initialInterval time.Duration

	// maxInterval 最大重试时间
	maxInterval time.Duration

	// maxRetries 最大重试时间
	maxRetries int32

	// retries 当前重试次数
	retries int32

	// maxIntervalReached 是否已经达到最大重试间隔
	maxIntervalReached atomic.Value
}

func NewExponentialBackoffRetry(initialInterval, maxInterval time.Duration, maxRetries int32) (*ExponentialBackoffRetry, error) {
	if initialInterval <= 0 {
		return nil, errs.NewErrInvalidIntervalValue(initialInterval)
	} else if initialInterval > maxInterval {
		return nil, errs.NewErrInvalidMaxIntervalValue(maxInterval, initialInterval)
	}
	return &ExponentialBackoffRetry{
		initialInterval: initialInterval,
		maxInterval:     maxInterval,
		maxRetries:      maxRetries,
	}, nil
}

func (retry *ExponentialBackoffRetry) Next() (time.Duration, bool) {
	retries := atomic.AddInt32(&retry.retries, 1)

	if retry.maxRetries <= 0 || retries <= retry.maxRetries {
		if reached, ok := retry.maxIntervalReached.Load().(bool); ok && reached {
			return retry.maxInterval, true
		}
		interval := retry.initialInterval * time.Duration(math.Pow(2, float64(retries-1)))

		if interval <= 0 || interval > retry.maxInterval {
			retry.maxIntervalReached.Store(true)
			return retry.maxInterval, true
		}

		return interval, true
	}

	return 0, false
}
