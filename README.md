# EffectiveMobile-test-task
Test task for EffectiveMobile

### Запуск тестового задания
Чтобы запустить тестовое задание, введите следующие команды в командную строку:
```
git clone https://github.com/Berlimor/EffectiveMobile-test-task.git
cd EffectiveMobile-test-task
docker-compose up --build -d
```
__Обратите внимание__, что для запуска на машине должен быть установлен __docker-compose__. Приложение развернётся локально с базой данных и миграциями.

### Тестирование
Для тестов зайдите на http://localhost:8080/swagger/index.html. Вашему вниманию будет представлен сваггер с RESTful API приложения. Для начала предлагается создать запись музыки _(POST /music)_, после чего добавить детали для этой записи _(POST /music/details)_.