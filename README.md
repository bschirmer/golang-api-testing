# GoLang Endpoint testing
A simple GoLang backend API to learn a little bit of GoLang.
It needs unit tests

## How to run
This is assuming you have GoLang and MongoDB installed.<br> 
If you need GoLang, this will help; https://golang.org/doc/install<br>
If you need MongoDB, this will help; https://treehouse.github.io/installation-guides/mac/mongo-mac.html<br>

1. Git clone this repo into your go workspace. `$GOPATH/src`
2. Open a terminal (if you didnt use it one to clone) at the repo folder
3. `go install`
4. If at this point a bunch of "cannot find package" errors pop up, run `go get`
5. start that mongodb daemon `mongod` (this may only be a mac thing, I am nervous about this DB)
6. `go run main.go`
7. In browser go to `localhost:8080/<endpoint>`
8. Realistically, you should test this with post man, thats where the good data is at
9. Test your little heart out

## Things to add/change
HTTPS instead of HTTP<br>
Better error handling<br>
Build proper resonses<br>
Add response codes to the actual response, instead of a 200 every response with a custom error. I thought for this exercise and for testing purposes, this approach would be easier<br>