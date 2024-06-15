package main

import (
	"fmt"
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

	body := responseRecorder.Body.String()
	list := strings.Split(body, ",")

	assert.Equal(t, len(list), totalCount)

}

// TestMainHandlerWhenCountMoreThanCiti.
// Город, который передаётся в параметре city, не поддерживается.
// Сервис возвращает код ответа 400 и ошибку wrong city value в теле ответа.
func TestMainHandlerWhenCountMoreThanCiti(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=4&city=sochi", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	status := responseRecorder.Code

	require.Equal(t, status, http.StatusBadRequest, "%d wrong city value", status)

	body := responseRecorder.Body.String()
	list := strings.Split(body, ",")

	assert.Equal(t, len(list), totalCount)
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

	responseBody := responseRecorder.Body.String()
	fmt.Println(responseBody)

	require.Equal(t, len(list), totalCount)
}

// Спасибо за разъяснения я значит не правильно понял задание по поводу функций тестирования
// я разобрался? смутило в задании я думаю не один я так понял код ответа 200 и тело ответа не пустое.
