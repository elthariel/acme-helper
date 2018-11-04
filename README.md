# acme-helper

A tiny http server that either serves an ACME http-01 request or
redirects to the requested URL but using the https protocol.

## Usage

To start `acme-helper` just provides a folder to store the acme
challenges (i.e. `.well-known`):

``` shellsession
$> acme-helper /path/to/.well-known
```
