module github.com/proconlon/redelivery-core

replace github.com/proconlon/redelivery-core => ./

go 1.24

require (
	github.com/emersion/go-imap v1.2.1
	github.com/joho/godotenv v1.5.1
	github.com/mnako/letters v0.2.4
)

require (
	github.com/emersion/go-sasl v0.0.0-20200509203442-7bfe0ed36a21 // indirect
	golang.org/x/net v0.35.0 // indirect
	golang.org/x/text v0.22.0 // indirect
)
