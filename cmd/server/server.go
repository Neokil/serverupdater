package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"serverupdater/internal/config"
)

func main() {
	configPath := flag.String("config", "./config.json", "path to the config file")
	if !flag.Parsed() {
		flag.Parse()
	}

	c, err := config.FromFile(*configPath)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Loaded config: %v\n", c)

	for _, h := range c.Handlers {
		http.HandleFunc(h.Uri, (func(h config.Handler) func(w http.ResponseWriter, req *http.Request) {
			return func(w http.ResponseWriter, r *http.Request) {
				if r.Method == "POST" {
					body, _ := ioutil.ReadAll(r.Body)
					err = verifySignature(body, r.Header.Get("X-Hub-Signature"), c.Secret)
					if err != nil {
						http.Error(w, "Unauthorized", http.StatusUnauthorized)
						return
					}

					for _, c := range h.Commands {
						fmt.Printf("Executing %s with args %v\n", c.Name, c.Args)
						out, err := exec.Command(c.Name, c.Args...).Output()
						if err != nil {
							if exitErr, ok := err.(*exec.ExitError); ok {
								fmt.Printf("-> Error-Message: %s\n", exitErr.Stderr)
								http.Error(w, string(exitErr.Stderr), http.StatusInternalServerError)
								return
							} else {
								fmt.Printf("-> Error: %s\n", err.Error())
								http.Error(w, err.Error(), http.StatusInternalServerError)
								return
							}
						}
						fmt.Printf("-> Output: %s\n", out)
					}
				}
			}
		})(h))
		fmt.Printf("Handler registered for %s\n", h.Uri)
	}

	fmt.Printf("Server started at %s:%d\n", c.Host, c.Port)
	http.ListenAndServe(fmt.Sprintf("%s:%d", c.Host, c.Port), nil)
}

func verifySignature(body []byte, signature string, secret string) error {
	if len(signature) == 0 {
		return errors.New("no signature found")
	}
	mac := hmac.New(sha1.New, []byte(secret))
	_, _ = mac.Write(body)
	expectedMAC := hex.EncodeToString(mac.Sum(nil))
	if !hmac.Equal([]byte(signature[5:]), []byte(expectedMAC)) {
		return errors.New("invalid signature")
	}
	return nil
}
