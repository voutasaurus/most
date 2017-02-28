# most

A tool that lists files in the current directory and all sub directories in order of line numbers.

The list cuts off at 50% of total line count, so by the pareto principle it's usually a
much shorter list than listing all of the files.

# install

Install go: https://golang.org/dl/

```
$ go get -u github.com/voutasaurus/most
```

# usage (examples run on the latest `less` source repo)

List files comprising 50% of the total line count:
```
$ most
    6755, configure
    2501, screen.c
    1806, command.c
    1774, less.nro
    1743, search.c
    1700, cmdbuf.c
    1640, less.man
    1257, line.c
    1253, regexp.c
    1131, filename.c
     974, ch.c
...
```

List files comprising 80% of the total line count:
```
$ most -p 80
    6755, configure
    2501, screen.c
    1806, command.c
    1774, less.nro
    1743, search.c
    1700, cmdbuf.c
    1640, less.man
    1257, line.c
    1253, regexp.c
    1131, filename.c
     974, ch.c
     902, NEWS
     874, lesskey.c
     842, decode.c
     826, edit.c
     823, charset.c
     796, version.c
     757, tags.c
     743, optfunc.c
     708, option.c
     698, configure.ac
     674, COPYING
     614, output.c
     608, opttbl.c
     586, prompt.c
     533, less.h
     470, linenum.c
     463, input.c
     446, forwback.c
     437, defines.h.in
     414, defines.ds
...
```
