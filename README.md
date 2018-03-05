# passgen 
[![Build Status](https://travis-ci.org/JonathanSudibya/passgen.svg?branch=master)](https://travis-ci.org/JonathanSudibya/passgen)
[![codecov](https://codecov.io/gh/JonathanSudibya/passgen/branch/master/graph/badge.svg)](https://codecov.io/gh/JonathanSudibya/passgen)

Simple password (or string generator) library for golang

# Disclaimer

This project is inspired by cmiceli's [password-generator-go](https://github.com/cmiceli/password-generator-go)

# Benchmark
```
BenchmarkNewPasswordAllChars-4   	     1000000	      1279 ns/op	      80 B/op	       2 allocs/op
BenchmarkNewPasswordCapsChars-4   	     1000000	      1276 ns/op	      80 B/op	       2 allocs/op
BenchmarkNewPasswordLowerChars-4   	     1000000	      1296 ns/op	      80 B/op	       2 allocs/op
BenchmarkNewPasswordNumberChars-4   	 1000000	      1263 ns/op	      80 B/op	       2 allocs/op
BenchmarkNewPasswordSymbolChars-4   	 1000000	      1240 ns/op	      80 B/op	       2 allocs/op
BenchmarkNewPasswordMixedChars-4   	     1000000	      1278 ns/op	      80 B/op	       2 allocs/op
```

# LICENSE

Please see [LICENSE](https://github.com/JonathanSudibya/passgen/blob/master/LICENSE)