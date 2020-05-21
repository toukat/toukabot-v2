package request

import (
	"github.com/toukat/toukabot-v2/config"
	"github.com/toukat/toukabot-v2/util/logger"
	"io"

	"fmt"
	"net/http"
)

var log = logger.GetLogger("HTTP Request")

func GetRequest(uri string) (io.ReadCloser, error) {
	c := config.GetConfig()
	log.Info(fmt.Sprintf("Making GET request to %s", uri))

	r, err := http.Get(c.APIHost + uri)
	if err != nil {
		log.Error(fmt.Sprintf("Unable to make GET request to %s", uri))
		log.Error(err)
		return nil, err
	}

	log.Info(fmt.Sprintf("Successfully made GET request to %s", uri))

	return r.Body, nil
}
