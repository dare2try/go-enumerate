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

import (
	"fmt"
	"log"
)

// Defines what it means to be iterable
type Iterable interface {
	Next() (interface{}, bool)
}

// **********************************************
// Constructor functions
// **********************************************

// Converts a slice into an iterable object
// Note: A copy of the slice will be create in
//       order to ensure that it is immutable
func Slice(slice []interface{}) Iterable {
	c := make([]interface{}, len(slice))
	copy(c, slice)

	return &sliceIterator{
		position: 0,
		length:   len(slice),
		slice:    c,
	}
}

type Projection func(interface{}) interface{}

// Applies a projection to each of the underlying iterators
// results
func Select(iterator Iterable, projection Projection) Iterable {
	return &selectIterator{
		iterator:   iterator,
		projection: projection,
	}
}

type Predicate func(interface{}) bool

// Tests each of the items provided by the underlying iterator
// with the boolean predicate and only returns those that are
// evaluated as true
func Where(iterator Iterable, predicate Predicate) Iterable {
	return &whereIterator{
		iterator:  iterator,
		predicate: predicate,
	}
}

// **********************************************
// Private stucts (no need to expose them, right now)
// **********************************************

// Allows the consumer to interate of a slice in a consistent
// fassion
type sliceIterator struct {
	position int
	length   int
	slice    []interface{}
}

// Returns the next item in the slice and a boolean result to
// identify if the result returned is the last item in the
// underlying collection. If there are no more items in the
// underlying collection, nil and flase is returned (respectively)
func (this *sliceIterator) Next() (interface{}, bool) {
	log.Println(fmt.Sprintf("Position: %d, Length: %d", this.position, this.length))
	if this.position < this.length {
		// log.Println(fmt.Sprintf("Position: %i, Length: %i", this.position, this.length))
		next := this.slice[this.position]
		this.position += 1
		return next, true
	}
	return nil, false
}

// Allows the consumer to apply a projection function to
// each of the items returned
type selectIterator struct {
	iterator   Iterable
	projection Projection
}

// Applies the projection function to each of the items
// returned by the underlying iterator
func (this *selectIterator) Next() (interface{}, bool) {
	if item, ok := this.iterator.Next(); ok {
		return this.projection(item), ok
	}
	return nil, false
}

// Allows the consumer to reduce the number of
// results provided by the iterator by first testing
// each of the items against the provided predicate
type whereIterator struct {
	iterator  Iterable
	predicate Predicate
}

// Tests the next item provided by the underlying iterator
// and returns it, if the predicate function returns true.
// Othrewise the function will continue to iterator over the
// results until if finds another or has no more items left
// to iterate
func (this *whereIterator) Next() (interface{}, bool) {
	for item, ok := this.iterator.Next(); ok; item, ok = this.iterator.Next() {
		if this.predicate(item) {
			return item, ok
		}
	}
	return nil, false
}
