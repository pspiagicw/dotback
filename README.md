# `dotback`

`dotback` backs up your dotfiles.

Instead of making confusing and outright dangerous symlinks, this simply backs up your dotfiles using a `toml` file.

<!-- TOC start (generated with https://github.com/derlin/bitdowntoc) -->

- [`dotback`](#dotback)
   * [Features](#features)
   * [Configuration](#configuration)
   * [Usage](#usage)
      + [Backup](#backup)
   * [Unix philosophy](#unix-philosophy)
   * [It does not](#it-does-not)
   * [It does](#it-does)
   * [Installation](#installation)
   * [Contribution](#contribution)

<!-- TOC end -->

## Features

- A single binary, no runtime environments or large dependencies.
- Designed with Unix philosophy, it only does backup.

## Configuration

It tries to find the default configuration in `$XDG_CONFIG_HOME/dotback/backup.toml`.
- It should also have a `storeLocation` variable defined. 
- It can optionally provide a `after-backup` list of commands to run.

A example file.

```toml

# A folder to store the backup, it will be created if it does not exist.
storeDir = "~/.local/state/backup"

# All commands should be defined by the user.
# It can be left empty or omitted.

after-backup = [
    "scp -r ...",
    "rsync ....",
    "tar -xvzf ..."
]


# A backup rule has [backup.<rule-name>] format.
# It should contain a `location` parameter.
[backup.nvim]
location = "~/.config/nvim"

[backup.neomutt]
location = "~/.config/neomutt"

# Backup location can also be a file.
[backup.gitconfig]
location = "~/.gitconfig"

```

</br>

> You can run `dotback config` to get info about the current config.

</br>

![config](./gifs/config.gif)

## Usage

### Backup

Simply run `dotback backup` to backup configured in the config file. It would backup everything in `storeDir`.


> This ignores `.git` and `.gitignore` files within any backup folder.

</br>

![demo](./gifs/backup.gif)

</br>

> You can also specify selective rules for backup. This only backs up `nvim`.

</br>
```sh
dotback backup nvim
```

> You can provide `--dry-run` flag to not actually backup anything.

## Unix philosophy

`dotback` is not a full fledged backup solution for your dotfiles. It is a simple tool with a simple use.

## It does not
- Provide incremental backup (Use rsync or borg)
- Provide automatic backup (Use systemd or cron for that) 
- Sync with different machines (Use syncthing for that)
- Install your dotfiles on a new machine.

## It does
- Copy the file or folder into the `backupLocation`.
- Provide ability to run scripts before and after running the backups.

You can see the backup.yml example file(in the repo) for more reference.

## Installation

You can either use `go` or [`gox`](https://github.com/pspiagicw/gox) to automatically install the binary.

```
go install github.com/pspiagicw/gox@latest

# Much better 

gox install github.com/pspiagicw/gox@latest
```

Or download the binary from the [releases](https://github.com/pspiagicw/dotback/releases) section.

## Contribution

It is a very highly opinated tool. It does not fit in everybody's workflow.
But if it does, please do share your support and love.

Anybody is welcome to contribute and extend this project. 
