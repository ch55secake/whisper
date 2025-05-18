# whisper
Terminal chat client and server built with grpc and bubbletea 


## Developing locally 

> [!NOTE]
> In order for the below-mentioned `wbuild` script, you need to have `fswatch` installed. You can get this by running: 
> `brew install fswatch`. <br>
> If on Linux, `apt install fswatch`.

For working on the project locally, and if on macOS, you can use the `wbuild` script, to watch for changes to the client package,
if any files are changed inside the directory, it will find the process running from the build directory and kill it, then rebuild and 
rerun the most recently built binary. 

I took this idea from this [article](https://leg100.github.io/en/posts/building-bubbletea-programs/) about building bubbletea projects. 

This project also requires that you have `golangci` installed in order to lint and apply proper formatting on any go code in
the project, to install `golangci`: 

```bash
macos: brew install golangci 
linux: brew install golangci 
```

To then run `golangci` run: 
```bash
golangci-lint run 
```
whilst at the root of the module. 

## Building and installing the project 

This project uses `v1.24.2` of `go`, and this is required in order to compile and install the project. To make compilation easier 
you can find a `Makefile` at the root of the project, this allows you to build both binaries for the client and the server. 

To properly install the client binary, run the below commands: 

```bash
  cd ~/Downloads
  chmod +x whisper
  mv ./whisper /usr/local/bin/
```

Once moved you should be able to run `whisper` from any terminal on your machine! 