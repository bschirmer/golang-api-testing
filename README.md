# A Jumbo Backend Test
Firstly, I would like to thank you for allowing me to do this test. I haven't touched this sort of tech in a while because my job doesnt need me to, but it was a lot of fun touching on things that challenge me. <br>

Andy Kidd told me that Jumbo are moving to GoLang so instead of coding this in a laguange I have 2 years experience in (PHP), I decided to show that no only am I eager to learn, but that I am determined to push myself. If I am being honest, instead of taking me 4 hours, this test would have taken me more lie 24 hours. This includes learning GoLang from scratch and finding all the information I needed<br>

## Implemented API
As I have never touched GoLang, I implemented all of the Pet endpoints except for upload image. This was because I had already spent a long time with the rest and really wanted to submit this. 

POST /pet <br>
PUT  /pet <br>
GET  /pet/findByStatus <br>
GET  /pet/{petId} <br>
POST /pet/{petId} <br>
DELETE /pet/{petId} <br>
<br>
All Models<br>
<br>

## Not Implemented 
All Store and User endpoints

## How to run
This is assuming you have GoLang and MongoDB installed. 
If you need GoLang, this will help; https://golang.org/doc/install
If you need MongoDB, this will help; https://treehouse.github.io/installation-guides/mac/mongo-mac.html

1. Git clone this repo into your go workspace. `$GOPATH/src`
2. Open a terminal (if you didnt use it one to clone) at the repo folder
3. `go install`
4. If at this point a bunch of "cannot find package" errors pop up, run `go get`
5. start that mongodb daemon `mongod` (this may only be a mac thing, I am nervous about this DB)
6. `go run main.go`
7. In browser go to `localhost:8080/<endpoint>`
8. Realistically, you should test this with post man, thats where the good data is at
9. Test your little heart out

I know how mundane testing is, so I have included the Postman collection I used to test this little api. You just need to import it into Postman and run it. Each api call has a test that should pass (fingers crossed)

### Testing Data
Because sending a DB is a little ticky in this scenario (as most DBs are in the cloud or managed by a team of people that isn't me), I have provided some test data that will be added everytime the api is started.

The test data includes:
* 4 pets
* 2 with the same status, 2 with the other 2 statuses
* It does not include goo Category, Tag or Image data as this wasn't really implemented, as explained above

The test data will allow a Read, Update, Find by status and Delete without any issues.
This is what the Postman collection covers. If you decide to run that collection, you will end up with 4 pets, but 2 will be updated and 1 will be new and 1 will be deleted. 

## Little tidbits
This is the first time I have touched GoLang<br>
Resources used:<br>
https://golang.org/doc/code.html - Go basics<br>
https://www.youtube.com/watch?v=W5b64DXeP0o - to get my webserver up and running<br>
https://www.youtube.com/watch?v=YMQUQ6XQgz8 - to get my head around routes and http verbs<br>
https://medium.com/@edwardpie/<br>parsing-json-request-body-return-json-response-with-golang-c4f862bbb19b - to get my head around getting http body data<br>
https://github.com/mlabouardy/movies-restapi - I used this to get my head aaround a datastore<br>

## Things I would add if I knew GO better and if I was making this a production application
HTTPS instead of HTTP<br>
Better error handling<br>
I would build proper resonses<br>
I would add response codes to the actual response, instead of a 200 every response with a custom error. I thought for this exercise and for testing purposes, this approach would be easier<br>