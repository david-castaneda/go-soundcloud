![go-souncloud](https://user-images.githubusercontent.com/21694364/61419517-93eba700-a8cc-11e9-8cd0-c379bd5a7129.png)

This Golang package is a hassle free way of interfacing with the SoundCloud Api. As simple as passing in a client id and your set to got with most of the resources SoundCloud has to offer.

## Directory
- [Getting Started](#getting-started)
	- [Installing](#installing)
	- [Importing](#importing)
	- [Authenticating](#authenticating)
- [Resources](#resources)
	- [Users](#users)
	- [Playlist](#playlist)
	- [Followers](#followers)
	- [Followings](#followings)
	- [Resolve](#resolve)
- [Technologies](#technologies)
- [Contributing](#contributing)

## Getting started
- ### Installing 
```bash
go get https://github.com/david-castaneda/go-soundcloud
```
- ### Importing
```golang
import (
	SoundCloud "github.com/go-soundcloud/soundcloud"
)
```
- ### Authenticating
```golang
func main() {
	client := SoundCloud.Auth(<SOUNCLOUD_CLIENT_ID>)
}

```

## Resources
The following is a list of all the endpoints currently available with this client.

- ### Users
Returns a SoundCloud User
```golang
func main() {
	client := SoundCloud.Auth(<SOUNCLOUD_CLIENT_ID>)
	usr, err := client.Resolve(<USERNAME>)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(usr)
}
```

- ### Playlist
Returns a users SoundCloud playlists
```golang
func main() {
	client := SoundCloud.Auth(<SOUNCLOUD_CLIENT_ID>)
	playlists, err := client.Playlists(username)

	if err != nil {
		log.Fatal(err)
	}
	
	fmtPlaylists, err := json.Marshal(playlists)
	
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Print(string(usr))
}
```

- ### Followers
Returns a list of users SoundCloud followers
```golang
func main() {
	client := SoundCloud.Auth(<SOUNCLOUD_CLIENT_ID>)
	followers, err := client.Followers(<USER_ID>)

	if err != nil {
		log.Fatal(err)
	}
	
	fmtFollowers, err := json.Marshal(followers)
	
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(string(fmtFollowers))
}
```

- ### Followings
Returns a list of users SoundCloud followings
```golang
func main() {
	client := SoundCloud.Auth(<SOUNCLOUD_CLIENT_ID>)
	followings, err := client.Followings(<USER_ID>)

	if err != nil {
		log.Fatal(err)
	}
	
	fmtFollowings, err := json.Marshal(followings)
	
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Print(string(fmtFollowings))
}
```

- ### Resolve
Given a SoundCloud url, this resource returns any data the url may have generated
- An example of when to use this resource is when trying to find out a **users id** with only the **username** available. This endpoint will return the same response [users](#users) would return, however can be used with any SoundCloud url.
```golang
func main() {
	url := "https://soundcloud.com/<USERNAME>"

	client := SoundCloud.Auth(<SOUNCLOUD_CLIENT_ID>)
	rslv, err := client.Resolve(url)

	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Print(string(*rslv))
}
```

## Technologies
This package is using Zero external dependencies.

## Contributing
If you would like to contribute, you can create a template and make a PR or fork the projects and make a PR. All contributions/suggestions are welcome.

### In the works 👨🏻‍💻
- All the missing resources 😅
