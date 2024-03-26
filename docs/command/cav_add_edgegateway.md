## cav add edgegateway

Add an edgeGateway

### Synopsis

Add an edgeGateway in a VDC. If the T0 is not specified, the first one will be used. No need to specify a name, the edgeGateway name is auto-generated.

```
cav add edgegateway [flags]
```

### Examples

```
add edgegateway --vdc <vdc name> [--t0 <t0 name>]
```

### Options

```
  -h, --help         help for edgegateway
      --t0 string    t0 name
      --vdc string   vdc name
```

### Options inherited from parent commands

```
  -t, --time   time elapsed for command
```

### SEE ALSO

* [cav add](cav_add.md)	 - Add resource to CloudAvenue.

