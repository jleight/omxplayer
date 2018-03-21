omxplayer
=========

[![Build Status](https://travis-ci.org/jleight/omxplayer.svg?branch=master)](https://travis-ci.org/jleight/omxplayer)
[![GoDoc](https://godoc.org/github.com/jleight/omxplayer?status.svg)](https://godoc.org/github.com/jleight/omxplayer)

> omxplayer is a simple library for controlling
> [omxplayer](https://github.com/popcornmix/omxplayer) on your
> [Raspberry Pi](http://www.raspberrypi.org/) from a
> [Go](https://golang.org/) application.


Table of Contents
-----------------

- [Requirements](#requirements)
- [Getting Started](#getting-started)
- [Note on Design Choices](#note-on-design-choices)
- [Usage](#usage)
- [Example](#example)
- [License](#license)


Requirements
------------

- The omxplayer library requires a build of the omxplayer application later than
  the [1d20cdc1be](https://github.com/popcornmix/omxplayer/commit/1d20cdc1be)
  commit.


Getting Started
---------------

The omxplayer library follows the standard Go package format, allowing you to
start using it with a single command:

    go get github.com/jleight/omxplayer

Once the library has been downloaded to your Go path, you can import it in your
project to start using it:

```go
import (
	"github.com/jleight/omxplayer"
)
```


Note on Design Choices
----------------------

The omxplayer application and D-Bus were designed such that:

1. The omxplayer application provides multi-user support by appending the name
  of the user that is running the process to the file containing the D-Bus
  connection information.
2. D-Bus requires authentication in which the user's name and home directory
  need to be provided.

The normal way of getting the current user's username and home directory would
be to use the `os/user` package. However, as of the time of writing this, the
`os/user` package does not work on the Raspberry Pi.

For this reason, we have opted to use `os.Getenv` to get the `USER` and `HOME`
environment variables in order to connect to omxplayer via D-Bus.

If, for any reason, either the `USER` or `HOME` environment variables are not
available, you must specify the user and the user's home directory before
attempting to start a new omxplayer instance. This can be done with the
`SetUser` method.

```go
omxplayer.SetUser("someuser", "/home/someuser")
```

If you are using this library to build an executable to run as a service, you
could choose to pass in the user and home directory as command line arguments:

```go
if len(os.Args) == 3 {
	omxplayer.SetUser(os.Args[1], os.Args[2])
}
```


Usage
-----

Now that you've downloaded and imported the omxplayer library in your
application, you can use it to start up a new instance of omxplayer, providing a
path to the file you would like to play:

```go
player, err := omxplayer.New("/path/to/video.mp4", "--no-osd")
```

This will start a new omxplayer process that will play the specified video file.
The original purpose of this library required that the omxplayer instance have
time to buffer the video before playing it, so this library starts the omxplayer
instance and immediately pauses it.

Sometimes it takes a while (a few hundred milliseconds) for omxplayer to write
its D-Bus information to a file. As a precaution, this library includes both an
`IsReady` and `WaitForReady` method. These can be used to check if the `Player`
instance is ready to start accepting D-Bus commands, or to wait until the
`Player` is ready, respectively. It is recommended that you either make sure
that the `Player` instance `IsReady` before issuing any other commands, or that
you `WaitForReady` if you cannot do anything else.

Now that you have a `Player` instance, you can control it through any of the
D-Bus methods described in the
[D-Bus Control](https://github.com/popcornmix/omxplayer#dbus-control) section
of the omxplayer application's README.


Example
-------

Here's an example to bring it all together:

```go
omxplayer.SetUser("root", "/root")
player, err := omxplayer.New("/root/testvideo.mp4", "--no-osd")

player.WaitForReady()
err = player.PlayPause()

time.Sleep(5 * time.Second)
err = player.ShowSubtitles()

time.Sleep(5 * time.Second)
err = player.Quit()
```

Of course, all of the D-Bus methods return `error`s, so you should make sure
handle them appropriately.


License
-------

The omxplayer library is available under the MIT license. See the
[LICENSE](https://github.com/jleight/omxplayer/blob/master/LICENSE) file for
the full license text.
