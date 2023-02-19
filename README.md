# Look into confidence intervals

This is an attempt to understand "confidence intervals"

My [reference](https://www.mathsisfun.com/data/confidence-interval.html).
I'm still hazy on how the Z values get calculated.

My code for [reservoir selection](https://github.com/bediger4000/reservoir-sampling) of N random lines from
an input stream, which is used in the `confidence` script.

## Generate Test Data

### `gennormal` - create a number of normally-distributed floats

```
$ go build gennormal.go
$ ./gennormal 10000 100. 12. > normals.dat
$ datamash count 1 mean 1 sstdev 1 < normals.dat
10000   100.0782444786  12.063269274679
```
### `genexp` - create a number of exponentially-distributed floats

```
$ go build genexp.go
$ ./genexp 10000 100. > exponentials.dat
$ datamash count 1 mean 1 sstdev 1 < exponentials.dat
10000   100.0782444786  12.063269274679
```
## Testing the data

The idea is to have a data set with a known mean value,
then randomly  choose some N number of samples from that data set.
Computing the std deviation of those N values lets us calculate
a confidence interval.
After that, we can check to see if the known mean value fits in the confidence interval.

Since a "confidence interval" is supposedly what percentage of samplings
the mean of all the values fits between the mean of the sampled values
plus or minus the confidence,
we should be able to pull some number of samples from the data set
and calculate the confidence interval, then see if the known mean value
fits in the confidence interval.
If we do that 100 times, the count of times that the known mean fits in
the interval calculated from the samples should be the confidence interval.

### script `confidence`

The script `confidence` does exactly that.
It finds the mean value of the data set,
then runs a procedure 100 times.

Each procedure takes 100 random samples of the data set,
calculates an interval that should give 95% confidencethat
the real mean lies in that interval.

It counts how many times of the 100 times it runs the procedure
the known mean lies within the confidence interval, and prints out that number.
```
$ ./gennormal 10000 100. 12. > normals.dat
$ ./confidence normals.dat
32.647932       34.997187       35.169150       IN
32.816231       34.997187       35.321460       IN
...
34.973445       34.997187       37.190662       IN
32.805576       34.997187       34.909615       OUT
35.005867       34.997187       37.606491       OUT
35.140604       34.997187       37.498228       OUT
35.454684       34.997187       37.779112       OUT
96
```

For 96% of the procedures,
the known mean of the entire data set lies within the confidence interval.
