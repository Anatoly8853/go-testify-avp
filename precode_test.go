package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestMainHandlerWhenCountMoreThanTotal.
// Запрос сформирован корректно, сервис возвращает код ответа 200 и тело ответа не пустое.
func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=4&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	status := responseRecorder.Code

	require.Equal(t, status, http.StatusOK, "%d city value", status)

	body := responseRecorder.Body
	// если body будет nil, то есть ли смысл дальше проверять????
	assert.NotEmpty(t, body)

	response := responseRecorder.Body.String()

	list := strings.Split(response, ",")

	assert.Equal(t, len(list), totalCount)

}

// TestMainHandlerWhenCountMoreThanCiti.
// Город, который передаётся в параметре city, не поддерживается.
// Сервис возвращает код ответа 400 и ошибку wrong city value в теле ответа.
func TestMainHandlerWhenCountMoreThanCiti(t *testing.T) {

	req := httptest.NewRequest("GET", "/cafe?count=4&city=sochi", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	status := responseRecorder.Code
	// тут нужно еще проверить ошибку wrong city value в теле ответа?????? или вывести
	require.Equal(t, status, http.StatusBadRequest, "%d wrong city value = %d", status, http.StatusBadRequest)
}

// TestMainHandlerWhenCountMoreThanValue.
// Если в параметре count указано больше, чем есть всего, должны вернуться все доступные кафе.
func TestMainHandlerWhenCountMoreThanValue(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	status := responseRecorder.Code

	require.Equal(t, status, http.StatusOK, "%d city value", status)

	body := responseRecorder.Body.String()

	list := strings.Split(body, ",")

	require.Equal(t, len(list), totalCount)
}
