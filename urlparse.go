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
    envPrefix     = ""
	delimiter     = "\t"
    port          = "" 
)

var Usage = func(exit_code int, msg string) {
	var fh = os.Stderr
	if exit_code == 0 {
		fh = os.Stdout
	}
	fmt.Fprintf(fh, `%s
 USAGE %s [OPTIONS] URL_TO_PARSE

 Display the parsed URL as delimited fields on one line.
 The default parts to show are protocol, host and path.

 EXAMPLES

 With no options returns "http\texample.com\t/my/page.html"

     %s http://example.com/my/page.html

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
		os.Args[0], os.Args[0], os.Args[0])

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

func exportEnv(varname, value string) {
    ppid := os.Getppid()
    fmt.Printf("DEBUG ppid: %d", ppid)
    //pprocess := os.FindProcess(ppid)
    //pprocess.SetEnv(varname, value)
}

func updateEnv(writeToEnv bool, varname string, value string) {
    if writeToEnv == true {
        fmt.Printf("DEBUG setting %s to %s\n", envPrefix+varname, value)
        exportEnv(envPrefix + varname, value)
    }
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
        env_usage       = "Set results as environment variables."
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
    flag.StringVar(&envPrefix, "environment", envPrefix, env_usage)
    flag.StringVar(&envPrefix, "E", envPrefix, env_usage)

	flag.BoolVar(&help, "help", help, help_usage)
	flag.BoolVar(&help, "h", help, help_usage)
}

func main() {
    var results []string
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

    writeToEnv := false
	use_delim := delimiter
    if envPrefix != "" {
        writeToEnv = true
    }
	if showProtocol == true {
        results = append(results, u.Scheme)
        updateEnv(writeToEnv, "PROTOCOL", u.Scheme)
	}
	if showHost == true {
		results = append(results, u.Host)
        updateEnv(writeToEnv, "HOST", u.Host)
	}
    if showPort == true && strings.Contains(u.Host, ":") == true {
        cur := strings.LastIndex(u.Host, ":")
        port = u.Host[cur+1:]
        results = append(results, port)
        updateEnv(writeToEnv, "PORT", port)
    }
	if showPath == true {
		results = append(results, u.Path)
        updateEnv(writeToEnv, "PATH", u.Path)
	}
	if showBase == true {
		results = append(results, path.Base(u.Path))
        updateEnv(writeToEnv, "BASE", path.Base(u.Path))
	}
	if showDir == true {
		results = append(results, path.Dir(u.Path))
        updateEnv(writeToEnv, "DIRECTORY", path.Dir(u.Path))
	}
	if showExtension == true {
		results = append(results, path.Ext(u.Path))
        updateEnv(writeToEnv, "EXTENSION", path.Ext(u.Path))
	}

	if len(results) == 0 {
        if writeToEnv == true {
            updateEnv(writeToEnv, "PROTOCOL", u.Scheme)
            updateEnv(writeToEnv, "HOST", u.Host)
            updateEnv(writeToEnv, "PATH", u.Path)
        } else {
		    fmt.Fprintf(os.Stdout, "%s%s%s%s%s%s",
			    u.Scheme, use_delim, u.Host, use_delim, u.Path)
	    }
    } else {
        fmt.Fprint(os.Stdout, strings.Join(results, use_delim))
    }
	os.Exit(0)
}
