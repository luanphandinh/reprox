example: server1 server2 proxy

proxy:
	go build -o ./build/proxy -v ./examples/proxy
	./build/proxy

server1:
	go build -o ./build/server1 -v ./examples/server1
	./build/server1

server2:
	go build -o ./build/server2 -v ./examples/server2
	./build/server2