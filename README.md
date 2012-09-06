## knit

**Note**: This is work in progress and subject to change.

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

Quantifiers determine how often a given stitch or group should
be repeated. These are made up of absolute numbers (`1`, `2`, etc).
These specify an absolute number of repetitions for the preceeding
stitch or group.

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

For example, pattern 'A' is defined as: `[P3 K3] 5`.
Pattern 'B' can incorporate 'A' by referencing it by name: `P10 $A 2 P10`.


### Reference Expansion

The parser does not expand the reference to 'A' during the parsing, but it
is left in there as a `Reference` node with the name `A`. It is up to the
host application to supply the actual contents of this reference during use,
or one can instruct the pattern to do so by calling the
`Pattern.Expand(ReferenceHandler)` method.

The host application must implement the `ReferenceHandler` and it should
return a valid, compiled pattern for the supplied reference name.
`Expand` expands all references in place.

After this call, the 'B' above would look like this:

	P10 [[P3 K3] 5] 2 P10


### Loop unrolling

The parser does not do loop unrolling by default. However, it can be
instructed to do so by calling the `Pattern.Unroll()` method.

For example:

	P2 K3 [P INC 2] 2

Will become:

	P P K K K P INC INC P INC INC

This is mostly practical when using the pattern to generate some other
form of output. Getting rid of the nested groupings like this makes
processing considerably easier.

After a call to `Unroll`, there should be no Number or Group nodes
left in the pattern. Only basic Stitch nodes.


### Usage

    go get github.com/jteeuwen/knit


### License

Unless otherwise stated, all of the work in this project is subject to a
1-clause BSD license. Its contents can be found in the enclosed LICENSE file.

