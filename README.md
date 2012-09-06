## knit

Knit parses knitting patterns into a parse tree.

To make this work reliably, we define a formal and relatively
strict pattern syntax. There are established knitting pattern formats
out there, but they are not strict enough to be suitable for parsing
in a practical and efficient manner.

The pattern constructs we do support are taken from the overview on
[craftyarncouncil.com](http://www.craftyarncouncil.com/knit.html).


### Pattern syntax

The lowest syntax level concerns the actual pattern constructs we
recognize and these are only a small handful of minimal
components we need to build everything else (see below).
All constructs are case insensitive.


### Stitch kinds

* `K`: Knit stitch. 
* `P`: Purl stitch. 
* `S` `Sl`: Slip stitch. 
* `Co`: Caston
* `Bo`: Bindoff
* `inc`: Increment
* `dec`: Decrement
* `Yo`: Yarn over

Stitches can be directly followed by a quantifier (see below), in order
to determine how often it should be repeated.

For example `P3 K2` means: Three Purl stitches, followed by two Knit stitches.

### Groupings

* Any sequence of stitches can be encased in `[` and `]`, to
  turn it into a distinct grouping. A group can have its own
  quantifiers (see below), to determine how often it should
  be repeated.

Groups can be directly followed by a quantifier, in order
to determine how often it should be repeated.

For example `[p3 k3] 10` means: Three Purl stitches, followed by three
knit stitches. And repeat the whole block ten times.


### Quantifiers

Quantifiers determine how often a given sitch or group should
be repeated. These can be made up of absolute numbers or special
tokens whch have a more abstract meaning.

* `0` `1` `2` etc: These give an absolute number of repetitions for
  the preceeding stitch or group.
* `+`: Repeat the preceeding stitch or group until the end of the row,
  regardless of how many stitches that might be.

Every stitch and group has an implicit quantifier of `1`.
There is therefore no need to specify it in the pattern if all you need is
for it to be knitted once.

E.g.: `PK` means one Purl Stitch, followed by one Knit stitch.
It is functionally identical to the pattern `P1K1`. 


### Pattern Nesting

In addition, we allow other patterns to be referenced inside a
given pattern by name. This allow us to split large patterns up into
smaller, managable chunks and build more complex patterns by embedding
the ready-made building blocks.

For example, pattern 'A' is defined as `[P3 K3] 10`.
This means: Purl three stitches and Knit three stitches; then repeat the
whole thing ten times.

Pattern 'B' can incorporate 'A' by referencing it by name: `P10 $A P10`.


### Reference Expansion

The parser does not expand the reference to 'A' during the parsing, but it
is left in there as a `Reference` node with the name `A`. It is up to the
host application to supply the actual contents of this reference during use,
or one can instruct the pattern to do so by calling the
`Pattern.Expand(ReferenceHandler)` method.

The host application must implement the `ReferenceHandler` and it should
return a valid, compiled pattern for the supplied reference name.
`Expand` returns a copy of the pattern where all references have been expanded.

After this call, the example pattern above would look like this:

	P10 [P3 K3] 10 P10


### Loop unrolling

The parser also does not do loop unrolling by default. However, it can be
instructed to do so by calling the `Pattern.Unwind()` method.

This returns a new pattern instance where all 'loops' are unrolled and any
references are replaced with the actual data from the referenced patterns.
For this purpose, it accepts a function pointer of type `ReferenceHandler`.

The host application must implement it and it should return a valid, compiled
pattern for the supplied reference name. After a call to `Expand`, the example
pattern above would look like this:

	P10 P3 K3 P3 K3 P3 K3 P3 K3 P3 K3 P3 K3 P3 K3 P3 K3 P3 K3 P3 K3 P10


### Usage

    go get github.com/jteeuwen/knit


### License

Unless otherwise stated, all of the work in this project is subject to a
1-clause BSD license. Its contents can be found in the enclosed LICENSE file.

