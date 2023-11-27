# `dotback`

This tool is used to backup your dotfiles.

Instead of making confusing and outright dangerous symlinks, this simply backs up your dotfiles using a `toml` file.

## Example

Assuming this as a example config for `dotback`.

```toml

backupLocation = "~/.local/state/backup"

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

It is a very highly opinated tool. It does not fit in everybody's workflow.
But if it does, please do share your support and love.

Anybody is welcome to contribute and extend this project.

