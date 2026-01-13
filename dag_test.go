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

type TestNode struct {
	name     string
	children []*TestNode
}

func NewTestNode(name string) *dag.Graph[string, *TestNode] {
	return dag.New(func(string) *TestNode {
		return &TestNode{
			name:     name,
			children: nil,
		}
	})
}

func (t *TestNode) Append(node *TestNode) {
	t.children = append(t.children, node)
}

// TestAppend calls [Graph.Append], to ensure that it correctly adds a new
// child to the parent node.
func TestAppend(t *testing.T) {
	var subject = NewTestNode("parent").Append(NewTestNode("child"))

	var expected = &TestNode{
		name: "parent",
		children: []*TestNode{
			{
				name: "child",
			},
		},
	}

	if !reflect.DeepEqual(subject.Evaluate("68yvwz"), expected) {
		t.Fatal("result should be equal to expected output")
	}
}

// TestLift calls [Graph.Lift], to ensure that dynamic nodes can be added
// using the provided build context.
func TestLift(t *testing.T) {
	var subject = NewTestNode("parent").Lift(func(str string) *dag.Graph[string, *TestNode] {
		return NewTestNode(str)
	})

	var expected = &TestNode{
		name: "parent",
		children: []*TestNode{
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
// context from the parent node.
func TestScope(t *testing.T) {
	var subject = dag.Scope(NewTestNode("parent").Lift(func(str string) *dag.Graph[string, *TestNode] {
		return NewTestNode(str)
	}), func(str string) string {
		return str + "!"
	})

	var expected = &TestNode{
		name: "parent",
		children: []*TestNode{
			{
				name: "3IjnT4!",
			},
		},
	}

	if !reflect.DeepEqual(subject.Evaluate("3IjnT4"), expected) {
		t.Fatal("result should be equal to expected output")
	}
}
