# Go, Little Drunken Bishop!

A small library for creating random art for the given bytes written in [Go](https://go.dev/).

## Table of Contents

- [How does it work?](#how-does-it-work)
- [Usage](#usage)
- [Example](#example)
- [License](#license)
- [Contributing](#contributing)

## How does it work?

There's a [great document](http://dirk-loss.de/sshvis/drunken_bishop.pdf) of Dirk Loss that explains in detail how the
art is generated. I definitely recommend checking it out.

## Usage

Import the library.

```go
import "github.com/elenaloonstra/go-little-drunken-bishop/pkg/drunkenbishop"
```

Then, create a new drunken bishop by calling the `New` function. When you've created a new drunken bishop, you can
generate a random art by calling the `Generate` function.

## Example

### Code

```go
package main

import (
	"github.com/eloonstra/go-little-drunken-bishop/pkg/drunkenbishop"
	"fmt"
)

func main() {
	fmt.Println(drunkenbishop.New(22, 8, true, "Drunken Bishop").Generate([]byte("Drunken Bishop")))
}

```

### Result

```
+---[Drunken Bishop]---+
| .o .    +.           |
| o + .  + ..          |
|. o o . .o ..         |
| . o   o  o.          |
|  . .   . . o         |
|   .     o =          |
|          + .         |
|           .          |
+----------------------+
```

## License

This project is licensed under the MIT license.
