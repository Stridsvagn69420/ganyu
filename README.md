# Ganyu
A tool to standardize commands across Linux distros and Windows

# Installation
## Downloading
### APT
WIP.

### Pacman
WIP.

### Kagero
WIP.

### DNF
WIP.

## Preparing
### Config file
You need to create a config file like this at `~/.config/ganyu/config.json`:
```json
{
    "sysupdate": {
        "root": true,
        "crosspkg": true
    },
    "yt-dl": [
        {
            "audio": "251",
            "video": "248",
            "audio+video": "22",
            "website": "www.youtube.com"
        },
        {
            "audio": "0",
            "video": "0",
            "audio+video": "0",
            "website": "www.instagram.com"
        },
        {
            "audio": "http-2176",
            "video": "http-2176",
            "audio+video": "http-2176",
            "website": "twitter.com"
        }
    ]
}
```
These are default settings, so change them how you'd like. See more [in the wiki](https://github.com/Stridsvagn69420/ganyu/wiki/Config).

# Usage
WIP.