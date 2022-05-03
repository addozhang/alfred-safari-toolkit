# Alfred Safari Toolkit

This is an [Alfred](https://www.alfredapp.com) workflow for Safari.

I have been using [tupton/alfred-safari-history](https://github.com/tupton/alfred-safari-history) for years, but it runs with Python2 which is removed from macOS 12.3.

So I decided to write one in Go by referring to it.

![](screenshot1.gif)

## How to install

Download the workflow from [release page](https://github.com/addozhang/alfred-safari-toolkit/releases) and double click.

## How to build

It's recommended to [go-alfred](https://github.com/jason0x43/go-alfred) for workflow packaging.

First, install is by executing `go install github.com/jason0x43/go-alfred/alfred@latest`.
After running `alfred pack`, you will find the workflow package under `workflow` folder. 

## Versioning

The current version is 1.0.0 which covers the above one. In the next version, planning to involve in the Safari tabs searching feature.
