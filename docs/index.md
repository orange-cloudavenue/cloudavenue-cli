# Getting Start

[![Go Report Card](https://goreportcard.com/badge/github.com/cloudavenue/cloudavenue-cli)](https://goreportcard.com/report/github.com/orange-cloudavenue/cloudavenue-cli)
[![GoDoc](https://godoc.org/github.com/orange-cloudavenue/cloudavenue-cli?status.svg)](https://godoc.org/github.com/orange-cloudavenue/cloudavenue-cli)

## Table of Contents
1. [Description](#Description)
2. [Installation](#Installation)
3. [Configuration](#Configuration)
4. [Usage](#Usage)

## Description

`cav` is the Command Line Interface for CloudAvenue Platform.
Is a terminal app built to give a basic view to manage your Cloud Avenue IaaS.
You can create, list or delete some basic infrastructure.

 -> Note : If you need more information about Cloud Avenue, please visit [Cloud Avenue documentation](https://wiki.cloudavenue.orange-business.com/w/index.php/Accueil).

## Installation

The binary `cav` will be installed in `/usr/local/bin/` directory. 

### Automatic
```bash
curl -sSfL https://raw.githubusercontent.com/orange-cloudavenue/cloudavenue-cli/main/scripts/install.sh | sudo sh
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

### Debian / Ubuntu

```bash
sudo dpkg -i cloudavenue-cli_0.0.5._386.deb
```

### RedHat / CentOS

```bash
sudo rpm -i cloudavenue-cli_0.0.5._386.rpm
```

### Mac OS X (comming soon)

```bash
brew install cloudavenue-cli
```

## Configuration

Two ways possible to configure your CLI. 

### Configuration File

The first is to use the config file, generate on the first use of your CLI, ex: 

```shell
cav
```

The first try, it's generate a config file locate in your home directory, under the following path `.cav/config.yaml`
You can set your credentials like this:

```yaml
cloudavenue_debug: false
cloudavenue_org: cav01exxxxxxxxxx
cloudavenue_password: YourStrongPassword
cloudavenue_username: yourname.surname
```

### Environment Variables

Credentials can be provided by using the `CLOUDAVENUE_ORG`, `CLOUDAVENUE_USER`, and `CLOUDAVENUE_PASSWORD` environment variables, respectively. Other environnement variables related to [List of Environment Variables](#list-of-environment-variables) can be configured.

For example:

```bash
export CLOUDAVENUE_ORG="my-org"
export CLOUDAVENUE_USER="my-user"
export CLOUDAVENUE_PASSWORD="my-password"
```

:warning: If all variables are set, the CLI override the file configuration to use credentials setted in variables.

## Usage

How to use the CLI

```shell
$> cav
cav is the Command Line Interface for CloudAvenue Platform

Usage:
  cav [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  create      Create resource to CloudAvenue.
  delete      Delete resource from CloudAvenue.
  get         Get resource to retrieve information from CloudAvenue.
  help        Help about any command

Flags:
  -h, --help      help for cav
  -t, --time      time elapsed for command
  -v, --version   version for cav

Use "cav [command] --help" for more information about a command.
```

### Demo

<img src="./multimedia/cli-cav.gif" alt="" width="150%" height="150%">
