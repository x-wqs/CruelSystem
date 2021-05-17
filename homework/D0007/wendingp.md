# Go language design notes

A Static Type Language with a Greater Level of Productivity
Go is a statically typed, garbage-collected, natively compiled programming language that mostly belongs to the C family of languages in terms of basic syntax. Go provides an expressive syntax with it’s lightweight type system and comes with concurrency as a built-in feature at the language level. Go provides performance comparable to C and C++ while offering rapid development.

Concurrency is Built In Feature at Language Level


especially with big programs (millions of lines of code) to know whether the attempts to build a scalable language have paid off. 

## goals
An overarching goal was that Go do more to help the working programmer by enabling tooling, automating mundane tasks such as code formatting, and removing obstacles to working on large code bases.

simplifications: There are no forward declarations and no header files; everything is declared exactly once. Initialization is expressive, automatic, and easy to use. Syntax is clean and light on keywords. Stuttering (foo.Foo* myFoo = new(foo.Foo)) is reduced by simple type derivation using the := declare-and-initialize construct. And perhaps most radically, there is no type hierarchy: types just are, they don't have to announce their relationships.

keep the concepts orthogonal:
 Methods can be implemented for any type; structures represent data while interfaces represent abstraction; and so on. Orthogonality makes it easier to understand what happens when things combine.