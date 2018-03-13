package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	// Version when binary build
	Version  string
	// Revision when binary build
	Revision string
)

var (
	cre = flag.String("credential", "", "credential file path. ex: credential.yml")
	sec = flag.String("secret", "", "client secret file path. ex: client_secret.json")
	sco = flag.String("scope", "", "Google OAuth scope. see: https://developers.google.com/identity/protocols/googlescopes")
	ver = flag.Bool("version", false, "print version")
	hel = flag.Bool("help", false, "print usage")
)

func printVersion() {
	fmt.Printf("version: %s(%s)\n", Version, Revision)
}

func getToken(config *oauth2.Config) (*oauth2.Token, error) {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)

	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var code string

	if _, err := fmt.Scan(&code); err != nil {
		fmt.Errorf("Unable to read authorization code %v", err)
		return nil, err
	}

	token, err := config.Exchange(oauth2.NoContext, code)

	if err != nil {
		fmt.Errorf("Unable to retrieve token from web %v", err)
		return nil, err
	}

	return token, nil
}

func saveToken(file string, token *oauth2.Token) error {
	fmt.Printf("Saving credential file to: %s\n", file)

	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)

	if err != nil {
		fmt.Errorf("Unable to cache oauth token: %v", err)
		return err
	}

	defer f.Close()

	err = json.NewEncoder(f).Encode(token)

	if err != nil {
		fmt.Errorf("Unable to write oauth token: %v", err)
		return err
	}

	return nil
}

func run(secret string, credential string, scope string) (int, error) {
	b, err := ioutil.ReadFile(secret)

	if err != nil {
		fmt.Errorf("Unable to read client secret file: %v", err)
		return 1, err
	}

	config, err := google.ConfigFromJSON(b, scope)

	if err != nil {
		fmt.Errorf("Unable to parse client secret file to config: %v", err)
		return 1, err
	}

	token, err := getToken(config)

	if err != nil {
		fmt.Errorf("Unable to get token: %v", err)
		return 1, err
	}

	if err = saveToken(credential, token); err != nil {
		fmt.Errorf("Unable to get token: %v", err)
		return 1, err
	}

	return 0, nil
}

func main() {
	flag.Parse()

	if *ver {
		printVersion()
		os.Exit(0)
	}

	if *hel {
		flag.Usage()
		printVersion()
		os.Exit(0)
	}

	if len(*sec) == 0 || len(*cre) == 0 || len(*sco) == 0 {
		fmt.Println("[ERROR] You must specify -secret, -credential and scope options.\n")
		flag.Usage()
		printVersion()
		os.Exit(1)
	}

	ret, err := run(*sec, *cre, *sco)

	if err != nil {
		fmt.Errorf("Failed to run: %v", err)
	}

	os.Exit(ret)
}
