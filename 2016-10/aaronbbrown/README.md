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

The client offers alternative strategies for playing the game:

* **random** (Default) - all throws are random
* **mirrorlast** - Client (Me) will repeat the opponent's (You) last move
* **mirrorwinner** - Client (Me) will repeat the previous winner's move

To pick a strategy, pass the `STRATEGY=strategy` env variable. e.g.:

```
STRATEGY=mirrorwinner ADDRESS="tcp://1.2.3.4:5555" script/client
```
