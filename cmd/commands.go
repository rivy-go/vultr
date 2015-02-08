package cmd

import (
	"fmt"
	"runtime"

	vultr "github.com/JamesClonk/vultr/lib"
	"github.com/jawher/mow.cli"
)

func (c *CLI) RegisterCommands() {
	// info
	c.Command("info", "display account information", accountInfo)

	// os
	c.Command("os", "list all available operating systems", osList)

	// iso
	c.Command("iso", "list all ISOs currently available on account", isoList)

	// plans
	c.Command("plans", "list all active plans", planList)

	// regions
	c.Command("regions", "list all active regions", regionList)

	// sshkeys
	c.Command("sshkey", "modify SSH public keys", func(cmd *cli.Cmd) {
		cmd.Command("create", "upload and add new SSH public key", sshKeysCreate)
		cmd.Command("update", "update an existing SSH public key", sshKeysUpdate)
		cmd.Command("delete", "remove an existing SSH public key", sshKeysDelete)
		cmd.Command("list", "list all existing SSH public keys", sshKeysList)
	})
	c.Command("sshkeys", "list all existing SSH public keys", sshKeysList)

	// ssh
	c.Command("ssh", "ssh into a virtual machine", printAPIKey)

	// servers
	c.Command("server", "modify virtual machines", func(cmd *cli.Cmd) {
		cmd.Command("create", "create a new virtual machine", serversCreate)
		cmd.Command("start", "start a virtual machine (restart if already running)", printAPIKey)
		cmd.Command("halt", "halt a virtual machine (hard power off)", printAPIKey)
		cmd.Command("reboot", "reboot a virtual machine (hard reboot)", printAPIKey)
		cmd.Command("reinstall", "reinstall OS on a virtual machine (all data will be lost)", printAPIKey)
		cmd.Command("change-os", "change OS on a virtual machine (all data will be lost)", printAPIKey)
		cmd.Command("delete", "delete a virtual machine", serversDelete)
		cmd.Command("bandwidth", "list bandwidth used by a virtual machine", printAPIKey)
		cmd.Command("list", "list all active or pending virtual machines on current account", serversList)
		cmd.Command("show", "list detailed information of a virtual machine", serversShow)
	})
	c.Command("servers", "list all active or pending virtual machines on current account", serversList)

	// snapshots
	c.Command("snapshot", "modify snapshots", func(cmd *cli.Cmd) {
		cmd.Command("create", "create a snapshot from an existing virtual machine", printAPIKey)
		cmd.Command("delete", "delete a snapshot", printAPIKey)
		cmd.Command("list", "list all snapshots on current account", printAPIKey)
	})
	c.Command("snapshots", "list all snapshots on current account", printAPIKey)

	// startup scripts
	c.Command("script", "modify startup scripts", func(cmd *cli.Cmd) {
		cmd.Command("create", "create a new startup script", printAPIKey)
		cmd.Command("update", "update an existing startup script", printAPIKey)
		cmd.Command("delete", "remove an existing startup script", printAPIKey)
		cmd.Command("list", "list all startup scripts on current account", printAPIKey)
	})
	c.Command("scripts", "list all startup scripts on current account", printAPIKey)

	// version
	c.Command("version", "vultr CLI version", func(cmd *cli.Cmd) {
		cmd.Action = func() {
			lengths := []int{24, 48}
			printTabbedLine(Columns{"Client version:", vultr.Version}, lengths)
			printTabbedLine(Columns{"Vultr API endpoint:", vultr.DefaultEndpoint}, lengths)
			printTabbedLine(Columns{"Vultr API version:", vultr.APIVersion}, lengths)
			printTabbedLine(Columns{"OS/Arch (client):", fmt.Sprintf("%v/%v", runtime.GOOS, runtime.GOARCH)}, lengths)
			printTabbedLine(Columns{"Go version:", runtime.Version()}, lengths)
			tabsFlush()
		}
	})
}

// for debugging..
func printAPIKey(cmd *cli.Cmd) {
	cmd.Action = func() {
		fmt.Println(*apiKey)
	}
}
