const std = @import("std");
const filepath = "days/day03/schematic.txt";
const allocator = std.heap.page_allocator;
const isDigit = std.ascii.isDigit;

const day03_result = struct {
    part1: usize,
    // part2: u64,
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

    // var part2_result: u64 = 0;

    while (try reader.readUntilDelimiterOrEof(buffer, '\n')) |line| {
        // std.debug.print("line {s}\n", .{line});

        // Need to use a copy of the current line, otherwise grid.append will be
        // simply overrridden with the final line in the buffer.
        var lineCopy = try alloc.dupe(u8, line);
        try grid.append(lineCopy);
        // part2_result += try part2.maxPowerGames(line);
    }

    const part1Result = sumParts(grid);

    const result = day03_result{
        .part1 = part1Result,
        // .part2 = part2_result,
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

fn isSymbol(char: u8) bool {
    return char != '.' and !std.ascii.isDigit(char);
}

fn isGear(char: u8) bool {
    return char == '*';
}
