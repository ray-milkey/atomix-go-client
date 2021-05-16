// Copyright 2019-present Open Networking Foundation.
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

package indexedmap

import (
	api "github.com/atomix/atomix-api/go/atomix/primitive/indexedmap"
	"github.com/atomix/atomix-go-framework/pkg/atomix/meta"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOptions(t *testing.T) {
	putRequest := &api.PutRequest{}
	assert.Equal(t, uint64(0), putRequest.Entry.Value.ObjectMeta.Revision.Num)
	IfMatch(meta.ObjectMeta{Revision: 1}).beforePut(putRequest)
	assert.Equal(t, uint64(1), putRequest.Entry.Value.ObjectMeta.Revision.Num)

	removeRequest := &api.RemoveRequest{}
	assert.Equal(t, uint64(0), removeRequest.Entry.Value.ObjectMeta.Revision.Num)
	IfMatch(meta.ObjectMeta{Revision: 2}).beforeRemove(removeRequest)
	assert.Equal(t, uint64(2), removeRequest.Entry.Value.ObjectMeta.Revision.Num)

	eventRequest := &api.EventsRequest{}
	assert.False(t, eventRequest.Replay)
	WithReplay().beforeWatch(eventRequest)
	assert.True(t, eventRequest.Replay)
}