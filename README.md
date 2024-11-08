<img align="right" alt="DiscordGo logo" src="/docs/img/discordgo.svg" width="200">

## DiscordGo Direct Message Ping Pong Example

This example demonstrates how to utilize DiscordGo to create a Ping Pong Bot
that sends the response through Direct Message.

This Bot will respond to "ping" in any server it's in with "Pong!" in the
sender's DM.

**Join [Discord Gophers](https://discord.gg/0f1SbxBZjYoCtNPP)
Discord chat channel for support.**

### Build

This assumes you already have a working Go environment setup and that
DiscordGo is correctly installed on your system.

From within the dm_pingpong example folder, run the below command to compile the
example.

```sh
go build
```

### Usage

This example uses bot tokens for authentication only. While user/password is
supported by DiscordGo, it is not recommended for bots.

```
./dm_pingpong --help
Usage of ./dm_pingpong:
  -t string
        Bot Token
```

The below example shows how to start the bot

```sh
./dm_pingpong -t YOUR_BOT_TOKEN
Bot is now running.  Press CTRL-C to exit.
```


## En utilisant comme point de départ votre projet ou celui fait ensemble aujourd'hui :
https://github.com/O-clock-Nornes/s04-go-introduction/tree/player

-Implémenter la commande delqui permet à un Player de se supprimer (en demandant une confirmation)
-Implémenter une commande météo dans votre bot discord qui fait :
*répond "mais de quelle ville ?" si elle est donné telle quelle
*et pour par exemple météo Toulouse répond avec la méto de toulouse
*Et pour ceux qui veulent aller plus loin vous implémenter différentes fonctionnalités :
-implémenter un jeux de rôle en mode texte multi joueur
-implémenter l'api d'un service d'IA pour avoir un bot capable de tenir une conversation.