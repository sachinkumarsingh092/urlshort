# urlshort
Make your own short URLs. Use a customized YAML file to make the paths and URLs.

## Usage

- **Cloning**
```bash
$ git clone https://github.com/sachinkumarsingh092/urlshort
```

- **Verify Checksums**
```
$ go mod verify
```

- **Make binaries**
```bash
$ go build -x main.go
```

- **Run the binaries**
```bash
$ ./main [OPTIONS]
```

The program will start on port 8080 locally.

## Options
```
-yaml string
    Sets the YAML file to create URL map. (default "urls.yaml")
-help
    Print help.
```

## Documentations
Docs for handler package at [GoDoc](https://godoc.org/github.com/sachinkumarsingh092/urlshort/handler).