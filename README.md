# GoCjdroute

## About
GoCjdroute is a tool, written in [Go](http://golang.org/), for editing your [cjdns](https://github.com/cjdelisle/cjdns) file `cjdroute.conf` with ease. For more information, view these links:

* [Project Meshnet](http://projectmeshnet.org/)
* [cjdns](https://github.com/cjdelisle/cjdns)
* [Go](http://golang.org/)

## How To Use
GoCjdroute is a command-line program. Some usage cases are presented below, followed by a more in-depth description of flags.

**GoCjdroute will remove the comments in your config file. They cannot be maintained or gotten back.** ***Please back up your config file. Even the best user can make a mistake.***

### The Commands
GoCjdroute has five primary commands. These are `auth`, `conn`, `lsa`, `lsc`, and `rm`. The syntaxes for any GoCjdroute command is `gocjdroute <command> [argument] [additional arguments]`. For some commands, `[argument]` is required, and in others, it is not.

#### auth

**This behavior writes to the file.**

```bash
gocjdroute auth [index]
```

`auth` is used to generate, add, and edit authorized passwords, based on their indexes in the *array* that is maintained in cjdroute.conf.

The argument `index` is not required, and if it is not supplied, it is implicitly `-1`, which means that a *new* authorized password will be appended to the array. If it is supplied, but out of the bounds of the array, it exhibits the same behavior as `-1`.

If it refers to an existing password, then the *edit* behavior will instead be invoked. This is just for editing identifying information, such as the `Name`, `Location`, or `IPv6` fields, as opposed to the password itself, which cannot be edited.

The JSON behavior is also available.
```bash
gocjdroute --json auth [index] <JSON-encoded-AuthPass>
```

Similarly to the above behavior, `index` is implicitly `-1` if it is not specified. If it is specified, then the password block specified by it will be replaced entirely, including the password.

The JSON can be specified in a single argument, or many, and may be indented or not. Whitespace has no effect on it. Please note that the `-j` option is equivalent to `--json`.

#### conn

**This behavior writes to the file.**

```bash
gocjdroute conn [details]
```

`conn` is used to add and edit connection blocks. Information can be entered via interactive prompt, or by supplied JSON.

If the argument `[details]`, which should usually be an IPv4 address, with a port number, supplied by the target peer, is not supplied, connection details will be asked for in the interactive prompt. Either a non-supplied `[details]` argument, or one which does not already exist in the configuration file will case a new connection block to be added to the configuration file.

If, instead, `[details]` is supplied, and refers to an existing connect block, *edit* mode will be invoked, in which the interactive prompt can be used to modify identifying information.

The JSON behavior is also available.
```bash
gocjdroute --json conn <JSON-encoded-map-of-Connections>
```

In this behavior, any `[details]` supplied separate of the JSON may interfere with the behavior. The JSON may be supplied as one argument, or many, and whitespace is ignored.

A valid invocation of this behavior is as follows, in two different forms. Please note that `--json` and `-j` are equivalent.
```bash
gocjdroute --json conn '{
    "75.75.75.75:9876": {
        "name": "Government Spies",
        "location": "Your Roof",
        "password": "123456",
        "publicKey": "averylongpublickey.k"
    }
}'
```

This is equivalent to the following.
```bash
gocjdroute --json conn { "75.75.75.75:9876": {"name": "Government Spies", "location": "Your Roof", "password": "123456", "publicKey": "averylongpublickey.k" } }
```

More than one connection block may be specified in the map, but bear in mind that it must be properly JSON-encoded, as follows.
```
{
    "75.75.75.75:9876": {
        "name": "Government Spies",
        "location": "Your Roof",
        "password": "123456",
        "publicKey": "averylongpublickey.k"
    },
    "76.76.76.76:6789": {
        "name": "Your Friendly Neighbor",
        "location": "Next Door",
        "password": "654321",
        "publicKey": "isntthatnedflanderschapgreat.k"
    }
}
```

#### lsa

```bash
gocjdroute lsa [search term]
```

`lsa` is a simple combination authorization listing and searching utility. It does not cause a write to the file, and so may be used non-destructively.

It goes through each authorized password, and displays it onscreen, JSON-encoded, and tagged with its index in the file.

If `[search term]` is specified, then it also checks every element, `Name`, `Location`, `IPv6`, and `Password` in each password block, and only displays it if at least one element *contains* the search term. They remain tagged with their actual indexes in the configuration. Regex is not allowed.

#### lsc

```bash
gocjdroute lsc [search term]
```

`lsc` behaves extremely similarly to `lsa`. The primary difference is that it displays connection blocks, rather than authorization blocks. It is also non-destructive.

It goes through each connection block, displaying it onscreen, JSON-encoded, in `map` form.

If `[search term]` is specified, then it checks all elements, as in `lsa`, including the connection details, `PublicKey` and `Password`. Regex is not allowed.

#### rm

**This behavior writes to the file. Please be careful with this one.**

```bash
gocjdroute rm <identifier>
```

`rm` has one of the most complex behaviors of the primary commands. It is able to remove *both* authorization and connection blocks, based entirely on whether it can parse `identifier` into an integer or not.

If `identifier` is an integer, it tries to remove that index from the authorized passwords array. If not, it deletes the element identified by it from the connection block map. If it fails to remove either, it does nothing.

The following command removes the zeroth authorized password.
```bash
gocjdroute rm 0
```

The following command, however, removes the connection to `75.75.75.75:9876`.
```bash
gocjdroute rm 75.75.75.75:9876
```

### The Flags
GoCjdroute has flags which behave similarly to flags in any other Unix script.

#### --file (-f)
Default: `/etc/cjdroute.conf`

`--flag` is used to specify the configuration file for viewing or editing, and . The syntax is as follows.

```bash
gocjdroute --file /path/to/cjdroute.conf <command> [argument] [additional arguments]
```

#### --json (-j)
`--json` is used in combination with certain other commands, discussed in each command's entry. The general syntax is as follows.

```bash
gocjdroute --json <command> [argument] <JSON-encoded-value>
```
 
## Installing
### Install Go
The very first thing that you need to do is install [Go](http://golang.org/).
#### Mac OS X
If you don't already have Homebrew installed, view [this](http://mxcl.github.com/homebrew/) and install it. You'll wonder how you ever lived without it once you install it. If for some reason you don't want to install Homebrew, you can download installers from [Go's website](http://golang.org/).
```bash
brew install go
```

#### Linux (Ubuntu)
The Ubuntu repositories have golang present in the repositories. For other Linux distributions, it may be necessary to get a package file and install it directly, or even to compile from source. (Go makes an effort to be easy to install, however.)
```bash
sudo apt-get install golang
```

#### Other
You can download installers from [Go's website](http://golang.org/).

### Get and Install GoCjdroute
#### Use the Gotool
If you are using the gotool, then you can use its `get` and `install` commands to install GoCjdroute. You may need to be root or to have configured your `GOPATH` in order to be able to install properly.
```bash
go get github.com/SashaCrofter/gocjdroute
go install github.com/SashaCrofter/gocjdroute
```

You should then be able to execute `gocjdroute` and be given usage information. If not, make sure that your `$GOPATH/bin` is on your `PATH`. This may require configuration of the `GOPATH`, as before.

#### Clone and Build with Git
GoCjdroute can be cloned using Git. The simplest way to do so is with the following command.
```bash
git clone git://github.com/SashaCrofter/gocjdroute.git
```

You may then use any Go compiler to compile it, and place it on your `PATH`.

## Contact
If you have problems, particularly with GoCjdroute itself, don't hesitate to contact me. It is usually simplest to just open an issue on the [GitHub repository](https://github.com/SashaCrofter/gocjdroute/issues). I will respond to them as quickly as possible, tag them, fix bugs, and implement improvements. (Pull requests are greatly appreciated!)

If, for some reason, you should wish to contact me directly, I can be reached at sasha@crofter.org.
