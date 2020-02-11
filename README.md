### The basics of Go

Tables of contents:

- [arrays](go-basics/arrays.md)

- [channels](go-basics/channels)

- [constants](go-basics/constants)

- [defer](go-basics/defer)

- [functions](go-basics/functions)

- [goroutines](go-basics/goroutines)

- [basic http server](go-basics/http-server)

- [if](go-basics/if)

- [interface](go-basics/interface)

- [loops](go-basics/loops)

- [map](go-basics/map)

- [panick](go-basics/panick)

- [pointers](go-basics/pointers)

- [primitives](go-basics/primitives)

- [slice](go-basics/slice)

- [struct](go-basics/struct)

- [switch](go-basics/switch)

- [variables](go-basics/variables)

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

References:

Channels: https://www.youtube.com/watch?v=f6kdp27TYZs

Go basics: https://www.youtube.com/watch?v=YS4e4q9oBaU
