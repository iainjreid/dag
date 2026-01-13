/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package dag_test

import (
	"reflect"
	"testing"

	"github.com/iainjreid/dag"
)

type TestVertex struct {
	name     string
	children []*TestVertex
}

func NewTestVertex(name string) *dag.Graph[string, TestVertex] {
	return dag.New(func(string) TestVertex {
		return TestVertex{
			name:     name,
			children: nil,
		}
	})
}

func (t TestVertex) Append(child TestVertex) {
	t.children = append(t.children, &child)
}

// TestAppend calls [Graph.Append], to ensure that it correctly adds a new
// child to the parent Vertex.
func TestAppend(t *testing.T) {
	var subject = NewTestVertex("parent").Append(NewTestVertex("child"))

	var expected = &TestVertex{
		name: "parent",
		children: []*TestVertex{
			{
				name: "child",
			},
		},
	}

	if !reflect.DeepEqual(subject.Evaluate("68yvwz"), expected) {
		t.Fatal("result should be equal to expected output")
	}
}

// TestLift calls [Graph.Lift], to ensure that dynamic Vertexs can be added
// using the provided build context.
func TestLift(t *testing.T) {
	var subject = NewTestVertex("parent").Lift(func(str string) *dag.Graph[string, TestVertex] {
		return NewTestVertex(str)
	})

	var expected = &TestVertex{
		name: "parent",
		children: []*TestVertex{
			{
				name: "x8azmu",
			},
		},
	}

	if !reflect.DeepEqual(subject.Evaluate("x8azmu"), expected) {
		t.Fatal("result should be equal to expected output")
	}
}

// TestScope calls [Scope], to ensure that it correctly modifies the execution
// context from the parent Vertex.
func TestScope(t *testing.T) {
	var subject = dag.Scope(NewTestVertex("parent").Lift(func(str string) *dag.Graph[string, TestVertex] {
		return NewTestVertex(str)
	}), func(str string) string {
		return str + "!"
	})

	var expected = &TestVertex{
		name: "parent",
		children: []*TestVertex{
			{
				name: "3IjnT4!",
			},
		},
	}

	if !reflect.DeepEqual(subject.Evaluate("3IjnT4"), expected) {
		t.Fatal("result should be equal to expected output")
	}
}

func BenchmarkAlloc(b *testing.B) {
	var subject = NewTestVertex("parent").Append(NewTestVertex("child"))

	for b.Loop() {
		subject.Evaluate("3IjnT4")
	}
}
