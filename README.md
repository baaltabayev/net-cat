# Net-cat

This project consists on recreating the **NetCat in a Server-Client Architecture** that can run in a server mode on a specified port listening for incoming connections, and it can be used in client mode, trying to connect to a specified port and transmitting information to the server.

- NetCat, `nc` system command, is a command-line utility that reads and writes data across network connections using TCP or UDP. It is used for anything involving TCP, UDP, or UNIX-domain sockets, it is able to open TCP connections, send UDP packages, listen on arbitrary TCP and UDP ports and many more.

- To see more information about NetCat inspect the manual `man nc`.

## Usage

```console
$ go run .
Listening on the port :8989
//Open minimum two new terminal and enter this code in each of them:
$ nc localhost 8989
//You can specify your port here:
$ go run . 2525
Listening on the port :2525
//There's an incorrect input:
$ go run . 2525 localhost
[USAGE]: ./TCPChat $port
$
```
## Authors
@abaltaba
@akaiyrgel