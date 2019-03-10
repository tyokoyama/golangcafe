package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"google.golang.org/api/gmail/v1"
)

func main() {
	token := &oauth2.Token{}
	client := &http.Client{}

	// Client IDとSecretはRefresh済みです。
	client_id := `328006125971-dold8i9smanincppfpc28ikhi2lomfg8.apps.googleusercontent.com`
	client_secret := `GRPBmfTb6YHZaXRS7zt5xnMy`

	conf := &oauth2.Config{
		ClientID:     client_id,
		ClientSecret: client_secret,
		Scopes:       []string{`https://www.googleapis.com/auth/gmail.readonly`},
		Endpoint: oauth2.Endpoint{
			AuthURL:  `https://accounts.google.com/o/oauth2/auth`,
			TokenURL: `https://accounts.google.com/o/oauth2/token`,
		},
		RedirectURL: `urn:ietf:wg:oauth:2.0:oob`,
	}

	ctx := context.Background()

	f, err := os.Open("token.json")
	if err != nil {
		// ファイルがない。
		authURL := conf.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
		fmt.Printf("authorization code\n%v\n", authURL)

		var code string
		if _, err := fmt.Scan(&code); err != nil {
			fmt.Errorf("Unable to read authorization code %v\n", err)
		}

		token, err = conf.Exchange(oauth2.NoContext, code)
		if err != nil {
			fmt.Errorf("Unable to retrieve token from web %v\n", err)
		}

		client = conf.Client(ctx, token)

		// tokenをファイルに保存
		f, err = os.Create("token.json")
		if err != nil {
			fmt.Errorf("file open error. %v\n", err)
		}
		json.NewEncoder(f).Encode(token)
	} else {
		json.NewDecoder(f).Decode(token)
		client = conf.Client(ctx, token)
	}
	f.Close()

	srv, err := gmail.New(client)
	if err != nil {
		fmt.Errorf("Unable to retrieve gmail client %v\n", err)
	}

	user := "me"
	r, err := srv.Users.Messages.List(user).Do()
	if err != nil {
		fmt.Errorf("Unable to retrieve messages. %v", err)
	}
	if len(r.Messages) > 0 {
		fmt.Print("Messages:\n")
		for _, m := range r.Messages {
			messages, err := srv.Users.Messages.Get(m.Id, user).Do()
			if err != nil {
				fmt.Errorf("Unable to retrieve messages get. %v", err)
			}
			fmt.Printf("- %s\n", messages)
		}
	} else {
		fmt.Print("No messages found.")
	}

	r, err := srv.Users.Labels.List(user).Do()
	if err != nil {
		fmt.Errorf("Unable to retrieve labels. %v", err)
	}
	if len(r.Labels) > 0 {
		fmt.Print("Labels:\n")
		for _, l := range r.Labels {
			fmt.Printf("- %s\n", l.Name)
		}
	} else {
		fmt.Print("No labels found.")
	}
}
