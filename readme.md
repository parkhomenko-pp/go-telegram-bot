# Go Telegram bot

[![Go](https://github.com/parkhomenko-pp/go-telegram-bot/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/parkhomenko-pp/go-telegram-bot/actions/workflows/go.yml?query=branch:master)

<img src="preview/icon.png" align="right" width=120 height=120/>

This is a Go Telegram bot project. It includes a bot that interacts with users on Telegram and a console game. Below you will find instructions on how to run the project and the current roadmap for future development.
<br><br>

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