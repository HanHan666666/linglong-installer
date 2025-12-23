set -e
for tag in none # dmesg libc.membrk libc.memgrind
do
	echo "-tags=$tag"

	echo "GOOS=linux GOARCH=386"
	GOOS=linux GOARCH=386 go build -tags=$tag -v ./...
	GOOS=linux GOARCH=386 go test -tags=$tag -c -o /dev/null

	echo "GOOS=linux GOARCH=arm"
	GOOS=linux GOARCH=arm go build -tags=$tag -v ./...
	GOOS=linux GOARCH=arm go test -tags=$tag -c -o /dev/null

	echo "GOOS=linux GOARCH=loong64"
	GOOS=linux GOARCH=loong64 go build -tags=$tag -v ./...
	GOOS=linux GOARCH=loong64 go test -tags=$tag -c -o /dev/null

	echo "GOOS=linux GOARCH=ppc64le"
	GOOS=linux GOARCH=ppc64le go build -tags=$tag -v ./...
	GOOS=linux GOARCH=ppc64le go test -tags=$tag -c -o /dev/null

	echo "GOOS=linux GOARCH=riscv64"
	GOOS=linux GOARCH=riscv64 go build -tags=$tag -v ./...
	GOOS=linux GOARCH=riscv64 go test -tags=$tag -c -o /dev/null

	echo "GOOS=linux GOARCH=s390x"
	GOOS=linux GOARCH=s390x go build -tags=$tag -v ./...
	GOOS=linux GOARCH=s390x go test -tags=$tag -c -o /dev/null
done
