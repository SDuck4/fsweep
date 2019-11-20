# fsweep - 🧹 Sweep old files
Sweep old files based on file's modified time and name.

## Install

### Binary
Download proper binary file from [github release page](https://github.com/SDuck4/fsweep/releases).

### Go
```sh
$ go get -u github.com/SDuck4/fsweep
```

## Usage
```sh
$ fsweep <path> <number-of-days> [flags]
```

### Arguments
- `<path>`: directory path to clean
- `<number-of-days>`: how old days files to delete

### Flags
- `-y`, `--assumeyes`: assume that the answer to any question which would be asked is yes
- `-h`, `--help`: help for fsweep
- `-n`, `--name <pattern>`: file name pattern to delete in regexp (default ".*")

## License

[MIT](./LICENSE)
