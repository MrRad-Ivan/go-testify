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
	require.Equal(t, responseRecorder.Code, http.StatusOK)

	//проверка на пустоту поля
	assert.NotEmpty(t, responseRecorder.Body)

	//конвертируем полученную строку в слайс
	body := responseRecorder.Body.String()
	list := strings.Split(body, ",")

	// проверка на общее кол-во
	assert.Len(t, list, totalCount)

}
