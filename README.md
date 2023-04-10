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