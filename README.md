# go-review

## profile

### ایجاد پروفایل CPU
```shell
go test -cpuprofile=cpu.prof -bench=.
```

### ایجاد پروفایل حافظه
```shell
go test -memprofile=mem.prof -bench=.
```

### تجزیه و تحلیل پروفایل با ابزار pprof
```shell
go tool pprof cpu.prof
```
