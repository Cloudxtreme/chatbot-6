package main

import (
    "strings"
    "testing"
    "net/http"
    "net/http/httptest"
    "github.com/stretchr/testify/assert"
)

func TestV1PostConversationHandler(t *testing.T) {
    body := strings.NewReader(`{"question": "hi", "response": "what's up ?"}`)

    req, err := http.NewRequest("POST", "/api/v1/conversation", body)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(V1PostConversationHandler)

    handler.ServeHTTP(rr, req)

    assert.Equal(t, rr.Code, http.StatusOK, "need to return StatusOK")
}
