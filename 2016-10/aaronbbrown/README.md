[![asciicast](https://asciinema.org/a/89391.png)](https://asciinema.org/a/89391)

## Requirements:

* Docker for Mac or docker-engine

## Building

```
script/dockerbuild
```

## Run the server

```
GAMES=5 PORT=5555 script/server
```


## Run the client

```
ADDRESS="tcp://1.2.3.4:5555" script/client
```

