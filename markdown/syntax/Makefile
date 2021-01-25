get-requireds:
	echo "TODO: install ragel and goyacc if need"

gen: get-requireds
	ragel -Z -G2 -o lexer.go lexer.rl
	goyacc -p md -o parser.go parser.y
	gofmt -w parser.go lexer.go
