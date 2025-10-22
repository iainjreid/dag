/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

// Package dag implements a minimally viable functional toolkit designed to
// build directed acyclic graphs in a composable manner.
package dag

type Builder[T any, N Node[N]] struct {
	Build func(T) N
}

type Node[N any] interface {
	Append(N)
}

// New creates a composition builder using the provided function
func New[T any, N Node[N]](build func(T) N) *Builder[T, N] {
	return &Builder[T, N]{
		build,
	}
}

// Append adds one or more children to the existing tree.
func (b *Builder[T, N]) Append(child *Builder[T, N]) *Builder[T, N] {
	return b.Tap(func(x T, parent N) {
		parent.Append(child.Build(x))
	})
}

// Lift allows for the dynamic insertion nodes into a tree.
func (b *Builder[T, N]) Lift(f func(context T) *Builder[T, N]) *Builder[T, N] {
	return b.Tap(func(x T, parent N) {
		parent.Append(f(x).Build(x))
	})
}

// Tap is a utility method that abstracts common behaviour required by
// [Builder.Append] and [Builder.Lift].
func (b *Builder[T, N]) Tap(f func(context T, parent N)) *Builder[T, N] {
	return New[T](func(x T) N {
		result := b.Build(x)
		f(x, result)
		return result
	})
}

// Scope maps the execution context from a parent tree, to be accepted by a
// subtree.
func Scope[T1, T2 any, N Node[N]](b *Builder[T2, N], f func(T1) T2) *Builder[T1, N] {
	return New[T1, N](func(x T1) N {
		return b.Build(f(x))
	})
}
