build:
	go build

run: build
	./MultithreadedPacman

clean:
	rm ./MultithreadedPacman