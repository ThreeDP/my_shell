FROM golang:1.18

WORKDIR /github.com/ThreeDP/shell/
ADD . /github.com/ThreeDP/shell/

EXPOSE 5050

CMD ["tail", "-f", "/dev/null"]