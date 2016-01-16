default: all

all:
	@go build ||:

server:
	@go build ||:
	@./icmp-tunnel server ||:

client:
	@go build ||:
	@./icmp-tunnel client ||:

