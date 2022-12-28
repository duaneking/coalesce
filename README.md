[![Go](https://github.com/duaneking/coalesce/actions/workflows/go.yml/badge.svg)](https://github.com/duaneking/coalesce/actions/workflows/go.yml)

# coalesce
Coalesce in Strongly Generic Typed GoLang.

No need for bloated modules or type specific functions.

Just use a generic one, and let GoLang use the type data it already has.

# Install

```
go get github.com/duaneking/coalesce
```

# Tests

This modules GoLang Native Unit Test Code Coverage is at 100% using 2 tests, by design.

# Migrating from Prior Legacy Code

If you used other golang modules before, you can remove it and replace it with this much smaller and fully tested module instead.

# Why

Golang does not support the "coalesce" operation by default as part of its language design, or its experimental syntax tree definition.  This is a common operation in other laguages such as C#, SQL, and others. See https://en.wikipedia.org/wiki/Null_coalescing_operator for more data on other languages and how they handle this.

# Ok, Really, why do this? The code is smaller than the docs.

Thats intentional. 

The lack of a language feature for an official coalesce operation in GoLang resulted in more complicated and bug filled code then I wanted to deal with for a personal project that I'm activly working on; so I built out this module so that I could just import it and reuse it for future projects, instead of needing to slow down my development process to deal with what I percive as a big design gap in GoLang.

From a language persective, I consider any modern language without native null coalesce support to have a severe design gap. And this was a big enough gap that before generics were supported in GoLang, your only option was to write a function that did the simple coalesce operation for the specific type you are working with.

This created the need for giant, bloated libraries that tried to define the same common operation typed for all the possible default supported types in GoLang in an effort to be a "complete solution".. and that creates an issue where if your custom type is not a part of that library, its not supported.

Your custom type is fully supported with this library, by design, unlike others that exist.

Not to menton, now you have less dead code from a library that your importing but not using 100% because you only need the operation for a few types. This fully solves dead code for this specific operation.

So you could import a module that didnt support generics, and deal with the cost and expence as a tradeoff.. or I could just create a generic version of the single operation that does not need to be duplicated all over the place, supports your custom types, and could replace thousands of lines of code.

So pondering this issue, I decided to simply solve the problem, once and for all.

This module is one single method. But it works. You import this, pass in the correct arguments, and it does everything that libraries several hundreds of times in size also claim to do... just slimmer and with 100% code fully tested, no matter how many times you use it, thanks to the fact that GoLang passes all type data as part of its internal structure.

Is it small? Yes.

Is it kinda dumb that its needed? Yes.

But is it a solved problem that I never have to think about now? Also, Yes. And thats a good thing.

# Is this code trustworthy enough to be used anywhere?

Yes. This code should be considered production quality as a module you can simply import and reuse. The code coverage is complete and 100%.
