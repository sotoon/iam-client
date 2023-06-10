package webhook

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type IAMWebhookServer struct {
	BepaSecret string

	IP                                net.IP
	Port                              uint16
	MaxAcceptableRequestTimestampDiff time.Duration
	IgnoreAuthentication              bool

	eventListeners []*EventListener
}

func (s *IAMWebhookServer) StartServer() error {
	router := gin.Default()
	router.POST("/", s.handleEvent)

	listenAddress := fmt.Sprintf("%s:%d", s.IP, s.Port)
	return router.Run(listenAddress)
}

func (s *IAMWebhookServer) handleEvent(c *gin.Context) {
	// TODO Check if update is not already handled
	// TODO use Sentry instead of rendering error in response
	payload, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if err := s.authenticateRequest(c.Request.Header, string(payload)); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if err := s.validateRequestTimestamp(c.Request.Header); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	event, err := parseRequestEvent(payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	for _, eventListener := range s.eventListeners {
		if eventListener.doesMatch(event) {
			if err := eventListener.handle(*event); err != nil {
				c.JSON(http.StatusInternalServerError, err.Error())
			}
		}
	}
	c.JSON(http.StatusOK, nil)
}

func (s *IAMWebhookServer) AddListener(el EventListener) {
	s.eventListeners = append(s.eventListeners, &el)
}

func parseRequestEvent(payload []byte) (*Event, error) {
	data := make(map[string]interface{})
	if err := json.Unmarshal(payload, &data); err != nil {
		return nil, err
	}

	meta := data["meta"].(map[string]interface{})
	action := EventAction(meta["action"].(string))
	objectType := IAMObjectType(meta["type"].(string))

	objectJson, err := json.Marshal(data["data"])
	if err != nil {
		return nil, err
	}

	object, err := parseIAMObjectJson(objectType, objectJson)
	if err != nil {
		return nil, err
	}

	return &Event{
		Action:     action,
		ObjectType: objectType,
		Object:     object,
	}, nil
}

func (s *IAMWebhookServer) authenticateRequest(h http.Header, payload string) error {
	if s.IgnoreAuthentication {
		return nil
	}

	signatureList, ok := h["X-Bepa-Signature"]
	if !ok {
		return errors.New("signature header not found")
	}
	if len(signatureList) != 1 {
		return errors.New("more than one signature header found")
	}
	signature := signatureList[0]

	timestampList, ok := h["X-Bepa-Timestamp"]
	if !ok {
		return errors.New("timestamp header not found")
	}
	if len(timestampList) != 1 {
		return errors.New("more than one timestamp header found")
	}
	timestamp := timestampList[0]

	if !isBepaSignatureValid(payload, s.BepaSecret, timestamp, signature) {
		return errors.New("invalid signature")
	}

	return nil
}

func (s *IAMWebhookServer) validateRequestTimestamp(h http.Header) error {
	if s.IgnoreAuthentication {
		return nil
	}

	timestampList, ok := h["X-Bepa-Timestamp"]
	if !ok {
		return errors.New("timestamp header not found")
	}
	if len(timestampList) != 1 {
		return errors.New("more than one timestamp header found")
	}

	timestamp, err := strconv.ParseInt(timestampList[0], 10, 64)
	if err != nil {
		return errors.New("timestamp is not 64-bit integer")
	}

	timeDiff := time.Since(time.Unix(timestamp, 0))
	if timeDiff < 0*time.Second || timeDiff > s.MaxAcceptableRequestTimestampDiff {
		return errors.New("request timestamp too old")
	}

	return nil
}
