
[![Go Report Card](http://goreportcard.com/badge/rsdoiel/urlparse)](http://goreportcard.com/report/rsdoiel/urlparse)
[![License](https://img.shields.io/badge/License-BSD%203--Clause-blue.svg)](https://opensource.org/licenses/BSD-3-Clause)

# urlparse

# Overview

A commandline utility that returns a delimited list of URL parts suitable
for use with other Unix utilities like _cut_.

## USAGE 

    urlparse [OPTIONS] URL_TO_PARSE

Display the parsed URL as delimited fields on one line.

## EXAMPLES


Get protocol. Returns "http".
 
     urlparse --protocol http://example.com/my/page.html


Get host or domain name.  Returns "example.com".
 
     urlparse --host http://example.com/my/page.html


Get path. Returns "/my/page.html".
 
     urlparse --path http://example.com/my/page.html


Get basename. Returns "page.html".
 
     urlparse --basename http://example.com/my/page.html


Get extension. Returns ".html".
 
     urlparse --extension http://example.com/my/page.html


## OPTIONS

Without options urlparse returns protocol, host and path fields 
separated by a tab.

+  -b, -basename	Display the base filename at the end of the path.
+  -D, -delimiter	Set the output delimited for parsed display. (defaults to tab)
+  -d, --directory    Display all but the last element of the path
+  -e, -extension	Display the filename extension (e.g. .html).
+  -H, -host	Display the host (domain name) in URL.
+  -p, -path	Display the path after the hostname.
+  -P, -protocol	Display the protocol of URL (defaults to http)

+  -h, -help	Display this help document.

## Installation

_urlparse_ can be installed with the *go get* command.

```
    go get github.com/rsdoiel/urlparse/...
```


## License

Copyright (c) 2014 All rights reserved.
Released under the Simplified BSD License
See: http://opensource.org/licenses/bsd-license.php 

