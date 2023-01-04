# bilisubdl

## Examples

Download 1049041 with english language

`bilisubdl dl 1049041 -l en`

Download 37738 and 1042594 with thai language

`bilisubdl dl 37738 1042594 -l th`

List 37738 subtitle language

`bilisubdl list 37738 -L`

Show today timeline

`bilisubdl timeline`

Show monday timeline

`bilisubdl timeline mon`

## Usage

```txt
Usage:
  bilisubdl [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  dl          Download subtitle from ID.
  help        Help about any command
  list        Show info
  search      Search anime
  timeline    Show timeline (sun|mon|tue|wed|thu|fri|sat)

Flags:
  -h, --help      help for bilisubdl
  -v, --version   version for bilisubdl

Use "bilisubdl [command] --help" for more information about a command.
```

## Installing

The `bilisubdl` command on Windows using [Scoop](https://scoop.sh/)

```txt
scoop install https://raw.githubusercontent.com/K0ng2/scoop-bucket/main/bucket/bilisubdl.json
```
