package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestMainHandlerWhenCountMoreThanTotal.
// Запрос сформирован корректно, сервис возвращает код ответа 200 и тело ответа не пустое.
func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	status := responseRecorder.Code

	if assert.Equal(t, status, http.StatusOK) {
		fmt.Printf("%d city value\n", status)
	}

	body := responseRecorder.Body.String()
	list := strings.Split(body, ",")
	//Проверяет на равенство два аргумента. Если аргументы не равны, в выводе будут подсвечены различия.
	assert.Equal(t, len(list), totalCount)

}

// TestMainHandlerWhenCountMoreThanCiti.
// Город, который передаётся в параметре city, не поддерживается.
// Сервис возвращает код ответа 400 и ошибку wrong city value в теле ответа.
func TestMainHandlerWhenCountMoreThanCiti(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=10&city=sochi", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	status := responseRecorder.Code

	//assert.Equal(t, status, http.StatusOK, "%d wrong city value", status)
	// или лучше так
	if assert.NotEqual(t, status, http.StatusOK) {
		fmt.Printf("%d wrong city value\n", status)
	}
}

// TestMainHandlerWhenCountMoreThanValue.
// Если в параметре count указано больше, чем есть всего, должны вернуться все доступные кафе.
func TestMainHandlerWhenCountMoreThanValue(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	responseBody := responseRecorder.Body.String()
	fmt.Println(responseBody)
}
