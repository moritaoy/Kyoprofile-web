sudo docker-compose exec app cd app
sudo docker-compose exec app go get -u -t ./...
sudo docker-compose exec app go run main.go
