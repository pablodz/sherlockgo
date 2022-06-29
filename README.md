# SherlockGo

[Original](https://github.com/sherlock-project/sherlock): Hunt down social media accounts by username across social networks

Sherlock but faster and lighter. GUI, CLI, client, Http request and more

![](https://imgur.com/KWminza.png)

## Documentation

Clic [here](https://sherlockgo.herokuapp.com/swagger/index.html) to see the documentation

## Usage

### Remote


```shell
# Then query the sherlock server
curl https://sherlockgo.herokuapp.com/api/v2/username/EXAMPLEUSER
curl https://sherlockgo.herokuapp.com/api/v2/username/EXAMPLEUSER/found/true
```

### Heroku

```shell
# Then query the sherlock server
https://sherlockgo.herokuapp.com/api/v2/username/EXAMPLEUSER
https://sherlockgo.herokuapp.com/api/v2/username/EXAMPLEUSER/found/true
```
### Not compiled

```shell
go run main.go
# Then query the sherlock server
curl localhost:6969/api/v2/username/EXAMPLEUSER
```

### Compiled

Download the binary from releases, they are automatically generated by the latest commit in the master branch.

```shell
./sherlockgo
# Then query the sherlock server
curl localhost:6969/api/v2/username/EXAMPLEUSER
```
## ROADMAP

Options:

- [ ] CLI
- [ ] GUI
- [ ] Tor
- [ ] Proxy
- [ ] HTTP1/2
- [ ] ipv4/ipv6 requests
- [ ] Format output
- [ ] Timeout request
- [ ] Timeout individual request
- [ ] Multiple usernames requests
- [x] Variations of username
- [x] Automatic documentation
- [x] Unit tests
- [x] Streaming requests
- [x] Automatic releaser of binaries for multiple architectures
- [x] Automatic build to Heroku

Error Types:

- [x] response_url
- [x] status_code
- [x] message
