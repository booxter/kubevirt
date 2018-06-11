package v1

import (
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Defaults", func() {

	It("should add ACPI feature if it is unspecified", func() {
		vm := &VirtualMachine{}
		SetObjectDefaults_VirtualMachine(vm)
		Expect(*vm.Spec.Domain.Features.ACPI.Enabled).To(BeTrue())
	})

	It("should not add non-ACPI feature by default", func() {
		vm := &VirtualMachine{}
		SetObjectDefaults_VirtualMachine(vm)
		Expect(vm.Spec.Domain.Features.APIC).To(BeNil())
		Expect(vm.Spec.Domain.Features.Hyperv).To(BeNil())
	})

	It("should add interface and pod network by default", func() {
		vm := &VirtualMachine{}
		SetObjectDefaults_VirtualMachine(vm)
		Expect(len(vm.Spec.Domain.Devices.Interfaces)).NotTo(BeZero())
		Expect(len(vm.Spec.Networks)).NotTo(BeZero())
	})

	It("should default to true to all defined features", func() {
		vm := &VirtualMachine{
			Spec: VirtualMachineSpec{
				Domain: DomainSpec{},
			},
		}
		vm.Spec.Domain.Features = &Features{
			ACPI: FeatureState{},
			APIC: &FeatureAPIC{},
			Hyperv: &FeatureHyperv{
				Relaxed:    &FeatureState{},
				VAPIC:      &FeatureState{},
				Spinlocks:  &FeatureSpinlocks{},
				VPIndex:    &FeatureState{},
				Runtime:    &FeatureState{},
				SyNIC:      &FeatureState{},
				SyNICTimer: &FeatureState{},
				Reset:      &FeatureState{},
				VendorID:   &FeatureVendorID{},
			},
		}
		SetObjectDefaults_VirtualMachine(vm)

		features := vm.Spec.Domain.Features
		hyperv := features.Hyperv

		Expect(*features.ACPI.Enabled).To(BeTrue())
		Expect(*features.APIC.Enabled).To(BeTrue())
		Expect(*hyperv.Relaxed.Enabled).To(BeTrue())
		Expect(*hyperv.VAPIC.Enabled).To(BeTrue())
		Expect(*hyperv.Spinlocks.Enabled).To(BeTrue())
		Expect(*hyperv.Spinlocks.Retries).To(Equal(uint32(4096)))
		Expect(*hyperv.VPIndex.Enabled).To(BeTrue())
		Expect(*hyperv.Runtime.Enabled).To(BeTrue())
		Expect(*hyperv.SyNIC.Enabled).To(BeTrue())
		Expect(*hyperv.SyNICTimer.Enabled).To(BeTrue())
		Expect(*hyperv.Reset.Enabled).To(BeTrue())
		Expect(*hyperv.VendorID.Enabled).To(BeTrue())
	})

	It("should not override defined feature states", func() {
		vm := &VirtualMachine{
			Spec: VirtualMachineSpec{
				Domain: DomainSpec{},
			},
		}
		vm.Spec.Domain.Features = &Features{
			ACPI: FeatureState{Enabled: _true},
			APIC: &FeatureAPIC{Enabled: _false},
			Hyperv: &FeatureHyperv{
				Relaxed:    &FeatureState{Enabled: _true},
				VAPIC:      &FeatureState{Enabled: _false},
				Spinlocks:  &FeatureSpinlocks{Enabled: _true},
				VPIndex:    &FeatureState{Enabled: _false},
				Runtime:    &FeatureState{Enabled: _true},
				SyNIC:      &FeatureState{Enabled: _false},
				SyNICTimer: &FeatureState{Enabled: _true},
				Reset:      &FeatureState{Enabled: _false},
				VendorID:   &FeatureVendorID{Enabled: _true},
			},
		}
		SetObjectDefaults_VirtualMachine(vm)

		features := vm.Spec.Domain.Features
		hyperv := features.Hyperv

		Expect(*features.ACPI.Enabled).To(BeTrue())
		Expect(*features.APIC.Enabled).To(BeFalse())
		Expect(*hyperv.Relaxed.Enabled).To(BeTrue())
		Expect(*hyperv.VAPIC.Enabled).To(BeFalse())
		Expect(*hyperv.Spinlocks.Enabled).To(BeTrue())
		Expect(*hyperv.Spinlocks.Retries).To(Equal(uint32(4096)))
		Expect(*hyperv.VPIndex.Enabled).To(BeFalse())
		Expect(*hyperv.Runtime.Enabled).To(BeTrue())
		Expect(*hyperv.SyNIC.Enabled).To(BeFalse())
		Expect(*hyperv.SyNICTimer.Enabled).To(BeTrue())
		Expect(*hyperv.Reset.Enabled).To(BeFalse())
		Expect(*hyperv.VendorID.Enabled).To(BeTrue())

		vm.Spec.Domain.Features = &Features{
			ACPI: FeatureState{Enabled: _false},
			APIC: &FeatureAPIC{Enabled: _true},
			Hyperv: &FeatureHyperv{
				Relaxed:    &FeatureState{Enabled: _false},
				VAPIC:      &FeatureState{Enabled: _true},
				Spinlocks:  &FeatureSpinlocks{Enabled: _false},
				VPIndex:    &FeatureState{Enabled: _true},
				Runtime:    &FeatureState{Enabled: _false},
				SyNIC:      &FeatureState{Enabled: _true},
				SyNICTimer: &FeatureState{Enabled: _false},
				Reset:      &FeatureState{Enabled: _true},
				VendorID:   &FeatureVendorID{Enabled: _false},
			},
		}
		SetObjectDefaults_VirtualMachine(vm)

		features = vm.Spec.Domain.Features
		hyperv = features.Hyperv

		Expect(*features.ACPI.Enabled).To(BeFalse())
		Expect(*features.APIC.Enabled).To(BeTrue())
		Expect(*hyperv.Relaxed.Enabled).To(BeFalse())
		Expect(*hyperv.VAPIC.Enabled).To(BeTrue())
		Expect(*hyperv.Spinlocks.Enabled).To(BeFalse())
		Expect(hyperv.Spinlocks.Retries).To(BeNil())
		Expect(*hyperv.VPIndex.Enabled).To(BeTrue())
		Expect(*hyperv.Runtime.Enabled).To(BeFalse())
		Expect(*hyperv.SyNIC.Enabled).To(BeTrue())
		Expect(*hyperv.SyNICTimer.Enabled).To(BeFalse())
		Expect(*hyperv.Reset.Enabled).To(BeTrue())
		Expect(*hyperv.VendorID.Enabled).To(BeFalse())
	})

	It("should set dis defaults", func() {
		vm := &VirtualMachine{
			Spec: VirtualMachineSpec{
				Domain: DomainSpec{
					Devices: Devices{
						Disks: []Disk{
							{
								Name: "cdrom_tray_unspecified",
								DiskDevice: DiskDevice{
									CDRom: &CDRomTarget{},
								},
							},
							{
								Name: "cdrom_tray_open",
								DiskDevice: DiskDevice{
									CDRom: &CDRomTarget{
										Tray:     TrayStateOpen,
										ReadOnly: _false,
									},
								},
							},
							{
								Name: "floppy_tray_unspecified",
								DiskDevice: DiskDevice{
									Floppy: &FloppyTarget{},
								},
							},
							{
								Name: "floppy_tray_open",
								DiskDevice: DiskDevice{
									Floppy: &FloppyTarget{
										Tray:     TrayStateOpen,
										ReadOnly: true,
									},
								},
							},
							{
								Name: "should_default_to_disk",
							},
						},
					},
				},
			},
		}
		SetObjectDefaults_VirtualMachine(vm)
		disks := vm.Spec.Domain.Devices.Disks

		Expect(disks[0].CDRom.Tray).To(Equal(TrayStateClosed))
		Expect(*disks[0].CDRom.ReadOnly).To(Equal(true))
		Expect(disks[1].CDRom.Tray).To(Equal(TrayStateOpen))
		Expect(*disks[1].CDRom.ReadOnly).To(Equal(false))
		Expect(disks[2].Floppy.Tray).To(Equal(TrayStateClosed))
		Expect(disks[2].Floppy.ReadOnly).To(Equal(false))
		Expect(disks[3].Floppy.Tray).To(Equal(TrayStateOpen))
		Expect(disks[3].Floppy.Tray).To(Equal(TrayStateOpen))
		Expect(disks[4].Disk).ToNot(BeNil())
	})

	It("should set the default watchdog and the default watchdog action", func() {
		vm := &VirtualMachine{
			Spec: VirtualMachineSpec{
				Domain: DomainSpec{
					Devices: Devices{
						Watchdog: &Watchdog{
							WatchdogDevice: WatchdogDevice{
								I6300ESB: &I6300ESBWatchdog{},
							},
						},
					},
				},
			},
		}
		SetObjectDefaults_VirtualMachine(vm)
		Expect(vm.Spec.Domain.Devices.Watchdog.I6300ESB.Action).To(Equal(WatchdogActionReset))
		vm.Spec.Domain.Devices.Watchdog.I6300ESB = nil
		SetObjectDefaults_VirtualMachine(vm)
		Expect(vm.Spec.Domain.Devices.Watchdog.I6300ESB).ToNot(BeNil())
		Expect(vm.Spec.Domain.Devices.Watchdog.I6300ESB.Action).To(Equal(WatchdogActionReset))
	})

	It("should set timer defaults", func() {
		vm := &VirtualMachine{
			Spec: VirtualMachineSpec{
				Domain: DomainSpec{
					Clock: &Clock{
						Timer: &Timer{
							HPET:   &HPETTimer{},
							KVM:    &KVMTimer{},
							PIT:    &PITTimer{},
							RTC:    &RTCTimer{},
							Hyperv: &HypervTimer{},
						},
					},
				},
			},
		}
		SetObjectDefaults_VirtualMachine(vm)
		timer := vm.Spec.Domain.Clock.Timer
		Expect(*timer.HPET.Enabled).To(BeTrue())
		Expect(*timer.KVM.Enabled).To(BeTrue())
		Expect(*timer.PIT.Enabled).To(BeTrue())
		Expect(*timer.RTC.Enabled).To(BeTrue())
		Expect(*timer.Hyperv.Enabled).To(BeTrue())
	})
})

var _ = Describe("Function GetNumberOfPodInterfaces()", func() {

	It("should work for empty network list", func() {
		spec := &VirtualMachineSpec{}
		Expect(GetNumberOfPodInterfaces(spec)).To(Equal(0))
	})

	It("should work for non-empty network list without pod network", func() {
		spec := &VirtualMachineSpec{}
		net := Network{}
		spec.Networks = []Network{net}
		Expect(GetNumberOfPodInterfaces(spec)).To(Equal(0))
	})

	It("should work for pod network with missing pod interface", func() {
		spec := &VirtualMachineSpec{}
		net := Network{
			NetworkSource: NetworkSource{
				Pod: &PodNetwork{},
			},
		}
		spec.Networks = []Network{net}
		Expect(GetNumberOfPodInterfaces(spec)).To(Equal(0))
	})

	It("should work for valid pod network / interface combination", func() {
		spec := &VirtualMachineSpec{}
		net := Network{
			NetworkSource: NetworkSource{
				Pod: &PodNetwork{},
			},
			Name: "testnet",
		}
		iface := Interface{Name: net.Name}
		spec.Networks = []Network{net}
		spec.Domain.Devices.Interfaces = []Interface{iface}
		Expect(GetNumberOfPodInterfaces(spec)).To(Equal(1))
	})

	It("should work for multiple pod network / interface combinations", func() {
		spec := &VirtualMachineSpec{}
		net1 := Network{
			NetworkSource: NetworkSource{
				Pod: &PodNetwork{},
			},
			Name: "testnet1",
		}
		iface1 := Interface{Name: net1.Name}
		net2 := Network{
			NetworkSource: NetworkSource{
				Pod: &PodNetwork{},
			},
			Name: "testnet2",
		}
		iface2 := Interface{Name: net2.Name}
		spec.Networks = []Network{net1, net2}
		spec.Domain.Devices.Interfaces = []Interface{iface1, iface2}
		Expect(GetNumberOfPodInterfaces(spec)).To(Equal(2))
	})
})

var _ = Describe("Function setDefaults_NetworkInterface()", func() {

	It("should append pod interface", func() {
		vm := &VirtualMachine{}
		net := Network{
			Name: "testnet",
		}
		iface := Interface{Name: net.Name}
		vm.Spec.Networks = []Network{net}
		vm.Spec.Domain.Devices.Interfaces = []Interface{iface}

		setDefaults_NetworkInterface(vm)
		Expect(len(vm.Spec.Domain.Devices.Interfaces)).To(Equal(2))
		Expect(vm.Spec.Domain.Devices.Interfaces[0].Name).To(Equal("testnet"))
		Expect(vm.Spec.Networks[0].Name).To(Equal("testnet"))
		Expect(vm.Spec.Networks[0].Pod).To(BeNil())

		defaultName := vm.Spec.Domain.Devices.Interfaces[1].Name
		Expect(strings.HasPrefix(defaultName, "default-")).To(BeTrue())
		Expect(vm.Spec.Networks[1].Name).To(Equal(defaultName))
		Expect(vm.Spec.Networks[1].Pod).ToNot(BeNil())
	})
})
