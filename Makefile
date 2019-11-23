example: server1 server2 proxy

proxy:
	go build -o ./build/reprox -v .
	./build/reprox

server1:
	go build -o ./build/reprox_server1 -v ./examples/server1
	./build/reprox_server1

server2:
	go build -o ./build/reprox_server2 -v ./examples/server2
	./build/reprox_server2