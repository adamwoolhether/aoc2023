pub fn extract_vals(line: []u8) u8 {
    var foundLower = false;
    var foundUpper = false;

    var start: usize = 0;
    var end: usize = line.len - 1;

    var lower: ?u8 = null;
    var upper: ?u8 = null;

    while (start <= end) {
        if (foundLower and foundUpper) {
            break;
        }

        if (!foundLower) {
            const lookup = isDigit(line[start]);

            if (lookup.found) {
                lower = lookup.num;
                foundLower = true;
            }
            start += 1;
        }

        if (!foundUpper) {
            const lookup = isDigit(line[end]);

            if (lookup.found) {
                upper = lookup.num;
                foundUpper = true;
            }
            end -= 1;
        }
    }

    if (!foundUpper) {
        return lower.? * 10 + lower.?;
    }

    if (!foundLower) {
        return upper.? * 10 + upper.?;
    }

    return lower.? * 10 + upper.?;
}

const isDigitResult = struct {
    found: bool,
    num: u8,
};

pub fn isDigit(b: u8) isDigitResult {
    var d: u8 = 0;
    var found = (b >= '0' and b <= '9');

    if (found) {
        d = b - '0';
    }

    return isDigitResult{
        .found = found,
        .num = @intCast(d),
    };
}
