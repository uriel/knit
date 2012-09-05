// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

/*
Knit parses knitting patterns into a parse tree.

To make this work reliably, we define a formal and relatively
strict pattern syntax. There are established knitting pattern formats
out there, but they are not strict enough to be suitable for parsing
in a practical and efficient manner.

In order to find the middle ground and not to inconvenience the user too
much by requiring them to learn a new pattern language, we incorporate as
many of the established pattern practises as possible, while still
maintaining reliability and effectiveness in an automated parsing
environment like this package.
*/
package knit
