### The basics of Go

Tables of contents:

- [arrays](go-basics/arrays.md)

- [channels](go-basics/channels.md)

- [constants](go-basics/constants.md)

- [defer](go-basics/defer.md)

- [functions](go-basics/functions.md)

- [goroutines](go-basics/goroutines.md)

- [basic web app](go-basics/go-wiki/README.md)

- [if](go-basics/if.md)

- [interface](go-basics/interface.md)

- [loops](go-basics/loops.md)

- [map](go-basics/map.md)

- [panick](go-basics/panick.md)

- [pointers](go-basics/pointers.md)

- [primitives](go-basics/primitives.md)

- [slice](go-basics/slice.md)

- [struct](go-basics/struct.md)

- [switch](go-basics/switch.md)

- [variables](go-basics/variables.md)

#### A. Setup:

1.  Dowload go binary : https://golang.org/dl/

2.  Setup paths:

    GOROOT - path to go folder

    GOPATH - path where your go packages will be downloaded

3.  Add bin folders to path:

    PATH=$PATH:$GOROOT/bin

    PATH:$PATH:$GOPATH/bin

4.  To build package on your current workspace

    GOPATH:\$GOPATH:/workspace-dir

#### B. Workspace

Create folder structure

    src/ - contains go source code
        /github.com - allows package to be fetch by `go get`
            /rodelrebucas
                /app-name
    pkg/ - contains intermediate binaries (for third parties)
    bin/ - contains binary files

#### C. Building package

Make sure GOPATH is setup correctly for building package inside workspace

Run `go build <complete-package-dir>`

`go build github.com/rodelrebucas/go-basics`

Run the generated executable

`./go-basics`

#### D. Installing built go package

`go install github.com/rodelrebucas/go-basics`

Generates executable inside `/bin folder`

**References**:

Go web application: https://golang.org/doc/articles/wiki/

Go basics: https://www.youtube.com/watch?v=YS4e4q9oBaU
