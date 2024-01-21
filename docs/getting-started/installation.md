## Installation

The binary `cav` will be installed in `/usr/local/bin/` directory. 

### Automatic
```bash
curl -sSfL https://raw.githubusercontent.com/orange-cloudavenue/cloudavenue-cli/main/scripts/install.sh | sudo sh
```

### Manual Debian / Ubuntu
```bash
curl -LO https://github.com/orange-cloudavenue/cloudavenue-cli/releases/download/v0.0.4/cloudavenue-cli_0.0.4_checksums.txt \
&& curl -LO https://github.com/orange-cloudavenue/cloudavenue-cli/releases/download/v0.0.4/cloudavenue-cli_0.0.4_386.deb \
&& sha256sum --ignore-missing -c cloudavenue-cli_0.0.4_checksums.txt
```

### Manual RedHat / CentOS
```bash
curl -LO https://github.com/orange-cloudavenue/cloudavenue-cli/releases/download/v0.0.4/cloudavenue-cli_0.0.4_checksums.txt \
&& curl -LO https://github.com/orange-cloudavenue/cloudavenue-cli/releases/download/v0.0.4/cloudavenue-cli_0.0.4_386.rpm \
&& sha256sum --ignore-missing -c cloudavenue-cli_0.0.4_checksums.txt
```

### Windows
```shell
Invoke-WebRequest https://github.com/orange-cloudavenue/cloudavenue-cli/releases/download/v0.0.4/cloudavenue-cli_0.0.4_checksums.txt -OutFile "cav-checksum.txt"
Invoke-WebRequest https://github.com/orange-cloudavenue/cloudavenue-cli/releases/download/v0.0.4/cloudavenue-cli_0.0.4_windows_amd64.zip -OutFile "cloudavenue-cli_0.0.4_windows_amd64.zip"
Expand-Archive -LiteralPath 'cloudavenue-cli_0.0.4_windows_amd64.zip' -DestinationPath 'c:\cloudavenue\' 
Get-FileHash 'c:\cloudavenue\cav.exe' -Algorithm SHA256 | Format-List
```

 -> Note : Check the value of checksum in the cav-checksum.txt file. And add 'c:\cloudavenue\' in your PATH

### Go env
```shell
go install github.com/orange-cloudavenue/cloudavenue-cli@latest
```

### Debian / Ubuntu

```bash
sudo dpkg -i cloudavenue-cli_0.0.5._386.deb
```

### RedHat / CentOS

```bash
sudo rpm -i cloudavenue-cli_0.0.5._386.rpm
```

### Mac OS X (comming soon)

```bash
brew install cloudavenue-cli
```