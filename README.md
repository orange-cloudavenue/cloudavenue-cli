:warning: This tool is under development, Big improvements or changes may appear.

# What is cav
`cav` is a CLI terminal app built to give a basic view to manage your Cloud Avenue IaaS.
It give you a quick view of your cloudavenue resources.
You can List, Create and Delete some resources.
Actually we can manage VMware resource based on pos VCD (Virtual Cloud Center), like VDC, VM, VAPP
Please Go to website: https://orange-cloudavenue.github.io/cloudavenue-cli/


# Install

Please see Documentation here: https://orange-cloudavenue.github.io/cloudavenue-cli/getting-started/installation/
# Configuration

Please see Documentation here: https://orange-cloudavenue.github.io/cloudavenue-cli/getting-started/configuration/

# Usage
* Navigation commands :
  
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

* list edge gateway:

```shell
$> cav get egw
NAME                        OWNER               
tn01e02ocbxxxxxxspt101     Shared     
tn01e02ocbxxxxxxspt102     PRODUCTION          
tn01e02ocbxxxxxxspt103     STAGING
```