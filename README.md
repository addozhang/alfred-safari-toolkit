# Alfred Safari Toolkit

This is an [Alfred](https://www.alfredapp.com) workflow for Safari.

I have been using [tupton/alfred-safari-history](https://github.com/tupton/alfred-safari-history) for years, but it runs with Python2 which is removed from macOS 12.3.

So I decided to write one in Go by referring to it.

![](screenshot1.gif)

## How to install

Download the workflow from [release page](https://github.com/addozhang/alfred-safari-toolkit/releases) and double click.

Make sure Alfred has full disk access otherwise it won't be able to access history.db file.

## How to build

It's recommended to [go-alfred](https://github.com/jason0x43/go-alfred) for workflow packaging.

1. First, install is by executing `go install github.com/jason0x43/go-alfred/alfred@latest`.
2. After running `CGO_ENABLED=1 alfred build` to build project, you will get the execution binary under `workflow` folder. 
3. At last, run `alfred pack` and the workflow package will present in root folder. 

## Versioning

The current version covers tupton/alfred-safari-history features. In the next version, planning to involve in the Safari tabs searching feature.

## Platform

The latest version has been tested below platform:

* macOS 12.3.1 with Apple Silicon CPU
* macOS 10.15.5, 12.4 with Intel CPU

If it works or not on other platforms, I'm very glad to know your feedback.

## FAQ

### Why use `CGO_ENABLED=1`?

This workflow depends on sqlite3 to query history from sqlist3 file. sqlite3 requires CGO support. 
