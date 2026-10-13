# Go Learning Journey

My personal repository for learning Go (Golang) through hands-on examples, practice programs, and notes.

> I will update this repository daily as I learn new concepts in Go, adding new directories, practice code, examples, and notes along the way.

## About this repository

This repository documents my learning journey, from understanding Go's fundamentals to exploring more advanced concepts and building small projects.

You can expect:
- Examples focused on individual concepts.
- Practice programs and experiments.
- Comments explaining how the code works and what I learned.
- New topics and projects as my understanding grows.

## Exploring the code

Browse the repository's folders to explore the available examples. Directories group related learning material, while file names and code comments provide context for each program.

The collection will keep growing. The [repository files](https://github.com/flutter-create-ani/go_learning) show what's available, and the [commit history](https://github.com/flutter-create-ani/go_learning/commits/main/) tracks my latest progress.

## Go setup

### 1. Install Go

Download Go for your operating system from [go.dev/dl](https://go.dev/dl/) and follow the [official installation instructions](https://go.dev/doc/install).

- **macOS:** Choose the package for your processor (Apple Silicon or Intel) and follow the installer.
- **Windows:** Run the MSI installer and reopen your terminal afterward.
- **Linux:** Follow the official archive installation instructions and add `/usr/local/go/bin` to your `PATH`.

Open a new terminal and verify the installation:

```sh
go version
```

If the command is not found, check your `PATH` and restart your terminal.

### 2. Clone the repository

With Git installed, run:

```sh
git clone https://github.com/flutter-create-ani/go_learning.git
cd go_learning
```

### 3. Run an example

For a standalone Go program, run its file from the repository root:

```sh
go run path/to/example.go
```

Replace `path/to/example.go` with the actual path of the example you want to explore.

Some folders contain independent programs, each with its own `main()` function. Run those files individually; compiling them together can cause duplicate declaration errors. Standalone programs that only use the standard library can be run by filename without creating a `go.mod` file.

For a project containing a `go.mod` file, follow any instructions provided with that project. Run commands from its module directory; `go run .` works when the current directory contains a runnable `main` package.

## Useful commands

Replace the example path below with a real file path.

| Command | Purpose |
| --- | --- |
| `go version` | Check the installed Go version |
| `go run path/to/example.go` | Compile and run a standalone program |
| `go build -o example-app path/to/example.go` | Build an executable on macOS/Linux |
| `go build -o example-app.exe path/to/example.go` | Build an executable on Windows |
| `gofmt -w path/to/example.go` | Format a source file in place |
| `go doc fmt.Println` | Read documentation for a function |

After building, run `./example-app` on macOS/Linux or `.\example-app.exe` in Windows PowerShell.

## My daily learning approach

1. Learn a new concept in Go.
2. Write an example or small project to understand it.
3. Add comments with explanations and important observations.
4. Run the code and experiment with different inputs.
5. Commit my progress to this repository.

This is an ongoing learning journal. Examples may evolve as I discover better ways to write Go.

## Learning resources

- [Go documentation](https://go.dev/doc/)
- [A Tour of Go](https://go.dev/tour/)
- [Effective Go](https://go.dev/doc/effective_go)
- [Go standard library](https://pkg.go.dev/std)

---

Learning by building, one concept at a time.
