setup:
	go get -u -t
	go mod tidy

me:
	go run ./main.go

wasm:
	gogio -target js .
	go get github.com/shurcooL/go-goon && go install github.com/shurcooL/goexec@latest
	goexec 'http.ListenAndServe(":3000", http.FileServer(http.Dir("gioui-experiment")))'