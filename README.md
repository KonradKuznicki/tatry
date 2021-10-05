# Tatry
Due to COVID-19 lock down I created this tool help me stay sane at home.

This tool is meant to update background with images from live cams, 

by default live cams from Polish mountains Tatry.

`currently runs on Ubuntu with Gnome only but pull requests are welcome :) `

## Usage
Fairly simple:
```bash
$ ./tatry
``` 

## Configuration
standard Cobra and Viper with env vars support

for list of possible options run with `-h`

default config file is stored in `$HOME/.tatry` 
remember to add proper file extension

available config standards JSON, TOML, YAML, HCL, envfile and Java properties config files

### custom config in .tatry.yaml
```yaml
URLs:
 - https://cache.bieszczady.live/thumbnails/portSolina-720x405.jpg
```


## Compilation
Go 1.14
```
$ go build
```

## Tests
Go 1.14
```
$ go test ./...
```
may fail due to race conditions while testing OS integration

to avoid that run with -short flag to skip them 

```
$ go test ./... -short
```

## Project Future

### Additional features
option to create collage from selected images instead of rotating them

### Other OSs
I plan to add support for MacOS nad Windows <br />
and support packets for Ubuntu
