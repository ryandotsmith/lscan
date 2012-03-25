# lscan

## Synopsis

Turn k=v log data into Go maps.

## Description

This pkg is used in wcld to parse log lines into maps to store in a PostgreSQL
database. See the test file for examples.

## Build

First you will need the latest Go1 RC.

```bash
$ cd $GOROOT
$ hg pull
$ hg update weekly
$ ./src/all.bash
```

After that you are all set!

```bash
$ go test .
```
