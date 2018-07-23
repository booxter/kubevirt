// Automatically generated by swagger-doc. DO NOT EDIT!

package v1

func (CloudInitNoCloudSource) SwaggerDoc() map[string]string {
	return map[string]string{
		"":               "Represents a cloud-init nocloud user data source.\nMore info: http://cloudinit.readthedocs.io/en/latest/topics/datasources/nocloud.html",
		"secretRef":      "UserDataSecretRef references a k8s secret that contains NoCloud userdata.\n+ optional",
		"userDataBase64": "UserDataBase64 contains NoCloud cloud-init userdata as a base64 encoded string.\n+ optional",
		"userData":       "UserData contains NoCloud inline cloud-init userdata.\n+ optional",
	}
}

func (DomainSpec) SwaggerDoc() map[string]string {
	return map[string]string{
		"resources": "Resources describes the Compute Resources required by this vmi.",
		"cpu":       "CPU allow specified the detailed CPU topology inside the vmi.\n+optional",
		"memory":    "Memory allow specifying the VMI memory features.\n+optional",
		"machine":   "Machine type.\n+optional",
		"firmware":  "Firmware.\n+optional",
		"clock":     "Clock sets the clock and timers of the vmi.\n+optional",
		"features":  "Features like acpi, apic, hyperv.\n+optional",
		"devices":   "Devices allows adding disks, network interfaces, ...",
	}
}

func (DomainPresetSpec) SwaggerDoc() map[string]string {
	return map[string]string{
		"resources": "Resources describes the Compute Resources required by this vmi.",
		"cpu":       "CPU allow specified the detailed CPU topology inside the vmi.\n+optional",
		"memory":    "Memory allow specifying the VMI memory features.\n+optional",
		"machine":   "Machine type.\n+optional",
		"firmware":  "Firmware.\n+optional",
		"clock":     "Clock sets the clock and timers of the vmi.\n+optional",
		"features":  "Features like acpi, apic, hyperv.\n+optional",
		"devices":   "Devices allows adding disks, network interfaces, ...\n+optional",
	}
}

func (ResourceRequirements) SwaggerDoc() map[string]string {
	return map[string]string{
		"requests": "Requests is a description of the initial vmi resources.\nValid resource keys are \"memory\" and \"cpu\".\n+optional",
		"limits":   "Limits describes the maximum amount of compute resources allowed.\nValid resource keys are \"memory\" and \"cpu\".\n+optional",
	}
}

func (CPU) SwaggerDoc() map[string]string {
	return map[string]string{
		"":      "CPU allows specifying the CPU topology.",
		"cores": "Cores specifies the number of cores inside the vmi.\nMust be a value greater or equal 1.",
		"model": "Model specifies the CPU model inside the VMI.\nList of available models https://github.com/libvirt/libvirt/blob/master/src/cpu/cpu_map.xml.\n+optional",
	}
}

func (Memory) SwaggerDoc() map[string]string {
	return map[string]string{
		"":          "Memory allows specifying the VirtualMachineInstance memory features.",
		"hugepages": "Hugepages allow to use hugepages for the VirtualMachineInstance instead of regular memory.\n+optional",
	}
}

func (Hugepages) SwaggerDoc() map[string]string {
	return map[string]string{
		"":         "Hugepages allow to use hugepages for the VirtualMachineInstance instead of regular memory.",
		"pageSize": "PageSize specifies the hugepage size, for x86_64 architecture valid values are 1Gi and 2Mi.",
	}
}

func (Machine) SwaggerDoc() map[string]string {
	return map[string]string{
		"type": "QEMU machine type is the actual chipset of the VirtualMachineInstance.",
	}
}

func (Firmware) SwaggerDoc() map[string]string {
	return map[string]string{
		"uuid": "UUID reported by the vmi bios.\nDefaults to a random generated uid.",
	}
}

func (Devices) SwaggerDoc() map[string]string {
	return map[string]string{
		"disks":      "Disks describes disks, cdroms, floppy and luns which are connected to the vmi.",
		"watchdog":   "Watchdog describes a watchdog device which can be added to the vmi.",
		"interfaces": "Interfaces describe network interfaces which are added to the vm.",
	}
}

func (Disk) SwaggerDoc() map[string]string {
	return map[string]string{
		"name":       "Name is the device name",
		"volumeName": "Name of the volume which is referenced.\nMust match the Name of a Volume.",
		"bootOrder":  "BootOrder is an integer value > 0, used to determine ordering of boot devices.\nLower values take precedence.\nDisks without a boot order are not tried if a disk with a boot order exists.\n+optional",
		"serial":     "Serial provides the ability to specify a serial number for the disk device.\n+optional",
	}
}

func (DiskDevice) SwaggerDoc() map[string]string {
	return map[string]string{
		"":       "Represents the target of a volume to mount.\nOnly one of its members may be specified.",
		"disk":   "Attach a volume as a disk to the vmi.",
		"lun":    "Attach a volume as a LUN to the vmi.",
		"floppy": "Attach a volume as a floppy to the vmi.",
		"cdrom":  "Attach a volume as a cdrom to the vmi.",
	}
}

func (DiskTarget) SwaggerDoc() map[string]string {
	return map[string]string{
		"bus":      "Bus indicates the type of disk device to emulate.\nsupported values: virtio, sata, scsi, ide.",
		"readonly": "ReadOnly.\nDefaults to false.",
	}
}

func (LunTarget) SwaggerDoc() map[string]string {
	return map[string]string{
		"bus":      "Bus indicates the type of disk device to emulate.\nsupported values: virtio, sata, scsi, ide.",
		"readonly": "ReadOnly.\nDefaults to false.",
	}
}

func (FloppyTarget) SwaggerDoc() map[string]string {
	return map[string]string{
		"readonly": "ReadOnly.\nDefaults to false.",
		"tray":     "Tray indicates if the tray of the device is open or closed.\nAllowed values are \"open\" and \"closed\".\nDefaults to closed.\n+optional",
	}
}

func (CDRomTarget) SwaggerDoc() map[string]string {
	return map[string]string{
		"bus":      "Bus indicates the type of disk device to emulate.\nsupported values: virtio, sata, scsi, ide.",
		"readonly": "ReadOnly.\nDefaults to true.",
		"tray":     "Tray indicates if the tray of the device is open or closed.\nAllowed values are \"open\" and \"closed\".\nDefaults to closed.\n+optional",
	}
}

func (Volume) SwaggerDoc() map[string]string {
	return map[string]string{
		"":     "Volume represents a named volume in a vmi.",
		"name": "Volume's name.\nMust be a DNS_LABEL and unique within the vmi.\nMore info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
	}
}

func (VolumeSource) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "Represents the source of a volume to mount.\nOnly one of its members may be specified.",
		"persistentVolumeClaim": "PersistentVolumeClaimVolumeSource represents a reference to a PersistentVolumeClaim in the same namespace.\nDirectly attached to the vmi via qemu.\nMore info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims\n+optional",
		"cloudInitNoCloud":      "CloudInitNoCloud represents a cloud-init NoCloud user-data source.\nThe NoCloud data will be added as a disk to the vmi. A proper cloud-init installation is required inside the guest.\nMore info: http://cloudinit.readthedocs.io/en/latest/topics/datasources/nocloud.html\n+optional",
		"registryDisk":          "RegistryDisk references a docker image, embedding a qcow or raw disk.\nMore info: https://kubevirt.gitbooks.io/user-guide/registry-disk.html\n+optional",
		"ephemeral":             "Ephemeral is a special volume source that \"wraps\" specified source and provides copy-on-write image on top of it.\n+optional",
		"emptyDisk":             "EmptyDisk represents a temporary disk which shares the vmis lifecycle.\nMore info: https://kubevirt.gitbooks.io/user-guide/disks-and-volumes.html\n+optional",
	}
}

func (EphemeralVolumeSource) SwaggerDoc() map[string]string {
	return map[string]string{
		"persistentVolumeClaim": "PersistentVolumeClaimVolumeSource represents a reference to a PersistentVolumeClaim in the same namespace.\nDirectly attached to the vmi via qemu.\nMore info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims\n+optional",
	}
}

func (EmptyDiskSource) SwaggerDoc() map[string]string {
	return map[string]string{
		"":         "EmptyDisk represents a temporary disk which shares the vmis lifecycle.",
		"capacity": "Capacity of the sparse disk.",
	}
}

func (RegistryDiskSource) SwaggerDoc() map[string]string {
	return map[string]string{
		"":                "Represents a docker image with an embedded disk.",
		"image":           "Image is the name of the image with the embedded disk.",
		"imagePullSecret": "ImagePullSecret is the name of the Docker registry secret required to pull the image. The secret must already exist.",
	}
}

func (ClockOffset) SwaggerDoc() map[string]string {
	return map[string]string{
		"":         "Exactly one of its members must be set.",
		"utc":      "UTC sets the guest clock to UTC on each boot. If an offset is specified,\nguest changes to the clock will be kept during reboots and are not reset.",
		"timezone": "Timezone sets the guest clock to the specified timezone.\nZone name follows the TZ environment variable format (e.g. 'America/New_York').",
	}
}

func (ClockOffsetUTC) SwaggerDoc() map[string]string {
	return map[string]string{
		"":              "UTC sets the guest clock to UTC on each boot.",
		"offsetSeconds": "OffsetSeconds specifies an offset in seconds, relative to UTC. If set,\nguest changes to the clock will be kept during reboots and not reset.",
	}
}

func (Clock) SwaggerDoc() map[string]string {
	return map[string]string{
		"":      "Represents the clock and timers of a vmi.",
		"timer": "Timer specifies whih timers are attached to the vmi.",
	}
}

func (Timer) SwaggerDoc() map[string]string {
	return map[string]string{
		"":       "Represents all available timers in a vmi.",
		"hpet":   "HPET (High Precision Event Timer) - multiple timers with periodic interrupts.",
		"kvm":    "KVM \t(KVM clock) - lets guests read the host’s wall clock time (paravirtualized). For linux guests.",
		"pit":    "PIT (Programmable Interval Timer) - a timer with periodic interrupts.",
		"rtc":    "RTC (Real Time Clock) - a continuously running timer with periodic interrupts.",
		"hyperv": "Hyperv (Hypervclock) - lets guests read the host’s wall clock time (paravirtualized). For windows guests.",
	}
}

func (RTCTimer) SwaggerDoc() map[string]string {
	return map[string]string{
		"tickPolicy": "TickPolicy determines what happens when QEMU misses a deadline for injecting a tick to the guest.\nOne of \"delay\", \"catchup\".",
		"present":    "Enabled set to false makes sure that the machine type or a preset can't add the timer.\nDefaults to true.\n+optional",
		"track":      "Track the guest or the wall clock.",
	}
}

func (HPETTimer) SwaggerDoc() map[string]string {
	return map[string]string{
		"tickPolicy": "TickPolicy determines what happens when QEMU misses a deadline for injecting a tick to the guest.\nOne of \"delay\", \"catchup\", \"merge\", \"discard\".",
		"present":    "Enabled set to false makes sure that the machine type or a preset can't add the timer.\nDefaults to true.\n+optional",
	}
}

func (PITTimer) SwaggerDoc() map[string]string {
	return map[string]string{
		"tickPolicy": "TickPolicy determines what happens when QEMU misses a deadline for injecting a tick to the guest.\nOne of \"delay\", \"catchup\", \"discard\".",
		"present":    "Enabled set to false makes sure that the machine type or a preset can't add the timer.\nDefaults to true.\n+optional",
	}
}

func (KVMTimer) SwaggerDoc() map[string]string {
	return map[string]string{
		"present": "Enabled set to false makes sure that the machine type or a preset can't add the timer.\nDefaults to true.\n+optional",
	}
}

func (HypervTimer) SwaggerDoc() map[string]string {
	return map[string]string{
		"present": "Enabled set to false makes sure that the machine type or a preset can't add the timer.\nDefaults to true.\n+optional",
	}
}

func (Features) SwaggerDoc() map[string]string {
	return map[string]string{
		"acpi":   "ACPI enables/disables ACPI insidejsondata guest.\nDefaults to enabled.\n+optional",
		"apic":   "Defaults to the machine type setting.\n+optional",
		"hyperv": "Defaults to the machine type setting.\n+optional",
	}
}

func (FeatureState) SwaggerDoc() map[string]string {
	return map[string]string{
		"":        "Represents if a feature is enabled or disabled.",
		"enabled": "Enabled determines if the feature should be enabled or disabled on the guest.\nDefaults to true.\n+optional",
	}
}

func (FeatureAPIC) SwaggerDoc() map[string]string {
	return map[string]string{
		"enabled":        "Enabled determines if the feature should be enabled or disabled on the guest.\nDefaults to true.\n+optional",
		"endOfInterrupt": "EndOfInterrupt enables the end of interrupt notification in the guest.\nDefaults to false.\n+optional",
	}
}

func (FeatureSpinlocks) SwaggerDoc() map[string]string {
	return map[string]string{
		"enabled":   "Enabled determines if the feature should be enabled or disabled on the guest.\nDefaults to true.\n+optional",
		"spinlocks": "Retries indicates the number of retries.\nMust be a value greater or equal 4096.\nDefaults to 4096.\n+optional",
	}
}

func (FeatureVendorID) SwaggerDoc() map[string]string {
	return map[string]string{
		"enabled":  "Enabled determines if the feature should be enabled or disabled on the guest.\nDefaults to true.\n+optional",
		"vendorid": "VendorID sets the hypervisor vendor id, visible to the vmi.\nString up to twelve characters.",
	}
}

func (FeatureHyperv) SwaggerDoc() map[string]string {
	return map[string]string{
		"":           "Hyperv specific features.",
		"relaxed":    "Relaxed relaxes constraints on timer.\nDefaults to the machine type setting.\n+optional",
		"vapic":      "VAPIC indicates whether virtual APIC is enabled.\nDefaults to the machine type setting.\n+optional",
		"spinlocks":  "Spinlocks indicates if spinlocks should be made available to the guest.\n+optional",
		"vpindex":    "VPIndex enables the Virtual Processor Index to help windows identifying virtual processors.\nDefaults to the machine type setting.\n+optional",
		"runtime":    "Runtime.\nDefaults to the machine type setting.\n+optional",
		"synic":      "SyNIC enable Synthetic Interrupt Controller.\nDefaults to the machine type setting.\n+optional",
		"synictimer": "SyNICTimer enable Synthetic Interrupt Controller timer.\nDefaults to the machine type setting.\n+optional",
		"reset":      "Reset enables Hyperv reboot/reset for the vmi.\nDefaults to the machine type setting.\n+optional",
		"vendorid":   "VendorID allows setting the hypervisor vendor id.\nDefaults to the machine type setting.\n+optional",
	}
}

func (Watchdog) SwaggerDoc() map[string]string {
	return map[string]string{
		"":     "Named watchdog device.",
		"name": "Name of the watchdog.",
	}
}

func (WatchdogDevice) SwaggerDoc() map[string]string {
	return map[string]string{
		"":         "Hardware watchdog device.\nExactly one of its members must be set.",
		"i6300esb": "i6300esb watchdog device.\n+optional",
	}
}

func (I6300ESBWatchdog) SwaggerDoc() map[string]string {
	return map[string]string{
		"":       "i6300esb watchdog device.",
		"action": "The action to take. Valid values are poweroff, reset, shutdown.\nDefaults to reset.",
	}
}

func (Interface) SwaggerDoc() map[string]string {
	return map[string]string{
		"name":       "Logical name of the interface as well as a reference to the associated networks.\nMust match the Name of a Network.",
		"model":      "Interface model.\nOne of: e1000, e1000e, ne2k_pci, pcnet, rtl8139, virtio.\nDefaults to virtio.",
		"ports":      "List of ports to be forwarded to the virtual machine.",
		"macAddress": "Interface MAC address. For example: de:ad:00:00:be:af or DE-AD-00-00-BE-AF.",
	}
}

func (InterfaceBindingMethod) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "Represents the method which will be used to connect the interface to the guest.\nOnly one of its members may be specified.",
	}
}

func (InterfaceBridge) SwaggerDoc() map[string]string {
	return map[string]string{
		"delegateNetworkToGuest": "Whether to delegate pod IP and MAC addresses to the guest.\nFor IP addresses, DHCP is used.\nDefaults to true.",
	}
}

func (InterfaceSlirp) SwaggerDoc() map[string]string {
	return map[string]string{}
}

func (Port) SwaggerDoc() map[string]string {
	return map[string]string{
		"":         "Port repesents a port to expose from the virtual machine.\nDefault protocol TCP.\nThe port field is mandatory",
		"name":     "If specified, this must be an IANA_SVC_NAME and unique within the pod. Each\nnamed port in a pod must have a unique name. Name for the port that can be\nreferred to by services.\n+optional",
		"protocol": "Protocol for port. Must be UDP or TCP.\nDefaults to \"TCP\".\n+optional",
		"port":     "Number of port to expose for the virtual machine.\nThis must be a valid port number, 0 < x < 65536.",
	}
}

func (Network) SwaggerDoc() map[string]string {
	return map[string]string{
		"":     "Network represents a network type and a resource that should be connected to the vm.",
		"name": "Network name.\nMust be a DNS_LABEL and unique within the vm.\nMore info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
	}
}

func (NetworkSource) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "Represents the source resource that will be connected to the vm.\nOnly one of its members may be specified.",
	}
}

func (PodNetwork) SwaggerDoc() map[string]string {
	return map[string]string{
		"":              "Represents the stock pod network interface.",
		"vmNetworkCIDR": "CIDR for vm network.\nDefault 10.0.2.0/24 if not specified.",
	}
}
