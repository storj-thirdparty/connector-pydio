# storj-pydio connector (uplink v1.0.5)

[![Go Report Card](https://goreportcard.com/badge/github.com/storj-thirdparty/connector-pydio)](https://goreportcard.com/report/github.com/storj-thirdparty/connector-pydio)

## Overview

Command line application (on Windows/Linux/Mac) for taking data backup from Pydio to Storj. Application connects to Pydio server and the souce code for interaction to Storj for cloud storage which is written in Golang.

Pydio, formerly known as AjaXplorer, is an open-source file-sharing and synchronisation software that runs on the user's own server or in the cloud.

```
Usage:
  connector-pydio [command] <flags>

Available Commands:
  help        Help about any command
  store       Command to upload data to a Storj V3 network.
  version     Prints the version of the tool
```  
  
```store``` - Connect to the specified Pydio Cells instance (default: ```pydio_property.json```). Backups of the Pydio storage are generated using tooling provided by Pydio and then uploaded to the Storj network. Connect to a Storj v3 network using the access specified in the Storj configuration file (default: ```storj_config.json```).


Sample configuration files are provided in the ```./config``` folder.

## Requirements and Install
To build from scratch, [install the latest Go](https://golang.org/doc/install#install).

> Note: Ensure go modules are enabled (GO111MODULE=on)

#### Option #1: clone this repo (most common)
To clone the repo
```
git clone https://github.com/storj-thirdparty/connector-pydio.git
```
Then, build the project using the following:
```
cd connector-pydio
go build
```
#### Option #2: go get into your gopath
To download the project inside your GOPATH use the following command:
```
go get github.com/storj-thirdparty/connector-pydio
```
## Run (short version)
Once you have built the project run the following commands as per your requirement:

##### Get help
```
$ ./connector-pydio --help
```
##### Check version
```
$ ./connector-pydio --version
```
##### Create backup from Pydio and upload to Storj
```
$ ./connector-pydio store
```
## Documentation
For more information on runtime flags, configuration, testing, and diagrams, check out the [Detail](//github.com/storj-thirdparty/connector-pydio/wiki/) or jump to:
* [Config Files](//github.com/storj-thirdparty/connector-pydio/wiki/#config-files)
* [Run (long version)](//github.com/storj-thirdparty/connector-pydio/wiki/#run)
* [Testing](//github.com/storj-thirdparty/connector-pydio/wiki/#testing)
* [Flow Diagram](//github.com/storj-thirdparty/connector-pydio/wiki/#flow-diagram)
