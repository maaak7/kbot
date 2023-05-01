# kbot
It's a simple [telegram bot](https://t.me/maaak7_bot) that knows a few commands and is able to reply to you.

## Setup
First of all, you must have a Telegram account, because you'll need to [create your own bot](https://core.telegram.org/bots/tutorial) there.

### Setup Telegram Bot Token
- Open terminal and run `read -s TELE_TOKEN`
- Copy token from @GodFather message and paste it in terminal
- Run command `export TELE_TOKEN`

### Create Go build
- Open terminal and run `go build -ldflags "-X="github.com/<your-github-account>/<repository-name>/cmd.appVersion=<your-app-version>`

## Instruction
There are three commands that the bot knows: **hello**, **date**, **day**. To run one of them, you need to enter `/start <command_name>`. For example - `/start date`.