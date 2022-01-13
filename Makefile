me:
	go run "./main.go"

# currently not working
wasm:
	gogio -target js ./main.go
	goexec 'http.ListenAndServe(":8080", http.FileServer(http.Dir("./")))'