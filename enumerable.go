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

// Defines what it means to be enumerable
// Since not everyone will like to createa tree
// of functions in order to combine projections
// and filtering, etc. the enumerable allows a
// simpler way to use them
type Enumerable interface {
	Iterable
	Select(Projection) Enumerable
	Where(Predicate) Enumerable
}

// **********************************************
// Constructor functions
// **********************************************

// Creates and enumerable over some iterable object
func ItE(iterator Iterable) Enumerable {
	return &iteratorEnumerable{
		iterator: iterator,
	}
}

// **********************************************
// Private stucts (no need to expose them, right now)
// **********************************************

type iteratorEnumerable struct {
	iterator Iterable
}

func (this *iteratorEnumerable) Select(projection Projection) Enumerable {
	return ItE(Select(this.iterator, projection))
}

func (this *iteratorEnumerable) Where(predicate Predicate) Enumerable {
	return ItE(Where(this.iterator, predicate))
}

func (this *iteratorEnumerable) Next() interface{} {
	return this.iterator.Next()
}
