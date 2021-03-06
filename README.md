# Go, Little Drunken Bishop!

```
go get github.com/eloonstra/go-little-drunken-bishop
```

## Table of Contents

- [How does it work?](#how-does-it-work)
- [Usage](#usage)
- [Example](#example)
- [License](#license)
- [Contributing](#contributing)
- [Future](#future)

## How does it work?

There's a [great document](http://dirk-loss.de/sshvis/drunken_bishop.pdf) of Dirk Loss that explains in detail how the
art is generated. I definitely recommend checking it out.

## Usage

Import the library.

```go
import "github.com/eloonstra/go-little-drunken-bishop/pkg/drunkenbishop"
```

Then call the `GenerateRandomArt` function as follows:
```go
drunkenbishop.GenerateRandomArt(17, 9, []byte("test"), true, "test title")
```

And that's it! The result is a string containing the art.

## Example
An example of the output of the `GenerateRandomArt` function with "DrunkenBishop" as bytes and as title:
```
+-[DrunkenBishop]-+
|      . o.. .    |
|       +   o o   |
|        .   @ E  |
|       .   = @   |
|        S   + =  |
|           . +   |
|                 |
|                 |
|                 |
+-----------------+
```

## License

This project is licensed under the MIT license.

## Contributing

All contributions are welcome. Feel free to open an issue or pull request.

## Future

The project is mainly finished; only the main function is not finished yet, so it's impossible to generate art from the command line.
