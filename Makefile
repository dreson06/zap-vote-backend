build-zapvote:
	@echo "compiling zapvote application..."
	@GOOS=linux GOARCH=amd64 go build -v -o build/zapvote ./cmd/zapvote
	@echo "compiling for zapvote done."

move: build-zapvote
	@dns="dreson6@100.119.230.42" && \
	remotePath="/home/dreson6/apps/zapvote/zapvote" && \
	localPath="./build/zapvote" && \
	pemPath="~/.ssh/mini-pc" && \
	ssh -i "$$pemPath" "$$dns" "rm -rf $$remotePath" && \
	echo "removing file from remote: $$?" && \
	scp -C -i "$$pemPath" "$$localPath" "$$dns:$$remotePath" && \
	echo "sending file to remote: $$?" && \
	ssh -i "$$pemPath" "$$dns" "sudo systemctl restart zapvote.service" &&\
	echo "restarting zapvote: $$?"

run:
	@go build -o build/zapvote ./cmd/zapvote && ./build/zapvote
