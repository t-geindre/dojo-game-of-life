03-bench:
	go test -count=5 -benchtime=50x -bench=. ./03-optimization

04-bench:
	go test -count=5 -benchtime=50x -bench=. ./04-zoom