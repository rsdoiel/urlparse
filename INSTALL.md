
# Installation

## Compiled version

*timefmt* is a command line program run from a shell like Bash. If you download the
repository a compiled version is in the dist directory. The compiled binary matching
your computer type and operating system can be copied to a bin directory in your PATH.

Compiled versions are available for Mac OS X (amd64 processor), Linux (amd64), Windows
(amd64) and Rapsberry Pi (both ARM6 and ARM7)

### Mac OS X

1. Download **timefmt-binary-release.zip** from [https://github.com/rsdoiel/timefmt/releases/latest](https://github.com/rsdoiel/timefmt/releases/latest)
2. Open a finder window, find and unzip **timefmt-binary-release.zip**
3. Look in the unziped folder and find *dist/macosx-amd64/timefmt*
4. Drag (or copy) *timefmt* to a "bin" directory in your path
5. Open and "Terminal" and run `timefmt -h` to confirm you were successful

### Windows

1. Download **timefmt-binary-release.zip** from [https://github.com/rsdoiel/timefmt/releases/latest](https://github.com/rsdoiel/timefmt/releases/latest)
2. Open the file manager find and unzip **timefmt-binary-release.zip**
3. Look in the unziped folder and find *dist/windows-amd64/timefmt.exe*
4. Drag (or copy) *timefmt.exe* to a "bin" directory in your path
5. Open Bash and and run `timefmt -h` to confirm you were successful

### Linux

1. Download **timefmt-binary-release.zip** from [https://github.com/rsdoiel/timefmt/releases/latest](https://github.com/rsdoiel/timefmt/releases/latest)
2. Find and unzip **timefmt-binary-release.zip**
3. In the unziped directory and find for *dist/linux-amd64/timefmt*
4. Copy *timefmt* to a "bin" directory (e.g. cp ~/Downloads/timefmt-binary-release/dist/linux-amd64/timefmt ~/bin/)
5. From the shell prompt run `timefmt -h` to confirm you were successful

### Raspberry Pi

If you are using a Raspberry Pi 2 or later use the ARM7 binary, ARM6 is only for the first generaiton Raspberry Pi.

1. Download **timefmt-binary-release.zip** from [https://github.com/rsdoiel/timefmt/releases/latest](https://github.com/rsdoiel/timefmt/releases/latest)
2. Find and unzip **timefmt-binary-release.zip**
3. In the unziped directory and find for *dist/raspberrypi-arm7/timefmt*
4. Copy *timefmt* to a "bin" directory (e.g. cp ~/Downloads/timefmt-binary-release/dist/raspberrypi-arm7/timefmt ~/bin/)
    + if you are using an original Raspberry Pi you should copy the ARM6 version instead
5. From the shell prompt run `timefmt -h` to confirm you were successful


## Compiling from source

If you have go v1.6.2 or better installed then should be able to "go get" to install all the **timefmt** utilities and
package. You will need the GOBIN environment variable set. In this example I've set it to $HOME/bin.

```
    GOBIN=$HOME/bin
    go get github.com/rsdoiel/timefmt/...
```

or

```
    git clone https://github.com/rsdoiel/timefmt src/github.com/rsdoiel/timefmt
    cd src/github.com/rsdoiel/timefmt
    make
    make test
    make install
```

