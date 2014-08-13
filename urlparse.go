/**
 * urlparse.go - a URL Parser library for use in Bash scripts.
 * @author R. S. Doiel, <rsdoiel@gmail.com>
 *
 * copyright (c) 2014 All rights reserved.
 * Released under the Simplified BSD License
 * See: http://opensource.org/licenses/bsd-license.php
 */
package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"
	"path"
    "strings"
)

var (
	help          bool
	showProtocol  bool
	showHost      bool
	showPort      bool
	showPath      bool
	showDir       bool
	showBase      bool
	showExtension bool
	showMimeType  bool
	delimiter     = "\t"
)

var Usage = func(exit_code int, msg string) {
	var fh = os.Stderr
	if exit_code == 0 {
		fh = os.Stdout
	}
	fmt.Fprintf(fh, `%s
 USAGE %s [OPTIONS] URL_TO_PARSE

 Display the parsed URL as delimited fields on one line.

 EXAMPLES

 Get protocol. Returns "http".
 
     %s --protocol http://example.com/my/page.html


 Get host or domain name.  Returns "example.com".
 
     %s --host http://example.com/my/page.html


 Get port.  Returns "8080".

     %s --port http://example.com:8080/my/page.html


 Get path. Returns "/my/page.html".
 
     %s --path http://example.com/my/page.html


 Get basename. Returns "page.html".
 
     %s --basename http://example.com/my/page.html


 Get extension. Returns ".html".
 
     %s --extension http://example.com/my/page.html


 Parse a URL setting environment variables beginning with 'WS_'

     %s --env=WS_ http://example.com/my/page.html

 The environment variables would be $WS_PROTOCOL, $WS_HOST,
 $WS_PORT, $WS_PATH, $WS_BASENAME, $WS_EXTENSION.

 Without options urlparse returns all fields separated by a tab.

 OPTIONS

`, msg, os.Args[0], os.Args[0], os.Args[0],
		os.Args[0], os.Args[0], os.Args[0],
		os.Args[0], os.Args[0])

	flag.VisitAll(func(f *flag.Flag) {
		fmt.Fprintf(fh, "  -%s\t%s\n", f.Name, f.Usage)
	})

	fmt.Fprintf(fh, `

 Copyright (c) 2014 All rights reserved.
 Released under the Simplified BSD License
 See: http://opensource.org/licenses/bsd-license.php 
`)
	os.Exit(exit_code)
}

func init() {
	const (
		delimiter_usage = "Set the output delimited for parsed display. (defaults to tab)"
		help_usage      = "Display this help document."
		protocol_usage  = "Display the protocol of URL (defaults to http)"
		host_usage      = "Display the host (domain name) in URL."
		port_usage      = "Display the port name in URL (defaults to space if no port is set.)"
		path_usage      = "Display the path after the hostname."
		dir_usage       = "Display all but the last element of the path"
		basename_usage  = "Display the base filename at the end of the path."
		extension_usage = "Display the filename extension (e.g. .html)."
	)

	flag.StringVar(&delimiter, "delimiter", delimiter, delimiter_usage)
	flag.StringVar(&delimiter, "D", delimiter, delimiter_usage)
	flag.BoolVar(&showProtocol, "protocol", false, protocol_usage)
	flag.BoolVar(&showProtocol, "T", false, protocol_usage)
	flag.BoolVar(&showHost, "host", false, host_usage)
	flag.BoolVar(&showHost, "H", false, host_usage)
	flag.BoolVar(&showPort, "port", false, port_usage)
	flag.BoolVar(&showPort, "P", false, port_usage)
	flag.BoolVar(&showPath, "path", false, path_usage)
	flag.BoolVar(&showPath, "p", false, path_usage)
	flag.BoolVar(&showDir, "directory", false, basename_usage)
	flag.BoolVar(&showDir, "d", false, basename_usage)
	flag.BoolVar(&showBase, "base", false, basename_usage)
	flag.BoolVar(&showBase, "b", false, basename_usage)
	flag.BoolVar(&showExtension, "extension", false, extension_usage)
	flag.BoolVar(&showExtension, "e", false, extension_usage)

	flag.BoolVar(&help, "help", help, help_usage)
	flag.BoolVar(&help, "h", help, help_usage)
}

func main() {
	flag.Parse()
	if help == true {
		Usage(0, "")
	}
	url_to_parse := flag.Arg(0)
	if url_to_parse == "" {
		Usage(1, "Missing URL to parse")
	}
	u, err := url.Parse(url_to_parse)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		os.Exit(1)
	}
	use_delim := delimiter
	full_display := true
	if showProtocol == true {
		fmt.Fprintf(os.Stdout, "%s%s", u.Scheme, use_delim)
		full_display = false
	}
	if showHost == true {
		fmt.Fprintf(os.Stdout, "%s%s", u.Host, use_delim)
		full_display = false
	}
    if showPort == true {
        if strings.Contains(u.Host, ":") == true {
            cur := strings.LastIndex(u.Host, ":")
            if  cur > -1 {
                fmt.Fprintf(os.Stdout, "%s%s", u.Host[cur+1], use_delim)
            } else {
                fmt.Fprintf(os.Stdout, " %s", use_delim)
            }
        } else {
            fmt.Fprintf(os.Stdout, " %s", use_delim)
        }
    }
	if showPath == true {
		fmt.Fprintf(os.Stdout, "%s%s", u.Path, use_delim)
		full_display = false
	}
	if showBase == true {
		fmt.Fprintf(os.Stdout, "%s%s", path.Base(u.Path), use_delim)
		full_display = false
	}
	if showDir == true {
		fmt.Fprintf(os.Stdout, "%s%s", path.Dir(u.Path), use_delim)
		full_display = false
	}
	if showExtension == true {
		fmt.Fprintf(os.Stdout, "%s%s", path.Ext(u.Path), use_delim)
		full_display = false
	}

	if full_display == true {
		fmt.Fprintf(os.Stdout, "%s%s%s%s%s",
			u.Scheme, use_delim, u.Host, use_delim, u.Path)
	}
	fmt.Fprintln(os.Stdout, "")
	os.Exit(0)
}
