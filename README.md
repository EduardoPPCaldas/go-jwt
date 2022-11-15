# Go JWT
This is a simple way to use JWT authentication in go

## Setup
for this application to run you need a postgres instance running, for that execute this commands:

```bash
docker pull postgres
```
and

```bash
docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -e POSTGRES_DB=gojwt -d postgres
```

create a .env file and fill with the information you need using the .env.example scaffold.

Then run 
```bash
go mod tidy
```
to download all required packages
## Run

to run this application with live reload run

```bash
CompileDaemon -command="./go-jwt"
```