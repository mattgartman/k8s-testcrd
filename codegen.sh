# ROOT_PACKAGE :: the package (relative to $GOPATH/src) that is the target for code generation
ROOT_PACKAGE="github.com/mattgartman/k8s-testcrd"
# CUSTOM_RESOURCE_NAME :: the name of the custom resource that we're generating client code for
CUSTOM_RESOURCE_NAME="coolresource"
# CUSTOM_RESOURCE_VERSION :: the version of the resource
CUSTOM_RESOURCE_VERSION="v1"

# retrieve the code-generator scripts and bins
go get -u k8s.io/code-generator/...
cd $GOPATH/src/k8s.io/code-generator

# run the code-generator entrypoint script
./generate-groups.sh all "$ROOT_PACKAGE/pkg/client" "$ROOT_PACKAGE/pkg/apis" "$CUSTOM_RESOURCE_NAME:$CUSTOM_RESOURCE_VERSION"

# view the newly generated files
tree $GOPATH/src/$ROOT_PACKAGE/pkg/client