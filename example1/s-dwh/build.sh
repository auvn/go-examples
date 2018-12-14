GOOS=linux GOARCH=amd64 go build -o service cmd/*/main.go && docker build -t auvn/go-examples/example1/cluster/services/sdwh . -f ../cluster/services/Dockerfile
