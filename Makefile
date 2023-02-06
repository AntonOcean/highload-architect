.PHONY: open-jmeter
open-jmeter:
	open /usr/local/bin/jmeter

.PHONY: run-jmeter-r
run-jmeter-r:
	jmeter -n -t ./HTTP_read.jmx -l ./results-r.csv -e -o ./report-r

.PHONY: run-jmeter-w
run-jmeter-w:
	jmeter -n -t ./HTTP_write.jmx -l ./results-w.csv -e -o ./report-w

.PHONY: show-nodes
show-nodes:
	PGPASSWORD=adminpassword psql -U postgres -h localhost -p 5432 -d postgres -c "show pool_nodes;"

.PHONY: show-repl
show-repl:
	PGPASSWORD=adminpassword psql -U postgres -h localhost -p 5432 -d postgres -c "SELECT pid,usename,application_name,state,sync_state FROM pg_stat_replication;"