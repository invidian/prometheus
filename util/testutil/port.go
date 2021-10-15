// Copyright 2021 The Prometheus Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package testutil

import (
	"crypto/rand"
	"math/big"
	"testing"
)

// RandomUnprivilegedPort returns valid unprivileged random port number which can be used for testing.
func RandomUnprivilegedPort(t *testing.T) int {
	t.Helper()

	min := 1024
	max := 65535

	i, err := rand.Int(rand.Reader, big.NewInt(int64(max-min+1)))
	if err != nil {
		t.Fatalf("Generating random port: %v", err)
	}

	return int(i.Int64()) + min
}
