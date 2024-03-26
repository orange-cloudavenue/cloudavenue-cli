## cav del

Delete resource(s) from CloudAvenue.

### Examples

```

	#Delete a Public IP
	cav del ip 192.168.0.2

	#Delete several vdc named xxxx and yyyy
	cav del vdc xxxx yyyy

	#Delete a edgegateway named zzzz
	cav del egw zzzz
```

### Options

```
  -h, --help   help for del
```

### Options inherited from parent commands

```
  -t, --time   time elapsed for command
```

### SEE ALSO

* [cav](cav.md)	 - cav is the Command Line Interface for CloudAvenue Platform
* [cav del edgegateway](cav_del_edgegateway.md)	 - Delete an edgeGateway (name or id).
* [cav del publicip](cav_del_publicip.md)	 - Delete public ip resource(s)
* [cav del s3](cav_del_s3.md)	 - Delete a s3 bucket
* [cav del vdc](cav_del_vdc.md)	 - Delete a vdc

