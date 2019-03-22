#!/usr/bin/env bash

set -e

image="k8s-multus-1.13.3@sha256:d037d12a7c847067b051a627925a344ffcfe191adf5b53dc84ccdabbac510995"

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
