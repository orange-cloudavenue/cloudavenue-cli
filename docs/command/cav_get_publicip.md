## cav get publicip

A brief list of your public ip resources

### Synopsis

A complete list information of your Public IP resources in your CloudAvenue account.
					The default output format print the minimal necessary information like name, status or group.
					You can use the -o flag to specify the output format.
					"wide" will print some additional information.
					"json" or "yaml" will print the result in the specified format.

```
cav get publicip [flags]
```

### Examples

```
get publicip
```

### Options

```
  -h, --help        help for publicip
      --ip string   A public ip4 adress (default "i")
```

### Options inherited from parent commands

```
  -o, --output string   Output format. One of: (wide, json, yaml)
  -t, --time            time elapsed for command
```

### SEE ALSO

* [cav get](cav_get.md)	 - Get resource to retrieve information from CloudAvenue.

