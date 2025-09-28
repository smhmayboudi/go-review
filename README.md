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

## package3

### رفتن به دایرکتوری package3

```shell
cd package3
```

### اجرای تست‌ها

```shell
go test -v
```

### اجرای تست‌ها با پوشش کامل

```shell
go test -v -cover
```

### اجرای بنچمارک‌ها

```shell
go test -bench=. -benchmem
```

### اجرای تست‌های خاص

```shell
go test -v -run TestUser3_Greet
go test -v -run TestAdd
```

### اجرای بنچمارک خاص

```shell
go test -bench=BenchmarkAdd -benchmem
```

### تولید پروفایل

```shell
go test -bench=. -cpuprofile=cpu.prof -memprofile=mem.prof
```

### مشاهده پوشش کد

```shell
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```