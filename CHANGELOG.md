## 0.1.0 (Unreleased)

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
