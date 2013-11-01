# image_analyzer

A simple tool that checks the dominant color percentage in an image.

It can be used to check if an image is probably a illustration rather than a
photograph for example.

## Installation

``` bash
go build
```

optionally move it to `/usr/local/bin/`

## Usage

Analyze a specific image:
``` bash
$ image_analyzer image.jpg
```

Run on the whole directory:
``` bash
$ image_analyzer
```

