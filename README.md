# Convert to Braille

A simple tool to convert images to braille ASCII art.

## Installation.

With Go installed and this repository cloned, we need to install the dependencies:

```bash
go mod tidy
```

And then we can build the project:

```bash
go build -o converter-to-braille
```

And finally, we can run the program:

```bash
./converter-to-braille --help
```

## Features.

This CLI tool supports the following features:

- Convert images to braille ASCII art and print them to the console. The images can be loaded from local files or URLs.
- Convert images to braille ASCII art and save them to a file.
- Convert GIFs to braille ASCII art and print them to the console.

