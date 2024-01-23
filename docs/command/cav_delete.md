## cav delete

Delete resource from CloudAvenue.

### Examples

```

	#Delete a Public IP
	cav del ip 192.168.0.2

	#Delete several vdc named xxxx and yyyy
	cav del vdc --name xxxx yyyy
```

### Options

```
  -h, --help   help for delete
```

### Options inherited from parent commands

```
  -t, --time   time elapsed for command
```

### SEE ALSO

* [cav](cav.md)	 - cav is the Command Line Interface for CloudAvenue Platform
* [cav delete edgegateway](cav_delete_edgegateway.md)	 - Delete an edgeGateway (name or id)
* [cav delete publicip](cav_delete_publicip.md)	 - Delete public ip resource(s)
* [cav delete s3](cav_delete_s3.md)	 - Delete a s3 bucket
* [cav delete vdc](cav_delete_vdc.md)	 - Delete a vdc

