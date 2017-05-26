package main

import (
    "testing"
    "net/http"
    "net/http/httptest"
    "github.com/stretchr/testify/assert"
)

func TestHealthCheckHandler(t *testing.T) {
    req, err := http.NewRequest("POST", "/api/v1/conversation", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(V1PostConversationHandler)

    handler.ServeHTTP(rr, req)

    assert.Equal(t, rr.Code, http.StatusOK, "need to be equals")

    assert.Equal(t, rr.Body.String(), `{"alive": true}`, "need to be equals")
}
