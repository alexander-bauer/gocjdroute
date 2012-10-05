# GoCjdroute

## About
GoCjdroute is a tool, written in [Go](http://golang.org/), for editing your [cjdns](https://github.com/cjdelisle/cjdns) file `cjdroute.conf` with ease. For more information, view these links:

* [Project Meshnet](http://projectmeshnet.org/)
* [cjdns](https://github.com/cjdelisle/cjdns)
* [Go](http://golang.org/)

## How To Use
GoCjdroute is a command-line program. Some usage cases are presented below, followed by a more in-depth description of flags. It will be assumed that you are operating on the file `/etc/cjdroute.conf`.

**GoCjdroute will remove the comments in your config file. They cannot be maintained or gotten back.**

List all of your authorized passwords (showing passwords)
```
gocjdroute --file /etc/cjdroute.conf --list-auth --show-pass
```

List all of your authorized passwords with non-essential fields (name, location, and IPv6 address) containing "DuoNoxSol," without showing the passwords.
```
gocjdroute --file /etc/cjdroute.conf --list-auth --show-pass --search DuoNoxSol
```

List all of your connections, without showing the passwords.
```
gocjdroute --file /etc/cjdroute.conf --list-conn
```

Add a new authorized password for Bob, who lives in Alice's House, and who has an IPv6 address `fce7:73bf:bdd6:13c4:051d:8840:8fa1:d639`. This will generate a new password for him, and print a block that you can send to him, for connecting to you. Name, location, and IPv6 address are optional.
```
gocjdroute --file /etc/cjdroute.conf --auth Bob "Alice's House" fce7:73bf:bdd6:13c4:051d:8840:8fa1:d639
```

Connect to Alice, who has given you her IPv4 connection details, `1.2.3.4:9876`, public key, and your authorization password. Her name, location, and IPv6 address are optional.
```
gocjdroute --file /etc/cjdroute.conf --connect 1.2.3.4:9876 alicespublicKey.k yourpasswordiskittens Alice "Bob's House"
```

Remove an authorization password based on its index, which can be found by listing and searching. In this case, the index is `16`. This will decrement the indexes of all authorized passwords a greater index.
```
gocjdroute --file /etc/cjdroute.conf --remove 16
```

Remove a connection based on its connection details. This will remove the connection to Alice mentioned earlier.
```
gocjdroute --file /etc/cjdroute.conf --remove 1.2.3.4:9876
```

*Don't forget to restart CJDNS once you've made your edits.*

## Installing
### Install Go
The very first thing that you need to do is install [Go](http://golang.org/).
##### Mac OS X
If you don't already have Homebrew installed, view [this](http://mxcl.github.com/homebrew/) and install it. You'll wonder how you ever lived without it once you install it. If for some reason you don't want to install Homebrew, you can download installers from [Go's website](http://golang.org/).
```bash
brew install go
```

##### Linux (Debian)
```bash
sudo apt-get install golang
```

##### Other
You can download installers from [Go's website](http://golang.org/).

### Clone The Repo!
Get a copy of GoCjdroute on your computer.
```bash
git clone git@github.com:lukevers/GoCjdroute.git
```

### Build The Files
Building the files is simple from the terminal.
```
go build
```
