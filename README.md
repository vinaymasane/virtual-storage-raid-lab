# Virtual Storage RAID Lab
Tech Stack : RAID Storage Virtualization

## Objective
Build a production-grade Linux image capable of:

* Booting under QEMU/KVM
* Supporting SSH access
* Being exposed as a host block device via qemu-nbd
* Participating in a software RAID1 mirror using mdadm
* Successfully booting as a RAID-backed virtual machine

## System Architecture

```
 disk_proto.qcow2
        |
     qemu-nbd
        |
    /dev/nbd0
        |
        +----------+
                   |
             secondary.raw
                   |
             /dev/loop0
                   |
             mdadm RAID1
                   |
                /dev/md0
                   |
               ext4 root
                   |
                GRUB2
                   |
             QEMU Virtual VM
```


## Host Environment

### Recommended Host OS:

Debian 13 (Trixie)

### Required Packages:

Refer to the PACKAGEs list under file: `./bootstrap/preinstall_host.sh`

### Config and Validation:
```
cd virtual-storage-raid-lab

chmod +x bootstrap/*.sh

sudo ./bootstrap/preinstall_host.sh

./bootstrap/validate_host.sh
```
