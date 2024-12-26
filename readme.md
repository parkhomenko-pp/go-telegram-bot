# Go Telegram bot

[![Go](https://github.com/parkhomenko-pp/go-telegram-bot/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/parkhomenko-pp/go-telegram-bot/actions/workflows/go.yml?query=branch:master)

<img src="preview/icon.png" align="right" width=150 height=150/>

This is a Go game project that allows two players to play Go via a Telegram bot. It also includes support for playing the game from the console. Below you will find instructions on how to run the project and the current roadmap for future development.
<br><br><br>

## Run project

```sh
go run src/bot/main.go # start telegram bot
```

```sh
go run src/console/main.go # start console game
```

## Roadmap
- Goban
    - [x] Themes support
    - [x] Place stones on board
    - [ ] Stones without breath determine
    - [ ] Stones without breath remove from goban
    - [ ] Captured areas count
    - [ ] Captured stones count
- Console
    - [ ] Full game support
- Bot
    - TODO