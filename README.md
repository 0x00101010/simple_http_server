# simple_http_server
A simple http server written in Go which supports some parts of the HTTP/1.0 protocol

## What do we support

* persistent connections from HTTP/1.1
* serving files with extensions *.jpg, *.html, *.txt, *.jpg, *.gif
* Some of HTTP Status Code (200, 404, 403, 400)
* Configuring document root and port

## Environment
* go version 1.13+

## How to build the project
Go to project root directory, run below command:
```
make build
```

## How to test
After building the project, please run the following command to start the server, and then try to type localhost:`port` in your browser.
```
make test
```

## One command execution
You could also try using the below command for both building and testing the project. I've commited the web resources under `scu_websites` folder
```
make run
```

## Author
Huiyi (Francis) Li
hli2@scu.edu

