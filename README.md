### nourybot-go

Lidl Twitch Bot in development as a way to learn Go/APIs/Databases.

Commands: [here](https://gist.github.com/lyx0/161913eb719afacea578b47239d0d969)

#### Dockerbuild (unused)
```
docker build -t nourybot .
docker run -it --rm -p 5051:5051 nourybot
``` 

#### Running it normally
```
go run main.go
```
or:
```
go build .
./nourybot-go
```