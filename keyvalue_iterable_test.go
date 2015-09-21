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

var _ = Describe("What is a key value iterable", func() {
	Describe("Map Iterator", func() {
		Describe("When creating a new map iterator", func() {
			var (
				source   map[interface{}]interface{}
				iterator enumerate.KeyValueIterable
			)

			BeforeEach(func() {
				source = map[interface{}]interface{}{"A": "a"}
				iterator = enumerate.Map(source)
			})

			It("should return nil after enumerating all values", func() {
				var key interface{}
				var value interface{}
				var ok bool

				key, value, ok = iterator.Next()
				Ω(key).ShouldNot(BeNil())
				Ω(value).ShouldNot(BeNil())
				Ω(ok).Should(BeTrue())

				key, value, ok = iterator.Next()
				Ω(key).Should(BeNil())
				Ω(value).Should(BeNil())
				Ω(ok).Should(BeFalse())
			})

			It("Should not be effected by changes to the slice after initialization", func() {
				source["A"] = "b"
				key, value, ok := iterator.Next()
				Ω(key).ShouldNot(BeNil())
				Ω(key).Should(Equal("A"))
				Ω(value).Should(Equal("a"))
				Ω(ok).Should(BeTrue())
			})
		})
	})

	Describe("KvSelect Iterator", func() {
		Describe("When creating a new key/value select iterator", func() {
			var (
				source   map[interface{}]interface{}
				iterator enumerate.KeyValueIterable
			)

			Context("with a contant projection function", func() {
				BeforeEach(func() {
					source = map[interface{}]interface{}{"A": "a", "B": "b", "C": "c"}
					iterator = enumerate.KvSelect(
						enumerate.Map(source),
						func(k interface{}, v interface{}) interface{} { return "a" })
				})

				It("should return the same number of elements as the original source", func() {
					var count int
					for _, _, ok := iterator.Next(); ok; _, _, ok = iterator.Next() {
						count += 1
					}
					Ω(count).Should(Equal(3))
				})

				It("should return the constant value provided in the projection", func() {
					result := make(map[interface{}]interface{}, 0)
					for key, value, ok := iterator.Next(); ok; key, value, ok = iterator.Next() {
						result[key] = value
					}
					Ω(result).Should(Equal(map[interface{}]interface{}{"A": "a", "B": "a", "C": "a"}))
				})
			})
		})
	})

	Describe("KvWhere Iterator", func() {
		Describe("When creating a new key/value where iterator", func() {
			var (
				source   map[interface{}]interface{}
				iterator enumerate.KeyValueIterable
			)

			BeforeEach(func() {
				source = map[interface{}]interface{}{"A": "a", "B": "b", "C": "a"}
				iterator = enumerate.KvWhere(
					enumerate.Map(source),
					func(k interface{}, v interface{}) bool { return v == "a" })
			})

			It("should return only items that match the predicate", func() {
				result := make(map[interface{}]interface{}, 0)
				for key, value, ok := iterator.Next(); ok; key, value, ok = iterator.Next() {
					result[key] = value
				}
				Ω(result).Should(Equal(map[interface{}]interface{}{"A": "a", "C": "a"}))
			})
		})
	})

})
