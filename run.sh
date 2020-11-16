if ! which CompileDaemon; then
    go get github.com/githubnemo/CompileDaemon
fi

CompileDaemon \
    -color=true \
    -graceful-kill=true \
    -pattern="^(\.env.+|\.env)|(.+\.go|.+\.c)$" \
    -build="go build -mod=vendor -o perpus-api ./cmd/perpus-api/..." \
    -command="./perpus-api"
