/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

// Package dag implements a minimally viable toolkit designed to build
// directed acyclic graphs in a composable manner.
package dag

type Graph[T any, N Node[N]] struct {
	Evaluate func(T) N
}

type Node[N any] interface {
	Append(N)
}

// New creates a [Graph] using the provided function as the [Node] factory.
func New[T any, N Node[N]](evaluate func(T) N) *Graph[T, N] {
	return &Graph[T, N]{
		evaluate,
	}
}

// Append adds a child [Graph] to the existing [Graph].
func (g *Graph[T, N]) Append(child *Graph[T, N]) *Graph[T, N] {
	return g.Tap(func(x T, parent N) {
		parent.Append(child.Evaluate(x))
	})
}

// Lift allows for the dynamic insertion a child [Graph] into the existing [Graph].
func (g *Graph[T, N]) Lift(fn func(context T) *Graph[T, N]) *Graph[T, N] {
	return g.Tap(func(x T, parent N) {
		parent.Append(fn(x).Evaluate(x))
	})
}

// Tap is a utility method that abstracts common behaviour required by both
// [Graph.Append] and [Graph.Lift].
func (g *Graph[T, N]) Tap(fn func(context T, parent N)) *Graph[T, N] {
	return New(func(x T) N {
		result := g.Evaluate(x)
		fn(x, result)
		return result
	})
}

// Scope maps the execution context from a parent [Graph] to a child [Graph].
func Scope[T1, T2 any, N Node[N]](graph *Graph[T2, N], fn func(T1) T2) *Graph[T1, N] {
	return New(func(x T1) N {
		return graph.Evaluate(fn(x))
	})
}
