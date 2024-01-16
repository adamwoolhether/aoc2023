const std = @import("std");
const shared = @import("./shared.zig");
const isDigit = std.ascii.isDigit;
const allocator = std.heap.page_allocator;

pub fn sumGearPowers(grid: std.ArrayList([]const u8)) !usize {
    var result: usize = 0;

    var gears = std.AutoHashMap([2]usize, usize).init(allocator);
    defer gears.deinit();

    for (grid.items, 0..) |row, rowIdx| {
        var colIdx: usize = 0;

        while (colIdx < row.len) : (colIdx += 1) {
            if (isDigit(row[colIdx])) {
                var start: usize = colIdx;
                var end: usize = colIdx;

                while (end < row.len) {
                    if (!isDigit(row[end])) break;

                    end += 1;
                }

                const check = nextToGear(grid, rowIdx, start, end);
                if (check.found) {
                    const num = std.fmt.parseInt(usize, row[start..end], 10) catch 0;
                    // std.debug.print("found valid gear int: {}\n", .{num});
                    // std.debug.print("gear loc: {},{}\n", .{ check.row, check.col });

                    var v = try gears.getOrPut([2]usize{ check.row, check.col });
                    if (!v.found_existing) {
                        v.value_ptr.* = num;
                    } else {
                        var found = v.value_ptr.*;
                        found *= num;

                        result += found;
                    }
                }

                colIdx = end;
            }
        }
    }

    return result;
}

const loc = struct {
    found: bool,
    row: usize,
    col: usize,
};

fn nextToGear(grid: std.ArrayList([]const u8), row: usize, low: usize, high: usize) loc {
    const rowCheck = [3]i32{ -1, 0, 1 };

    for (rowCheck) |delta| {
        const r = @as(i32, @intCast(row)) + delta;
        if (r < 0 or r >= grid.items.len) continue;

        var l = std.math.sub(usize, low, 1) catch 0;
        const checkRow = grid.items[@as(usize, @intCast(r))];

        while (l < high) : (l += 1) {
            if (isGear(checkRow[l])) {
                const location = loc{ .found = true, .row = @as(usize, @intCast(r)), .col = l };
                return location;
            }

            if (high + 1 > checkRow.len) continue;
            if (shared.isSymbol(checkRow[high])) {
                const location = loc{ .found = true, .row = @as(usize, @intCast(r)), .col = high };
                return location;
            }
        }
    }

    const res = loc{ .found = false, .col = undefined, .row = undefined };
    return res;
}

fn isGear(char: u8) bool {
    return char == '*';
}
