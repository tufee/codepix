/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/tufee/codepix/application/grpc"
	"github.com/tufee/codepix/infrastructure/db"
)

var grpcPortNumber int

// grpcCmd represents the grpc command
var grpcCmd = &cobra.Command{
	Use:   "grpc",
	Short: "Start gRPC server",
	Run: func(cmd *cobra.Command, args []string) {
		database := db.ConnectDB(os.Getenv("env"))
		grpc.StartGrpcServer(database, grpcPortNumber)
	},
}

func init() {
	rootCmd.AddCommand(grpcCmd)
	grpcCmd.Flags().IntVarP(&grpcPortNumber, "port", "p", 50051, "gRPC Port")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// grpcCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// grpcCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
