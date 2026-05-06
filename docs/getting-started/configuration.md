## Configuration

Two ways possible to configure your CLI.

### Configuration File

On the first try, a config file locate in your home directory, under the following path `.cav/config.yaml` will be generated.

```{ .shell .no-copy }
> cav
Configuration file created in /root/.cav/config.yaml
Please fill it with your credentials and re-run the command.
```

You can set your credentials like this:

```{ .yaml .no-copy }
cloudavenue:
    username: "yourUserName"
    password: "YourStrongPassword"
    org: "cav01exxxxxxxxxx"
    url: ""
    debug: false
    core_api: ""  # Override the backend API endpoint (optional)
```

### Environment Variables

The following environment variables are supported:

| Variable | Description | Required | Default |
|---|---|---|---|
| `CLOUDAVENUE_USERNAME` | CloudAvenue username | Yes | — |
| `CLOUDAVENUE_PASSWORD` | CloudAvenue password | Yes | — |
| `CLOUDAVENUE_ORG` | CloudAvenue organization (format: `cav01exxxxxxxxxx`) | Yes | — |
| `CLOUDAVENUE_URL` | CloudAvenue console URL | No | Computed from org |
| `CLOUDAVENUE_DEBUG` | Enable debug mode | No | `false` |
| `CLOUDAVENUE_CORE_API` | Override the backend API endpoint. Must be an `https://` URL. Useful for development or staging environments. | No | `https://api1.cloudavenue.orange-business.com` |

For example:

```{ .shell }
export CLOUDAVENUE_ORG="my-org"
export CLOUDAVENUE_USERNAME="my-user"
export CLOUDAVENUE_PASSWORD="my-password"
```

!!! warning
    If all variables are set, the CLI override the file configuration to use credentials setted in variables.
