# `dotback`

This tool is used to backup your dotfiles.

Instead of making confusing and outright dangerous symlinks, this simply backs up your dotfiles using a `toml` file.

## Usage

### Backup

In it's simplest form, running `dotback backup` would backup all the files into a specified folder `storeDir`.

This ignores `.git` and `.gitignore` files within any backup folder. This is to avoid any submodules
caveats appearing when it.

### Restore

You can run the restore command to reverse the backup (Copy the backed up file to the original location) thus erasing any changes made at source.

## Example

Assuming this as a example config for `dotback`.

```toml

storeDir = "~/.local/state/backup"
after-backup = [
    "scp -r ...",
    "rsync ....",
    "tar -xvzf ..."
]

[backup.nvim]
location = "~/.config/nvim"

[backup.neomutt]
location = "~/.config/neomutt"

[backup.gitconfig]
location = "~/.gitconfig"

```

When you run 
```sh
dotback backup
```
This would backup the all the above mentioned locations to the configured location. 

## It does not
- Provide incremental backup (Use rsync or borg)
- Provide automatic backup (Use systemd or cron for that) 
- Sync with different machines (Use syncthing for that)
- Install your dotfiles on a new machine.

## It does
- Copy the file or folder into the `backupLocation`.
- Provide ability to run scripts before and after running the backups.

## Configuration

It tries to find the default configuration in `$XDG_CONFIG_HOME/dotback/backup.yml`. This file should have the below format.

```
[backup.name-of-backup]
location = "<location-of-backup-file/folder>
```

- It should also have a `storeLocation` variable defined. 
- It can optionally provide a `after-backup` list of commands to run.

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
