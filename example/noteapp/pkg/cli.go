package pkg

import (
	"fmt"
	"os"

	"github.com/sourcenetwork/sourcehub/x/acp/types"
	"github.com/spf13/cobra"
)

var NoteCmds *NoteCommands
var PermMan *PermissionManager
var Querier *NoteQuerier
var ACPCl *ACPClient

var RootCmd = cobra.Command{}

var cmdGenAccount = cobra.Command{
	Use: "gen-account",
	RunE: func(cmd *cobra.Command, args []string) error {
		addr, key := NewActor()
		keyStr := DumpKey(key)

		fmt.Println(addr)
		fmt.Println(keyStr)

		return nil
	},
}

var cmdNewNote = cobra.Command{
	Use:  "new-note key-file title note",
	Args: cobra.ExactArgs(3),
	RunE: func(cmd *cobra.Command, args []string) error {
		keyFile := args[0]
		title := args[1]
		body := args[2]

		keyBytes, err := os.ReadFile(keyFile)
		if err != nil {
			return err
		}

		session := NewSession(string(keyBytes))

		note, err := NoteCmds.Create(cmd.Context(), session, title, body)
		if err != nil {
			return err
		}

		cmd.Printf("note %v created\n", note.ID)
		fmt.Printf("%v\n", note.ID)

		return nil
	},
}

var cmdShareNote = cobra.Command{
	Use:  "share key-file noteID collaboratorId",
	Args: cobra.ExactArgs(3),
	RunE: func(cmd *cobra.Command, args []string) error {
		keyFile := args[0]
		noteId := args[1]
		collaboratorId := args[2]

		keyBytes, err := os.ReadFile(keyFile)
		if err != nil {
			return err
		}

		session := NewSession(string(keyBytes))

		promise, err := PermMan.ShareNote(cmd.Context(), session, noteId, collaboratorId)
		if err != nil {
			return err
		}

		cmd.Printf("Broadcast Tx %v\n", promise.GetTxHash())

		cmd.Println("Awaiting Tx")
		_, err = promise.Await()
		if err != nil {
			return err
		}

		return nil
	},
}

var cmdListNotes = cobra.Command{
	Use:  "list-notes key-file",
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		keyFile := args[0]

		keyBytes, err := os.ReadFile(keyFile)
		if err != nil {
			return err
		}

		session := NewSession(string(keyBytes))

		notes, err := Querier.FetchReadableNotes(cmd.Context(), session)
		if err != nil {
			return err
		}

		for _, note := range notes {
			fmt.Printf("%v: %v: %v\n", note.ID, note.Title, note.Body)
		}
		return nil
	},
}

var cmdCreatePolicy = cobra.Command{
	Use:  "new-policy key-file policy-file",
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		keyFile := args[0]
		policyFile := args[1]

		keyBytes, err := os.ReadFile(keyFile)
		if err != nil {
			return err
		}

		policyStr, err := os.ReadFile(policyFile)
		if err != nil {
			return err
		}

		session := NewSession(string(keyBytes))

		msg := types.NewMsgCreatePolicyNow(session.Actor, string(policyStr), types.PolicyMarshalingType_SHORT_YAML)

		promise, err := ACPCl.TxCreatePolicy(cmd.Context(), session, msg)
		if err != nil {
			return err
		}

		polId, err := promise.Await()
		if err != nil {
			return err
		}

		cmd.Printf("Policy %v created\n", polId)

		return nil
	},
}

var cmdUnshareNote = cobra.Command{
	Use:  "unshare key-file noteID collaboratorId",
	Args: cobra.ExactArgs(3),
	RunE: func(cmd *cobra.Command, args []string) error {
		keyFile := args[0]
		noteId := args[1]
		collaboratorId := args[2]

		keyBytes, err := os.ReadFile(keyFile)
		if err != nil {
			return err
		}

		session := NewSession(string(keyBytes))

		promise, err := PermMan.ShareNote(cmd.Context(), session, noteId, collaboratorId)
		if err != nil {
			return err
		}

		cmd.Printf("Broadcast Tx %v\n", promise.GetTxHash())

		cmd.Println("Awaiting Tx")
		_, err = promise.Await()
		if err != nil {
			return err
		}

		return nil
	},
}

var cmdListLocalNotes = cobra.Command{
	Use:  "list-local-notes key-file",
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		keyFile := args[0]

		keyBytes, err := os.ReadFile(keyFile)
		if err != nil {
			return err
		}

		session := NewSession(string(keyBytes))

		notes, err := Querier.FetchLocalNodes(cmd.Context(), session)
		if err != nil {
			return err
		}

		for _, note := range notes {
			fmt.Printf("%v: %v: %v\n", note.ID, note.Title, note.Body)
		}

		return nil
	},
}

func init() {
	RootCmd.AddCommand(
		&cmdGenAccount,
		&cmdNewNote,
		&cmdShareNote,
		&cmdListNotes,
		&cmdCreatePolicy,
		&cmdUnshareNote,
		&cmdListLocalNotes,
	)
}
