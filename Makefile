.PHONY: open-jmeter
open-jmeter:
	open /usr/local/bin/jmeter

.PHONY: run-jmeter
run-jmeter:
	jmeter -n -t ./HTTP_Request_1.jmx -l ./results.csv -e -o ./report

.PHONY: show-nodes
show-nodes:
	PGPASSWORD=adminpassword psql -U postgres -h localhost -p 5432 -d postgres -c "show pool_nodes;"
