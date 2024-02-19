/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var s3Cmd = &cobra.Command{
	Use: "s3",
	Short: "The s3 subcommand downloads and installs certificates from S3",
	
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Not implemented: download certs from s3.")
	},
}

func init() {
	rootCmd.AddCommand(s3Cmd)

	s3Cmd.Flags().String("bucket", "", "Path of s3 bucket where certs are installed")
}
