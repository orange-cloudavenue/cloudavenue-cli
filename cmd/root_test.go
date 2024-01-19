package cmd_test

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/orange-cloudavenue/cloudavenue-cli/cmd"
)

const (
	get         string = "get"
	create      string = "create"
	delete      string = "delete"
	help        string = "help"
	version     string = "version"
	vdc         string = "vdc"
	s3          string = "s3"
	edgeGateway string = "edgegateway"
	publicip    string = "publicip"
)

type tt struct {
	name string
	args []string
	fail bool
}
type tts []tt

func TestRootCmd(t *testing.T) {
	// ? Test configuration
	// Bad configuration
	t.Run("Bad Configuration", func(t *testing.T) {
		os.Setenv("ENDPOINT", "")
		if err := cmd.Execute(); err != nil {
			check := err.Error()
			if !strings.Contains(check, "Error in CloudAvenue parameter") {
				t.Errorf("Fail %v", err)
			}
		}
	})

	// Good configuration
	t.Run("Configuration", func(t *testing.T) {
		os.Setenv("CLOUDAVENUE_URL", "https://console1.cloudavenue.orange-business.com")
		if err := cmd.Execute(); err != nil {
			t.Errorf("Fail %v", err)
		}
	})

	// ? Test all commands
	allCmd := cmd.NewRootCmd().Commands()
	if len(allCmd) == 0 {
		panic("No command found")
	}
	for _, oneCmd := range allCmd {

		tests := tts{}
		switch oneCmd.Use {
		// ? Test list argument
		case get:
			// ? Test all subcommands
			subCmd := oneCmd.Commands()
			for _, cmdSubCmd := range subCmd {
				tests = tts{
					{
						name: oneCmd.Use + "_" + cmdSubCmd.Use + " without args",
						args: []string{oneCmd.Use, cmdSubCmd.Use},
					},
					{
						name: oneCmd.Use + "_" + cmdSubCmd.Use + " with a whatever flag",
						args: []string{oneCmd.Use, cmdSubCmd.Use, "--whatever"},
						fail: true,
					},
					{
						name: oneCmd.Use + "_" + cmdSubCmd.Use + " with ouput flag",
						args: []string{oneCmd.Use, cmdSubCmd.Use, "--output", "wide"},
					},
					{
						name: oneCmd.Use + "_" + cmdSubCmd.Use + " with output flag and a whatever args",
						args: []string{oneCmd.Use, cmdSubCmd.Use, "--output", "whatever"},
						fail: true,
					},
					{
						name: oneCmd.Use + "_" + cmdSubCmd.Use + " with time flag",
						args: []string{oneCmd.Use, cmdSubCmd.Use, "--time"},
					},
				}
				startTest(tests, t)
			}
		case create:
			// ? Test all subcommands
			subCmd := oneCmd.Commands()
			for _, cmdSubCmd := range subCmd {
				switch cmdSubCmd.Use {
				case vdc, s3, publicip:
					tests = tts{
						{
							name: oneCmd.Use + "_" + cmdSubCmd.Use + " without args",
							args: []string{oneCmd.Use, cmdSubCmd.Use},
							fail: true,
						},
						{
							name: oneCmd.Use + "_" + cmdSubCmd.Use + " with a whatever flag",
							args: []string{oneCmd.Use, cmdSubCmd.Use, "--whatever"},
							fail: true,
						},
						{
							name: oneCmd.Use + "_" + cmdSubCmd.Use + " with name flag",
							args: []string{oneCmd.Use, cmdSubCmd.Use, "--name", "whatever"},
							fail: false,
						},
						{
							name: oneCmd.Use + "_" + cmdSubCmd.Use + " with name flag and an empty args",
							args: []string{oneCmd.Use, cmdSubCmd.Use, "--name"},
							fail: true,
						},
						{
							name: oneCmd.Use + "_" + cmdSubCmd.Use + " with time flag",
							args: []string{oneCmd.Use, cmdSubCmd.Use, "--time"},
						},
					}
				case edgeGateway:
					tests = tts{
						{
							name: oneCmd.Use + "_" + cmdSubCmd.Use + " without args",
							args: []string{oneCmd.Use, cmdSubCmd.Use},
							fail: true,
						},
						{
							name: oneCmd.Use + "_" + cmdSubCmd.Use + " with a whatever flag",
							args: []string{oneCmd.Use, cmdSubCmd.Use, "--whatever"},
							fail: true,
						},
						{
							name: oneCmd.Use + "_" + cmdSubCmd.Use + " with vdc flag",
							args: []string{oneCmd.Use, cmdSubCmd.Use, "--vdc", "whatever"},
							fail: false,
						},
						{
							name: oneCmd.Use + "_" + cmdSubCmd.Use + " with vdc flag and an empty args",
							args: []string{oneCmd.Use, cmdSubCmd.Use, "--vdc"},
							fail: true,
						},
						{
							name: oneCmd.Use + "_" + cmdSubCmd.Use + " with time flag",
							args: []string{oneCmd.Use, cmdSubCmd.Use, "--time"},
						},
					}
				}
				startTest(tests, t)
			}

		// ? Test help argument
		case help:
			tests = tts{
				{
					name: oneCmd.Use,
					args: []string{oneCmd.Use},
					fail: false,
				},
			}

		// ? Test version argument
		case version:
			tests = tts{
				{
					name: oneCmd.Use,
					args: []string{oneCmd.Use},
					fail: false,
				},
			}
		}
		startTest(tests, t)
	}

	// ? Test Delete command
	tests := tts{
		{
			name: delete + " without args",
			args: []string{delete},
			fail: true,
		},
		{
			name: delete + " with a whatever flag",
			args: []string{delete, "--whatever"},
			fail: true,
		},
		{
			name: delete + "a bucket",
			args: []string{delete, "s3", "whatever"},
			fail: false,
		},
		{
			name: delete + "a vdc",
			args: []string{delete, "vdc", "whatever"},
			fail: false,
		},
		{
			name: delete + "an edgegateway",
			args: []string{delete, "edgegateway", "whatever"},
			fail: true,
		},
		{
			name: delete + "a publicip",
			args: []string{delete, "publicip", "whatever"},
			fail: true,
		},
	}
	startTest(tests, t)
	// if strings.Contains(oneCmd.Use, cmd.CmdEdgeGateway) {
	// 	tests[3] = tt{
	// 		name: oneCmd.Use + "_" + cmdSubCmd.Use + " with good flags with argument",
	// 		args: []string{oneCmd.Use, cmdSubCmd.Use, "--vdc", "whatever", "--t0", "1"},
	// 		fail: true,
	// 	}
	// }

}

func startTest(tests tts, t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// create the root command
			x := cmd.NewRootCmd()

			// redirect the output to a buffer
			var stdout bytes.Buffer
			x.SetOut(&stdout)

			// set the args
			x.SetArgs(test.args)

			// execute the command
			err := x.Execute()
			if err != nil && !test.fail {
				t.Errorf("Fail %v", err)
			}

		})
	}
}
