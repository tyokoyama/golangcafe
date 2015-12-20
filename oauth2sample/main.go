package main

import (
	"fmt"

	"golang.org/x/net/context"
	"google.golang.org/api/gmail/v1"
	"golang.org/x/oauth2"
)

func main() {
	// Client IDとSecretはRefresh済みです。
	client_id := `328006125971-dold8i9smanincppfpc28ikhi2lomfg8.apps.googleusercontent.com`
	client_secret := `aGwy2V20jfpgSGpz0FzMrf7A`

	conf := &oauth2.Config {
		ClientID: client_id,
		ClientSecret: client_secret,
		Scopes: []string{`https://www.googleapis.com/auth/gmail.readonly`},
		Endpoint: oauth2.Endpoint {
			AuthURL: `https://accounts.google.com/o/oauth2/auth`,
			TokenURL: `https://accounts.google.com/o/oauth2/token`,
		},
		RedirectURL: `urn:ietf:wg:oauth:2.0:oob`,
	}

	ctx := context.Background()

	authURL := conf.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("authorization code\n%v\n", authURL)

	var code string
	if _, err := fmt.Scan(&code); err != nil {
		fmt.Errorf("Unable to read authorization code %v\n", err)
	}

	token, err := conf.Exchange(oauth2.NoContext, code)
	if err != nil {
		fmt.Errorf("Unable to retrieve token from web %v\n", err)
	}

	client := conf.Client(ctx, token)

	srv, err := gmail.New(client)
	if err != nil {
		fmt.Errorf("Unable to retrieve gmail client %v\n", err)
	}

	  user := "me"
	  r, err := srv.Users.Labels.List(user).Do()
	  if err != nil {
	    fmt.Errorf("Unable to retrieve labels. %v", err)
	  }
	  if (len(r.Labels) > 0) {
	    fmt.Print("Labels:\n")
	    for _, l := range r.Labels {
	      fmt.Printf("- %s\n",  l.Name)
	    }
	  } else {
	    fmt.Print("No labels found.")
	  }
}