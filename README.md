# Dag

A minimally viable toolkit designed to build [Directed acyclic graphs] using
functional principles.

## Introduction

While developing [Dagger], I found that it was helpful to maintain a working
copy—also implemented in Go—of the fundamentals that underpinned the package due
to the difficulties I encountered whilst porting the library from an untyped
language.

As such, this repository holds a simplified implementation of the core behaviour
for demonstration and learning purposes for both myself and anyone else that
might be interested.

Before looking at the code itself, it would be useful to understand the motive
behind this body of work and the scope for interesting applications that you
might draw from it.

In 2019 I embarked on a personal project to deliver simple web applications
using fewer over-the-air bytes than comparable emerging micro-frameworks, whos
code representation (the transpiled output of JSX) offered little uniformality
for compression tools to take advantage of during transmission.

It became apparent to me that by writing a factory that applied the rules of
monad composition over that of the DOM API, it was possible to describe a DOM
tree plainly, through a series of component (blocks of web elements) definitions
that could then be further combined as one would with a traditional monad.

The result was [Fui], or "Functional User Interfaces", that achieved the goal
above albeit in a user-unfriendly way to some, due to the strict nature of how
the framework must be used.

Applications structured this way are, by design, repetitive in their form, and
more precisely mirror the structures of a tree, rather than a loose heirarchy of
components and effects that one might often see in React or Angular codebase.

This naturally leads to more compression friendly code, at the cost of a
dramatic concentration in the underlying framework complexity. Whether this
trade-off is acceptable is up to the programmer using it, but nevertheless
offers a novel alternative to the accepted norms of application development.

## License

This software is made available under the terms of the Mozilla Public License,
Version 2.0.

If a copy of this license was not distributed with this software, one can be
obtained at <https://mozilla.org/MPL/2.0/>.

[Dagger]: https://github.com/iainjreid/dagger
[Directed acyclic graphs]: https://wikipedia.org/wiki/Directed_acyclic_graph
[Fui]: https://github.com/iainjreid/fui
[encoding]: https://github.com/iainjreid/encoding
