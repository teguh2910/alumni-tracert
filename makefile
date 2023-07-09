init:
	cd backend && go mod init tracert

gen:
	protoc --proto_path=proto --go_out=./backend/proto --go-grpc_out=require_unimplemented_servers=false:./backend/proto --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative proto/*/*.proto

deploy:
	cd backend && CGO_ENABLED=0 GOOS=linux go build -o alumni server.go && docker build --tag=jackyhtg/alumni:$(TAG) . && docker push jackyhtg/alumni:$(TAG)
	ssh rijal /root/bin/anter-api $(TAG)

server:
	cd backend && go run server.go

install-fe:
	cd frontend && npm install

deploy-fe:
	cd frontend && npm run build && cd public && scp -r ./* rijal:/home/smart/public_html/anter_rijalasepnugroho_com/

frontend:
	cd frontend && npm run dev

.PHONY: init gen deploy server install-fe deploy-fe frontend