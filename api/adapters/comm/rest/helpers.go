// Package rest is responsible for rest communication layer
package rest

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/rs/zerolog/log"

	apierr "github.com/serdarkalayci/carpool/api/adapters/comm/rest/errors"
)

// ToJSON serializes the given interface into a string based JSON format
func ToJSON(i interface{}, w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(i)
}

// FromJSON deserializes the object from JSON string
// in an io.Reader to the given interface
func FromJSON(i interface{}, r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(i)
}

func readPayload(r *http.Request) (payload []byte, e error) {
	payload, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		e = &apierr.ErrCannotReadPayload{}
		log.Error().Err(err).Msg(e.Error())
		return
	}
	if len(payload) == 0 {
		e = &apierr.ErrPayloadMissing{}
		log.Error().Err(err).Msg(e.Error())
		return
	}
	return
}
