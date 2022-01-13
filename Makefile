me:
	go run "./main.go"

wasm:
	gogio -target js .
	goexec 'http.ListenAndServe(":8080", http.FileServer(http.Dir("gioui-experiment")))'