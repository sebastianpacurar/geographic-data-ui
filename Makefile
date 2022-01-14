me:
	go run "./main.go"

wasm:
	gogio -target js .
	goexec 'http.ListenAndServe(":3000", http.FileServer(http.Dir("gioui-experiment")))'