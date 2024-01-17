.PHONY: benchmark benchmark-5x upgrade-deps count-files

benchmark:
		rm -rf benchmark/

		go test -bench=. -cpu=1 -benchtime 1000000x

benchmark-5x:
	@for i in 1 2 3 4 5; do \
		rm -rf benchmark/; \
		go test -bench=. -cpu=1 -benchtime 1000000x; \
	done

upgrade-deps:
	go get -u ./...
	go mod tidy

count-files:
	@find benchmark -type d -exec sh -c 'echo -n "{}: "; find "{}" -maxdepth 1 -type f | wc -l' \;

