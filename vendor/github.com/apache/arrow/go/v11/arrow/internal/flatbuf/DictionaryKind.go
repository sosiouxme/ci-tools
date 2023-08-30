// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flatbuf

import "strconv"

/// ----------------------------------------------------------------------
/// Dictionary encoding metadata
/// Maintained for forwards compatibility, in the future
/// Dictionaries might be explicit maps between integers and values
/// allowing for non-contiguous index values
type DictionaryKind int16

const (
	DictionaryKindDenseArray DictionaryKind = 0
)

var EnumNamesDictionaryKind = map[DictionaryKind]string{
	DictionaryKindDenseArray: "DenseArray",
}

var EnumValuesDictionaryKind = map[string]DictionaryKind{
	"DenseArray": DictionaryKindDenseArray,
}

func (v DictionaryKind) String() string {
	if s, ok := EnumNamesDictionaryKind[v]; ok {
		return s
	}
	return "DictionaryKind(" + strconv.FormatInt(int64(v), 10) + ")"
}