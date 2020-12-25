## Silika File Manager - Go Back
Api для файлового менеджера Silika  
Реализация хранения 

main.go  
Пример реализации Api
## Тест
curl -X POST http://localhost:3447/fsapi/runKod

## Websocket
 Отправляем тестовое сообщение
 Запускаем SiFileBack
 ```sh
go run .
```
Если все прошло правильно, то выведется сообщение 
```
....
...  Starting, HTTP on:3447
....
```
Это значит, что приложение слушает порт 33447

Запустим SiFileBack - TestTool

```
При подключении клиента надо распаковать в темповую папку архив и запустить докер. После того как пришла команда сохранять или нет, мы закрываем докер и либо сохраняем, либо удаляем файлы.

go mod init https://github.com/xela07ax/Silika-FileManager/tree/main/sifileback