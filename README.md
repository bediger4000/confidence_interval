# Look into confidence intervals

My [reference](https://www.mathsisfun.com/data/confidence-interval.html).

My code for [reservoir selection](https://github.com/bediger4000/reservoir-sampling) of N random lines from
an input stream, which is used in the `confidence` script.

## `gennormal` - create a number of normally-distributed floats

```
$ go build gennormal.go
$ ./gennormal 10000 100. 12. > normals.dat
$ datamash count 1 mean 1 sstdev 1 < normals.dat
10000   100.0782444786  12.063269274679
```

## script `confidence`

```
$ ./confidence normals.dat
```
