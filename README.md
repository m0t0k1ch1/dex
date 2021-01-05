# dex

a decimal <-> hex converter written in Go

dex can handle big numbers.

## decimal -> hex

``` sh
$ ./dex 100000000000000000000
0x056bc75e2d63100000
```

## hex -> decimal

``` sh
$ ./dex 0x056bc75e2d63100000
100000000000000000000
```
