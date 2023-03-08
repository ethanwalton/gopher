/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// showRbacCmd represents the showRbac command
//clusterContext := ""

var showRbacCmd = &cobra.Command{
	Use:   "show-rbac",
	Short: "show cluster rbac and users/groups associated",
	Long: `This command gives the ability to be able to further describe rbac 
	inside of a Kubernetes cluster. It further enables us to be able to easily output cluster
	rbac and determine users and groups that are associated with clusterroles and 
	clusterrolebindings within the cluster`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("showRbac has been called")
		clusterScope := ""
		if clusterScope != "all-clusters" && clusterScope != "cluster" && clusterScope != "namespace" {
			fmt.Println("Cluster scope must be: namespace, cluster or all-clusters")
		}

	},
}

func init() {
	rootCmd.AddCommand(showRbacCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showRbacCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showRbacCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	clusterContext := ""
	clusterScope := ""
	showRbacCmd.Flags().StringVarP(&clusterContext, "cluster-context", "c", "", "Use kubectl config get-context to determine cluster context")
	showRbacCmd.Flags().StringVarP(&clusterScope, "cluster-scope", "s", "", "Scope of Kubernetes -- namespace, cluster, all-clusters")

}
