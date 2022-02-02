package mcobra

import (
	coral "github.com/muesli/coral"
	"github.com/muesli/mango"
	mpflag "github.com/muesli/mango-pflag"
)

// NewManPage creates a new mango.ManPage from a coral.Command.
func NewManPage(section uint, c *coral.Command) (*mango.ManPage, error) {
	manPage := mango.NewManPage(section, c.Name(), c.Short).
		WithLongDescription(c.Long)

	if err := AddCommand(manPage, c); err != nil {
		return nil, err
	}
	return manPage, nil
}

// AddCommand adds a coral.Command to a mango.ManPage.
func AddCommand(m *mango.ManPage, c *coral.Command) error {
	return addCommandTree(m, c, nil)
}

func addCommandTree(m *mango.ManPage, c *coral.Command, parent *mango.Command) error {
	var item *mango.Command
	if parent == nil {
		// set root command
		item = mango.NewCommand(c.Name(), "", "")
		m.Root = *item
	} else {
		item = mango.NewCommand(c.Name(), c.Short, c.Use)
		if err := parent.AddCommand(item); err != nil {
			return err
		}
	}

	// add commands
	if c.HasSubCommands() {
		for _, sub := range c.Commands() {
			if sub.Hidden {
				// ignore hidden commands
				continue
			}

			if err := addCommandTree(m, sub, item); err != nil {
				return err
			}
		}
	}

	// add flags
	if c.HasFlags() {
		c.Flags().VisitAll(mpflag.PFlagCommandVisitor(item))
	}
	if c.HasPersistentFlags() {
		c.PersistentFlags().VisitAll(mpflag.PFlagCommandVisitor(item))
	}

	return nil
}