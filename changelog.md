# Changelog

All notable changes to this project, will be documented here.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## Unreleased (v0.1.0)

### Added

- Added --config option
- Added --example-config option.
- Added --ignore option to backup subcommand.
- Added `install.sh` to install `dotback` into `/usr/local/bin` or `~/.local/bin`.
- Parse config and backup files accordingly
- Run after-backup procedure.
- Add selective backups
- Confirmation questions
- `config` subcommand to print config info.
- Help printing for subcommands.

### Changed

- Documentation updates in `README.md`:
- Fixed install commands to use `github.com/pspiagicw/dotback@latest`.
- Clarified actual backup behavior (`.git` and `.gitignore` are skipped during copy).
- Added config schema, usage examples, limitations, and safety notes.
- Renamed example config from `backup.yml` to `backup.example.toml`.
- CLI polish:
- Added `list` command to print configured backup rule names.
- Standardized help usage text and backup flag descriptions.
- Sorted displayed backup rules in `dotback config`.
- Installer updates:
- `install.sh` now installs prebuilt release binaries (no local clone/build required).
- Supports optional `VERSION=<tag>` pinning during install.
