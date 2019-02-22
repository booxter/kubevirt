#!/bin/sh

# This script enables VFs for all available SR-IOV capable PFs, and registers
# all of them with vfio subsystem. This is so that SR-IOV device plugin can
# then allocate these VFIO enabled devices to pods and have libvirt pass them
# into qemu using vfio device type.

# first, load the kernel module if not already
modprobe vfio-pci

for file in $(find /sys/devices/ -name *sriov_totalvfs*); do
    pfroot=$(dirname $file)

    # enable all available VFs
    echo 0 > $pfroot/sriov_numvfs
    cat $file > $pfroot/sriov_numvfs

    # bind all VFs with vfio
    for virtfn in $(ls -d $pfroot/virtfn*); do
        pciid=$(basename $(readlink $virtfn))
        if [ -e $virtfn/driver/unbind ]; then
            echo $pciid > $virtfn/driver/unbind
        fi
        echo $(lspci -n -s $pciid | sed 's/:/ /g' | awk -e '{print $4 " " $5}') > /sys/bus/pci/drivers/vfio-pci/new_id
        echo "Registered $pciid with vfio-pci"
    done
done
