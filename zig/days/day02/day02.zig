const std = @import("std");
const part1 = @import("./part1.zig");
const part2 = @import("./part2.zig");
const filepath = "days/day02/games.txt";
const allocator = std.heap.page_allocator;

const day02_result = struct {
    part1: u32,
    part2: u64,
};

pub fn run() anyerror!day02_result {
    const file = try std.fs.cwd().openFile(filepath, .{});
    defer file.close();

    var buf_reader = std.io.bufferedReaderSize(1024, file.reader());
    const reader = buf_reader.reader();

    var buffer = try allocator.alloc(u8, 1024);
    defer allocator.free(buffer);

    var part1_result: u32 = 0;
    var part2_result: u64 = 0;

    while (try reader.readUntilDelimiterOrEof(buffer, '\n')) |line| {
        part1_result += try part1.possibleGames(line);
        part2_result += try part2.maxPowerGames(line);
    }

    const result = day02_result{
        .part1 = part1_result,
        .part2 = part2_result,
    };

    return result;
}
