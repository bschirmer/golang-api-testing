# Jumbo Backend Test




## Implemented API


## Not Implemented 


## Instructions to run


## Little tidbits
This is the first time I have touched GoLang
Resources used:
https://golang.org/doc/code.html - Go basics
https://www.youtube.com/watch?v=W5b64DXeP0o - to get my webserver up and running
https://www.youtube.com/watch?v=YMQUQ6XQgz8 - to get my head around routes and http verbs
https://medium.com/@edwardpie/parsing-json-request-body-return-json-response-with-golang-c4f862bbb19b - to get my head around getting http body data
https://github.com/mlabouardy/movies-restapi - I used this to get my head aaround a datastore

## Things I would add if I knew GO better and if I was making this a production application
HTTPS instead of HTTP
Better error handling
I would add response codes to the actual response, instead of a 200 every response with a custom error. I thought for this exercise and for testing purposes, this approach would be easier

## How to run
This is assuming you have GoLang and MongoDB installed. 
If you need GoLang, this will help; https://golang.org/doc/install
If you need MongoDB, this will help; https://treehouse.github.io/installation-guides/mac/mongo-mac.html

1. Git clone this repo into your go workspace. $GOPATH/src
2. Open a terminal (if you didnt use it one to clone) at the repo folder
3. go install
4. start that mongodb daemon (this may only be a mac thing, I am nervous about this DB)
5. go run main.go
6. Test your little heart out

I know how mundane testing is, so I have included the Postman collection I used to test this little api. You just need to import it into Postman and run it. Each api call has a test that should pass (fingers crossed)

## TODO Dev notes
- uninstall mongo and see what happens
- push to git
    - clone git
    - see who happens
- write running ops

- default data
    - at elast 1 pet
    - update that pet id
    - 5 pets, 2 with same status, 1 and 1
        - this is for find with status
    - a deleteable pet