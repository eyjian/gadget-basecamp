ifeq ($(OS),Windows_NT)
	target=gcts.exe
else
	target=gcts
endif

all: ${target}

${target}: generate_create_table_sql.go
ifeq ($(OS),Windows_NT)
	set GOOS=windows
	set GOARCH=amd64
endif
	go mod tidy && go build -o $@ $<

.PHONY: clean

clean:
	rm -f ${target}

install: ${target}
ifeq ($(OS),Windows_NT)
	copy ${target} %GOPATH%\bin\
else
	cp ${target} $$GOPATH/bin/
endif
