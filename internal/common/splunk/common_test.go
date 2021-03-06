// Copyright 2020, OpenTelemetry Authors
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

package splunk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetValues(t *testing.T) {
	metric := Metric{
		Fields: map[string]interface{}{},
	}
	assert.Equal(t, map[string]interface{}{}, metric.GetValues())
	metric.Fields["metric_name:foo"] = "bar"
	assert.Equal(t, map[string]interface{}{"foo": "bar"}, metric.GetValues())
	metric.Fields["metric_name:foo2"] = "foobar"
	assert.Equal(t, map[string]interface{}{"foo": "bar", "foo2": "foobar"}, metric.GetValues())
}
