# Create a module
- mkdir modules
- cd modules
- go mod init github.com/<username>/<module_name>
- create a new file, i.e: utils.go with a package name different from "main"
- upload you code to the specified github repository

# Install 3rd party modules
- go get github.com/<source_username>/<source_module>
  go get github.com/donvito/hellomod
- go get github.com/<source_username>/<source_module>/<version>
  go get github.com/donvito/hellomod/v2
- go.sum is generated with the hash of the current version of the 3rd party software,
  if something change in the source code the checksum will be different and you can check
  if there's no malicious software there.

**Path to the downloaded dependencies**
- go mod download -json

# Build
- go build .\main.go

# Run
- .\main.exe
- go run .\main.go

# Remove unused dependencies
- go mod tidy