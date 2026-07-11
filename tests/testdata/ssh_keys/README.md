SSH keys are intentionally not committed.

Generate a temporary key pair before running integration tests:

mkdir -p tests/testdata/ssh_keys

ssh-keygen \
    -t ed25519 \
    -N "" \
    -f tests/testdata/ssh_keys/id_ed25519.pub

The bootstrap process should copy id_ed25519.pub into the VM's
authorized_keys before Ansible configuration.