const std = @import("std");
const part1 = @import("./part1.zig");
const part2 = @import("./part2.zig");
const filePath = "days/day01/calibration.txt";
const allocator = std.heap.page_allocator;

const day01_result = struct {
    part1: u32,
    part2: u32,
};

pub fn run() anyerror!day01_result {
    const file = try std.fs.cwd().openFile(filePath, .{});
    defer file.close();

    var buf_reader = std.io.bufferedReaderSize(1024, file.reader());
    const reader = buf_reader.reader();

    var buffer = try allocator.alloc(u8, 100);
    defer allocator.free(buffer);

    var part1_result: u32 = 0;
    var part2_result: u32 = 0;

    while (try reader.readUntilDelimiterOrEof(buffer, '\n')) |line| {
        part1_result += part1.extract_vals(line);
        part2_result += part2.extract_vals(line);
    }

    const result = day01_result{
        .part1 = part1_result,
        .part2 = part2_result,
    };

    return result;
}
