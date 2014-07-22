getpost
=======

getpost is a simple webserver that captures HTTP POST requests. 
The last request is then available via GET request (regular 
browser request) at http://[address]:[port]/request and http://[address]:[port]/body.

This might be useful for some debugging scenarios.

### Installation

The Go way (install in GOPATH/bin):

```
go get github.com/sontags/getpost
```

### Usage

```
ADDRESS=0.0.0.0 PORT=8765 $GOPATH/bin/getpost
```

The variables ADDRESS and PORT used in the example are the default, use whatever required or leave away
