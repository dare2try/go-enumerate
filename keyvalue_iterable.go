// Copyright 2015 Carlos Perez

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package enumerate

type KeyValueIterable interface {
	Next() (interface{}, interface{}, bool)
}

// **********************************************
// Constructor functions
// **********************************************

func Map(source map[interface{}]interface{}) KeyValueIterable {
	length := len(source)
	keys := make([]interface{}, 0, length)
	copyOfSource := make(map[interface{}]interface{}, length)

	for k, v := range source {
		keys = append(keys, k)
		copyOfSource[k] = v
	}

	// We create the slice iterator directly to avoid an unneccessary
	// copy from being created
	return &mapIterator{
		position: 0,
		length:   length,
		keys: &sliceIterator{
			position: 0,
			length:   length,
			source:   keys,
		},
		source: copyOfSource,
	}
}

// **********************************************
// Private stucts (no need to expose them, right now)
// **********************************************

// Allows the consumer to iterate over a map
type mapIterator struct {
	position int
	length   int
	keys     Iterable
	source   map[interface{}]interface{}
}

// Returns the next key/value pair and true if one exists, otherwise
// nil, nil, false is returned
func (this *mapIterator) Next() (interface{}, interface{}, bool) {
	if key, ok := this.keys.Next(); ok {
		value, _ := this.source[key]
		return key, value, ok
	}
	return nil, nil, false
}
