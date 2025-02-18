/*
 * Copyright 2021 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Code generated by mockery v2.9.4. DO NOT EDIT.

package limiter

import mock "github.com/stretchr/testify/mock"

// ConcurrencyLimiter is an autogenerated mock type for the ConcurrencyLimiter type
type ConcurrencyLimiter struct {
	mock.Mock
}

// Acquire provides a mock function with given fields:
func (_m *ConcurrencyLimiter) Acquire() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Release provides a mock function with given fields:
func (_m *ConcurrencyLimiter) Release() {
	_m.Called()
}

// Status provides a mock function with given fields:
func (_m *ConcurrencyLimiter) Status() (int, int) {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 int
	if rf, ok := ret.Get(1).(func() int); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(int)
	}

	return r0, r1
}
