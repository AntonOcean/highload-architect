.PHONY: open-jmeter
open-jmeter:
	open /usr/local/bin/jmeter

.PHONY: run-jmeter-p
run-jmeter-p:
	jmeter -n -t ./HTTP_read_chat.jmx -l ./results-p.csv -e -o ./report-p

.PHONY: run-jmeter-t
run-jmeter-t:
	jmeter -n -t ./HTTP_read_chat.jmx -l ./results-t.csv -e -o ./report-t

.PHONY: show-nodes
show-nodes:
	PGPASSWORD=adminpassword psql -U postgres -h localhost -p 5432 -d postgres -c "show pool_nodes;"

.PHONY: show-repl
show-repl:
	PGPASSWORD=adminpassword psql -U postgres -h localhost -p 5432 -d postgres -c "SELECT pid,usename,application_name,state,sync_state FROM pg_stat_replication;"