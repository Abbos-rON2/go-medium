# Social Network Pet Project

## Structure of project 

- api 
- cmd 
  - main.go
- config 
- errs 
- migrations
- models 
- rest 
  - handlers
  - server.go
- storage
  - repo 
  - storage.go

---

## To run the project
1. Create database
2. Set enviromental variables
3. `go mod tidy`
4. `go run cmd/main.go`
5. `http://localhost:8080/swagger/index.html`


## Entities

1. User
2. Post
   1. Like
   2. Comment

3. Chat
   1. Message
   2. Participants