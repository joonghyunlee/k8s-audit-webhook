package main

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type eventList struct {
	Kind       string                 `json:"kind"`
	APIVersion string                 `json:"apiVersion"`
	Metadata   map[string]interface{} `json:"metadata"`
	Items      []event                `json:"items"`
}

type event struct {
	Level      string `json:"level"`
	AuditID    string `json:"auditID"`
	Stage      string `json:"stage"`
	RequestURI string `json:"requestURI"`
	Verb       string `json:"verb"`
	User       struct {
		Username string   `json:"username"`
		Groups   []string `json:"groups"`
	} `json:"user"`
	SourceIPs []string `json:"sourceIPs"`
	UserAgent string   `json:"userAgent"`
	ObjectRef struct {
		Resource   string `json:"resource"`
		APIVersion string `json:"apiVersion"`
	} `json:"objectRef"`
	ResponseStatus struct {
		Metadata struct {
		} `json:"metadata"`
		Status string `json:"status"`
		Reason string `json:"reason"`
		Code   int    `json:"code"`
	} `json:"responseStatus"`
	RequestReceivedTimestamp time.Time `json:"requestReceivedTimestamp"`
	StageTimestamp           time.Time `json:"stageTimestamp"`
	Annotations              struct {
		AuthorizationK8SIoDecision string `json:"authorization.k8s.io/decision"`
		AuthorizationK8SIoReason   string `json:"authorization.k8s.io/reason"`
	} `json:"annotations"`
}

func handler(c echo.Context) error {
	events := new(eventList)
	if err := c.Bind(events); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Could not parse the audit data")
	}

	for i, item := range events.Items {
		messages := []string{}

		messages = append(messages, "Item: "+strconv.Itoa(i))
		messages = append(messages, "\n\tSource IPs: "+strings.Join(item.SourceIPs, ","))
		messages = append(messages, "\n\tUser Name: "+item.User.Username)
		messages = append(messages, "\n\tUser Groups: "+strings.Join(item.User.Groups, ","))
		messages = append(messages, "\n\tUser Agent: "+item.UserAgent)

		messages = append(messages, "\n\tStage: "+item.Stage)
		messages = append(messages, "\n\tRequest URI: "+item.RequestURI)
		messages = append(messages, "\n\tVerb: "+item.Verb)
		messages = append(messages, "\n\tRequest Received Timestamp: "+item.RequestReceivedTimestamp.String())

		messages = append(messages, "\n\tObject Resource: "+item.ObjectRef.Resource)
		messages = append(messages, "\n\tObject API Version: "+item.ObjectRef.APIVersion)

		messages = append(messages, "\n\tResponse Code: "+strconv.Itoa(item.ResponseStatus.Code))
		messages = append(messages, "\n\tResponse Status: "+item.ResponseStatus.Status)
		messages = append(messages, "\n\tResponse Reason: "+item.ResponseStatus.Reason+"\n")

		log.Debug(messages)
	}

	return c.NoContent(http.StatusOK)
}

func main() {
	e := echo.New()

	log.SetLevel(log.DEBUG)
	log.SetHeader("${time_rfc3339_nano} ${level} ${method} ${uri} ${status}")

	e.POST("/audits", handler)

	e.Logger.Fatal(e.Start(":8080"))
}
