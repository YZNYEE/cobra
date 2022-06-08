/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"cobradir/show"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
)

// lsCmd represents the ls command

var islist bool
var dir string
var tree bool
var deep *int64
var hide bool

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "A brief description ofyour command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(islist)
		//fmt.Println("路径:", dir)
		var path string
		if dir == "" {
			path, _ = os.Getwd()
		} else {
			path = dir
		}
		_, err := os.Stat(path)
		if err != nil {
			fmt.Println("路径错误，请仔细检查:" + path)
		}
		fileInfoList, _ := ioutil.ReadDir(path)
		fmt.Println(len(fileInfoList), tree, hide)
		show.Showlist(fileInfoList)
	},
}

var treeCmd = &cobra.Command{Use: "tree",
	Short: "执行树状结构",
	Long:  "打印树状结构的目录列表，--L控制树的最大深度",
	Run: func(cmd *cobra.Command, args []string) {
		var path string
		if dir == "" {
			path, _ = os.Getwd()
		} else {
			path = dir
		}
		fmt.Println(path)
		_, err := os.Stat(path)
		if err == nil {
			fmt.Println("路径错误，请仔细检查:" + path)
			fileInfoList, _ := ioutil.ReadDir(path)
			fmt.Println(len(fileInfoList), tree, hide)
			show.Showtree(fileInfoList, int(*deep), path, hide)
		}
	}}

func init() {
	rootCmd.AddCommand(lsCmd)
	//lsCmd.Flags().BoolP("islist", "l", true, "是否列表展示")
	lsCmd.PersistentFlags().BoolVarP(&islist, "islist", "l", false, "test bool")
	//lsCmd.PersistentFlags().BoolVarP(&tree, "tree", "t", false, "test bool")
	lsCmd.PersistentFlags().BoolVarP(&hide, "hide", "a", false, "test bool")
	lsCmd.PersistentFlags().StringVarP(&dir, "dir", "d", "", "路径")
	//lsCmd.Flags().StringVarP(&dir, "dir", "d", "", "指定路径")
	lsCmd.AddCommand(treeCmd)
	deep = treeCmd.Flags().Int64("L", 1, "test int")
	//fmt.Println("deep:", deep)
	//b := lsCmd.HasAlias("islist")
	//fmt.Println("b:", b)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// lsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
