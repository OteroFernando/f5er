package main

import (
	"fmt"
	"github.com/jmcvetta/napping"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

var (
	f5Host    string
	username  string
	passwd    string
	cfgFile   string = "f5.json"
	f5Input   string
	f5Pool    string
	session   napping.Session
	transport *http.Transport
	client    *http.Client
	now       bool
)

func InitialiseConfig() {
	viper.SetConfigFile(cfgFile)
	viper.AddConfigPath(".")
	viper.SetDefault("f5", "192.168.0.1")
	viper.SetDefault("debug", false)
	viper.SetDefault("force", false)

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Can't find your config file: %s\n", cfgFile)
	}
	viper.AutomaticEnv()

	viper.BindPFlag("f5", f5Cmd.PersistentFlags().Lookup("f5"))
	viper.BindPFlag("debug", f5Cmd.PersistentFlags().Lookup("debug"))
	viper.BindPFlag("pool", onlinePoolMemberCmd.Flags().Lookup("pool"))
	viper.BindPFlag("pool", offlinePoolMemberCmd.Flags().Lookup("pool"))
	viper.BindPFlag("input", f5Cmd.PersistentFlags().Lookup("input"))

	if f5Cmd.PersistentFlags().Lookup("f5").Changed {
		// use cmdline f5 flag if supplied
		viper.Set("f5", f5Host)
	}
	if f5Cmd.PersistentFlags().Lookup("debug").Changed {
		viper.Set("debug", true)
	}
	if f5Cmd.PersistentFlags().Lookup("input").Changed {
		viper.Set("input", f5Input)
	}
	if offlinePoolMemberCmd.Flags().Lookup("pool").Changed {
		viper.Set("pool", f5Pool)
	}
	if offlinePoolMemberCmd.Flags().Lookup("now").Changed {
		viper.Set("now", true)
	}
	if onlinePoolMemberCmd.Flags().Lookup("pool").Changed {
		viper.Set("pool", f5Pool)
	}

	debug = viper.GetBool("debug")
	now = viper.GetBool("now")
	username = viper.GetString("username")
	passwd = viper.GetString("passwd")
	f5Host = viper.GetString("f5")

	if username == "" {
		log.Fatalf("\nerror: missing required username configurable in %s\n\n", cfgFile)
	}
	if passwd == "" {
		log.Fatalf("\nerror: missing required passwd configurable in %s\n\n", cfgFile)
	}
	// finally check that f5 is not an empty string (default)
	//	if f5Host != "" {
	//		viper.Set("f5", f5Host)
	//	} else {
	//		f5Host = viper.GetString("f5")
	//	}
	if f5Host == "" {
		log.Fatalf("\nerror: missing required option --f5\n\n")
	}

}

func checkRequiredFlag(flg string) {
	if !viper.IsSet(flg) {
		log.SetFlags(0)
		log.Fatalf("\nerror: missing required option --%s\n\n", flg)
	}
}

func init() {

	f5Cmd.PersistentFlags().StringVarP(&f5Host, "f5", "f", "", "IP or hostname of F5 to poke")
	f5Cmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "debug output")
	f5Cmd.PersistentFlags().StringVarP(&f5Input, "input", "i", "", "input json f5 configuration")
	offlinePoolMemberCmd.Flags().StringVarP(&f5Pool, "pool", "p", "", "F5 pool name")
	offlinePoolMemberCmd.Flags().BoolVarP(&now, "now", "n", false, "force member offline immediately")
	onlinePoolMemberCmd.Flags().StringVarP(&f5Pool, "pool", "p", "", "F5 pool name")

	// show
	f5Cmd.AddCommand(showCmd)
	showCmd.AddCommand(showPoolCmd)
	showCmd.AddCommand(showPoolMemberCmd)
	showCmd.AddCommand(showVirtualCmd)
	showCmd.AddCommand(showNodeCmd)
	showCmd.AddCommand(showDeviceCmd)
	showCmd.AddCommand(showRuleCmd)
	showCmd.AddCommand(showStackCmd)

	// add
	f5Cmd.AddCommand(addCmd)
	addCmd.AddCommand(addPoolCmd)
	addCmd.AddCommand(addPoolMemberCmd)
	addCmd.AddCommand(addNodeCmd)
	addCmd.AddCommand(addVirtualCmd)
	addCmd.AddCommand(addRuleCmd)
	addCmd.AddCommand(addStackCmd)

	// update
	f5Cmd.AddCommand(updateCmd)
	updateCmd.AddCommand(updatePoolCmd)
	updateCmd.AddCommand(updatePoolMemberCmd)
	updateCmd.AddCommand(updateNodeCmd)
	updateCmd.AddCommand(updateVirtualCmd)
	updateCmd.AddCommand(updateRuleCmd)
	updateCmd.AddCommand(updateStackCmd)

	// delete
	f5Cmd.AddCommand(deleteCmd)
	deleteCmd.AddCommand(deletePoolCmd)
	deleteCmd.AddCommand(deletePoolMemberCmd)
	deleteCmd.AddCommand(deleteNodeCmd)
	deleteCmd.AddCommand(deleteVirtualCmd)
	deleteCmd.AddCommand(deleteRuleCmd)
	deleteCmd.AddCommand(deleteStackCmd)

	// offline
	f5Cmd.AddCommand(offlineCmd)
	offlineCmd.AddCommand(offlinePoolMemberCmd)

	// online
	f5Cmd.AddCommand(onlineCmd)
	onlineCmd.AddCommand(onlinePoolMemberCmd)

	//	log.SetFlags(log.Ltime | log.Lshortfile)
	log.SetFlags(0)

}

func main() {
	InitialiseConfig()
	InitSession()
	//f5Cmd.DebugFlags()
	f5Cmd.Execute()
}
