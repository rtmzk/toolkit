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

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

type Server struct {
	port   uint
	scheme string
}

func WithPort(port uint) Option[Server] {
	return func(s *Server) {
		s.port = port
	}
}

func WithPortErr(port uint) OptionErr[Server] {
	return func(s *Server) error {
		if port < 1024 {
			return errors.New("建议端口使用1024以上的端口")
		}

		s.port = port
		return nil
	}
}

func WithScheme(scheme string) Option[Server] {
	return func(s *Server) {
		s.scheme = scheme
	}
}

func WithSchemeErr(scheme string) OptionErr[Server] {
	return func(s *Server) error {
		if scheme == "http" || scheme == "https" {
			s.scheme = scheme
			return nil
		}

		return errors.New("未知协议")
	}
}

func TestApply(t *testing.T) {
	s := &Server{}
	Apply[Server](s, WithPort(80), WithScheme("http"))
	assert.Equal(t, s, &Server{port: 80, scheme: "http"})
}

func TestApplyErr(t *testing.T) {
	s := &Server{}
	err := ApplyErr[Server](s, WithPortErr(1025), WithSchemeErr("https"))
	require.NoError(t, err)
	assert.Equal(t, s, &Server{port: 1025, scheme: "https"})
}
