package utils

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/shgysd/hash/api/types"
)

// IsAllowMethod 許可してないメソッドを弾く
func IsAllowMethod(w http.ResponseWriter, r string, m string) bool {
	if r != m {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Write([]byte("405 Method Not Allowed"))
		return true
	}

	return false
}

// UnmarshalBody bodyをJSON化
func UnmarshalBody(r io.ReadCloser, d *types.SignIn) error {
	b, err := ioutil.ReadAll(r)
	defer r.Close()
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, &d)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}
