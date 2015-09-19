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

package enumerate_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/dare2try/go-enumerate"
)

var _ = Describe("Where Iterator", func() {
	Describe("When creating a new where iterator", func() {
		var (
			slice    []interface{}
			iterator enumerate.Iterable
		)

		BeforeEach(func() {
			slice = []interface{}{"a", "b", "a"}
			iterator = enumerate.Where(
				enumerate.Slice(slice),
				func(x interface{}) bool { return x == "a" })
		})

		It("should return only items that match the predicate", func() {
			result := make([]interface{}, 0)
			for item := iterator.Next(); item != nil; item = iterator.Next() {
				result = append(result, item)
			}
			Ω(result).Should(Equal([]interface{}{"a", "a"}))
		})
	})
})