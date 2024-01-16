const std = @import("std");
const filepath = "days/day03/schematic.txt";
const allocator = std.heap.page_allocator;
const isDigit = std.ascii.isDigit;

const day03_result = struct {
    part1: usize,
    part2: usize,
};

pub fn run() anyerror!day03_result {
    const file = try std.fs.cwd().openFile(filepath, .{});
    defer file.close();

    var buf_reader = std.io.bufferedReader(file.reader());
    const reader = buf_reader.reader();

    var buffer = try allocator.alloc(u8, 1024);
    defer allocator.free(buffer);

    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    const alloc = gpa.allocator();
    defer {
        const deinit = gpa.deinit();
        if (deinit == .leak) {
            @panic("dealloc fail");
        }
    }

    var grid = std.ArrayList([]const u8).init(alloc);
    defer {
        for (grid.items) |line| {
            alloc.free(line);
        }
        grid.deinit();
    }

    while (try reader.readUntilDelimiterOrEof(buffer, '\n')) |line| {
        // std.debug.print("line {s}\n", .{line});

        // Need to use a copy of the current line, otherwise grid.append will be
        // simply overrridden with the final line in the buffer.
        var lineCopy = try alloc.dupe(u8, line);
        try grid.append(lineCopy);
    }

    const part1Result = sumParts(grid);
    const part2Result = sumGearPowers(grid);

    const result = day03_result{
        .part1 = part1Result,
        .part2 = try part2Result,
    };

    return result;
}

fn sumParts(grid: std.ArrayList([]const u8)) usize {
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
            if (isSymbol(checkRow[l])) {
                return true;
            }

            if (high + 1 > checkRow.len) continue;
            if (isSymbol(checkRow[high])) return true;
        }
    }

    return false;
}

fn sumGearPowers(grid: std.ArrayList([]const u8)) !usize {
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
            if (isSymbol(checkRow[high])) {
                const location = loc{ .found = true, .row = @as(usize, @intCast(r)), .col = high };
                return location;
            }
        }
    }

    const res = loc{ .found = false, .col = undefined, .row = undefined };
    return res;
}

fn isSymbol(char: u8) bool {
    return char != '.' and !std.ascii.isDigit(char);
}

fn isGear(char: u8) bool {
    return char == '*';
}
