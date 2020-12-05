# urlshort
Make your own short URLs. Use a customized YAML file to make the paths and URLs.

## Installation

  
**Cloning**
```
$ git clone https://github.com/sachinkumarsingh092/urlshort
```

**Verifing Checksums**
```
$ go mod verify
```

- ### Build locally
  
  - **Making binaries**
  ```
  $ go build -x main.go
  ```
  
  - **Running binaries**
  ```
  $ ./main [OPTIONS]
  ```
  
  The program will start on port 8080 locally.

  - **Test using `curl`**:
  ```
  curl -i localhost:8080/<SOME_PATH>
  ```

- ### Build using Docker

  From the root of the directory:
  
  - **Building image from the Dockerfile**
  ```
  docker build -t <IMAGE_NAME>:latest .
  ```
  
  - **Running the container and port forwarding**
  ```
  docker run --publish 80:8080 -t <IMAGE_NAME>:latest 
  ```
  
  The host can now communicate with the container on port 80.

  - **Test using `curl`**:
  ```
  curl -i localhost:80/<SOME_PATH>
  ```

## Options
```
-yaml=FILENAME
    Sets the YAML file to create URL map. (default "urls.yaml")
-help
    Print help.
```

## Documentations
Docs for handler package at [GoDoc](https://godoc.org/github.com/sachinkumarsingh092/urlshort/handler).