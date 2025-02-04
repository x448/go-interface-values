/*
Copyright 2022

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package lem_test

import (
	"testing"

	"go-interface-values/tests/lem/lem"
)

var lemFuncs = map[string]func(*testing.B){}

func TestLem(t *testing.T) {
	lem.SetBenchmem("true")
	lem.SetBenchtime("1000x")
	lem.B{
		Benchmarks:  lemFuncs,
		IncludeDirs: []string{lem.MyDirectory()},
	}.Run(t)
}
