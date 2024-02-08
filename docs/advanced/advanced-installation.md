## Advanced Installation

### Package
=== "Debian/Ubuntu"

        curl -LO https://github.com/orange-cloudavenue/cloudavenue-cli/releases/download/v0.0.9/cloudavenue-cli_0.0.9_checksums.txt \
        && curl -LO https://github.com/orange-cloudavenue/cloudavenue-cli/releases/download/v0.0.9/cloudavenue-cli_0.0.9_386.deb \
        && sha256sum --ignore-missing -c cloudavenue-cli_0.0.9_checksums.txt \
        && sudo dpkg -i cloudavenue-cli_0.0.9_386.deb

=== "Redhat/Centos/Fedora"

        curl -LO https://github.com/orange-cloudavenue/cloudavenue-cli/releases/download/v0.0.9/cloudavenue-cli_0.0.9_checksums.txt \
        && curl -LO https://github.com/orange-cloudavenue/cloudavenue-cli/releases/download/v0.0.9/cloudavenue-cli_0.0.9_386.rpm \
        && sha256sum --ignore-missing -c cloudavenue-cli_0.0.9_checksums.txt \
        && sudo rpm -i cloudavenue-cli_0.0.9_386.rpm

=== "Alpine"

        curl -LO https://github.com/orange-cloudavenue/cloudavenue-cli/releases/download/v0.0.9/cloudavenue-cli_0.0.9_checksums.txt \
        && curl -LO https://github.com/orange-cloudavenue/cloudavenue-cli/releases/download/v0.0.9/cloudavenue-cli_0.0.9_386.apk \
        && sha256sum --ignore-missing -c cloudavenue-cli_0.0.9_checksums.txt \
        && sudo apk add cloudavenue-cli_0.0.9_386.apk
        
!!! Note
        For other Arch please replace word `386` by `amd64` or `arm64`

### Windows
``` powershell
Invoke-WebRequest https://github.com/orange-cloudavenue/cloudavenue-cli/releases/download/v0.0.9/cloudavenue-cli_0.0.9_checksums.txt -OutFile "cav-checksum.txt"
Invoke-WebRequest https://github.com/orange-cloudavenue/cloudavenue-cli/releases/download/v0.0.9/cloudavenue-cli_0.0.9_windows_amd64.zip -OutFile "cloudavenue-cli_0.0.9_windows_amd64.zip"
Expand-Archive -LiteralPath 'cloudavenue-cli_0.0.9_windows_amd64.zip' -DestinationPath 'c:\cloudavenue\' 
Get-FileHash 'c:\cloudavenue\cav.exe' -Algorithm SHA256 | Format-List
```
!!! Note
        Check the value of checksum in the cav-checksum.txt file. And add 'c:\cloudavenue\' in your PATH

### MacOS X
=== "MacOS (ARM64)"

        curl -LO https://github.com/orange-cloudavenue/cloudavenue-cli/releases/download/v0.0.9/cloudavenue-cli_0.0.9_darwin_arm64.tar.gz \
        && tar xvf cloudavenue-cli_0.0.9_darwin_arm64.tar.gz \
        && sudo chmod 755 ./cloudavenue-cli_0.0.9_darwin_arm64/cav \
        && sudo mv ./cloudavenue-cli_0.0.9_darwin_arm64/cav /usr/local/bin

=== "MacOS (AMD64)"

        curl -LO https://github.com/orange-cloudavenue/cloudavenue-cli/releases/download/v0.0.9/cloudavenue-cli_0.0.9_darwin_amd64.tar.gz \
        && tar xvf cloudavenue-cli_0.0.9_darwin_amd64.tar.gz \
        && sudo chmod 755 ./cloudavenue-cli_0.0.9_darwin_amd64/cav \
        && sudo mv ./cloudavenue-cli_0.0.9_darwin_amd64/cav /usr/local/bin

=== "MacOS X(brew)"

        Comming soon

### Linux tar.gz
=== "Linux (386)"

        curl -LO https://github.com/orange-cloudavenue/cloudavenue-cli/releases/download/v0.0.9/cloudavenue-cli_0.0.9_linux_386.tar.gz \
        && tar xvf cloudavenue-cli_0.0.9_linux_386.tar.gz \
        && sudo chmod 755 ./cloudavenue-cli_0.0.9_linux_386/cav \
        && sudo mv ./cloudavenue-cli_0.0.9_linux_386/cav /usr/local/bin

=== "Linux (amd64)"

        curl -LO https://github.com/orange-cloudavenue/cloudavenue-cli/releases/download/v0.0.9/cloudavenue-cli_0.0.9_linux_amd64.tar.gz \
        && tar xvf cloudavenue-cli_0.0.9_linux_amd64.tar.gz \
        && sudo chmod 755 ./cloudavenue-cli_0.0.9_linux_amd64/cav \
        && sudo mv ./cloudavenue-cli_0.0.9_linux_amd64/cav /usr/local/bin

=== "Linux (arm64)"

        curl -LO https://github.com/orange-cloudavenue/cloudavenue-cli/releases/download/v0.0.9/cloudavenue-cli_0.0.9_linux_arm64.tar.gz \
        && tar xvf cloudavenue-cli_0.0.9_linux_arm64.tar.gz \
        && sudo chmod 755 ./cloudavenue-cli_0.0.9_linux_arm64/cav \
        && sudo mv ./cloudavenue-cli_0.0.9_linux_arm64/cav /usr/local/bin

### Go env
``` go
go install github.com/orange-cloudavenue/cloudavenue-cli@latest
```