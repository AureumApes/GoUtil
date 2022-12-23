GoUtils
=======
A collection of handy Go functions used by some Amiraxoba projects

Installation
------------
1. Install [Go](https://go.dev)
2. Run `go get -u -x github.com/AureumApes/GoUtils/` or `go get github.com/AureumApes/GoUtils/`
3. Start using in your projects

Functions
---------
### Generate ID
Usage: `GoUtils#GenerateId(lenght)`<br />
Returns: Random string of numbers, with a special leght
### Is Numeric
Usage: `GoUtils#IsNumeric(stringToTest)`<br />
Returns: Bool, if the string only contains a convertable 64-bit float
### Remove From Array
Usage: `GoUtils#RemoveFromArray(slice, index)`<br />
Returns: The slice, that is given to it in the Parameters, but without the entry with the given index
