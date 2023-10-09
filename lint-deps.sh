set -ex
rm -rf bin/build
mkdir -p bin/build
cd bin/build
git clone https://github.com/golangci/golangci-lint .
git reset --hard 411e0bbbd3096aa0ee2b924160629bdf2bc81d40 # v1.54.2
export CGO_ENABLED=1
make build
cd ..; cp build/golangci-lint .
rm -rf build
mkdir aslint-build; cd aslint-build
git clone https://github.com/AlexSSD7/aslint .
go build -buildmode=plugin -o aslint.so plugin/main.go
mv aslint.so ..; cd ..; rm -rf aslint-build
rm -f .golangci.yml
wget -q https://raw.githubusercontent.com/AlexSSD7/aslint/master/.golangci.yml