# Go Wrapper for prefix.cc

This Go package is intended to provide functions for fetching the commonly used namespace of a given prefix name 
and for finding the commonly used prefix name of a given namespace or IRI respectively. 

It is a Go wrapper for the API provided by [prefix.cc](prefix.cc).

# Build

```
go build
```
or
```
make build
```

Both of these commands create a binary named `goprefixcc`.

# Running

```
./goprefixcc (string [-reverse]) | -version

Usage of ./goprefixcc:
  -reverse
    	performs a reverse lookup
  -version
    	prints version of this app
```


## Example

```
./goprefixcc foaf
```

The result would be `http://xmlns.com/foaf/0.1/`.

```
./goprefixcc -reverse "http://purl.org/dc/terms/title"
```

The result would be `dc`.
