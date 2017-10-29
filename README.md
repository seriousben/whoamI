# whoamI

Tiny Go webserver that prints os information and HTTP request to output

```sh
$ docker run -it --rm -p "8083:8080" seriousben/whoami
$ curl -v 0.0.0.0:8083/
*   Trying ::1...
* TCP_NODELAY set
* Connected to localhost (::1) port 8083 (#0)
> GET / HTTP/1.1
> Host: localhost:8083
> User-Agent: curl/7.54.0
> Accept: */*
>
< HTTP/1.1 200 OK
< Date: Sun, 29 Oct 2017 13:44:29 GMT
< Content-Length: 130
< Content-Type: text/plain; charset=utf-8
<
Hostname: 287826a54272
IP: 127.0.0.1
IP: 172.17.0.2
GET / HTTP/1.1
Host: localhost:8083
User-Agent: curl/7.54.0
Accept: */*

* Connection #0 to host localhost left intact
$ curl -v localhost:8083/health
*   Trying ::1...
* TCP_NODELAY set
* Connected to localhost (::1) port 8083 (#0)
> GET /health HTTP/1.1
> Host: localhost:8083
> User-Agent: curl/7.54.0
> Accept: */*
>
< HTTP/1.1 200 OK
< Date: Sun, 29 Oct 2017 13:43:35 GMT
< Content-Length: 0
< Content-Type: text/plain; charset=utf-8
<
* Connection #0 to host localhost left intact
$ curl -v http://localhost:8083/api
*   Trying ::1...
* TCP_NODELAY set
* Connected to localhost (::1) port 8083 (#0)
> GET /api HTTP/1.1
> Host: localhost:8083
> User-Agent: curl/7.54.0
> Accept: */*
>
< HTTP/1.1 200 OK
< Date: Sun, 29 Oct 2017 13:45:14 GMT
< Content-Length: 118
< Content-Type: text/plain; charset=utf-8
<
{"hostname":"287826a54272","ip":["127.0.0.1","172.17.0.2"],"headers":{"Accept":["*/*"],"User-Agent":["curl/7.54.0"]}}
* Connection #0 to host localhost left intact
```
