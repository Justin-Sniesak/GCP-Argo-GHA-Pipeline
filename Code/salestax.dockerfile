FROM golang:alpine3.22
WORKDIR /app
COPY salesTaxWesternWa.go .
CMD ["go", "run", "salesTaxWesternWa.go"]