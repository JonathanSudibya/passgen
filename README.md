# passgen

Simple password (or string generator) library for golang

# Disclaimer

This project is based on cmiceli's [password-generator-go](https://github.com/cmiceli/password-generator-go)

# Benchmark
```
BenchmarkNewPasswordAllChars-4   	  500000	      2485 ns/op	     328 B/op	       6 allocs/op
BenchmarkNewPasswordCapsChars-4   	 1000000	      1366 ns/op	     144 B/op	       4 allocs/op
BenchmarkNewPasswordMixedChars-4   	  500000	      3373 ns/op	     208 B/op	       5 allocs/op
```

# LICENSE

Please see [LICENSE](https://github.com/JonathanSudibya/passgen/blob/master/LICENSE)