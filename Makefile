default: all

all:
	@go build -race ||:

server: all
	@sudo ./icmp-tunnel server ||:

client: all
	@sudo ./icmp-tunnel client ||:

