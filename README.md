<h1 align="center">
    Discord Downloader <i>Go</i>
</h1>
<p align="center">
    <a href="https://travis-ci.com/get-got/discord-downloader-go" alt="Travis Build">
        <img src="https://travis-ci.com/get-got/discord-downloader-go.svg?branch=master" />
    </a>
    <a href="https://hub.docker.com/r/getgot/discord-downloader-go" alt="Docker Build">
        <img src="https://img.shields.io/docker/cloud/build/getgot/discord-downloader-go" />
    </a>
    <a href="https://goreportcard.com/report/github.com/get-got/discord-downloader-go" alt="Go Report Card">
        <img src="https://goreportcard.com/badge/github.com/get-got/discord-downloader-go" />
    </a>
    <br>
    <a href="https://github.com/get-got/discord-downloader-go/releases" alt="All Releases">
        <img src="https://img.shields.io/github/downloads/get-got/discord-downloader-go/total?label=all-releases&logo=GitHub" />
    </a>
    <a href="https://hub.docker.com/r/getgot/discord-downloader-go" alt="Docker Pulls">
        <img src="https://img.shields.io/docker/pulls/getgot/discord-downloader-go?label=docker-pulls&logo=Docker" />
    </a>
    <a href="https://github.com/get-got/discord-downloader-go/releases/latest" alt="Latest Release">
        <img src="https://img.shields.io/github/downloads/get-got/discord-downloader-go/latest/total?label=latest-release&logo=GitHub" />
    </a>
    <br>
    <a href="https://discord.gg/6Z6FJZVaDV">
        <img src="https://img.shields.io/discord/780985109608005703?logo=discord"alt="Join the Discord">
    </a>
</p>
<h2 align="center">
    <a href="https://discord.com/invite/6Z6FJZVaDV">
        <b>Need help? Have suggestions? Join the Discord server!</b>
    </a>
    <br/><br/>
    <a href="https://github.com/get-got/discord-downloader-go/releases/latest">
        <b>DOWNLOAD LATEST RELEASE</b>
    </a>
</h2>
<div align="center">

| Operating System  | Architectures _( ? = available but untested )_    |
| -----------------:|:----------------------------------------------- |
| Windows           | **amd64**, arm64 _(?)_, armv7/6/5 _(?)_, 386 _(?)_
| Linux             | **amd64**, **arm64**, **armv7/6/5**,<br/>risc-v64 _(?)_, mips64/64le _(?)_, s390x _(?)_, 386 _(?)_
| Darwin (Mac)      | amd64 _(?)_, arm64 _(?)_
| FreeBSD           | amd64 _(?)_, arm64 _(?)_, armv7/6/5 _(?)_, 386 _(?)_
| OpenBSD           | amd64 _(?)_, arm64 _(?)_, armv7/6/5 _(?)_, 386 _(?)_
| NetBSD            | amd64 _(?)_, arm64 _(?)_, armv7/6/5 _(?)_, 386 _(?)_

</div><br>

This project is a cross-platform cli single-file program to interact with a Discord Bot (genuine bot application or user account, limitations apply to both respectively) to locally download files posted from Discord in real-time as well as a full archive of old messages. It can download any directly sent Discord attachments or linked files and supports fetching highest possible quality files from specific sources _([see list below](#supported-download-sources))._ It also supports **very extensive** settings configurations and customization, applicable globally or per-server/category/channel/user. Tailor the bot to your exact needs and runtime environment. See the [Features](#-features) list below for the full list. See the [List of Settings](#-list-of-settings) below for a settings breakdown. See [Getting Started](#%EF%B8%8F-getting-started) or anything else in the table of contents right under this to learn more!

<h3 align="center">
    <b>Originally a fork of <a href="https://github.com/Seklfreak/discord-image-downloader-go">Seklfreak's <i>discord-image-downloader-go</i></a></b>
</h3>
<h4 align="center">
    The original project was abandoned, for a list of differences and why I made an independent project, <a href="#differences-from-seklfreaks-discord-image-downloader-go--why-i-made-this"><b>see below</b></a>
</h4>

---

- [⚠️ **WARNING!** Discord does not allow Automated User Accounts (Self-Bots/User-Bots)](#️-warning-discord-does-not-allow-automated-user-accounts-self-botsuser-bots)
- [🤖 Features](#-features)
  - [Supported Download Sources](#supported-download-sources)
  - [Commands](#commands)
- [✔️ Getting Started](#️-getting-started)
  - [Getting Started Step-by-Step](#getting-started-step-by-step)
  - [Bot Login Credentials](#bot-login-credentials)
  - [Permissions in Discord](#permissions-in-discord)
    - [Bot Intents for Discord Application Bots](#bot-intents-for-discord-application-bots)
      - [NECESSARY IF USING A GENUINE DISCORD APPLICATION](#necessary-if-using-a-genuine-discord-application)
  - [How to Find Discord IDs](#how-to-find-discord-ids)
  - [Differences from Seklfreak's _discord-image-downloader-go_ \& Why I made this](#differences-from-seklfreaks-discord-image-downloader-go--why-i-made-this)
- [📚 Guide: Downloading History (Old Messages)](#-guide-downloading-history-old-messages)
  - [Command Arguments](#command-arguments)
    - [Examples](#examples)
- [🔨 Guide: Settings / Configuration](#-guide-settings--configuration)
- [🛠 List of Settings](#-list-of-settings)
  - [🛠 ... UNDER CONSTRUCTION ... 🔨](#--under-construction--)
  - [ALL SETTINGS WITH DEFAULT VALUES](#all-settings-with-default-values)
  - [settings - main](#settings---main)
  - [settings - credentials group](#settings---credentials-group)
  - [settings - adminChannels group](#settings---adminchannels-group)
  - [settings - source group](#settings---source-group)
  - [settings - source / filters group](#settings---source--filters-group)
  - [settings - source / log group](#settings---source--log-group)
- [🤖 Settings Examples](#-settings-examples)
  - [example - minimum bot app](#example---minimum-bot-app)
  - [example - minimum user account without 2FA](#example---minimum-user-account-without-2fa)
  - [example - minimum user account with 2FA](#example---minimum-user-account-with-2fa)
  - [example - server with friends](#example---server-with-friends)
  - [example - scraping public servers (user)](#example---scraping-public-servers-user)
  - [example - scraping public server (bot app, as admin)](#example---scraping-public-server-bot-app-as-admin)
- [❔ FAQ](#-faq)
- [⚙️ Development](#️-development)

---

## ⚠️ **WARNING!** Discord does not allow Automated User Accounts (Self-Bots/User-Bots)

[Read more in Discord Trust & Safety Team's Official Statement...](https://support.discordapp.com/hc/en-us/articles/115002192352-Automated-user-accounts-self-bots-)

While this project works for user logins, I do not reccomend it as you risk account termination. If you can, [use a proper Discord Bot user for this program.](https://discord.com/developers/applications)

> _NOTE: This only applies to real User Accounts, not Bot users. This program currently works for either._

Now that that's out of the way...

---

## 🤖 Features

### Supported Download Sources

- Direct Links to Files
- Discord File Attachments
- Twitter _(requires API key, see config section)_
- Instagram _(requires account login, see config section)_
- Reddit
- Imgur
- Streamable
- Gfycat
- Tistory
- Flickr _(requires API key, see config section)_
- _I'll always welcome requests but some sources can be tricky to parse..._
  
### Commands

Commands are used as `ddg <command> <?arguments?>` _(unless you've changed the prefix)_
Command     | Arguments? | Description
---         | ---   | ---
`help`, `commands`  | No    | Lists all commands.
`ping`, `test`      | No    | Pings the bot.
`info`      | No    | Displays relevant Discord info.
`status`    | No    | Shows the status of the bot.
`stats`     | No    | Shows channel stats.
`history`   | [**SEE HISTORY SECTION**](#-guide-downloading-history-old-messages) | **(BOT AND SERVER ADMINS ONLY)** Processes history for old messages in channel.
`exit`, `kill`, `reload`    | No    | **(BOT ADMINS ONLY)** Exits the bot _(or restarts if using a keep-alive process manager)_.
`emojis`    | Optionally specify server IDs to download emojis from; separate by commas | **(BOT ADMINS ONLY)** Saves all emojis for channel.

---

## ✔️ Getting Started

_Confused? Try looking at [the step-by-step list.](#getting-started-step-by-step)_

Depending on your purpose for this program, there are various ways you can run it.

- [Run the executable file for your platform. _(Process managers like **pm2** work well for this)_](https://github.com/get-got/discord-downloader-go/releases/latest)
- [Run the executable file via command prompt. _(`discord-downloader-go.exe settings2` or similar to run multiple instances sharing a database with separate settings files)_](https://github.com/get-got/discord-downloader-go/releases/latest)
- [Run automated image builds in Docker.](https://hub.docker.com/r/getgot/discord-downloader-go) _(Google it)._
  - Mount your settings.json to ``/root/settings.json``
  - Mount a folder named "database" to ``/root/database``
  - Mount your save folders or the parent of your save folders within ``/root/``
    - _i.e. ``X:\My Folder`` to ``/root/My Folder``_
- Install Golang and compile/run the source code yourself. _(Google it)_

You can either create a `settings.json` following the examples & variables listed below, or have the program create a default file (if it is missing when you run the program, it will make one, and ask you if you want to enter in basic info for the new file).

- [Ensure you follow proper JSON syntax to avoid any unexpected errors.](https://www.w3schools.com/js/js_json_syntax.asp)
- [Having issues? Try this JSON Validator to ensure it's correctly formatted.](https://jsonformatter.curiousconcept.com/)

[![Tutorial Video](http://img.youtube.com/vi/06UUXDQ80f8/0.jpg)](http://www.youtube.com/watch?v=06UUXDQ80f8)

### Getting Started Step-by-Step

1. Download & put executable within it's own folder.
2. Configure Main Settings (or run once to have settings generated). [_(SEE BELOW)_](#-list-of-settings)
3. Enter your login credentials in the `"credentials"` section. [_(SEE BELOW)_](#-list-of-settings)
4. Put your Discord User ID as in the `"admins"` list of the settings. [_(SEE BELOW)_](#-list-of-settings)
5. Put a Discord Channel ID for a private channel you have access to into the `"adminChannels"`. [_(SEE BELOW)_](#-list-of-settings)
6. Put your desired Discord Channel IDs into th6e `"channels"` section. [_(SEE BELOW)_](#-list-of-settings)
   - I know it can be confusing if you don't have experience with programming or JSON in general, but this was the ideal setup for extensive configuration like this. Just be careful with comma & quote placement and you should be fine. [See examples below for help.](#-settings-examples)

### Bot Login Credentials

- If using a **Bot Application,** enter the token into the `"token"` setting. Remove the lines for `"username"` and `"password"` or leave blank (`""`). **To create a Bot User,** go to [discord.com/developers/applications](https://discord.com/developers/applications) and create a `New Application`. Once created, go to `Bot` and create. The token can be found on the `Bot` page. To invite to your server(s), go to `OAuth2` and check `"bot"`, copy the url, paste into browser and follow prompts for adding to server(s).
- If using a **User Account (Self-Bot) WITHOUT 2FA (2-Factor Authentication),** fill out the `"username"` and `"password"` settings. Remove the line for `"token"` or leave blank (`""`).
- If using a **User Account (Self-Bot) WITH 2FA (2-Factor Authentication),** enter the token into the `"token"` setting. Remove the lines for `"username"` and `"password"` or leave blank (`""`). Your account token can be found by `opening the browser console / dev tools / inspect element` > `Network` tab > `filter for "library"` and reload the page if nothing appears. Assuming there is an item that looks like the below screenshot, click it and `find "Authorization" within the "Request Headers" section` of the Headers tab. The random text is your token.

<img src="https://i.imgur.com/2BdaJSH.png"> <img src="https://i.imgur.com/i9DItcH.png">

### Permissions in Discord

- In order to perform basic downloading functions, the bot will need `Read Message` permissions in the server(s) of your designated channel(s).
- In order to respond to commands, the bot will need `Send Message` permissions in the server(s) of your designated channel(s). If executing commands via an Admin Channel, the bot will only need `Send Message` permissions for that channel, and that permission will not be required for the source channel.
- In order to process history commands, the bot will need `Read Message History` permissions in the server(s) of your designated channel(s).

#### Bot Intents for Discord Application Bots

##### NECESSARY IF USING A GENUINE DISCORD APPLICATION

- Go to the Discord Application management page, choose your application, go to the `Bot` category, and ensure `Message Content Intent` is enabled.

<img src="https://i.imgur.com/2GcyA2B.png"/>

### How to Find Discord IDs

- **_Use the info command!_**
- **Discord Developer Mode:** Enable `Developer Mode` in Discord settings under `Appearance`.
- **Finding Channel ID:** _Enable Discord Developer Mode (see above),_ right click on the channel and `Copy ID`.
- **Finding User ID:** _Enable Discord Developer Mode (see above),_ right click on the user and `Copy ID`.
- **Finding Emoji ID:** _Enable Discord Developer Mode (see above),_ right click on the emoji and `Copy ID`.
- **Finding DM/PM ID:** Inspect Element on the DM icon for the desired user. Look for `href="/channels/@me/CHANNEL_ID_HERE"`. Using this ID in place of a normal channel ID should work perfectly fine.

---

### Differences from [Seklfreak's _discord-image-downloader-go_](https://github.com/Seklfreak/discord-image-downloader-go) & Why I made this

- _Better command formatting & support_
- Configuration is JSON-based rather than ini to allow more elaborate settings and better organization. With this came many features such as channel-specific settings.
- Channel-specific control of downloaded filetypes / content types (considers things like .mov as videos as well, rather than ignore them), Optional dividing of content types into separate folders.
- (Optional) Reactions upon download success.
- (Optional) Discord messages upon encountered errors.
- Extensive bot status/presence customization.
- Consistent Log Formatting, Color-Coded Logging
- Somewhat different organization than original project; initially created from scratch then components ported over.
- _Various fixes, improvements, and dependency updates that I also contributed to Seklfreak's original project._

> I've been a user of Seklfreak's project since ~2018 and it's been great for my uses, but there were certain aspects I wanted to expand upon, one of those being customization of channel configuration, and other features like message reactions upon success, differently formatted statuses, etc. If some aspects are rudimentary or messy, please make a pull request, as this is my first project using Go and I've learned everything from observation & Stack Overflow.

---

## 📚 Guide: Downloading History (Old Messages)

> This guide is to show you how to make the bot go through all old messages in a channel and catalog them as though they were being sent right now, in order to download them all.

### Command Arguments

If no channel IDs are specified, it will try and use the channel ID for the channel you're using the command in.

Argument / Flag         | Details
---                     | ---
**channel ID(s)**       | One or more channel IDs, separated by commas if multiple.
`all`                   | Use all available registered channels.
`cancel` or `stop`      | Stop downloading history for specified channel(s).
`list` or `status`      | Output running history jobs in Discord & program.
`--since=YYYY-MM-DD`    | Will process messages sent after this date.
`--since=message_id`    | Will process messages sent after this message.
`--before=YYYY-MM-DD`   | Will process messages sent before this date.
`--before=message_id`   | Will process messages sent before this message.

**_Order of arguments does not matter_**

#### Examples

- `ddg history`
- `ddg history cancel`
- `ddg history all`
- `ddg history stop all`
- `ddg history 000111000111000`
- `ddg history 000111000111000, 000222000222000`
- `ddg history 000111000111000,000222000222000,000333000333000`
- `ddg history 000111000111000, 000333000333000 cancel`
- `ddg history 000111000111000 --before=000555000555000`
- `ddg history 000111000111000 --since=2020-01-02`
- `ddg history 000111000111000 --since=2020-10-12 --before=2021-05-06`
- `ddg history 000111000111000 --since=000555000555000 --before=2021-05-06`
- `ddg history status`
- `ddg history list`

---

## 🔨 Guide: Settings / Configuration

> I tried to make the configuration as user friendly as possible, though you still need to follow proper JSON syntax **(watch those commas)**. Most settings are optional and will use default values or be unused if missing from your settings file.

When initially launching the bot it will create a default settings file if you do not create your own `settings.json` manually. All JSON settings follow camelCase format.

**If you have a ``config.ini`` from _Seklfreak's discord-image-downloader-go_, it will import settings if it's in the same folder as the program.**

The bot accepts `.json` or `.jsonc` for comment-friendly json.

---

## 🛠 List of Settings

### 🛠 ... UNDER CONSTRUCTION ... 🔨

### ALL SETTINGS WITH DEFAULT VALUES

```json
{
    "credentials": {
        "token": "YOUR_USER_OR_BOT_TOKEN",
        "email": "YOUR_USER_EMAIL_NO_2FA",
        "password": "YOUR_USER_PASSWORD_NO_2FA",
        "twitterAccessToken": "",
        "twitterAccessTokenSecret": "",
        "twitterConsumerKey": "",
        "twitterConsumerSecret": "",
        "instagramUsername": "",
        "instagramPassword": "",
        "flickrApiKey": ""
    },

    "admins": [
        "YOUR_DISCORD_USER_ID"
    ],
    "adminChannels": [
        {
            "channel": "DISCORD_CHANNEL_ID_FOR_COMMANDS",
            "logProgram": false,
            "logStatus": true,
            "logErrors": true,
            "unlockCommands": false
        }
    ],

    "processLimit": 32,

    "debug": false,
    "settingsOutput": true,
    "messageOutput": true,
    "messageOutputHistory": false,

    "discordLogLevel": 0,
    "discordTimeout": 180,
    "downloadTimeout": 60,
    "downloadRetryMax": 3,
    "exitOnBadConnection": false,
    "githubUpdateChecking": true,

    "commandPrefix": "ddg ",
    "scanOwnMessages": false,
    "allowGeneralCommands": true,
    "inflateDownloadCount": 0,
    "europeanNumbers": false,

    "checkupRate": 30,
    "connectionCheckRate": 5,
    "presenceRefreshRate": 3,

    "save": true,
    "allowCommands": true,
    "scanEdits": true,
    "ignoreBots": true,

    "sendErrorMessages": true,
    "sendFileToChannel": "",
    "sendFileToChannels": [ "" ],
    "sendFileDirectly": true,
    "sendFileCaption": "",

    "filenameDateFormat": "2006-01-02_15-04-05",
    "filenameFormat": "{{date}} {{file}}",

    "presenceEnabled": true,
    "presenceStatus": "idle",
    "presenceType": 0,
    "presenceLabel": "{{timeSavedShort}} - {{countShort}} files",
    "presenceDetails": "{{timeSavedLong}}",
    "presenceDetails": "{{count}} files total",

    "reactWhenDownloaded": true,
    "reactWhenDownloadedEmoji": "",
    "reactWhenDownloadedHistory": false,
    "historyTyping": true,
    "embedColor": "",

    "historyMaxJobs": 3,
    "autoHistory": false,
    "autoHistoryBefore": "",
    "autoHistorySince": "",
    "sendAutoHistoryStatus": false,
    "sendHistoryStatus": true,

    "divideByYear": false,
    "divideByMonth": false,
    "divideByServer": false,
    "divideByChannel": false,
    "divideByUser": false,
    "divideByType": true,
    "divideFoldersUseID": false,
    "saveImages": true,
    "saveVideos": true,
    "saveAudioFiles": true,
    "saveTextFiles": false,
    "saveOtherFiles": false,
    "savePossibleDuplicates": true,
    "filters": {
        "blockedPhrases": [ "" ],
        "allowedPhrases": [ "" ],       
        "blockedUsers": [ "" ],
        "allowedUsers": [ "" ],       
        "blockedRoles": [ "" ],
        "allowedRoles": [ "" ],       
        "blockedExtensions": [
            ".htm",
            ".html",
            ".php",
            ".exe",
            ".dll",
            ".bin",
            ".cmd",
            ".sh",
            ".py",
            ".jar"
        ],
        "allowedExtensions": [ "" ],

        "blockedDomains": [ "" ],
        "allowedDomains": [ "" ]
    },

    "all": {
        "destination": "FOLLOW_CHANNELS_BELOW_FOR_REST"
    },
    "allBlacklistUsers": [ "" ],
    "allBlacklistServers": [ "" ],
    "allBlacklistCategories": [ "" ],
    "allBlacklistChannels": [ "" ],

    "users": [
        {
            "user": "SOURCE_DISCORD_USER_ID",
            "users": [ "SOURCE_DISCORD_USER_ID" ],
            "destination": "FOLLOW_CHANNELS_BELOW_FOR_REST"
        }
    ],

    "servers": [
        {
            "server": "SOURCE_DISCORD_SERVER_ID",
            "servers": [ "SOURCE_DISCORD_SERVER_ID" ],
            "serverBlacklist": [ "DISCORD_CHANNELS_TO_BLOCK" ],
            "destination": "FOLLOW_CHANNELS_BELOW_FOR_REST"
        }
    ],

    "categories": [
        {
            "category": "SOURCE_DISCORD_CATEGORY_ID",
            "categories": [ "SOURCE_DISCORD_CATEGORY_ID" ],
            "categoryBlacklist": [ "DISCORD_CHANNELS_TO_BLOCK" ],
            "destination": "FOLLOW_CHANNELS_BELOW_FOR_REST"
        }
    ],

    "channels": [
        {
            "channel": "SOURCE_DISCORD_CHANNEL_ID",
            "channels": [ "SOURCE_DISCORD_CHANNEL_ID" ],
            "destination": "files/example-folder",

            "enabled": true,
            "save": true,
            "allowCommands": true,
            "scanEdits": true,
            "ignoreBots": true,

            "sendErrorMessages": true,
            "sendFileToChannel": "",
            "sendFileToChannels": [ "" ],
            "sendFileDirectly": true,
            "sendFileCaption": "",

            "filenameDateFormat": "2006-01-02_15-04-05",
            "filenameFormat": "{{date}} {{file}}",

            "presenceEnabled": true,
            "reactWhenDownloaded": true,
            "reactWhenDownloadedEmoji": "",
            "reactWhenDownloadedHistory": false,
            "blacklistReactEmojis": [ "" ],
            "historyTyping": true,
            "embedColor": "",

            "autoHistory": false,
            "autoHistoryBefore": "",
            "autoHistorySince": "",
            "sendAutoHistoryStatus": false,
            "sendHistoryStatus": true,

            "divideByYear": false,
            "divideByMonth": false,
            "divideByServer": false,
            "divideByChannel": false,
            "divideByUser": false,
            "divideByType": true,
            "divideFoldersUseID": false,
            "saveImages": true,
            "saveVideos": true,
            "saveAudioFiles": true,
            "saveTextFiles": false,
            "saveOtherFiles": false,
            "savePossibleDuplicates": true,
            "filters": {
                "blockedPhrases": [ "" ],
                "allowedPhrases": [ "" ],

                "blockedUsers": [ "" ],
                "allowedUsers": [ "" ],

                "blockedRoles": [ "" ],
                "allowedRoles": [ "" ],

                "blockedExtensions": [
                    ".htm",
                    ".html",
                    ".php",
                    ".exe",
                    ".dll",
                    ".bin",
                    ".cmd",
                    ".sh",
                    ".py",
                    ".jar"
                ],
                "allowedExtensions": [ "" ],

                "blockedDomains": [ "" ],
                "allowedDomains": [ "" ]
            },

            "logLinks": {
                "destination": "",
                "destinationIsFolder": false,

                "divideLogsByServer": true,
                "divideLogsByChannel": true,
                "divideLogsByUser": false,
                "divideLogsByStatus": false,
                
                "logDownloads": true,
                "logFailures": true,

                "filterDuplicates": false,
                "prefix": "",
                "suffix": "",
                "userData": false
            },

            "logMessages": {
                "destination": "",
                "destinationIsFolder": false,

                "divideLogsByServer": true,
                "divideLogsByChannel": true,
                "divideLogsByUser": false,

                "filterDuplicates": false,
                "prefix": "",
                "suffix": "",
                "userData": false
            }

        }

    ]
}
```

### settings - main

| SETTING KEY         | TYPE                                    | DEFAULT    | DESCRIPTION                                                          | EXAMPLE                           |
| :-----------------: | --------------------------------------- | :--------: | -------------------------------------------------------------------- | --------------------------------- |
| credentials         | `credentials group`                     |            | See `credentials group` below.                                       |                                   |
| admins              | array of <br/>strings                   | None       | Discord IDs of users<br/> to use admin commands.                     | `"admins": [ "0", "0" ],`         |
| adminChannels       | array of <br/>`adminChannel groups`     | None       | See `adminChannel group` below.                                      |                                   |
| discordLogLevel     | int (whole number)                      | 0 (errors) | 0 = Errors, <br/>1 = Warning, <br/>2 = Informational, <br/>3 = Debug | `"discordLogLevel": 2,`           |
| debug         | boolean <br/>(true or false)            | false      | Enables extra output for narrowing down problems.                    | `"debug": true,`            |
| messageOutput       | boolean <br/>(true or false)            | true       | Enables discord message output.                                      | `"messageOutput": true,`          |

### settings - credentials group

### settings - adminChannels group

### settings - source group

### settings - source / filters group

### settings - source / log group

---

## 🤖 Settings Examples

TODO: UNDER CONSTRUCTION

### example - minimum bot app

### example - minimum user account without 2FA

### example - minimum user account with 2FA

### example - server with friends

### example - scraping public servers (user)

### example - scraping public server (bot app, as admin)

---

## ❔ FAQ

- **_Q: How do I install?_**
- **A: [SEE #getting-started](#%EF%B8%8F-getting-started)**

---

- **_Q: How do I convert from Seklfreak's discord-image-downloader-go?_**
- **A: Place your config.ini from that program in the same directory as this program and delete any settings.json file if present. The program will import your settings from the old project and make a new settings.json. It will still re-download files that DIDG already downloaded, as the database layout is different and the old database is not imported.**

---

## ⚙️ Development

- I'm a complete amateur with Golang. If anything's bad please make a pull request.
- Follows Semantic Versioning: `[MAJOR].[MINOR].[PATCH]` <https://semver.org/>
- [github.com/Seklfreak/discord-image-downloader-go - the original project this was founded on](https://github.com/Seklfreak/discord-image-downloader-go)
