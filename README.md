# pterogo: A Go library for the Pterodactyl API

[![Go Report Card](https://goreportcard.com/badge/github.com/elasthicc/pterodactyl-go)](https://goreportcard.com/report/github.com/elasthicc/pterodactyl-go)

![pterogo-maskot](media/pterogo_maskot.png)


> **Note**: This library is under active development as we expand it to cover
> the Pterodactyl API. Consider the public API of this package a little
> unstable as we work towards a v1.0.

Package pterogo is a library for the Pterodactyl API.

## Example

```golang
func main() {
	pteroApp := pteroapp.NewApplication(pteroapp.WithEndpoint(url), pteroapp.WithToken(token))

	user, _, err := pteroApp.Users.GetByID(context.Background(), 1)
	if err != nil {
		log.Fatalf("error retrieving user: %s\n", err)
	}

	fmt.Printf("user email is: %s\n", user.Attributes.Email)
}
```

## Pterodactyl API Version Support

The library supports the [Pterodactyl v1 API](https://dashflo.net/docs/api/pterodactyl/v1/).

## Go Version Support

The library supports the latest two Go minor versions, e.g. at the time Go 1.21 is released, it supports Go 1.20 and 1.21.

This matches the official [Go Release Policy](https://go.dev/doc/devel/release#policy).

When the minimum required Go version is changed, it is announced in the release notes for that version.


## Contributing

Gin is the work of hundreds of contributors. We appreciate your help!

Please see [CONTRIBUTING](CONTRIBUTING.md) for details on submitting patches and the contribution workflow.

## License
MIT License