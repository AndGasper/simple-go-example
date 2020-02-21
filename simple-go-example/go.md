# Go 

### (Un)Marshalling Mapping LP
``` 
Exported identifiers
An identifier may be exported to permit access to it from another package. An identifier is exported if both:

the first character of the identifier's name is a Unicode upper case letter (Unicode class "Lu"); and
the identifier is declared in the package block or it is a field name or method name.
All other identifiers are not exported.
```


### Dependency Management (I'm starting to get why npm is so popular)
Q. What's in: `curl https://raw.githubusercontent.com/golang/dep/master/install.sh`
A. See `notes/go/install.sh`


environment variables
```
INSTALL_DIRECTORY
DEP_RELEASE_TAG
DEP_OS
DEP_ARCH
```

"Error" Message 1:
```
Installation requires your GOBIN directory C/bin to exist. Please create it.
```

```
findGoBinDirectory() {
    EFFECTIVE_GOPATH=$(go env GOPATH)
    # CYGWIN: Convert Windows-style path into sh-compatible path
    if [ "$OS_CYGWIN" = "1" ]; then
	EFFECTIVE_GOPATH=$(cygpath "$EFFECTIVE_GOPATH")
    fi
    if [ -z "$EFFECTIVE_GOPATH" ]; then
        echo "Installation could not determine your \$GOPATH."
        exit 1
    fi
    if [ -z "$GOBIN" ]; then
        GOBIN=$(echo "${EFFECTIVE_GOPATH%%:*}/bin" | sed s#//*#/#g)
    fi
    if [ ! -d "$GOBIN" ]; then
        echo "Installation requires your GOBIN directory $GOBIN to exist. Please create it."
        exit 1
    fi
    eval "$1='$GOBIN'"
}
```

- Command:
```
go env
```