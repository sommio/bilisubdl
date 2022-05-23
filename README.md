# bilisubdl

## Examples

```bash
# Download 1049041 with english language
bilisubdl 1049041 -l en

# Download 37738 and 1042594 with thai language
bilisubdl 37738 1042594 -l th

# list 37738 subtitle language
bilisubdl 37738 --list-subs
```

## Usage

```bash
bilisubdl [id] [flags]

Flags:
  -h, --help              help for bilisubdl
  -l, --language string   Subtitle language to download (e.g. en)
      --list-subs         List available subtitles language
  -w, --overwrite         Force overwrite downloaded subtitles
  -v, --version           version for bilisubdl
```
