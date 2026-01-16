// В этом пакете вы реализуете два хендлера.

// Для корневого эндпоинта / нужно реализовать хендлер, который возвращает HTML из файла index.html.
// Второй хендлер для эндпоинта /upload должен выполнять следующие действия:
// 		Парсить html-форму из файла index.html.
// 		Получить файл из формы (не забудьте его закрыть).
// 		Прочитать данные из файла.
// 		Передать эти данные в функцию автоопределения из пакета service, которую вы создали, чтобы получить переконвертируемую строку.
// 		Создать локальный файл.
// 		Записать в локальный файл результат конвертации строки. Для генерации имени файла вы можете использовать
// 		время с помощью time.Now().UTC().String().
// 			Чтобы получить расширения файла, используйте filepath.Ext().
// 		Вернуть результат конвертации строки.
// Там, где это необходимо, обработайте возможные ошибки. Статус при возникновении ошибок http.StatusInternalServerError.

package handlers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "invalid request method", http.StatusInternalServerError)
		return
	}
	txt, err := os.ReadFile("index.html")
	if err != nil {
		http.Error(w, "error read html file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write(txt)
}

func DownoloadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "invalid request method", http.StatusInternalServerError)
		return
	}

	dwnFile, header, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "file upload error", http.StatusInternalServerError)
		return
	}
	defer dwnFile.Close()

	reqStr, err := io.ReadAll(dwnFile)
	if err != nil {
		http.Error(w, "read file error", http.StatusInternalServerError)
		return
	}

	resultStr, err := service.AutoConvertor(string(reqStr))
	if err != nil {
		http.Error(w, "convert data error", http.StatusInternalServerError)
		return
	}

	err = os.WriteFile(time.Now().UTC().Format("2006-01-02_15-04-05")+filepath.Ext(header.Filename), []byte(resultStr), 0755)
	if err != nil {
		http.Error(w, "save data to local file error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("conversion result: " + resultStr))

}
