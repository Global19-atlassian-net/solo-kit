package kube

import (
	"text/template"

	"github.com/solo-io/solo-kit/pkg/code-generator/codegen/templates"
)

// TODO(marco): replace hardcoded types
var GenerateScriptTemplate = template.Must(template.New("kube_generate").Funcs(templates.Funcs).Parse(`
#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

SCRIPT_ROOT=$(dirname "${BASH_SOURCE[0]}")/..
ROOT_PKG={{ .ProjectConfig.GoPackage }}
CLIENT_PKG=${ROOT_PKG}/kube/client
APIS_PKG=${ROOT_PKG}/kube/apis

# Below code is copied from https://github.com/weaveworks/flagger/blob/master/hack/update-codegen.sh
# Grab code-generator version from go.sum.
CODEGEN_PKG=vendor/k8s.io/code-generator

#if [[ ! -d ${CODEGEN_PKG} ]]; then
#    echo "${CODEGEN_PKG} is missing. Run 'go mod vendor'."
#    exit 1
#fi


echo ">> Using ${CODEGEN_PKG}"

# code-generator does work with go.mod but makes assumptions about
# the project living in ` + "$GOPATH/src" + `. To work around this and support
# any location; create a temporary directory, use this as an output
# base, and copy everything back once generated.
TEMP_DIR=$(mktemp -d)
cleanup() {
    echo ">> Removing ${TEMP_DIR}"
    rm -rf ${TEMP_DIR}
}
trap "cleanup" EXIT SIGINT

echo ">> Temporary output directory ${TEMP_DIR}"

# Ensure we can execute.
chmod +x ${CODEGEN_PKG}/generate-groups.sh


${CODEGEN_PKG}/generate-groups.sh all \
    ${CLIENT_PKG} \
    ${APIS_PKG} \
    appmesh:v1beta1 \
    --output-base "${TEMP_DIR}"
# Copy everything back.
cp -a "${TEMP_DIR}/${ROOT_PKG}/." "${SCRIPT_ROOT}/"

`))