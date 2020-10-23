package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var eventListJSON = `{
	"kind": "EventList",
	"apiVersion": "audit.k8s.io/v1",
	"metadata": {},
	"items": [
	  {
		"level": "Request",
		"auditID": "b7699b7e-e876-4c97-9b18-f7e7adc18841",
		"stage": "ResponseComplete",
		"requestURI": "/api/v1/nodes?limit=500",
		"verb": "list",
		"user": {
		  "username": "https://sts.windows.net/147a2b71-5ce9-4933-94c4-2054328de565/#a7b4eb91-181f-4b91-a405-c4d904f1af0f",
		  "groups": [
			"29243834-aec2-4872-b903-661512b6ec08",
			"0e22bf18-6762-4e50-b489-6eb39b652962",
			"system:authenticated"
		  ]
		},
		"sourceIPs": [
		  "74.203.144.5"
		],
		"userAgent": "kubectl.exe/v1.10.11 (windows/amd64) kubernetes/637c7e2",
		"objectRef": {
		  "resource": "nodes",
		  "apiVersion": "v1"
		},
		"responseStatus": {
		  "metadata": {},
		  "status": "Failure",
		  "reason": "Forbidden",
		  "code": 403
		},
		"requestReceivedTimestamp": "2019-02-06T14:45:25.277447Z",
		"stageTimestamp": "2019-02-06T14:45:25.277756Z",
		"annotations": {
		  "authorization.k8s.io/decision": "forbid",
		  "authorization.k8s.io/reason": ""
		}
	  },
	  {
		"level": "Request",
		"auditID": "2cf8ad1e-25b1-49d7-bb0d-b56341587c12",
		"stage": "ResponseComplete",
		"requestURI": "/api/v1/nodes?limit=500",
		"verb": "list",
		"user": {
		  "username": "https://sts.windows.net/147a2b71-5ce9-4933-94c4-2054328de565/#a7b4eb91-181f-4b91-a405-c4d904f1af0f",
		  "groups": [
			"29243834-aec2-4872-b903-661512b6ec08",
			"0e22bf18-6762-4e50-b489-6eb39b652962",
			"system:authenticated"
		  ]
		},
		"sourceIPs": [
		  "74.203.144.5"
		],
		"userAgent": "kubectl.exe/v1.10.11 (windows/amd64) kubernetes/637c7e2",
		"objectRef": {
		  "resource": "nodes",
		  "apiVersion": "v1"
		},
		"responseStatus": {
		  "metadata": {},
		  "status": "Failure",
		  "reason": "Forbidden",
		  "code": 403
		},
		"requestReceivedTimestamp": "2019-02-06T14:46:36.735833Z",
		"stageTimestamp": "2019-02-06T14:46:36.735963Z",
		"annotations": {
		  "authorization.k8s.io/decision": "forbid",
		  "authorization.k8s.io/reason": ""
		}
	  }
	]
  }`

func TestAudit(t *testing.T) {
	// given
	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/audits", strings.NewReader(eventListJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// when
	assert.NoError(t, handler(c))

	// then
	assert.Equal(t, http.StatusOK, rec.Code)
}
