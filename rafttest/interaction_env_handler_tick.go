// Copyright 2019 The etcd Authors
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

package rafttest

import (
	"strconv"
	"testing"

	"github.com/cockroachdb/datadriven"
)

func (env *InteractionEnv) handleTickElection(t *testing.T, d datadriven.TestData) error {
	idx := firstAsNodeIdx(t, d)
	return env.Tick(idx, env.Nodes[idx].Config.ElectionTick)
}

func (env *InteractionEnv) handleTickHeartbeat(t *testing.T, d datadriven.TestData) error {
	idx := firstAsNodeIdx(t, d)
	return env.Tick(idx, env.Nodes[idx].Config.HeartbeatTick)
}

func (env *InteractionEnv) handleTick(t *testing.T, d datadriven.TestData) error {
	idx := firstAsNodeIdx(t, d)

	if len(d.CmdArgs) != 2 || len(d.CmdArgs[1].Vals) > 0 {
		t.Fatalf("expected exactly one key with no vals: %+v", d.CmdArgs[1:])
	}

	n, err := strconv.Atoi(d.CmdArgs[1].Key)
	if err != nil {
		t.Fatal(err)
	}

	return env.Tick(idx, n)
}

// Tick the node at the given index the given number of times.
func (env *InteractionEnv) Tick(idx int, num int) error {
	for i := 0; i < num; i++ {
		env.Nodes[idx].Tick()
	}
	return nil
}
