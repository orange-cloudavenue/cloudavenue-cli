## 0.1.0 (Unreleased)

### :tada: **Improvements**

* cmd: replace command create by add and delete by del. (GH-85)

### :dependabot: **Dependencies**

* deps: bumps actions/download-artifact from 4.1.2 to 4.1.4 (GH-98)
* deps: bumps dependabot/fetch-metadata from 1.6.0 to 2.0.0 (GH-102)
* deps: bumps github.com/aws/aws-sdk-go from 1.50.15 to 1.50.20 (GH-96)
* deps: bumps github.com/aws/aws-sdk-go from 1.50.20 to 1.50.25 (GH-97)
* deps: bumps github.com/creativeprojects/go-selfupdate from 1.1.3 to 1.1.4 (GH-101)
* deps: bumps github.com/orange-cloudavenue/cloudavenue-sdk-go from 0.9.1 to 0.10.0 (GH-100)
* deps: bumps google.golang.org/protobuf from 1.32.0 to 1.33.0 (GH-99)
* deps: update golang from 1.20 to 1.21 (GH-80)

## 0.0.10 (February 13, 2024)

### :rocket: **New Features**

* `cmd/update`: Now an `update` command is available to update the `cav` binary to the latest version. (GH-76)
* `ouput`: Add json and yaml format for get command. (GH-66)
### :information_source: **Notes**

* `ci`: Add github action dependabot for lib dependencies. (GH-68)

### :dependabot: **Dependencies**

* deps: bumps actions/cache from 3 to 4 (GH-72)
* deps: bumps actions/download-artifact from 3.0.2 to 4.1.1 (GH-74)
* deps: bumps actions/download-artifact from 4.1.1 to 4.1.2 (GH-90)
* deps: bumps actions/setup-go from 4.0.1 to 5.0.0 (GH-73)
* deps: bumps actions/setup-python from 4 to 5 (GH-71)
* deps: bumps actions/upload-artifact from 3 to 4 (GH-75)
* deps: bumps github.com/aws/aws-sdk-go from 1.49.16 to 1.50.8 (GH-70)
* deps: bumps github.com/aws/aws-sdk-go from 1.50.10 to 1.50.15 (GH-88)
* deps: bumps github.com/aws/aws-sdk-go from 1.50.8 to 1.50.10 (GH-81)
* deps: bumps github.com/orange-cloudavenue/cloudavenue-sdk-go from 0.7.0 to 0.9.0 (GH-69)
* deps: bumps github.com/orange-cloudavenue/cloudavenue-sdk-go from 0.9.0 to 0.9.1 (GH-87)
* deps: bumps golangci/golangci-lint-action from 3.7.0 to 4.0.0 (GH-89)

## 0.0.9 (January 25, 2024)

### :rocket: **New Features**

* `completion`: Add completion command. (GH-59)

### :tada: **Improvements**

* `documentation`: Add github pages documentation. (GH-59)

## 0.0.8 (January 19, 2024)

### :rocket: **New Features**

* `New config` -  Add a config file configuration. (GH-57)


## 0.0.7 (January 17, 2024)

### :tada: **Improvements**

* doc: Improve README and Getting Start. (GH-56)
* installation: Add automatic installation. (GH-56)

### :bug: **Bug Fixes**

* goreleaser: Fix - Darwin binary name with `cav`name and right arch. (GH-51)
* version: Fix - Now return the right version compiled. (GH-51)

## 0.0.6 (January 15, 2024)

### :tada: **Improvements**

* `command` - Add animation during command execution time. (GH-47)
* `command` - Add new option `--output` to improve result in `cav` command. (GH-43)
* `command` - improve and change order of command management. (GH-44)
* `command` - improve speed of printed result for vdc list command. (GH-43)

### :dependabot: **Dependencies**

* deps: bumps github.com/cloudflare/circl from 1.3.3 to 1.3.7 (GH-49)
* deps: bumps orange-cloudavenue/cloudavenue-sdk-go from 0.5.6 to 0.7.0 (GH-43)

## 0.0.5 (December  21, 2023)

### :rocket: **New Features**

* `New t0` - Add List operation for t0 (Internet Gateway). (GH-33)

### :tada: **Improvements**

* `Refacto` - rename CLI: cav. (GH-35)

## 0.0.4 (December  4, 2023)

### :rocket: **New Features**

* `New s3` - Add List, Create, Delete operations for s3 bucket (Object Storage). (GH-22)
* `New vdc` - Add List, Create, Delete operations for vdc (Virtual Data Center). (GH-22)
* `binary` - Add packaging distribution. (GH-26)
* `edgegateway` - Add List, Create, Delete operations for edgegateway. (GH-22)
* `publicip` - Add List, Create, Delete operations for publicip. (GH-22)
* `test` - Add Unit Test. (GH-22)
* `timetracking` - Add elapsed time per command. (GH-22)

### :dependabot: **Dependencies**

* deps: bumps orange-cloudavenue/cloudavenue-sdk-go from 0.5.5 to 0.5.6 (GH-22)
