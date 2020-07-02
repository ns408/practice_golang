
```shell
export GOBIN=~/go/bin/

# Build and install the binary in GOBIN path
go install

# Build the binary in the relative directory
go build

# Compiles the file to a temporary location and runs it:
go run main.go

## shows location of temporary path
go run --work main.go
```

## Formatting

```shell
go fmt variables.go
```

## Vim

Vim-go plugin: https://github.com/fatih/vim-go/blob/master/README.md#install

```shell
# This works on mine because I already created the dir structure when install vim-ruby plugin
git clone https://github.com/fatih/vim-go.git ~/.vim/bundle/vim-go
```

Notes:
- package main  - every go file must start with the `package name`. The main function should always reside in the main package.
- func main() - special function from which execution starts.
- fmt.Println("Hello") - `Println` function of the `fmt` package. `package.function()` is the syntax to **call a function** in a package.


