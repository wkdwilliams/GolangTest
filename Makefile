default:
	go build -o build/main main.go
dev:
	compiledaemon -command="./GolangTest" -color=true
