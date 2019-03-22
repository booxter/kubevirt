#!/usr/bin/env bash

set -e

image="os-3.11.0@sha256:2d0a8f59dfebe181f550c4fbcd90d491a56a7d642d761c32a3c7732644325c0b"

source cluster/ephemeral-provider-common.sh

function up() {
    ${_cli} run --reverse $(_add_common_params)
    prepare_client openshift

    # Update Kube config to support unsecured connection
    ${KUBEVIRT_PATH}cluster/$KUBEVIRT_PROVIDER/.kubectl config set-cluster node01:8443 --server=https://$(_main_ip):$(_port ocp)
    ${KUBEVIRT_PATH}cluster/$KUBEVIRT_PROVIDER/.kubectl config set-cluster node01:8443 --insecure-skip-tls-verify=true

    # Make sure that local config is correct
    prepare_config
}
