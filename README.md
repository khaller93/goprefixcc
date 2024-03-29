# Go Wrapper for prefix.cc

This Go application and module is intended to provide functions for fetching the namespace of a given prefix
name and for finding the prefix name of a given namespace or IRI respectively. 

It is a Go wrapper for the API provided by [prefix.cc](https://prefix.cc).

## Build

```bash
go build
```

This command creates a binary named `goprefixcc`.

## Running

```
./goprefixcc ([-reverse] string) | -version

Usage of ./goprefixcc:
  -reverse
    	performs a reverse lookup
  -version
    	prints version of this app
```


### Example

```bash
./goprefixcc foaf
```

The result would be `http://xmlns.com/foaf/0.1/`.

```bash
./goprefixcc -reverse "http://purl.org/dc/terms/title"
```

The result would be `dc`.
