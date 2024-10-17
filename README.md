# Errorhole
Convenient error handling

## Installation

To install errorhole, use `go get`:

```bash
go get github.com/Toolnado/errorhole
```
## Usage

Here's a basic example of how to use errorhole:

```go
func main() {
	defer errorhole.Catch() // Waiting for an error
	// do something...
	app, err := NewApp()
	errorhole.Nil(err) // Check error
	app.Run()
}
```
## Contributing

Contributions are welcome! Please fork the repository and submit a pull request with your changes. Be sure to follow the project's coding guidelines and include tests for any new features or bug fixes.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.