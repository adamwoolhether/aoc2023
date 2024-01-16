const std = @import("std");
const shared = @import("./shared.zig");
const isDigit = std.ascii.isDigit;

pub fn sumParts(grid: std.ArrayList([]const u8)) usize {
    var result: usize = 0;

    for (grid.items, 0..) |row, rowIdx| {
        // std.debug.print("grid row: {s}\n", .{row});
        var colIdx: usize = 0;

        while (colIdx < row.len) : (colIdx += 1) {
            if (isDigit(row[colIdx])) {
                var start: usize = colIdx;
                var end: usize = colIdx;

                // find the end
                while (end < row.len) {
                    if (!isDigit(row[end])) break;

                    end += 1;
                }

                // std.debug.print("found int {s} at {},{}\n", .{ row[start..end], start, end - 1 });

                if (nextToSymbol(grid, rowIdx, start, end)) {
                    // std.debug.print("found valid int {s} at {},{}\n", .{ row[start..end], start, end - 1 });
                    const num = std.fmt.parseInt(usize, row[start..end], 10) catch 0;
                    result += num;
                }

                colIdx = end;
            }
        }
    }

    return result;
}

fn nextToSymbol(grid: std.ArrayList([]const u8), row: usize, low: usize, high: usize) bool {
    const rowCheck = [3]i32{ -1, 0, 1 };

    for (rowCheck) |delta| {
        const r = @as(i32, @intCast(row)) + delta;
        if (r < 0 or r >= grid.items.len) continue;

        var l = std.math.sub(usize, low, 1) catch 0;
        const checkRow = grid.items[@as(usize, @intCast(r))];

        while (l < high) : (l += 1) {
            if (shared.isSymbol(checkRow[l])) {
                return true;
            }

            if (high + 1 > checkRow.len) continue;
            if (shared.isSymbol(checkRow[high])) return true;
        }
    }

    return false;
}
