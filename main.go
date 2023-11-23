package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"

	"github.com/pkg/browser"
	"gopkg.in/yaml.v2"
)

type Profiles struct {
	Profiles map[string]Profile `yaml:"profiles"`
}

type Profile struct {
	AuthorizationUrl string            `yaml:"authorizationUrl"`
	ClientId         string            `yaml:"clientId"`
	ClientSecret     string            `yaml:"clientSecret"`
	Scopes           string            `yaml:"scopes"`
	TokenUrl         string            `yaml:"tokenUrl"`
	Params           map[string]string `yaml:"params"`
}

func isFlagPassed(name string) (bool, string) {
	found := false
	configFile := ""
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
			configFile = f.Value.String()
		}
	})
	return found, configFile
}

func createKeyValuePairs(m map[string]string) string {
	b := new(bytes.Buffer)
	for key, value := range m {
		fmt.Fprintf(b, "&%s=%s", key, value)
	}
	return b.String()
}

func main() {
	prof := flag.String("profile", "default", "The configuration profile to use.")
	flag.String("config", ".getto", "The configuration file to use.")
	init := flag.Bool("init", false, "Used to create a configuration file with example configuration profile.")
	printConfig := flag.Bool("print-config", false, "Prints out current configuration file.")
	flag.CommandLine.Usage = func() {
		fmt.Println("GetTo(ken) utility to get access tokens through oauth2 authorization flow")
		flag.Usage()
	}
	flag.Parse()
	if !flag.Parsed() {
		flag.Usage()
		os.Exit(1)
	}
	proffy := *prof

	if *init {
		profiles := Profiles{
			Profiles: map[string]Profile{},
		}

		profile := Profile{
			Params: map[string]string{},
		}

		profile.AuthorizationUrl = "https://accounts.google.com/o/oauth2/v2/auth"
		profile.ClientId = "589478345472-afh60s6u44ikncmonmtk4h6rl23ipdr8.apps.googleusercontent.com"
		profile.ClientSecret = "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"
		profile.Scopes = "https%3A//www.googleapis.com/auth/drive.metadata.readonly"
		profile.TokenUrl = "https://oauth2.googleapis.com/token"
		profile.Params["prompt"] = "consent"
		profiles.Profiles["google"] = profile
		fBytes, e := yaml.Marshal(&profiles)
		if e != nil {
			fmt.Printf("Error while Marshaling. %v", e)
		}
		dirname, er := os.UserHomeDir()
		if er != nil {
			log.Fatal(er)
		}
		filename := fmt.Sprintf("%s/.getto", dirname)
		er = os.WriteFile(filename, fBytes, 0644)
		if er != nil {
			panic("Unable to create configuration file.")
		}
		fmt.Println("Configuration file created at", filename)
		os.Exit(0)
	}

	dirname, er := os.UserHomeDir()
	if er != nil {
		log.Fatal(er)
	}
	configFile := fmt.Sprintf("%s/.getto", dirname)

	configFileSet, setConfigFile := isFlagPassed("config")
	if configFileSet {
		configFile = setConfigFile
	}

	var config Profiles

	fBytes, err := os.ReadFile(configFile)
	if err != nil {
		panic(fmt.Sprintf("Unable to find configuration file %s", configFile))
	}

	if *printConfig {
		fmt.Println(string(fBytes))
		os.Exit(0)
	}

	err = yaml.Unmarshal(fBytes, &config)
	if err != nil {
		panic(fmt.Sprintf("Unable to parse configuration file %s", configFile))
	}

	chosen, ok := config.Profiles[proffy]
	if !ok {
		fmt.Println("Profile", proffy, "does not exist in configuration file.")
		os.Exit(42)
	} else {
		fmt.Println("Using profile", proffy)
	}

	fmt.Println(fetchUserToken(chosen))
}

type AuthResponse struct {
	AccessToken string `json:"access_token"`
}

func doPostRequest(url string, body string) ([]byte, error) {
	bBody := []byte(body)
	bodyReader := bytes.NewReader(bBody)
	req, err := http.NewRequest(http.MethodPost, url, bodyReader)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		os.Exit(1)
	}
	fmt.Println("able to create request to", url)

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("client: could not make request: %s\n", err)
		os.Exit(1)
	}
	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}
	println("able to make request")
	return resBody, nil
}

func fetchUserToken(profile Profile) string {
	fmt.Println(profile.AuthorizationUrl)

	const (
		redirectURL = "http://localhost:4200"
	// 	authorizationUrl = "https://login-demo.curity.io/oauth/v2/oauth-authorize?response_type=code&client_id=%s&redirect_uri=%s&scope=%s&state=%s&prompt=consent&ui_locales=en&nonce=%s"
	)

	// var (
	// 	clientID     = "demo-web-client"
	// 	clientSecret = "6koyn9KpRuofYt2U"
	// )

	// if clientID == "" && clientSecret == "" {
	// 	panic(fmt.Errorf("spotify client ID and secret missing"))
	// }

	// authorization code - received in callback
	code := ""
	// local state parameter for cross-site request forgery prevention
	state := fmt.Sprintf("%d-fFA", rand.Int())

	nonce := fmt.Sprintf("%d-dv4", rand.Int())

	// scope of the access: we want to modify user's playlists
	scope := profile.Scopes

	// extra params
	eparams := createKeyValuePairs(profile.Params)
	// loginURL
	path := fmt.Sprintf("%s?response_type=code&client_id=%s&redirect_uri=%s&scope=%s&state=%s&nonce=%s%s", profile.AuthorizationUrl, profile.ClientId, redirectURL, scope, state, nonce, eparams)

	// channel for signaling that server shutdown can be done
	messages := make(chan bool)

	// callback handler, redirect from authentication is handled here
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// check that the state parameter matches
		if s, ok := r.URL.Query()["state"]; ok && s[0] == state {
			// code is received as query parameter
			if codes, ok := r.URL.Query()["code"]; ok && len(codes) == 1 {
				// save code and signal shutdown
				code = codes[0]
				messages <- true
			}
		}
		// redirect user's browser to spotify home page
		http.Redirect(w, r, "https://www.spotify.com/", http.StatusSeeOther)
	})

	// open user's browser to login page
	if err := browser.OpenURL(path); err != nil {
		panic(fmt.Errorf("failed to open browser for authentication %s", err.Error()))
	}

	server := &http.Server{Addr: ":4200"}
	// go routine for shutting down the server
	go func() {
		okToClose := <-messages
		if okToClose {
			if err := server.Shutdown(context.Background()); err != nil {
				log.Println("Failed to shutdown server", err)
			}
		}
	}()
	// start listening for callback - we don't continue until server is shut down
	log.Println(server.ListenAndServe())

	// authentication complete - fetch the access token
	params := url.Values{}
	params.Add("grant_type", "authorization_code")
	params.Add("code", code)
	params.Add("redirect_uri", redirectURL)
	params.Add("client_id", profile.ClientId)
	params.Add("client_secret", profile.ClientSecret)

	data, err := doPostRequest(
		profile.TokenUrl,
		params.Encode())

	if err == nil {
		print(string(data))
		response := AuthResponse{}
		if err = json.Unmarshal(data, &response); err == nil {
			// happy end: token parsed successfully
			return response.AccessToken
		}
	}
	panic(fmt.Errorf("unable to acquire Spotify user token"))
}
