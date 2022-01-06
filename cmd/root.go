/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/0xAX/notificator"
	"github.com/spf13/cobra"
)

func notify(title, body string) error {
	notifier := notificator.New(notificator.Options{
		AppName: "LMK",
	})
	return notifier.Push(title, body, "", notificator.UR_NORMAL)
}

func timer() func() time.Duration {
	start := time.Now()
	return func() time.Duration {
		end := time.Now()
		return end.Sub(start)
	}
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "lmk",
	Short: "Let me know (LMK) notifies when a command finishes execution",
	Long:  ``,
	DisableFlagParsing: true,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		// Execute the args that are passed in
		c := exec.Command(args[0], args[1:]...)
		c.Stderr = os.Stderr
		c.Stdout = os.Stdout
		cmdTimer := timer()
		cmdErr := c.Run()
		totalTime := cmdTimer().Seconds()

		title := fmt.Sprintf("Finished running: %s", args[0])
		body := "Successfully completed command"
		if cmdErr != nil {
			body = fmt.Sprintf("Failed to run command: %s", cmdErr.Error())
		}
		body = fmt.Sprintf("%s\nTook %f seconds", body, totalTime)
		err := notify(title, body)
		if err != nil {
			log.Fatal(err)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.lmk.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
