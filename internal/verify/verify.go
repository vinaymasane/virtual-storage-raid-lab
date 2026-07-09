package verify

// Run executes a series of verification checks, including checking the serial console, SSH service, RAID configuration, and collecting artifacts. It returns an error if any of the checks fail.
func Run() error {

	checks := []func() error{

		CheckSerial,

		CheckSSH,

		CheckRaid,

		CollectArtifacts,
	}

	for _, c := range checks {

		if err := c(); err != nil {
			return err
		}

	}

	return nil
}
