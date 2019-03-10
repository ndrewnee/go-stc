# go-tsc

## Description

Port of [The Super Tiny Compiler](https://github.com/jamiebuilds/the-super-tiny-compiler) written in Go

## Usage

```bash
go get -u github.com/ndrewnee/go-stc/...
stc "(add 10 (subtract 10 6))"
# Output: add(10, subtract(10, 6));
```

## Tests

```bash
cd $GOPATH/src/github.com/ndrewnee/go-stc
go test ./...
```
