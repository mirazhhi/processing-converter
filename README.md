``` docker build -t converter_local . ``` Создаем образ 
``` docker run -v /your_local_path:/app -ti converter_local go build main.go

- Описываем .env file 
- Запускаем бинарник ./main
- Читаем наш XML