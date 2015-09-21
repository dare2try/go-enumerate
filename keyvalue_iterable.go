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

// Converts a map into an iterable object
// Note: A copy of the map will be create in
//       order to ensure that it is immutable
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

type KeyValueProjection func(interface{}, interface{}) interface{}

func KvSelect(iterator KeyValueIterable, projection KeyValueProjection) KeyValueIterable {
	return &kvSelectIterator{
		iterator:   iterator,
		projection: projection,
	}
}

type KeyValuePredicate func(interface{}, interface{}) bool

// Tests each of the items provided by the underlying iterator
// with the boolean predicate and only returns those that are
// evaluated as true
func KvWhere(iterator KeyValueIterable, predicate KeyValuePredicate) KeyValueIterable {
	return &kvWhereIterator{
		iterator:  iterator,
		predicate: predicate,
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

// Allows the consumer to apply a projection function over
// each of the key/value pairs
type kvSelectIterator struct {
	iterator   KeyValueIterable
	projection KeyValueProjection
}

// Applies the projection function over the kev/value pair returned
// by the underlying iterator and returns the result
func (this *kvSelectIterator) Next() (interface{}, interface{}, bool) {
	if key, value, ok := this.iterator.Next(); ok {
		return key, this.projection(key, value), ok
	}
	return nil, nil, false
}

// Allows the consumer to reduce the number of
// results provided by the iterator by first testing
// each of the key/value pairs against the provided predicate
type kvWhereIterator struct {
	iterator  KeyValueIterable
	predicate KeyValuePredicate
}

// Tests the next item provided by the underlying iterator
// and returns it, if the predicate function returns true.
// Othrewise the function will continue to iterator over the
// results until if finds another or has no more items left
// to iterate
func (this *kvWhereIterator) Next() (interface{}, interface{}, bool) {
	for key, value, ok := this.iterator.Next(); ok; key, value, ok = this.iterator.Next() {
		if this.predicate(key, value) {
			return key, value, ok
		}
	}
	return nil, nil, false
}
