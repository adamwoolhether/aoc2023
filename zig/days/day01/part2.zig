const std = @import("std");
const part1 = @import("./part1.zig");

pub fn extract_vals(line: []u8) u8 {
    var foundLower = false;
    var foundUpper = false;
    var lower: ?u8 = null;
    var upper: ?u8 = null;

    var headStart: usize = 0;
    var headEnd: usize = 0;
    var tailStart = line.len - 1;
    var tailEnd = line.len;

    while (headStart < tailStart) {
        if (foundLower and foundUpper) {
            break;
        }

        if (!foundLower) {
            const lookup = part1.isDigit(line[headEnd]);
            if (lookup.found) {
                lower = lookup.num;
                foundLower = true;
            }

            const strLookup = isStringDigit(line[headStart..headEnd]);
            if (strLookup.found) {
                lower = strLookup.num;
                foundLower = true;
            }

            headEnd += 1;
        }

        if (!foundUpper) {
            const lookup = part1.isDigit(line[tailStart]);
            if (lookup.found) {
                upper = lookup.num;
                foundUpper = true;
            }

            const strLookup = isStringDigit(line[tailStart..tailEnd]);
            if (strLookup.found) {
                upper = strLookup.num;
                foundUpper = true;
            }

            tailStart -= 1;
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

const isStringDigitResult = struct {
    found: bool,
    num: u8,
};

fn isStringDigit(b: []u8) isStringDigitResult {
    if (std.mem.containsAtLeast(u8, b, 1, "one")) {
        return isStringDigitResult{ .found = true, .num = 1 };
    }
    if (std.mem.containsAtLeast(u8, b, 1, "two")) {
        return isStringDigitResult{ .found = true, .num = 2 };
    }
    if (std.mem.containsAtLeast(u8, b, 1, "three")) {
        return isStringDigitResult{ .found = true, .num = 3 };
    }
    if (std.mem.containsAtLeast(u8, b, 1, "four")) {
        return isStringDigitResult{ .found = true, .num = 4 };
    }
    if (std.mem.containsAtLeast(u8, b, 1, "five")) {
        return isStringDigitResult{ .found = true, .num = 5 };
    }
    if (std.mem.containsAtLeast(u8, b, 1, "six")) {
        return isStringDigitResult{ .found = true, .num = 6 };
    }
    if (std.mem.containsAtLeast(u8, b, 1, "seven")) {
        return isStringDigitResult{ .found = true, .num = 7 };
    }
    if (std.mem.containsAtLeast(u8, b, 1, "eight")) {
        return isStringDigitResult{ .found = true, .num = 8 };
    }
    if (std.mem.containsAtLeast(u8, b, 1, "nine")) {
        return isStringDigitResult{ .found = true, .num = 9 };
    }

    return isStringDigitResult{ .found = false, .num = 0 };
}
