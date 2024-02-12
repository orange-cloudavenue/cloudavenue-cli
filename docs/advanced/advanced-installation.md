## Advanced Installation

### Package
=== "Debian/Ubuntu"

        curl -LO https://github.com/orange-cloudavenue/cloudavenue-cli/releases/download/v__TAG__/cloudavenue-cli___TAG___checksums.txt \
        && curl -LO https://github.com/orange-cloudavenue/cloudavenue-cli/releases/download/v__TAG__/cloudavenue-cli___TAG___386.deb \
        && sha256sum --ignore-missing -c cloudavenue-cli___TAG___checksums.txt \
        && sudo dpkg -i cloudavenue-cli___TAG___386.deb

=== "Redhat/Centos/Fedora"

        curl -LO https://github.com/orange-cloudavenue/cloudavenue-cli/releases/download/v__TAG__/cloudavenue-cli___TAG___checksums.txt \
        && curl -LO https://github.com/orange-cloudavenue/cloudavenue-cli/releases/download/v__TAG__/cloudavenue-cli___TAG___386.rpm \
        && sha256sum --ignore-missing -c cloudavenue-cli___TAG___checksums.txt \
        && sudo rpm -i cloudavenue-cli___TAG___386.rpm

=== "Alpine"

        curl -LO https://github.com/orange-cloudavenue/cloudavenue-cli/releases/download/v__TAG__/cloudavenue-cli___TAG___checksums.txt \
        && curl -LO https://github.com/orange-cloudavenue/cloudavenue-cli/releases/download/v__TAG__/cloudavenue-cli___TAG___386.apk \
        && sha256sum --ignore-missing -c cloudavenue-cli___TAG___checksums.txt \
        && sudo apk add cloudavenue-cli___TAG___386.apk
        
!!! Note
        For other Arch please replace word `386` by `amd64` or `arm64`

### Windows
``` powershell
Invoke-WebRequest https://github.com/orange-cloudavenue/cloudavenue-cli/releases/download/v__TAG__/cloudavenue-cli___TAG___checksums.txt -OutFile "cav-checksum.txt"
Invoke-WebRequest https://github.com/orange-cloudavenue/cloudavenue-cli/releases/download/v__TAG__/cloudavenue-cli___TAG___windows_amd64.zip -OutFile "cloudavenue-cli___TAG___windows_amd64.zip"
Expand-Archive -LiteralPath 'cloudavenue-cli___TAG___windows_amd64.zip' -DestinationPath 'c:\cloudavenue\' 
Get-FileHash 'c:\cloudavenue\cav.exe' -Algorithm SHA256 | Format-List
```
!!! Note
        Check the value of checksum in the cav-checksum.txt file. And add 'c:\cloudavenue\' in your PATH

### MacOS X
=== "MacOS (ARM64)"

        curl -LO https://github.com/orange-cloudavenue/cloudavenue-cli/releases/download/v__TAG__/cloudavenue-cli___TAG___darwin_arm64.tar.gz \
        && tar xvf cloudavenue-cli___TAG___darwin_arm64.tar.gz \
        && sudo chmod 755 ./cloudavenue-cli___TAG___darwin_arm64/cav \
        && sudo mv ./cloudavenue-cli___TAG___darwin_arm64/cav /usr/local/bin

=== "MacOS (AMD64)"

        curl -LO https://github.com/orange-cloudavenue/cloudavenue-cli/releases/download/v__TAG__/cloudavenue-cli___TAG___darwin_amd64.tar.gz \
        && tar xvf cloudavenue-cli___TAG___darwin_amd64.tar.gz \
        && sudo chmod 755 ./cloudavenue-cli___TAG___darwin_amd64/cav \
        && sudo mv ./cloudavenue-cli___TAG___darwin_amd64/cav /usr/local/bin

=== "MacOS X(brew)"

        Comming soon

### Linux tar.gz
=== "Linux (386)"

        curl -LO https://github.com/orange-cloudavenue/cloudavenue-cli/releases/download/v__TAG__/cloudavenue-cli___TAG___linux_386.tar.gz \
        && tar xvf cloudavenue-cli___TAG___linux_386.tar.gz \
        && sudo chmod 755 ./cloudavenue-cli___TAG___linux_386/cav \
        && sudo mv ./cloudavenue-cli___TAG___linux_386/cav /usr/local/bin

=== "Linux (amd64)"

        curl -LO https://github.com/orange-cloudavenue/cloudavenue-cli/releases/download/v__TAG__/cloudavenue-cli___TAG___linux_amd64.tar.gz \
        && tar xvf cloudavenue-cli___TAG___linux_amd64.tar.gz \
        && sudo chmod 755 ./cloudavenue-cli___TAG___linux_amd64/cav \
        && sudo mv ./cloudavenue-cli___TAG___linux_amd64/cav /usr/local/bin

=== "Linux (arm64)"

        curl -LO https://github.com/orange-cloudavenue/cloudavenue-cli/releases/download/v__TAG__/cloudavenue-cli___TAG___linux_arm64.tar.gz \
        && tar xvf cloudavenue-cli___TAG___linux_arm64.tar.gz \
        && sudo chmod 755 ./cloudavenue-cli___TAG___linux_arm64/cav \
        && sudo mv ./cloudavenue-cli___TAG___linux_arm64/cav /usr/local/bin

### Go env
``` go
go install github.com/orange-cloudavenue/cloudavenue-cli@latest
```