.PHONY: open-jmeter
open-jmeter:
	open /usr/local/bin/jmeter

.PHONY: run-jmeter
run-jmeter:
	jmeter -n -t ./HTTP_Request_1.jmx -l ./results.csv -e -o ./report

