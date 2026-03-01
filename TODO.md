# TODO

## Remaining improvements

- [ ] Prevent destination collisions when multiple rules share the same basename in `storeDir`.
- [ ] Add tests:
- [ ] unit tests for config decode/validation.
- [ ] unit tests for selective backup and `--ignore` behavior.
- [ ] unit tests for destination path and ignore matching.
- [ ] integration tests for `backup --dry-run` and config loading.
- [ ] CLI polish:
- [ ] standardize help/usage output wording and examples.
- [ ] add a `list` subcommand to print configured rule names.
- [ ] Security hardening:
- [ ] warn when backing up sensitive paths to weakly-permissioned `storeDir`.
- [ ] optional allowlist/strict mode for `after-backup` commands.
