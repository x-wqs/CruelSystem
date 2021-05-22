## source
https://talks.golang.org/2012/splash.article

## abstract
> designed to make working in today's environment more productive, built-in concurrency, garbage collection, rigorous dependency management, adaptablility of software architecture as system grow, robustness across the boundaries between components

## introduction
> compiled, concurrent, garbage-collected, statically typed language

> efficient, scalable, productive

## go at google
> designed to eliminate slowness and clumsiness of software development at google, thereby to make process more productive and scalable

## pain points
- slow builds

- uncontrolled dependencies

- different subset of a language

- poor program understanding

- duplication of effort ? 

- cost of updates

- version skew

- difficulty of writing automatic tools

- cross-language builds

## dependencies in c
> header file be read many times

## attention for go
> work at scale

> be familiar with procedural languages

> be modern

## dependencies in go
> unused dependencies are compile-time error

> don't import packages not used directly

> so it guranteens each file opened exactly once

## packages
> paths are unique, names not need to be unique, usually concise 
## remote packages
> go get $package

## naming
> identifier start with upper case letter is public, otherwise not. This rule applies to variables, types, functions, methods, constants, fields and everything

## concurrency
> go is not purely memory safe in presence of concurrency, sharing is legal and passing a pointer over a channel is idiomatic

## composition not inheritance
> go don't has inheritance, only have interface

## tools
- Gofmt

- gofix

- godoc

## summary
- clear dependencies

- clear syntax

- clear semantics

- composition over inheritance

- simplicity provided by programming model(garbage collection, concurrency)

- easy tooling

## terms
conceive 构想

rigorous 严格的

clumsy 笨拙的

skew 歪斜

concise 简明

caveat 警告

idiomatic 惯用的

controversial 有争议的

distort 扭曲

