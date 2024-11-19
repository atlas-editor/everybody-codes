Benchmark on Apple M1 Pro.

```
quest    part 1    part 2    part 3
---------------------------------------
01       0.084 ms  0.109 ms  0.722 ms
02       0.225 ms  40.59 ms  605.3 ms
03       0.108 ms  0.638 ms  3.420 ms
04       0.402 ms  0.205 ms  0.270 ms
05       0.165 ms  366.3 ms  0.687 ms
06       0.252 ms  0.638 ms  3.831 ms
07       0.087 ms  0.214 ms  1337. ms
08       0.081 ms  0.067 ms  40.19 ms
09       0.155 ms  0.674 ms  68.29 ms
10       0.078 ms  0.511 ms  2.685 ms
11       0.069 ms  0.102 ms  41.40 ms
```

Run as:
```
go run x.go input.txt P
```
where `input.txt` contains the input for the problem and `P` is the part (`1`, `2` or `3`).