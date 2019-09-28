FROM ubuntu AS compiler
RUN apt-get update
RUN apt-get install -y golang
COPY  hellogo.go /
RUN GOPATH=/ && go build -o hellogo .

FROM ubuntu 
COPY  --from=compiler hellogo.go /
CMD /hellogo
