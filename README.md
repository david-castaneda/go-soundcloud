![go-souncloud](https://user-images.githubusercontent.com/21694364/61419517-93eba700-a8cc-11e9-8cd0-c379bd5a7129.png)

This Golang package is a hassle free way of interfacing with the SoundCloud Api. As simple as passing in a client id and your set to got with most of the services SoundCloud has to offer.

## Installation

```bash
go get https://github.com/david-castaneda/go-soundcloud
```
## Examples

### Basic Usage
```golang
func main() {
	username := <username>

	sc := SoundCloud.Auth(SOUNCLOUD_CLIENT_ID)

	usr, err := sc.Resolve(username)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(usr)
}
```

### In the works 👨🏻‍💻
- All the missing services 😅