Simple test go projects

notes:

1.3
$ go run main.go

1.4
$ go run main.go t1 t2 t3

1.5
& go build
./1.5 > out.gif

1.6
& go build
./1.6 > out.gif

1.7
$ go build
./1.7 http://gopl.io

1.8
$ go build
./1.8 http://gopl.io
./1.8 gopl.io

1.9
$ go build
./1.9 http://gopl.io
./1.9 gopl.io

1.10
$ go build
./1.10 http://gopl.io

1.11
./1.10 http://Google.Com http://Youtube.Com http://Facebook.Com http://Wikipedia.Org http://Yahoo.Com http://Twitter.Com http://Amazon.Com http://Instagram.Com

1.12
$ go run main.go
http://localhost:8000/?cycles=20;res=0.001;size=100;nframes=64;delay=8

2.1
//add $GOPATH variable:
GOPATH="/home/alex/go"
//test the package
tempconv$ go build
//install the package
tempconv$ go install
//run code from the package
go run main.go

2.2
//using package without compiling:
go build main.go
./main -40
//altarnative way without build:
go run main.go 30 50



