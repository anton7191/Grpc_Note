PHONY: generate

generate:
		mkdir -p pkg\note_v1
		/mnt/c/protoc-21.5-win64/bin/protoc.exe --proto_path api/note_v1 --go_out=pkg/note_v1 --go_opt=paths=import --go-grpc_out=pkg/note_v1 --go-grpc_opt=paths=import api/note_v1/note.proto
		mv pkg/note_v1/github.com/anton7191/Note-server-api/pkg/note_v1/* pkg/note_v1/
		rm -rf pkg/note_v1/github.com/