export GOPATH="$(go env GOPATH)"
export PATH="${PATH}:${GOPATH}/bin"

echo "Starting go fumpt" && gofumpt -l -w . && \
echo "Starting fieldalignment" && fieldalignment -fix ./**  && \
echo "Starting gci" && gci write -s standard -s default  . && \
echo "Starting golangci-lint" &&   golangci-lint run && \

