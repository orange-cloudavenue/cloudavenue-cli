## cav completion fish

Generate the autocompletion script for fish

### Synopsis

Generate the autocompletion script for the fish shell.

To load completions in your current shell session:

	cav completion fish | source

To load completions for every new session, execute once:

	cav completion fish > ~/.config/fish/completions/cav.fish

You will need to start a new shell for this setup to take effect.


```
cav completion fish [flags]
```

### Options

```
  -h, --help              help for fish
      --no-descriptions   disable completion descriptions
```

### Options inherited from parent commands

```
  -t, --time   time elapsed for command
```

### SEE ALSO

* [cav completion](cav_completion.md)	 - Generate the autocompletion script for the specified shell

