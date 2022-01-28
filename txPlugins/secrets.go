// Copyright (c) 2022 Kells Kearney. All rights reserved.
//
// Use of this source code is governed by the MIT License that can be found
// in the LICENSE file.
//
package pluginMeta

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

"github.com/rs/zerolog"
)

func GetSecret(cipherPhrase string) (string, error) {
	splits := strings.SplitN(cipherPhrase, ":", 2)
	if splits == nil || len(splits) != 2 { // Just plain text, nothing to do
		return cipherPhrase, nil
	}
	var err error
	var plaintext string

	fetchMethod := splits[0]
	fetchArg := splits[1]

	switch fetchMethod {
	case "filename": // Look up secret according to file path eg Kubernetes secrets
		plaintext, err = fetchFromFile(fetchArg)
	case "env":
		plaintext = os.Getenv(fetchArg)
	default:
		return "", fmt.Errorf("Unable to decode secret for %s password: %s", fetchMethod, fetchArg)
	}

	return plaintext, err
}

func fetchFromFile(filename string) (string, error) {
	data, err := ioutil.ReadFile(filepath.Clean(filename))
	if err != nil {
		return "", fmt.Errorf("Unable to read secret from file %s: %s", filename, err)
	}
	return strings.TrimSuffix(string(data), "\n"), nil
}

// MergeSecrets takes a key/value pair and updates with secrets
//
func MergeSecrets(pluginDataMapping map[string]string, log *zerolog.Logger) {
	for key, value := range pluginDataMapping {
		if strings.Contains(key, "secret") ||
			strings.Contains(key, "password") {
			plaintext, err := GetSecret(value)
			if err != nil {
				log.Warn().Err(err).Str("secret", key).Str("cipher_text", value).Msg("Unable to decode secret")
			} else {
				pluginDataMapping[key] = plaintext
			}
		}
	}
}
