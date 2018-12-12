// Copyright 2018 HouseCanary, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package schema

var _ WrappedType = (*NotNilType)(nil)

// A NotNilType expresses Type! in GraphQL terms
type NotNilType struct {
	of Type
}

func (t *NotNilType) isType() {}

// Unwrap returns the wrapped type (i.e. Type in Type!)
func (t *NotNilType) Unwrap() Type {
	return t.of
}

// InputListCreator returns a factory of lists of the contained type
func (t *NotNilType) InputListCreator() InputListCreator {
	return t.of.(InputableType).InputListCreator()
}

func (t *NotNilType) signature() string {
	return t.of.signature() + "!"
}
