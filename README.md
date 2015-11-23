# bench_write

## Macbook Pro

```
BenchmarkWriteEveryOpenFile-4      30000             48164 ns/op              72 B/op          3 allocs/op
BenchmarkWriteOpeningFile-4      1000000              1365 ns/op               0 B/op          0 allocs/op
ok      github.com/catatsuy/bench_write 4.502s
```

## Debian7(on VirtualBox)

```
BenchmarkWriteEveryOpenFile       500000              3053 ns/op              72 B/op          3 allocs/op
BenchmarkWriteOpeningFile        1000000              1382 ns/op               0 B/op          0 allocs/op
ok      _/home/vagrant/bench_write      2.969s
```
