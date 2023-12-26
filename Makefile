
new:
	@[ "${DAY}" ] || ( echo ">> DAY var is not set"; exit 1 )
	mkdir -p go/days/"day"${DAY}
	echo "package day${DAY}" >> go/days/"day"${DAY}/main.go
	mkdir -p zig/days/"day"${DAY}
	touch zig/days/"day"${DAY}.zig