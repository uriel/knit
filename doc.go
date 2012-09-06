// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

/*
Knit parses knitting patterns into a parse tree.

To make this work reliably, we define a simple and relatively
strict pattern syntax. There are established knitting pattern formats
out there, but they are not strict enough to be suitable for parsing
in a practical and efficient manner.

The pattern constructs we do support are taken from the overview on
[craftyarncouncil.com](http://www.craftyarncouncil.com/knit.html).

Usage example:

	pat, err := knit.Parse("MyPattern", "[P3 K3] 10")

or:

	pat := knit.MustParse("MyPattern", "[P3 K3] 10")
*/
package knit
