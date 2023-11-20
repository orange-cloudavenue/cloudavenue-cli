# cloudavenue-cli
`cloudavenue-cli` is a terminal app built to give a basic view to manage your Cloud Avenue IaaS.

# install

```
go install github.com/orange-cloudavenue/cloudavenue-cli@latest
```

# usage
* Navigation commands :
  
```shell
$> cloudavenue-cli
cloudavenue-cli is the Command Line Interface for CloudAvenue Platform

Usage:
  cloudavenue-cli [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  edgegateway Option to manage your edgeGateway NSX on CloudAvenue.
  help        Help about any command
  s3          Option to manage your s3 (Object Storage) on CloudAvenue.
  vdc         Option to manage your vdc (Virtual Data Center) on CloudAvenue.

Flags:
  -h, --help   help for cloudavenue-cli

Use "cloudavenue-cli [command] --help" for more information about a command.
```

* list edge gateway:

```shell
$> cloudavenue-cli edgegateway list
| edgeName                | edgeId                               | ownerType | ownerName       | rateLimit | description                                      |
| ----------------------- | ------------------------------------ | --------- | --------------- | --------- | ------------------------------------------------ |
| tn01e02xxx00062xxspt101 | dde5d31a-2f32-xxxx-b3b3-127245958298 | vdc-group | Shared          | 250       | Edge Gateway for customer with BSS ID ocb0006205 |
| tn01e02xxx00062xxspt103 | 4c76e96e-12e2-xxxx-b998-d9c4aa197999 | vdc       | ModuleTF        | 5         | Edge Gateway for customer with BSS ID ocb0006205 |
``````