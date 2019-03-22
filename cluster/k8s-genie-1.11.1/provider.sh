#!/usr/bin/env bash

set -e

image="k8s-genie-1.11.1@sha256:bd3c56acfd0bdad24e204a49e41d2192eab4cad282a1bb9bed01790874ba8b58"

source cluster/ephemeral-provider-common.sh

function up() {
    ${_cli} run $(_add_common_params)
    prepare_client k8s

    # Set server and disable tls check
    ${KUBEVIRT_PATH}cluster/$KUBEVIRT_PROVIDER/.kubectl config set-cluster kubernetes --server=https://$(_main_ip):$(_port k8s)
    ${KUBEVIRT_PATH}cluster/$KUBEVIRT_PROVIDER/.kubectl config set-cluster kubernetes --insecure-skip-tls-verify=true

    # Make sure that local config is correct
    prepare_config
}
