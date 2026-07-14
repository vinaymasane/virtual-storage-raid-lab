# Virtual Storage RAID Lab

## Requirements

- Debian 13 Host
- Go 1.24+
- Packer
- QEMU/KVM
- qemu-nbd
- mdadm
- Ansible

## Build

```bash
make bootstrap
make build
```

## Execution

```bash
make image
make mirror
make raid
make launch
make configure
make verify
make collect
```

## Architecture

```
Bootstrap
     |
Packer Build
     |
disk_proto.qcow2
     |
qemu-nbd
     |
RAID1
     |
Launch VM
     |
ttyS0
     |
SSH
     |
Ansible
     |
Validation
     |
Artifacts
```

## Deliverables

- disk_proto.qcow2
- RAID1
- ttyS0
- SSH
- mdadm
- Validation
- Artifact Collection