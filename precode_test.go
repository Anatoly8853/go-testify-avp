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
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	status := responseRecorder.Code
	//если так, то выводится в консоль ошибка
	assert.NotEqual(t, status, http.StatusOK, "%d city value", status)

	//если так, то мы ни чего в консоли не видим, а должны по условию видеть код ответа 200 и тело ответа не пустое
	//можно по подробнее объяснить что мы хотим увидеть и как ?????
	assert.Equal(t, status, http.StatusOK, "%d city value", status)
	/*   в примерах пишут так
	if assert.Equal(t, status, http.StatusOK) {
		fmt.Printf("%d city value\n", status)
	}
	*/
	body := responseRecorder.Body.String()
	list := strings.Split(body, ",")
	//Проверяет на равенство два аргумента. Если аргументы не равны, в выводе будут подсвечены различия.
	require.Equal(t, len(list), totalCount)

}

// TestMainHandlerWhenCountMoreThanCiti.
// Город, который передаётся в параметре city, не поддерживается.
// Сервис возвращает код ответа 400 и ошибку wrong city value в теле ответа.
func TestMainHandlerWhenCountMoreThanCiti(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=10&city=sochi", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	status := responseRecorder.Code

	assert.Equal(t, status, http.StatusOK, "%d wrong city value", status)

	body := responseRecorder.Body.String()
	list := strings.Split(body, ",")

	//Проверяет на равенство два аргумента. Если аргументы не равны, в выводе будут подсвечены различия.
	require.Equal(t, len(list), totalCount)
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

	assert.Equal(t, status, http.StatusOK, "%d city value", status)

	body := responseRecorder.Body.String()
	list := strings.Split(body, ",")

	responseBody := responseRecorder.Body.String()
	fmt.Println(responseBody)

	//Проверяет на равенство два аргумента. Если аргументы не равны, в выводе будут подсвечены различия.
	require.Equal(t, len(list), totalCount)
}
