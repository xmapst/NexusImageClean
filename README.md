# NexusImageClean
Nexus CLI for Docker Registry v2 🐳


[![CircleCI](https://circleci.com/gh/xmapst/NexusImageClean.svg?style=svg)](https://circleci.com/gh/xmapst/NexusImageClean) [![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)

<div align="center">
<img src="logo.png" width="60%"/>
</div>

Nexus CLI for Docker Registry

## Usage
```text
$ ./NexusImageClean
NAME:
   Nexus docker image clean CLI - Manage Docker Private Registry on Nexus

USAGE:
   NexusImageClean [global options] command [command options] [arguments...]

VERSION:
   1.0.0-beta

AUTHOR:
   XMapst <xmapst@gmail.com>

COMMANDS:
   configure  Configure Nexus Credentials
   image      Mange Docker Images
   help, h    Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)
```

## Download

Below are the available downloads for the latest version of Nexus CLI (1.0.0-beta). Please download the proper package for your operating system and architecture.

### Linux:

```
wget https://github.com/xmapst/NexusImageClean/releases/download/1.0.0-beta/NexusImageClean
```

## Installation

To install the library and command line program, use the following:

```
go get -u github.com/xmapst/NexusImageClean
```

## Available Commands

```
$ NexusImageClean configure
```

```
$ NexusImageClean image ls
```

```
$ NexusImageClean image tags -name mlabouardy/nginx
```

```
$ NexusImageClean image info -name mlabouardy/nginx -tag 1.2.0
```

```
$ NexusImageClean image delete -name mlabouardy/nginx -tag 1.2.0
```

```
$ NexusImageClean image delete -name mlabouardy/nginx -keep 4
```

```
$ NexusImageClean image delete -keep 4
```
