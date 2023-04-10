# hello-flat
A sample of flat buffers using Go.

## Prepare
Install flatc.

## Generate
```bash
flatc --go --gen-object-api schema/fbs/user.fbs
```

## Run
```bash
go run main.go
```

result
```
=====
no pack
length=144, msg=`100000000C001C0018000800070014000C0000000000000A0000304100004041000050411800000004000000080000004A6F686E20446F6500000000030000000C0000001C00000034000000D6FFFFFF040000000500000061726D6F72000000EAFFFFFF0400000006000000736869656C640000000006000800040006000000040000000500000073776F7264000000`
name=`John Doe`
color=Color(10)
position{x, y, z} = {11, 12, 13}
item[0]:name=armor
item[1]:name=shield
item[2]:name=sword
=====
pack
length=144, msg=`100000000C001C0018000C000B0004000C000000180000000000000A00003041000040410000504158000000030000003C0000001C00000004000000D6FFFFFF040000000500000061726D6F72000000EAFFFFFF0400000006000000736869656C640000000006000800040006000000040000000500000073776F7264000000080000004A6F686E20446F6500000000`
name=`John Doe`
color=Color(10)
position{x, y, z} = {11, 12, 13}
item[0]:name=sword
item[1]:name=shield
item[2]:name=armor
```

## Bench
```bash
go test -bench . -benchmem 
```

result
```
goos: linux
goarch: amd64
pkg: github.com/blck-snwmn/hello-flat
cpu: 11th Gen Intel(R) Core(TM) i7-11700 @ 2.50GHz
BenchmarkCreate-16               1324420               937.0 ns/op          1108 B/op          4 allocs/op
BenchmarkCreateWithFunc-16       1300110               939.9 ns/op          1108 B/op          4 allocs/op
BenchmarkCreateWithUserT-16       822213              1451 ns/op            1272 B/op         15 allocs/op
BenchmarkProto-16                 587275              2028 ns/op             880 B/op         19 allocs/op
PASS
ok      github.com/blck-snwmn/hello-flat        8.144s
```