### Go basics

Dowload go binary : https://golang.org/dl/

#### A. Setup:

1. Setup paths:

    GOROOT - path to go folder
    GOPATH - path where your go packages will be downloaded

2. Add bin folders to path:

    PATH=$PATH:$GOROOT/bin
    PATH:$PATH:$GOPATH/bin

3. To build package on your current workspace

    GOPATH:$GOPATH:/workspace-dir

    

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


#### E. Contents

- [variables](variables)

- [primitives](primitives)

- [arrays](arrays)

- [defer](defer)

- [functions](functions)

- [constants](constants)

- [if](if)

- [loops](loops)

- [map](map)

- [panick](panick)

- [pointers](pointers)

- [struct](struct)

- [switch](switch)

