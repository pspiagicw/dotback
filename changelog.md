# Changelog

All notable changes to this project, will be documented here.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## Unreleased (v0.1.0)

### Added

- Added --config option
- Added --example-config option.
- Added --ignore option to backup subcommand.
- Parse config and backup files accordingly
- Run after-backup procedure.
- Add selective backups
- Confirmation questions
- `config` subcommand to print config info.
- Help printing for subcommands.

### Changed

- Documentation updates in `README.md`:
- Fixed install commands to use `github.com/pspiagicw/dotback@latest`.
- Clarified actual backup behavior (`.git` skipped, `.gitignore` not skipped by default).
- Added config schema, usage examples, limitations, and safety notes.
- Renamed example config from `backup.yml` to `backup.example.toml`.
