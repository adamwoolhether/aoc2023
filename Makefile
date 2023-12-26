
new:
	@[ "${DAY}" ] || ( echo ">> DAY var is not set"; exit 1 )
	mkdir -p days/"day"${DAY}
	echo "package day${DAY}" >> days/"day"${DAY}/main.go