# nourybot-go

Lidl Twitch Bot in development

Commands: [here](https://gist.github.com/lyx0/161913eb719afacea578b47239d0d969)

### Build it with Docker

```
docker build -t nourybot .
docker run -it --rm -p 5051:5051 nourybot
``` 

### Run it normally
```
go run main.go
```