package main

import (
	"fmt"
	"os"

	coral "github.com/muesli/coral"
	mcoral "github.com/muesli/mango-coral"
	"github.com/muesli/roff"
)

var (
	rootCmd = &coral.Command{
		Use:   "mango",
		Short: "A man-page generator",
		Long: "mango is a man-page generator for the Go flag, pflag, and cobra packages.\n" +
			"Features:\n" +
			"* User-friendly\n" +
			"* Plugable",
		RunE: func(cmd *coral.Command, agrs []string) error {
			return nil
		},
		Example: "mango --help",
	}

	oneCmd = &coral.Command{
		Use:     "1 [arg]",
		Example: "1 foobar",
		Short:   "The first command",
		RunE: func(cmd *coral.Command, agrs []string) error {
			return nil
		},
	}

	twoCmd = &coral.Command{
		Use:   "2",
		Short: "The second command",
		RunE: func(cmd *coral.Command, agrs []string) error {
			return nil
		},
	}

	config string
	one    string
	two    string
)

func init() {
	rootCmd.PersistentFlags().StringVar(&config, "config", "", "config file (default is $HOME/.mango.yaml)")
	oneCmd.Flags().StringVar(&one, "one", "", "first value")
	oneCmd.Flags().StringVar(&two, "two", "", "second value")

	rootCmd.AddCommand(oneCmd)
	rootCmd.AddCommand(twoCmd)
}

func main() {
	manPage, err := mcoral.NewManPage(1, rootCmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	manPage = manPage.WithSection("Authors", "mango was written by Christian Muehlhaeuser <https://github.com/muesli/mango>").
		WithSection("Copyright", "Copyright (C) 2022 Christian Muehlhaeuser.\n"+
			"Released under MIT license.")

	fmt.Println(manPage.Build(roff.NewDocument()))
}
