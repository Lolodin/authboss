<<<<<<< HEAD
# Readme
В тестовом не говорилось о реализации регистрации, при создании стора
создаются два юзера с разными ролям: 
* login `User` password `1234`
* login `Admin` password `1234`

На тестовой странице `localhost:8080/` будут ссылки на ресурсы [foo bar sigma], 
а так же на роуты /login и /logout, для доступа к ресурсам нужно пройти аунтефикация в /login

## Запуск

`$ git clone https://github.com/Lolodin/authboss`

`$ cd authboss`

`$ go build cmd/main.go `

`$ /main`

Откройте браузер по адресу http://localhost:8080/ 

## Запуск в Docker

`$ git clone https://github.com/Lolodin/authboss`

`$ cd authboss`

`$ docker build -t test .`

`$ docker run -it -p 8080:8080 test`

Откройте браузер по адресу http://localhost:8080/ 
=======
# Readme
В тестовом не говорилось о реализации регистрации, при создании стора
создаются два юзера с разными ролям: 
* login `User` password `1234`
* login `Admin` password `1234`

На тестовой странице `localhost:8080/` будут ссылки на ресурсы [foo bar sigma], 
а так же на роуты /login и /logout, для доступа к ресурсам нужно пройти аунтефикация в /login

## Запуск

`$ git clone https://github.com/Lolodin/authboss`

`$ cd authboss`

`$ go build cmd/main.go `

`$ /main`

Откройте браузер по адресу http://localhost:8080/ 

## Запуск в Docker

`$ git clone https://github.com/Lolodin/authboss`

`$ cd authboss`

`$ docker build -t test .`

`$ docker run -it -p 8080:8080 test`

Откройте браузер по адресу http://localhost:8080/ 
>>>>>>> ada51a8d2b1bb559aaa80881ef661d3f9aed25e1
