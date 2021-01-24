test:
	go test -v -count 1 -race -timeout 1m ./...

build-render-server:
	go build -o ./bin/render-server ./renderserver/main.go

heroky-build-and-deploy:
	heroku apps:create --region eu pikchr-render-server
	heroku apps
	heroku git:remote -a pikchr-render-server
	heroku container:push web
	heroku container:release web
