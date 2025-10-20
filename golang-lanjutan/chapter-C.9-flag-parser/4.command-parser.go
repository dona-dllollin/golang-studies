package main 

import (
	"fmt"
	"os"

	"github.com/alecthomas/kingpin/v2"
)

var app = kingpin.New("App", "Simple app")

var (
	commandAdd = app.Command("add", "add new user")
	commandAddFlagOverride = commandAdd.Flag("override", "override existing user").Short('o').Bool()
	commandAddArgUser = commandAdd.Arg("user", "username").Required().String()	
)

var (
	commandUpdate           = app.Command("update", "update user")
	commandUpdateArgOldUser = commandUpdate.Arg("old", "old username").Required().String()
	commandUpdateArgNewUser = commandUpdate.Arg("new", "new username").Required().String()
)

var (
	commandDelete          = app.Command("delete", "delete user")
	commandDeleteFlagForce = commandDelete.Flag("force", "force deletion").Short('f').Bool()
	commandDeleteArgUser   = commandDelete.Arg("user", "username").Required().String()
)

func main() {
	commandAdd.Action(func(ctx *kingpin.ParseContext) error {
		user := *commandAddArgUser
		override := *commandAddFlagOverride
		fmt.Printf("Adding user %s, override: %v\n", user, override)
		return nil
	})

	commandUpdate.Action(func(ctx *kingpin.ParseContext) error {
		oldUser := *commandUpdateArgOldUser
		newUser := *commandUpdateArgNewUser
		fmt.Printf("Updating user from %s to %s\n", oldUser, newUser)
		return nil
	})

	commandDelete.Action(func(ctx *kingpin.ParseContext) error {
		user := *commandDeleteArgUser
		force := *commandDeleteFlagForce
		fmt.Printf("Deleting user %s, force: %v\n", user, force)
		return nil
	})

	kingpin.MustParse(app.Parse(os.Args[1:]))
}