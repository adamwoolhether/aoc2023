
new:
	@[ "${DAY}" ] || ( echo ">> DAY var is not set"; exit 1 ); \
	DAYDIR=day${DAY} && \
	mkdir -p go/days/$${DAYDIR} && \
	echo "package $${DAYDIR}" >> go/days/$${DAYDIR}/$${DAYDIR}.go && \
	mkdir -p zig/days/$${DAYDIR} && \
	echo "const std = @import(\"std\");" >> zig/days/$${DAYDIR}/$${DAYDIR}.zig

fmt:
	zig fmt .