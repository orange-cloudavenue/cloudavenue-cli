# Getting Start

[![Go Report Card](https://goreportcard.com/badge/github.com/cloudavenue/cloudavenue-cli)](https://goreportcard.com/report/github.com/orange-cloudavenue/cloudavenue-cli)
[![GoDoc](https://godoc.org/github.com/orange-cloudavenue/cloudavenue-cli?status.svg)](https://godoc.org/github.com/orange-cloudavenue/cloudavenue-cli)

## Table of Contents
1. [Prerequisites](#Prerequisites)
2. [Download](#download)
3. [Installation](#installation)
4. [Use](#use)

## Description

`cav` is the Command Line Interface for CloudAvenue Platform.
Is a terminal app built to give a basic view to manage your Cloud Avenue IaaS.
You can create, list or delete some basic infrastructure.

 -> Note : If you need more information about Cloud Avenue, please visit [Cloud Avenue documentation](https://wiki.cloudavenue.orange-business.com/w/index.php/Accueil).

## Configuration

### Environment Variables

Credentials can be provided by using the `CLOUDAVENUE_ORG`, `CLOUDAVENUE_USER`, and `CLOUDAVENUE_PASSWORD` environment variables, respectively. Other environnement variables related to [List of Environment Variables](#list-of-environment-variables) can be configured.

For example:

```bash
export CLOUDAVENUE_ORG="my-org"
export CLOUDAVENUE_USER="my-user"
export CLOUDAVENUE_PASSWORD="my-password"
```

## Install

### Automatic
```bash
curl -sSfL https://raw.githubusercontent.com/orange-cloudavenue/cloudavenue-cli/main/scripts/install.sh | sh
```

### Manual Debian / Ubuntu
```bash
curl -LO https://github.com/orange-cloudavenue/cloudavenue-cli/releases/download/v0.0.4/cloudavenue-cli_0.0.4_checksums.txt \
&& curl -LO https://github.com/orange-cloudavenue/cloudavenue-cli/releases/download/v0.0.4/cloudavenue-cli_0.0.4_386.deb \
&& sha256sum --ignore-missing -c cloudavenue-cli_0.0.4_checksums.txt
```

### Manual RedHat / CentOS
```bash
curl -LO https://github.com/orange-cloudavenue/cloudavenue-cli/releases/download/v0.0.4/cloudavenue-cli_0.0.4_checksums.txt \
&& curl -LO https://github.com/orange-cloudavenue/cloudavenue-cli/releases/download/v0.0.4/cloudavenue-cli_0.0.4_386.rpm \
&& sha256sum --ignore-missing -c cloudavenue-cli_0.0.4_checksums.txt
```

### Windows
```shell
Invoke-WebRequest https://github.com/orange-cloudavenue/cloudavenue-cli/releases/download/v0.0.4/cloudavenue-cli_0.0.4_checksums.txt -OutFile "cav-checksum.txt"
Invoke-WebRequest https://github.com/orange-cloudavenue/cloudavenue-cli/releases/download/v0.0.4/cloudavenue-cli_0.0.4_windows_amd64.zip -OutFile "cloudavenue-cli_0.0.4_windows_amd64.zip"
Expand-Archive -LiteralPath 'cloudavenue-cli_0.0.4_windows_amd64.zip' -DestinationPath 'c:\cloudavenue\' 
Get-FileHash 'c:\cloudavenue\cav.exe' -Algorithm SHA256 | Format-List
```

 -> Note : Check the value of checksum in the cav-checksum.txt file. And add 'c:\cloudavenue\' in your PATH

### Go env
```shell
go install github.com/orange-cloudavenue/cloudavenue-cli@latest
```

## Installation

The binary is installed on `/usr/bin/` PATH.

### Debian / Ubuntu

```bash
sudo dpkg -i cloudavenue-cli_0.0.5._386.deb
```

### RedHat / CentOS

```bash
sudo rpm -i cloudavenue-cli_0.0.5._386.rpm
```

### Mac OS X

```bash
brew install cloudavenue-cli
```

## Use

How to use the CLI

```shell
$> cav
cav is the Command Line Interface for CloudAvenue Platform

Usage:
  cav [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  edgegateway Option to manage your edgeGateway NSX on CloudAvenue.
  help        Help about any command
  s3          Option to manage your s3 (Object Storage) on CloudAvenue.
  vdc         Option to manage your vdc (Virtual Data Center) on CloudAvenue.

Flags:
  -h, --help   help for cav

Use "cav [command] --help" for more information about a command.
```

### Demo

<img src="./multimedia/cli-cav.gif" alt="" width="150%" height="150%">
