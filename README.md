# lazzys

lazzys is a tool for custom implementations.

The server listen for specified route on specified port and return a response and log the informations related to the served request.

## Install

You can install lazzys with `go install`

```
▶ go install github.com/ahmedid/lazzys@latest
```

## Usage

```
▶ lazzys --help
usage: lazzys [-p port] [-path] [-c code] [-d body]
  -c int
    	HTTP status code (default 200)
  -d string
    	The response body
  -p int
    	Server port to listen (default 3080)
  -path string
    	URL Path
```

## Example

```
▶ lazzys -p 3080 -path home -d hello-world                              | ▶ curl -i http://localhost:3080/home
[2023-08-05 10:56:17] GET /home - (curl/7.54.0) [::1]:61401 - 200 11    |   HTTP/1.1 200 OK
                                                                        |   Date: Sat, 05 Aug 2023 09:56:17 GMT
                                                                        |   Content-Length: 11
                                                                        |   Content-Type: text/plain; charset=utf-8
                                                                        |
                                                                        |    hello-world
```

```
▶ lazzys -p 3080 -path admin -d admin-area -c 403                       | ▶ curl -i http://localhost:3080/admin
[2023-08-05 10:56:17] GET /admin - (curl/7.54.0) [::1]:61401 - 200 11   |   HTTP/1.1 403 Forbidden
                                                                        |   Date: Sat, 05 Aug 2023 10:01:05 GMT
                                                                        |   Content-Length: 10
                                                                        |   Content-Type: text/plain; charset=utf-8
                                                                        |
                                                                        |    admin-area
```

