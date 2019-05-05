# SmartStatster
Statistic bot

Текущая версия - создана заглушка для обхода блокировки исходящих запросов к мессенджеру по ip
Комплект содержит:
1. Тестовое приложение на go
2. Сервер для обработки запросов браузером

Схема работы:

(1) --GET request local--> (2) --selenium GET/Chrome request--> мессенджер
мессенджер --GET response json--> (2) --GET response json--> (1)(json parsing)

Требования:
selenium webdriver
maven
Chrome
запущенный VPN
jetty;
  зависимости java:
  jetty-webapp
  webdrivermanager

Инструкция:
устанавливаем maven плагин, jetty, chromedriver.exe
импортируем проект java и зависимости,
создаём файл src/resources/config/default.proprties с параметрами:
------------------------------------------------------
webdriver.chrome.driver=driver/chromedriver.exe (ссылка на драйвер)
api.url=https://api.telegram.org/bot
bot.token=<токен бота>
------------------------------------------------------

запускаем jetty сервер с java приложением

запускаем go проект из SmartStatsterCore с тестовым запросом 
(например http://127.0.0.1:8085/bot_get?action=getupdates, экшены будут добавляться в java приложение,
сейчас поддерживается getupdates и sendmessage), обрабатываем запросы от мессенджера
  
