# Go Learning Journey

My personal repository for learning Go (Golang) through hands-on examples, practice programs, and notes.

> I will update this repository daily as I learn new concepts in Go, adding practice code, examples, and notes along the way.

## What's inside

| Path | Topic |
| --- | --- |
| [`hello.go`](hello.go) | My first Go program |
| [`1section/values.go`](1section/values.go) | Basic values |
| [`1section/variable.go`](1section/variable.go) | Variables, types, zero values, and short declarations |
| [`1section/constant.go`](1section/constant.go) | Constants |
| [`1section/shadowing.go`](1section/shadowing.go) | Variable shadowing |
| [`1section/enum.go`](1section/enum.go) | Enum-style constants |
| [`1section/project.go`](1section/project.go) | Log levels using `iota`, a custom type, and a method receiver |
| [`2section/forloop.go`](2section/forloop.go) | Loops, `break`, `continue`, arrays, and `range` |
| [`2section/ifelse.go`](2section/ifelse.go) | Conditions and map lookups |

## Go setup

### 1. Install Go

Download Go for your operating system from [go.dev/dl](https://go.dev/dl/) and follow the [official installation instructions](https://go.dev/doc/install).

- **macOS:** Choose the package for your processor (Apple Silicon or Intel), open it, and follow the installer.
- **Windows:** Run the MSI installer and reopen your terminal afterward.
- **Linux:** Follow the official archive installation instructions and add `/usr/local/go/bin` to your `PATH`.

Open a new terminal and verify the installation:

```sh
go version
```

If the command is not found, check your `PATH` and restart your terminal.

### 2. Clone this repository

With Git installed, run:

```sh
git clone https://github.com/flutter-create-ani/go_learning.git
cd go_learning
```

### 3. Run your first example

From the repository root:

```sh
go run hello.go
```

Expected output:

```text
Hello, world!
```

Run other examples individually:

```sh
go run 1section/variable.go
go run 1section/project.go
go run 2section/forloop.go
go run 2section/ifelse.go
```

**Run one example file at a time.** The section folders contain independent programs with their own `main()` functions. Running all files in a section together can cause duplicate declaration errors. These standalone examples can be run by filename without creating a `go.mod` file.

## Useful commands

| Command | Purpose |
| --- | --- |
| `go version` | Check the installed Go version |
| `go run hello.go` | Compile and run an example |
| `go build -o hello-example hello.go` | Build an executable on macOS/Linux |
| `go build -o hello-example.exe hello.go` | Build an executable on Windows |
| `gofmt -w hello.go` | Format a source file in place |
| `go doc fmt.Println` | Read documentation for a function |

After building, run `./hello-example` on macOS/Linux or `.\hello-example.exe` in Windows PowerShell.

## My daily learning approach

1. Learn a new concept in Go.
2. Write a small example to understand how it works.
3. Add comments with explanations and important observations.
4. Run the example and experiment with different inputs.
5. Commit my progress to this repository.

This is a work in progress and will grow as I continue learning.

## Learning resources

- [Go documentation](https://go.dev/doc/)
- [A Tour of Go](https://go.dev/tour/)
- [Effective Go](https://go.dev/doc/effective_go)
- [Go standard library](https://pkg.go.dev/std)

---

Learning by building, one concept at a time.

