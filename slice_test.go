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

var _ = Describe("Slice Iterator", func() {
	Describe("When creating a new slice iterator", func() {
		var (
			slice    []interface{}
			iterator enumerate.Iterable
		)

		BeforeEach(func() {
			slice = []interface{}{"a"}
			iterator = enumerate.Slice(slice)
		})

		It("should return nil after enumerating all values", func() {
			Ω(iterator.Next()).ShouldNot(BeNil())
			Ω(iterator.Next()).Should(BeNil())
		})

		It("Should not be effected by changes to the slice after initialization", func() {
			slice[0] = "b"
			item := iterator.Next()
			Ω(item).ShouldNot(BeNil())
			Ω(item).Should(Equal("a"))
		})
	})
})
