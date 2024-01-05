package main

import (
	"context"
	"log"

	"github.com/sourcenetwork/sourcehub/example/noteapp/pkg"
)

func main() {
	repo, err := pkg.NewSQLite("./notes.db")
	if err != nil {
		log.Fatal(err)
	}

	listener, err := pkg.NewTxListener("tcp://127.0.0.1:26657")
	if err != nil {
		log.Fatal(err)
	}

	err = listener.Listen(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	client, err := pkg.NewACPClient("127.0.0.1:9090", listener)
	if err != nil {
		log.Fatal(err)
	}

	querier, err := pkg.NewACPQueryClient("127.0.0.1:9090")
	if err != nil {
		log.Fatal(err)
	}

	permMan := pkg.NewPermissionManager(&client)
	noteCmd := pkg.NewNoteCommands(repo, permMan)
	acpQuerier := pkg.NewNoteQuerier(&client, repo, querier)

	pkg.PermMan = permMan
	pkg.NoteCmds = &noteCmd
	pkg.Querier = &acpQuerier
	pkg.ACPCl = &client

	cmd := pkg.RootCmd
	err = cmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
