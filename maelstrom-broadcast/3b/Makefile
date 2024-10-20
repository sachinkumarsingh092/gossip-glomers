install:
	go install .

test:
	cd ../../maelstrom && \
	./maelstrom test -w broadcast --bin ~/go/bin/3b --node-count 5 --time-limit 20 --rate 10 && \
	cd -

debug:
	cd ../../maelstrom && \
	./maelstrom serve && \
	cd -
