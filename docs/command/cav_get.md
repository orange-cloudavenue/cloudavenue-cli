## cav get

Get resource to retrieve information from CloudAvenue.

### Examples

```

	#List all T0
	cav get t0

	#List all T0 in wide format
	cav get t0 -o wide

	#List all Public IP
	cav get publicip

	#List all VDC in yaml format
	cav get vdc -o yaml

	#List all S3 in json format
	cav get s3 -o json
```

### Options

```
  -h, --help            help for get
  -o, --output string   Output format. One of: (wide, json, yaml)
```

### Options inherited from parent commands

```
  -t, --time   time elapsed for command
```

### SEE ALSO

* [cav](cav.md)	 - cav is the Command Line Interface for CloudAvenue Platform
* [cav get edgegateway](cav_get_edgegateway.md)	 - A brief list of your edgegateway resources
* [cav get publicip](cav_get_publicip.md)	 - A brief list of your public ip resources
* [cav get s3](cav_get_s3.md)	 - A brief list of your s3 resources
* [cav get t0](cav_get_t0.md)	 - A brief list of your t0 resources
* [cav get vdc](cav_get_vdc.md)	 - A brief list of your vdc resources

