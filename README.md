# pswdgen

pswdgen generates a random password.

## Build

`go build .`

## Install

`go install github.com/chevaliersoft/pswdgen@v1.0.0`

## Man

```man
pswdgen generates a password for you

Usage:
  pswdgen [flags]

Flags:
  -h, --help         help for pswdgen
  -l, --length int   password length (default 16)
  -w, --lower        use lower characters
  -n, --numbers      use numbers
  -s, --specials     use special characters
  -u, --upper        use upper characters
```
