# checkout-system

<a href="https://pkg.go.dev/github.com/billiem/checkout-system/checkout"><img src="https://pkg.go.dev/badge/GitHub.com/billiem/checkout-system.svg" alt="Go Reference"></a>

checkout-system provides utilities to calculate a total checkout price from JSON data.

Documentation can be viewed at: <a href="https://pkg.go.dev/github.com/billiem/checkout-system/checkout"> godoc </a>

# Usage

The checkout package can be used in an external project by installing it by calling `go get github.com/billiem/checkout-system/checkout`,
and importing it with:

    import (
        "github.com/billiem/checkout-system/checkout"
    )

The CLI functionality can also be used by cloning the github repository at <a href = 'https://github.com/billiem/checkout-system'>github.com/billiem/checkout-system</a>.

Once cloned, the program can be executed either via `go run .` followed by the checkout argument/ products flag, or by building the binary with `go build .`, and running it by calling `./checkout-system`.

checkout-system accepts a positional command line argument for the checkout_data json location, and uses the -products flag to accept the product_data json location (used for pricing). If the argument/ flag is not given/ default relative locations will be checked as defined in the constants section of `/checkout/cli.go`

Filepaths accepted by these arguments can either be absolute or relative, if they are relative, they will be relative to the location of the binary. If the binary is installed via `go install`, this will be relative to `$GOBIN`

Example calls:

from the ./checkout-system directory using `go run .` with no argument/ flag
`go run .`

from the ./checkout-system directory after using go build . with a relative checkout_data argument
`./checkout-system checkout_data.json`
