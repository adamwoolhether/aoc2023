const std = @import("std");

const filePath = "days/day01/calibration.txt";

pub fn day01() anyerror!u32 {
    const file = try std.fs.cwd().openFile(filePath, .{});
    defer file.close();

    var buf_reader = std.io.bufferedReaderSize(1024, file.reader());
    const reader = buf_reader.reader();

    var buffer: [100]u8 = undefined;
    var result: u32 = 0;

    while (try reader.readUntilDelimiterOrEof(&buffer, '\n')) |line| {
        const val = extract_value(line);
        result += val;
        // std.debug.print("{}\n", .{val});
    }

    return result;
}

fn extract_value(line: []u8) u8 {
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
            if (isDigit(line[start])) {
                lower = @intCast(line[start] - '0');
                foundLower = true;
            }
            start += 1;
        }

        if (!foundUpper) {
            if (isDigit(line[end])) {
                upper = @intCast(line[end] - '0');
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

fn isDigit(b: u8) bool {
    return (b >= '0' and b <= '9');
}
