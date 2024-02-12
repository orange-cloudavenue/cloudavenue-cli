# ![Orange Logo](multimedia/orange.png) Presentation

[![Releases](https://img.shields.io/github/release-pre/orange-cloudavenue/cloudavenue-cli.svg?sort=semver)](https://github.com/orange-cloudavenue/cloudavenue-cli/releases)
[![LICENSE](https://img.shields.io/github/license/orange-cloudavenue/cloudavenue-cli.svg)](https://github.com/orange-cloudavenue/cloudavenue-cli/blob/main/LICENSE)
![Static Badge](https://img.shields.io/badge/codecov-69%25-orange)
[![CI](https://github.com/orange-cloudavenue/cloudavenue-cli/actions/workflows/deploy_documentation.yml/badge.svg)](https://github.com/orange-cloudavenue/cloudavenue-cli/actions/workflows/deploy_documentation.yml)

![Go version](https://img.shields.io/github/go-mod/go-version/furiko-io/furiko)
[![Go Report Card](https://goreportcard.com/badge/github.com/orange-cloudavenue/cloudavenue-cli)](https://goreportcard.com/report/github.com/orange-cloudavenue/cloudavenue-cli)
[![GoDoc](https://godoc.org/github.com/orange-cloudavenue/cloudavenue-cli?status.svg)](https://godoc.org/github.com/orange-cloudavenue/cloudavenue-cli)

!!! warning
    This tool is under development, big improvements or changes may appear.

[cav](command/cav.md) is a Command Line Interface for Orange CloudAvenue Platform.  
Is supported by the most used operating systems.  
Is a terminal app built to give a basic view to your Cloud Avenue IaaS.  
You can list, delete and create some basic infrastructure.  
Is based on the Terraform [Provider Cloud Avenue](https://registry.terraform.io/providers/orange-cloudavenue/cloudavenue/latest/docs)

## When to use it

[cav](command/cav.md) is able to give you a quik view of your Cloud Avenue Infrastructure.  
It can quickly create an object with a simple template defined in the application coding (in futur release with a parameter template).  
No needs to use terraform to see your IAAS state.  
It is recommended to use it when you want to perform a quick operation on your infrastructure.  
See [demo below](index.md#demo) 

!!! info
    If you need more information about Cloud Avenue, please visit [Cloud Avenue documentation](https://wiki.cloudavenue.orange-business.com/wiki/Welcome).

## Usage

!!! example "How to use the CLI"

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
      update      Check for updates and update the application
      version     Print the version number of cav

    Flags:
      -h, --help      help for cav
      -t, --time      time elapsed for command

    Use "cav [command] --help" for more information about a command.
    ```

[See All Commands reference](command/cav.md)

## Demo

<img src="multimedia/cli-cav.gif" alt="" width="150%" height="150%">

## License

Cloud Avenue CLI [cav](command/cav.md) is licensed under the Mozilla Public License 2.0.

Permissions of this weak copyleft license are conditioned on making available source code of licensed files and modifications of those files under the same license (or in certain cases, one of the GNU licenses).  
Copyright and license notices must be preserved. Contributors provide an express grant of patent rights.  
However, a larger work using the licensed work may be distributed under different terms and without source code for files added in the larger work.

See [LICENSE](https://github.com/orange-cloudavenue/cloudavenue-cli/blob/main/LICENSE).

## Contributing

See [CONTRIBUTING.md](https://github.com/orange-cloudavenue/cloudavenue-cli/blob/main/CONTRIBUTING.md).