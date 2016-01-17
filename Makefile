default: all

all:
	@go build -race ||:

server: all
	@sudo ./icmp-tunnel server ||:

client:
	@sudo ./icmp-tunnel client ||:

