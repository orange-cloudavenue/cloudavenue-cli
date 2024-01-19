## Configuration

Two ways possible to configure your CLI. 

### Configuration File

On the first try, a config file locate in your home directory, under the following path `.cav/config.yaml` will be generated.

```shell
> cav
Configuration file created in /root/.cav/config.yaml 
Please fill it with your credentials and re-run the command.
```

You can set your credentials like this:

```yaml
cloudavenue_debug: false
cloudavenue_org: cav01exxxxxxxxxx
cloudavenue_password: YourStrongPassword
cloudavenue_username: yourname.surname
```

### Environment Variables

Credentials can be provided by using the `CLOUDAVENUE_ORG`, `CLOUDAVENUE_USER`, and `CLOUDAVENUE_PASSWORD` environment variables, respectively. Other environnement variables related to [List of Environment Variables](#list-of-environment-variables) can be configured.

For example:

```bash
export CLOUDAVENUE_ORG="my-org"
export CLOUDAVENUE_USER="my-user"
export CLOUDAVENUE_PASSWORD="my-password"
```

:warning: If all variables are set, the CLI override the file configuration to use credentials setted in variables.
