# bilisubdl

## Examples


Download 1049041 with english language

`bilisubdl 1049041 -l en`

Download 37738 and 1042594 with thai language

`bilisubdl 37738 1042594 -l th`

List 37738 subtitle language

`bilisubdl 37738 --list-subs`

Show today timeline

`bilisubdl --timeline`

Show monday timeline

`bilisubdl --timeline=mon`

## Usage

```bash
Usage:
  bilisubdl [id] [flags]

Flags:
      --dlepisode                   Download subtitle from episode id
      --episode stringArray         Episode select (e.g. 5,8-10)
      --filename string             Set subtitle filename (e.g. Abc %d = Abc 1, Abc %02d = Abc 02)
                                    (This option only works in combination with --dlepisode flag) 
  -h, --help                        help for bilisubdl
  -l, --language string             Subtitle language to download (e.g. en)
  -L, --list-language               List available subtitle language
      --list-section                List available section
  -o, --output string               Set output directory (default "./")
  -w, --overwrite                   Force overwrite downloaded subtitles
  -s, --search string               Search anime
      --section stringArray         Section select (e.g. 5,8-10)
  -T, --timeline string[="today"]   Show timeline (sun|mon|tue|wed|thu|fri|sat)
  -v, --version                     version for bilisubdl
```
