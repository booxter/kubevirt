#!/usr/bin/env bash

set -e

image="k8s-1.10.11@sha256:b97e556795a56b9aa1763ddf3a49322b49f96877dccb7a164bbca779df078536"

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
