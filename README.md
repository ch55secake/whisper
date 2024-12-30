# whisper
> Terminal chat client built with websockets and bubbletea 


## Developing locally 

> [!IMPORTANT] In order for the below-mentioned `wbuild` script, you need to have `fswatch` installed. You can get this by running: 
> `brew install fswatch`. If on Linux, you can replace the `fswatch` part of the script with, `inotifywait -e attrib $(find . -name '*.go')`, 
> although this is not something that I tested myself. So it's not guaranteed to work. 

For working on the project locally, and if on macOS, you can use the `wbuild` script, to watch for changes to the client package,
if any files are changed inside the directory, it will find the process running from the build directory and kill it, then rebuild and 
rerun the most recently built binary. 

I took this idea from this [article](https://leg100.github.io/en/posts/building-bubbletea-programs/) about building bubbletea projects. 