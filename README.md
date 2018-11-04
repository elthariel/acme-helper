# acme-helper

A tiny http server that either serves an ACME http-01 request or
redirects to the requested URL but using the https protocol.

## Usage

To start `acme-helper` just provides a folder to store the acme
challenges (i.e. `.well-known`):

``` shellsession
$> acme-helper /path/to/.well-known
```

## Docker image

A distroless [docker image is
available](https://hub.docker.com/r/elthariel/acme-helper/). The
`.well-known` acme challenges are served from the `/well-known` folder
in the image.

To use it, run:

``` shellsession
$> docker run -v /your/acme/well-known:/well-known elthariel/acme-helper
```
