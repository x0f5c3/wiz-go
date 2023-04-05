# wiz-go
wiz-go

## Flags
|Flag|Usage|
|----|-----|
|`-d, --debug`|enable debug messages|
|`--no-update`|disables update checks|
|`--raw`|print unstyled raw output (set it if output is written to a file)|

## Commands
|Command|Usage|
|-------|-----|
|`wiz-go completion`|Generate the autocompletion script for the specified shell|
|`wiz-go discover`||
|`wiz-go help`|Help about any command|
# ... completion
`wiz-go completion`

## Usage
> Generate the autocompletion script for the specified shell

wiz-go completion

## Description

```
Generate the autocompletion script for wiz-go for the specified shell.
See each sub-command's help for details on how to use the generated script.

```

## Commands
|Command|Usage|
|-------|-----|
|`wiz-go completion bash`|Generate the autocompletion script for bash|
|`wiz-go completion fish`|Generate the autocompletion script for fish|
|`wiz-go completion powershell`|Generate the autocompletion script for powershell|
|`wiz-go completion zsh`|Generate the autocompletion script for zsh|
# ... completion bash
`wiz-go completion bash`

## Usage
> Generate the autocompletion script for bash

wiz-go completion bash

## Description

```
Generate the autocompletion script for the bash shell.

This script depends on the 'bash-completion' package.
If it is not installed already, you can install it via your OS's package manager.

To load completions in your current shell session:

	source <(wiz-go completion bash)

To load completions for every new session, execute once:

#### Linux:

	wiz-go completion bash > /etc/bash_completion.d/wiz-go

#### macOS:

	wiz-go completion bash > $(brew --prefix)/etc/bash_completion.d/wiz-go

You will need to start a new shell for this setup to take effect.

```

## Flags
|Flag|Usage|
|----|-----|
|`--no-descriptions`|disable completion descriptions|
# ... completion fish
`wiz-go completion fish`

## Usage
> Generate the autocompletion script for fish

wiz-go completion fish

## Description

```
Generate the autocompletion script for the fish shell.

To load completions in your current shell session:

	wiz-go completion fish | source

To load completions for every new session, execute once:

	wiz-go completion fish > ~/.config/fish/completions/wiz-go.fish

You will need to start a new shell for this setup to take effect.

```

## Flags
|Flag|Usage|
|----|-----|
|`--no-descriptions`|disable completion descriptions|
# ... completion powershell
`wiz-go completion powershell`

## Usage
> Generate the autocompletion script for powershell

wiz-go completion powershell

## Description

```
Generate the autocompletion script for powershell.

To load completions in your current shell session:

	wiz-go completion powershell | Out-String | Invoke-Expression

To load completions for every new session, add the output of the above command
to your powershell profile.

```

## Flags
|Flag|Usage|
|----|-----|
|`--no-descriptions`|disable completion descriptions|
# ... completion zsh
`wiz-go completion zsh`

## Usage
> Generate the autocompletion script for zsh

wiz-go completion zsh

## Description

```
Generate the autocompletion script for the zsh shell.

If shell completion is not already enabled in your environment you will need
to enable it.  You can execute the following once:

	echo "autoload -U compinit; compinit" >> ~/.zshrc

To load completions in your current shell session:

	source <(wiz-go completion zsh); compdef _wiz-go wiz-go

To load completions for every new session, execute once:

#### Linux:

	wiz-go completion zsh > "${fpath[1]}/_wiz-go"

#### macOS:

	wiz-go completion zsh > $(brew --prefix)/share/zsh/site-functions/_wiz-go

You will need to start a new shell for this setup to take effect.

```

## Flags
|Flag|Usage|
|----|-----|
|`--no-descriptions`|disable completion descriptions|
# ... discover
`wiz-go discover`
wiz-go discover
# ... help
`wiz-go help`

## Usage
> Help about any command

wiz-go help [command]

## Description

```
Help provides help for any command in the application.
Simply type wiz-go help [path to command] for full details.
```


---
> **Documentation automatically generated with [PTerm](https://github.com/pterm/cli-template) on 05 April 2023**
