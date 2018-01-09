# licrep

[![GoDoc](https://godoc.org/github.com/kopoli/licrep?status.svg)](https://godoc.org/github.com/kopoli/licrep)

Report and embed the licenses of your go dependencies.

## The problem

The common Open Source licenses have language which requires the license and
copyright notice be included in all copies of the software.

In MIT license

```
The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.
```

In BSD license

```
* Redistributions in binary form must reproduce the above copyright notice,
  this list of conditions and the following disclaimer in the documentation
  and/or other materials provided with the distribution.
```

Go software can be distributed as a single binary and this tool is for
embedding the licenses to that binary.

## Installation

```
$ go get github.com/kopoli/licrep
```

## Usage

Display the summary of the dependencies of your go package by running:

```
$ licrep -s
```

Embed the licenses as a data structure to your program by adding the following
comment:

```
//go:generate licrep -o licenses.go
```

This will create a type called `License` and a function `GetLicenses`, which
returns `[]License, error`.

See all options by running:
```
$ licrep --help
```

For a complete example see `_example/example.go`.

## License

MIT license
