package verify

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
