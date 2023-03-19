# AppGraph

## Commands

```sh
    ./ag svc create myserviceOne --csi 373468 --type component --repo=bitbucket/myapp/src/myserviceOne -> serviceId : 1000
    ./ag svc create myserviceTwo --csi 373468 --type component --repo=bitbucket/myapp/src/myserviceTwo -> serviceId : 2000
    ./ag svc ls --csi 373468
    ./ag svc gen-sbom --id 1000 ./src/myserviceOne
    ./ag permit request --csi 373468 --type build 1000 2000 
    ./ag permit ls --csi 373468
    ./ag permit request --csi 373468 --type deploy 1

``