# Rockset Go Client
[![CircleCI](https://circleci.com/gh/rockset/rockset-go-client/tree/master.svg?style=svg)](https://circleci.com/gh/rockset/rockset-go-client/tree/master)
[![Documentation](https://godoc.org/github.com/getlantern/rockset-go-client?status.svg)](http://godoc.org/github.com/getlantern/rockset-go-client)
[![License](https://img.shields.io/github/license/rockset/rockset-go-client.svg?maxAge=2592000)](https://github.com/getlantern/rockset-go-client/LICENSE)
[![GitHub issues](https://img.shields.io/github/issues/rockset/rockset-go-client.svg)](https://github.com/getlantern/rockset-go-client/issues)
[![Release](https://img.shields.io/github/release/rockset/rockset-go-client.svg?label=Release)](https://github.com/getlantern/rockset-go-client/releases)
[![Coverage Status](https://coveralls.io/repos/github/rockset/rockset-go-client/badge.svg?branch=master)](https://coveralls.io/github/rockset/rockset-go-client?branch=master)

Official Go client library for Rockset

## Installation

Install the Rockset Go client from Github:

```
go get github.com/getlantern/rockset-go-client
```

or install it from a source code checkout:

```
cd $GOPATH/src/github.com
mkdir rockset
cd rockset
git clone git@github.com:rockset/rockset-go-client.git
go install rockset-go-client/rockclient.go
```

## Usage

You can see a few [examples](https://godoc.org/github.com/getlantern/rockset-go-client/#pkg-overview) in the godoc
of how to create a collection, how to put documents in a collection and how to use SQL to query your collections.

```
client, err := rockset.NewClient()
if err != nil {
    log.Fatal(err)
}

ctx := context.TODO()
q := `SELECT * FROM "_events" LIMIT 1`

res, _, err := client.Query(ctx, q)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("%+v\n", res)
```

## Testing

There are a number of testable examples which require an API key, i.e. set the `ROCKSET_APIKEY` 
environment variable. 

To run tests:
```
go test ./...
```

### Code Coverage

```
go test ./... -coverprofile cover.out
go tool cover -func=cover.out
go tool cover -html=cover.out -o coverage.html
```

## Support

Feel free to log issues against this client through GitHub.

## License

The Rockset Go Client is licensed under the [Apache 2.0 License](https://github.com/getlantern/rockset-go-client/blob/master/LICENSE)
