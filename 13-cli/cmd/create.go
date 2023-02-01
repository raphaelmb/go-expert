/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/raphaelmb/go-expert/13-cli/internal/database"
	"github.com/spf13/cobra"
)

func newCreatecCmd(categoryDb database.Category) *cobra.Command {
	return &cobra.Command{
		Use: "create",
		Short: "Create a new category",
		Long: `Create a new category`,
		RunE: runCreate(categoryDb),
	}
}

func runCreate(categoryDb database.Category) RunEFunc {
	return func(cmd *cobra.Command, args []string) error {
		name, _ := cmd.Flags().GetString("name")
		description, _ := cmd.Flags().GetString("description")
		_, err := categoryDb.Create(name, description)
		if err != nil {
			return err
		}
		return nil
	}
}

func init() {
	createCmd := newCreatecCmd(GetCategoryDB(GetDb()))
	categoryCmd.AddCommand(createCmd)
	createCmd.Flags().StringP("name", "n", "", "Name of the category")
	createCmd.Flags().StringP("description", "d", "", "Description of the category")
}
