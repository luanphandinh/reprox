# reprox

Simple reverse proxy.

## Up and running
```bash
make -j3
```
```bash
go build -o ./build/reprox_server1 -v ./examples/server1
go build -o ./build/reprox_server2 -v ./examples/server2
go build -o ./build/reprox -v .
github.com/luanphandinh/reprox/examples/server2
github.com/luanphandinh/reprox/examples/server1
./build/reprox
2019/11/23 21:39:09 Reverse Proxy server start in port 8080
./build/reprox_server1
./build/reprox_server2
2019/11/23 21:39:10 Server 1 start in port 3000
2019/11/23 21:39:11 Server 2 start in port 3001

```

* Proxy server start in port 8080
    * /1 will forward request to Server 1
    * /2 will forward request to Server 2
* Server 1 start in port 3000
* Server 2 start in port 3001
