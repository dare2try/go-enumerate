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

var _ = Describe("Select Iterator", func() {
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
