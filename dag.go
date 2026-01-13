/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

// Package dag implements a minimally viable toolkit designed to build
// directed acyclic graphs in a composable manner.
package dag

type Graph[T any, V Vertex[V]] struct {
	Evaluate func(T) V
}

type Vertex[V any] interface {
	Append(V)
}

// New creates a [Graph] using the provided function as the [Vertex] factory.
func New[T any, V Vertex[V]](evaluate func(T) V) *Graph[T, V] {
	return &Graph[T, V]{
		evaluate,
	}
}

// Append adds a child [Graph] to the existing [Graph].
func (g *Graph[T, V]) Append(child *Graph[T, V]) *Graph[T, V] {
	return g.Tap(func(x T, parent V) {
		parent.Append(child.Evaluate(x))
	})
}

// Lift allows for the dynamic insertion a child [Graph] into the existing [Graph].
func (g *Graph[T, V]) Lift(fn func(context T) *Graph[T, V]) *Graph[T, V] {
	return g.Tap(func(x T, parent V) {
		parent.Append(fn(x).Evaluate(x))
	})
}

// Tap is a utility method that abstracts common behaviour required by both
// [Graph.Append] and [Graph.Lift].
func (g *Graph[T, V]) Tap(fn func(context T, parent V)) *Graph[T, V] {
	return New(func(x T) V {
		result := g.Evaluate(x)
		fn(x, result)
		return result
	})
}

// Scope maps the execution context from a parent [Graph] to a child [Graph].
func Scope[T1, T2 any, V Vertex[V]](graph *Graph[T2, V], fn func(T1) T2) *Graph[T1, V] {
	return New(func(x T1) V {
		return graph.Evaluate(fn(x))
	})
}
