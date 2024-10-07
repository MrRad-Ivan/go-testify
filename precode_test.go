package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// здесь нужно добавить необходимые проверки
	// проверка на статус 200 ок
	require.Equal(t, http.StatusOK, responseRecorder.Code)

	//Проверяем кол-во выдаваемых городов
	body := responseRecorder.Body.String()
	list := strings.Split(body, ",")

	assert.Len(t, list, totalCount)

}

func ErrorStatusGerNotcity(t *testing.T) {
	erorrReq := httptest.NewRequest("GET", "/cafe?count=10&city=notcity", nil) //запрос к сервису с несуществующим городом

	responseRecorder := httptest.NewRecorder()
	erorrHandlerr := http.HandlerFunc(mainHandle)
	erorrHandlerr.ServeHTTP(responseRecorder, erorrReq)

	// проверяем статус 400
	require.Equal(t, http.StatusBadRequest, responseRecorder.Code)

	// проверяем запись при вводе несуществующего города
	expectedErrorMessage := "wrong city value"
	assert.Equal(t, expectedErrorMessage, responseRecorder.Body.String())

}
