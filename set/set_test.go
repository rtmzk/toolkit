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

package set

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetAdd(t *testing.T) {
	addval := []int{1, 10, 20, 30, 1}
	s := NewMapSet[int](10)

	t.Run("Add", func(t *testing.T) {
		for _, val := range addval {
			s.Add(val)
		}

		assert.Equal(t, s.m, map[int]struct{}{
			1:  struct{}{},
			10: struct{}{},
			20: struct{}{},
			30: struct{}{},
		})
	})
}

func TestSetDelete(t *testing.T) {
	testCase := []struct {
		name    string
		delVal  int
		setSet  map[int]struct{}
		wantSet map[int]struct{}
		isExist bool
	}{
		{
			name:   "delete value exist",
			delVal: 2,
			setSet: map[int]struct{}{
				2: struct{}{},
			},
			wantSet: map[int]struct{}{},
			isExist: true,
		},
		{
			name:   "delete value not found",
			delVal: 3,
			setSet: map[int]struct{}{
				30: struct{}{},
			},
			wantSet: map[int]struct{}{
				30: struct{}{},
			},
			isExist: false,
		},
	}

	for _, tcase := range testCase {
		t.Run(tcase.name, func(t *testing.T) {
			s := NewMapSet[int](10)
			s.m = tcase.setSet
			s.Delete(tcase.delVal)
			assert.Equal(t, tcase.wantSet, s.m)
		})
	}
}

func TestSetExist(t *testing.T) {
	s := NewMapSet[int](1)
	s.Add(1)

	testCase := []struct {
		name    string
		val     int
		isExist bool
	}{
		{
			name:    "key exist",
			val:     1,
			isExist: true,
		},
		{
			name:    "key not found",
			val:     2,
			isExist: false,
		},
	}

	for _, tcase := range testCase {
		t.Run(tcase.name, func(t *testing.T) {
			ok := s.Exist(tcase.val)
			assert.Equal(t, tcase.isExist, ok)
		})
	}
}

func TestSet_Keys(t *testing.T) {
	testcase := []struct {
		name    string
		setVal  map[int]struct{}
		wantVal map[int]struct{}
	}{
		{
			name: "found value",
			setVal: map[int]struct{}{
				1: struct{}{},
				2: struct{}{},
				3: struct{}{},
			},
			wantVal: map[int]struct{}{
				1: struct{}{},
				2: struct{}{},
				3: struct{}{},
			},
		},
	}

	for _, tcase := range testcase {
		t.Run(tcase.name, func(t *testing.T) {
			s := NewMapSet[int](10)
			s.m = tcase.setVal
			vals := s.Keys()
			ok := equal(vals, tcase.wantVal)
			assert.Equal(t, true, ok)
		})
	}
}

func equal(nums []int, m map[int]struct{}) bool {
	for _, num := range nums {
		_, ok := m[num]
		if !ok {
			return false
		}

		delete(m, num)
	}

	return len(m) == 0
}
