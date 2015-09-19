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

var _ = Describe("Iterator Enumerable", func() {
	Describe("When creating a new iterator enumerable", func() {
		var (
			slice      []interface{}
			enumerable enumerate.Enumerable
		)

		BeforeEach(func() {
			slice = []interface{}{"a", "b", "a"}
			enumerable = enumerate.ItE(enumerate.Slice(slice))
		})

		It("should return only items that match the predicate when appending a where", func() {
			result := enumerable.Where(func(x interface{}) bool { return x == "a" })

			var count int
			for item := result.Next(); item != nil; item = result.Next() {
				Ω(item).ShouldNot(BeNil())
				Ω(item).Should(Equal("a"))
				count += 1
			}
			Ω(count).Should(Equal(2))
		})

		It("should return the constant in the projection func when appending a select", func() {
			result := enumerable.Select(func(x interface{}) interface{} { return "a" })

			var count int
			for item := result.Next(); item != nil; item = result.Next() {
				Ω(item).ShouldNot(BeNil())
				Ω(item).Should(Equal("a"))
				count += 1
			}
			Ω(count).Should(Equal(3))
		})
	})
})
