package cmd

import (
	"strconv"
	"strings"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

// samCmd represents the sam command
var samCmd = &cobra.Command{
	Use:   "sam",
	Short: "square-and-multiply",
	Long:  `Calculates the square-and-multiply algorithm.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		mod, err := cmd.Flags().GetInt("mod")
		if err != nil {
			return err
		}
		base, err := cmd.Flags().GetInt("base")
		if err != nil {
			return err
		}
		exp, err := cmd.Flags().GetInt("exp")
		if err != nil {
			return err
		}

		pterm.Println(strconv.FormatInt(int64(exp), 2))
		binSplit := strings.Split(strconv.FormatInt(int64(exp), 2), "")
		var sOrm string
		var output string

		xMod := base

		for i, s := range binSplit {
			if i >= 1 {
				xMod = (xMod * xMod) % mod
				sOrm += pterm.Sprintln(i, ": sq ")
				output += pterm.Sprintln(xMod, " mod ", mod)
				if s == "1" {
					xMod = (xMod * base) % mod
					sOrm += pterm.Sprintln(i, ": mul ")
					output += pterm.Sprintln(xMod, " mod ", mod)
				}
			}
		}

		panels := pterm.Panels{
			{{sOrm}, {output}},
		}

		_ = pterm.DefaultPanel.WithPanels(panels).WithPadding(2).Render()

		return nil
	},
}

func init() {
	rootCmd.AddCommand(samCmd)

	samCmd.Flags().IntP("mod", "m", 0, "modulo")
	samCmd.MarkFlagRequired("mod")
	samCmd.Flags().IntP("exp", "e", 0, "exponent")
	samCmd.MarkFlagRequired("exp")
	samCmd.Flags().IntP("base", "b", 0, "base")
	samCmd.MarkFlagRequired("base")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// samCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// samCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
