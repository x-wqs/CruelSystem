## introduction
Go: a compiled, concurrent, garbage-collected, statically typed language

## Dependencies in Go
The first step to making Go scale, dependency-wise, is that the language defines that unused dependencies are a compile-time error (not a warning, an error). If the source file imports a package it does not use, the program will not compile.

Another feature of the Go dependency graph is that it has no cycles. The language defines that there can be no circular imports in the graph, and the compiler and linker both check that they do not exist. 

## Packages
package paths are unique, but there is no such requirement for package names

```
import "log"                          // Standard package
import googlelog "google/base/go/log" // Google-specific package
```
## Remote packages
```
go get github.com/4ad/doozer // Shell command to fetch package
import "github.com/4ad/doozer" // Doozer client's import statement
var client doozer.Conn         // Client's use of package
```

## Syntax
Go does not support default function arguments. This was a deliberate simplification.

## Naming
In Go the name itself carries the visibility of an identifier: 

uppercase initial letter: Name is visible to clients of package
otherwise: name (or _Name) is not visible to clients of package

## Semantics
It is a compiled, statically typed, procedural language with pointers and so on. 

## Concurrency
Go embodies a variant of CSP with first-class channels.

## Garbage collection
Moreover, in a concurrent object-oriented language it's almost essential to have automatic memory management because the ownership of a piece of memory can be tricky to manage as it is passed around among concurrent executions.

GC brings significant costs: general overhead, latency, and complexity of the implementation

The current design is a parallel mark-and-sweep collector.

## Composition not inheritance
There is no type hierarchy. This was an intentional design choice.
All data types that implement these methods satisfy this interface implicitly; there is no implemented declaration. That said, interface satisfaction is statically checked at compile time so despite this decoupling interfaces are type-safe.

## Errors
There is no control structure associated with error handling.
## Tools
Gofmt, godoc, and gofix.
