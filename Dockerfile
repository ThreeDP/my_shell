FROM golang:1.18

WORKDIR /prod
ADD . /prod/

EXPOSE 5050

CMD ["tail", "-f", "/dev/null"]