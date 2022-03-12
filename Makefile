setup:
	go mod download gioui.org
	go get -u
	go mod tidy

me:
	go run ./main.go

wasm:
	go install gioui.org/cmd/gogio@latest
	gogio -target js .
	go get github.com/shurcooL/go-goon && go install github.com/shurcooL/goexec@latest
	goexec 'http.ListenAndServe(":3000", http.FileServer(http.Dir("gioui-experiment")))'