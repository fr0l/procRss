# procRss

A utility to calculate resident set size (RSS) for a process and it's descendants.

Useful to get a hint of how much RAM is occupied by a process with a lot of children (for example iOS Simulator has more than 100 child processes).

Under the hood it uses ```ps``` utility, so make sure it is installed (package ```procps``` on popular linux distributions). 

## Usage:

```bash
$ ./procRss 1234
```

```bash
$ pgrep -f launchd_sim | xargs ./procRss # iOS Simulator

# sample output
[procRss 86116]
PID:   86116, process RSS       5 MB, descendants RSS    1769 MB, total RSS    1774 MB

$ pgrep -f java | xargs ./procRss # a couple of busy Jenkins nodes

# sample output
[./procRss_linux64 5789 12889]
PID:    5789, process RSS      70 MB, descendants RSS    1587 MB, total RSS    1657 MB
PID:   12889, process RSS      74 MB, descendants RSS     685 MB, total RSS     759 MB
```

## Build:

Use go compiler
```bash
$ go build
```

it will produce binary procRss in current directory

or use a build script

```bash
$ ./build.sh
```
it will produce two binaries in ```build``` directory - one for MacOS and one for Linux