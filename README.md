# mackerel-linux-socket-queue-size

A Mackerel plugin that parses /proc/net/tcp and outputs the number of listen-queues and accept-queues for the specified IP and port, respectively.

## Usage

```
Usage:
  mackerel-plugin-linux-socket-queue-size [OPTIONS]

Application Options:
  -i, --ip=     process name
  -p, --port=   release directory
      --prefix= process port name
```

## exmaple

```
$ mackerel-plugin-linux-socket-queue-size --ip 192.168.1.1 --port 9000 --prefix nginx-ssl
nginx-ssl-socket-queue.syn-queue	0	1670762508
nginx-ssl-socket-queue.accept-queue	2	1670762508
```
