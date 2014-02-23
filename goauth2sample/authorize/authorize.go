package authorize

import (
	"encoding/json"
	"fmt"
	"os"
	"io/ioutil"
)

type Auth struct {
	ClientID string			`json:"id"`
	Secret string			`json:"secret"`
	RedirectUrl string		`json:"redirect"`
}

func GetAuthInfo() (auth Auth, err error) {
	var text []byte
	auth = Auth {
		ClientID: "",
		Secret: "",
		RedirectUrl: "",
	}

	file, err := os.Open("auth.json")
	defer file.Close()
	if os.IsNotExist(err) {
		// ファイルなし

		inputNewData(&auth)

		err = writeFile(auth)

	} else if err == nil {
		text, err = ioutil.ReadAll(file)
		if err != nil {
			return auth, err
		}

		err = json.Unmarshal(text, &auth)
		if err != nil || auth.ClientID == "" || auth.Secret == "" || auth.RedirectUrl == "" {
			// データがおかしい->新規入力
			inputNewData(&auth)
			err = writeFile(auth)
		}
	}

	return
}

func inputNewData(auth *Auth) {
	// 新規入力
	fmt.Println("Input ClientID")
	fmt.Scanf("%s\n", &auth.ClientID)
	fmt.Println("Input Secret")
	fmt.Scanf("%s\n", &auth.Secret)
	fmt.Println("Input Redirect URL")
	fmt.Scanf("%s\n", &auth.RedirectUrl)
}

func writeFile(auth Auth) error {
	text, err := json.Marshal(auth)
	err = ioutil.WriteFile("auth.json", text, 0777)

	return err
}