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

- Convert images to braille ASCII art and print them to the console.
- Convert images to braille ASCII art and save them to a file.
- Convert GIFs to braille ASCII art and print them to the console.

## Future improvements.

- Add support to get the image from a URL.
- Add support to get the GIF from a URL.
- Save the braille ASCII art as a image file or a GIF file (as appropriate).
