// +build windows

package cmd

import "github.com/spf13/cobra"

func getHelpOsConstrained() string {
	return `
	
Some examples:
  check_service.exe --name audiosrv
    Checks for the service to exist and shows the service state and user.
  check_service.exe --name audiosrv --state running
    Checks for the service in the running state.
  check_service.exe --name audiosrv --state running --user "NT AUTHORITY\LocalService"
    Checks for the service in the running state and running as user.
  check_service.exe --name audiosrv --user "NT AUTHORITY\LocalService"
    Checks for the service to exist and would be run as user.
`
}

func addFlagsOsConstrained(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&state, "state", "s", "", "the desired state of the service")
	cmd.Flags().StringVarP(&user, "user", "u", "", "the user the service should run as")
	cmd.Flags().StringVarP(&manager, serviceManagerFlag, "m", "wmi", "Service manager. Allowed options are: \"wmi\" and \"svcmgr\"")
}
