package providers

import (
	"bytes"
	"github.com/go-faster/errors"
	"github.com/rs/zerolog/log"
	"io"
	"net/http"
	"net/url"
)

// AuthenticatedJsonClient is a client that sends JSON requests with basic authentication and logging.
type AuthenticatedJsonClient struct {
	client   *http.Client
	username string
	password string
	base     *url.URL
}

func NewAuthenticatedJson(baseURL *url.URL, username, password string) *AuthenticatedJsonClient {
	return &AuthenticatedJsonClient{
		client:   new(http.Client),
		username: username,
		password: password,
		base:     baseURL,
	}
}

func (d AuthenticatedJsonClient) Do(r *http.Request) (*http.Response, error) {
	r.URL = d.base.ResolveReference(r.URL)
	r.SetBasicAuth(d.username, d.password)
	r.Header.Add("Content-Type", "application/json")
	err := logRequestBody(r)
	if err != nil {
		return nil, errors.Wrap(err, "could not log request body")
	}
	response, err := d.client.Do(r)
	if err != nil {
		return nil, errors.Wrap(err, "request failed")
	}
	err = logResponseBody(response)
	if err != nil {
		return nil, errors.Wrap(err, "could not log response body")
	}
	return response, err
}

func logRequestBody(r *http.Request) error {
	logger := log.Logger
	if r.Body != nil {
		bodyReader, err := r.GetBody()
		if err != nil {
			return err
		}
		body, err := io.ReadAll(bodyReader)
		if err != nil {
			return err
		}
		logger = log.With().RawJSON("body", body).Logger()
	}
	logger.Info().
		Str("method", r.Method).
		Str("endpoint", r.URL.String()).
		Msg("request")
	return nil
}

func logResponseBody(response *http.Response) error {
	dup := new(bytes.Buffer)
	tee := io.TeeReader(response.Body, dup)
	responseBody, err := io.ReadAll(tee)
	if err != nil {
		return err
	}
	response.Body = io.NopCloser(dup)
	log.Info().
		RawJSON("body", responseBody).
		Msg("response")
	return nil
}
