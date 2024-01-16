const std = @import("std");
const part1 = @import("./part1.zig");
const part2 = @import("./part2.zig");
const allocator = std.heap.page_allocator;
const filepath = "days/day03/schematic.txt";

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

    const part1Result = part1.sumParts(grid);
    const part2Result = part2.sumGearPowers(grid);

    const result = day03_result{
        .part1 = part1Result,
        .part2 = try part2Result,
    };

    return result;
}
