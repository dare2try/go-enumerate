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

var _ = Describe("What is an iterable", func() {
	Describe("Slice Iterator", func() {
		Describe("When creating a new slice iterator", func() {
			var (
				source   []interface{}
				iterator enumerate.Iterable
			)

			BeforeEach(func() {
				source = []interface{}{"a"}
				iterator = enumerate.Slice(source)
			})

			It("should return nil after enumerating all values", func() {
				var item interface{}
				var ok bool

				item, ok = iterator.Next()
				Ω(item).ShouldNot(BeNil())
				Ω(ok).Should(BeTrue())

				item, ok = iterator.Next()
				Ω(item).Should(BeNil())
				Ω(ok).Should(BeFalse())
			})

			It("Should not be effected by changes to the slice after initialization", func() {
				source[0] = "b"
				item, ok := iterator.Next()
				Ω(item).ShouldNot(BeNil())
				Ω(item).Should(Equal("a"))
				Ω(ok).Should(BeTrue())
			})
		})
	})

	Describe("Select Iterator", func() {
		Describe("When creating a new select iterator", func() {
			var (
				source   []interface{}
				iterator enumerate.Iterable
			)

			Context("with a contant projection function", func() {
				BeforeEach(func() {
					source = []interface{}{"a", "b", "c"}
					iterator = enumerate.Select(
						enumerate.Slice(source),
						func(x interface{}) interface{} { return "a" })
				})

				It("should return the same number of elements as the original source", func() {
					var count int
					for _, ok := iterator.Next(); ok; _, ok = iterator.Next() {
						count += 1
					}
					Ω(count).Should(Equal(3))
				})

				It("should return the constant value provided in the projection", func() {
					result := make([]interface{}, 0)
					for item, ok := iterator.Next(); ok; item, ok = iterator.Next() {
						result = append(result, item)
					}
					Ω(result).Should(Equal([]interface{}{"a", "a", "a"}))
				})
			})
		})
	})

	Describe("Where Iterator", func() {
		Describe("When creating a new where iterator", func() {
			var (
				source   []interface{}
				iterator enumerate.Iterable
			)

			BeforeEach(func() {
				source = []interface{}{"a", "b", "a"}
				iterator = enumerate.Where(
					enumerate.Slice(source),
					func(x interface{}) bool { return x == "a" })
			})

			It("should return only items that match the predicate", func() {
				result := make([]interface{}, 0)
				for item, ok := iterator.Next(); ok; item, ok = iterator.Next() {
					result = append(result, item)
				}
				Ω(result).Should(Equal([]interface{}{"a", "a"}))
			})
		})
	})

})
