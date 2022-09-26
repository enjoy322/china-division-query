profile:
	go test ./ -coverprofile=data.out

html:
	go tool cover -html=data.out

func:
	go tool cover -func=data.out