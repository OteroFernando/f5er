package main

import (
	"crypto/tls"
	"errors"
	"fmt"
	//	"github.com/kr/pretty"
	"github.com/jmcvetta/napping"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"net/url"
)

var (
	f5Host      string
	username    string
	passwd      string
	credentials map[string]string
	debug       bool
	cfgFile     string = "f5.json"
	f5Input     string
	f5Pool      string
	transport   *http.Transport
	client      *http.Client
	session     napping.Session
	now         bool
)

type httperr struct {
	Message string
	Errors  []struct {
		Resource string
		Field    string
		Code     string
	}
}

func InitialiseConfig() {
	viper.SetConfigFile(cfgFile)
	viper.AddConfigPath(".")
	viper.SetDefault("debug", false)
	viper.SetDefault("force", false)

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Can't find your config file: %s\n", cfgFile)
	}
	viper.AutomaticEnv()

	if !viper.IsSet("credentials") {
		log.Fatal("no login credentials defined in config")
	}
	credentials = viper.GetStringMapString("credentials")
	var ok bool
	username, ok = credentials["username"]
	if !ok {
		log.Fatal("no username defined in config")
	}
	passwd, ok = credentials["passwd"]
	if !ok {
		log.Fatal("no passwd defined in config")
	}
	f5Host, ok = credentials["f5"]
	if !ok {
		log.Fatal("no f5 defined in config")
	}

	viper.BindPFlag("pool", onlinePoolMemberCmd.Flags().Lookup("pool"))
	viper.BindPFlag("pool", offlinePoolMemberCmd.Flags().Lookup("pool"))
	if f5Cmd.PersistentFlags().Lookup("f5").Changed {
		viper.Set("f5", f5Host)
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
}

func checkRequiredFlag(flg string) {
	if !viper.IsSet(flg) {
		log.SetFlags(0)
		log.Fatalf("\nerror: missing required option --%s\n\n", flg)
	}
}

func bail(msg string) {
	log.SetFlags(0)
	log.Fatalf("\n%s\n\n", msg)
}

func GetRequest(u string, res interface{}) (error, *napping.Response) {

	// REST connection setup
	transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client = &http.Client{Transport: transport}
	//
	// Setup HTTP Basic auth for this session (ONLY use this with SSL).  Auth
	// can also be configured on a per-request basis when using Send().
	//
	session = napping.Session{
		Client:   client,
		Log:      debug,
		Userinfo: url.UserPassword(username, passwd),
	}
	//
	// Send request to server
	//
	e := httperr{}
	resp, err := session.Get(u, nil, &res, &e)
	if err != nil {
		return err, resp
	}
	if resp.Status() == 401 {
		return errors.New("unauthorised - check your username and passwd"), resp
	}
	if resp.Status() >= 300 {
		return errors.New(e.Message), resp
	} else {

		// all is good in the world
		return nil, resp
	}
}

func DeleteRequest(u string, res interface{}) (error, *napping.Response) {

	// REST connection setup
	transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client = &http.Client{Transport: transport}
	//
	// Setup HTTP Basic auth for this session (ONLY use this with SSL).  Auth
	// can also be configured on a per-request basis when using Send().
	//
	session = napping.Session{
		Client:   client,
		Log:      debug,
		Userinfo: url.UserPassword(username, passwd),
	}
	//
	// Send request to server
	//
	e := httperr{}
	resp, err := session.Delete(u, &res, &e)
	if err != nil {
		return err, resp
	}
	if resp.Status() == 401 {
		return errors.New("unauthorised - check your username and passwd"), resp
	}
	if resp.Status() >= 300 {
		return errors.New(e.Message), resp
	} else {
		// all is good in the world
		return nil, resp
	}
}

func PostRequest(u string, pload interface{}, res interface{}) (error, *napping.Response) {

	// REST connection setup
	transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client = &http.Client{Transport: transport}
	//
	// Setup HTTP Basic auth for this session (ONLY use this with SSL).  Auth
	// can also be configured on a per-request basis when using Send().
	//
	session = napping.Session{
		Client:   client,
		Log:      debug,
		Userinfo: url.UserPassword(username, passwd),
	}
	//
	// Send request to server
	//
	e := httperr{}
	resp, err := session.Post(u, &pload, &res, &e)
	if err != nil {
		return err, resp
	}
	if resp.Status() == 401 {
		return errors.New("unauthorised - check your username and passwd"), resp
	}
	if resp.Status() >= 300 {
		return errors.New(e.Message), resp
	} else {

		// all is good in the world
		return nil, resp
	}
}

func PutRequest(u string, pload interface{}, res interface{}) (error, *napping.Response) {

	// REST connection setup
	transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client = &http.Client{Transport: transport}
	//
	// Setup HTTP Basic auth for this session (ONLY use this with SSL).  Auth
	// can also be configured on a per-request basis when using Send().
	//
	session = napping.Session{
		Client:   client,
		Log:      debug,
		Userinfo: url.UserPassword(username, passwd),
	}
	//
	// Send request to server
	//
	e := httperr{}
	resp, err := session.Put(u, &pload, &res, &e)
	if err != nil {
		return err, resp
	}
	if resp.Status() == 401 {
		return errors.New("unauthorised - check your username and passwd"), resp
	}
	if resp.Status() >= 300 {
		return errors.New(e.Message), resp
	} else {
		// all is good in the world
		return nil, resp
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

	// add
	f5Cmd.AddCommand(addCmd)
	addCmd.AddCommand(addPoolCmd)
	addCmd.AddCommand(addPoolMemberCmd)
	addCmd.AddCommand(addNodeCmd)
	addCmd.AddCommand(addVirtualCmd)

	// update
	f5Cmd.AddCommand(updateCmd)
	updateCmd.AddCommand(updatePoolCmd)
	updateCmd.AddCommand(updatePoolMemberCmd)
	updateCmd.AddCommand(updateNodeCmd)
	addCmd.AddCommand(updateVirtualCmd)

	// delete
	f5Cmd.AddCommand(deleteCmd)
	deleteCmd.AddCommand(deletePoolCmd)
	deleteCmd.AddCommand(deletePoolMemberCmd)
	deleteCmd.AddCommand(deleteNodeCmd)
	addCmd.AddCommand(deleteVirtualCmd)

	// offline
	f5Cmd.AddCommand(offlineCmd)
	offlineCmd.AddCommand(offlinePoolMemberCmd)

	// online
	f5Cmd.AddCommand(onlineCmd)
	onlineCmd.AddCommand(onlinePoolMemberCmd)

	//	log.SetFlags(log.Ltime | log.Lshortfile)
	log.SetFlags(0)
	InitialiseConfig()

}

func main() {
	//	f5Cmd.DebugFlags()
	f5Cmd.Execute()
}
