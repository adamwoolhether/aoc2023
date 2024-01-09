const std = @import("std");
const day01 = @import("days/day01/day01.zig");
const day02 = @import("days/day02/day02.zig");
const day03 = @import("days/day03/day03.zig");

pub fn main() anyerror!void {
    const day01Result = try day01.run();
    const day02Result = try day02.run();
    const day03Result = try day03.run();

    try std.io.getStdOut().writer().print("Day 01 Part 1 Result: {}\n", .{day01Result.part1});
    try std.io.getStdOut().writer().print("Day 01 Part 2 Result: {}\n", .{day01Result.part2});
    try std.io.getStdOut().writer().print("Day 02 Part 1 Result: {}\n", .{day02Result.part1});
    try std.io.getStdOut().writer().print("Day 02 Part 2 Result: {}\n", .{day02Result.part2});
    try std.io.getStdOut().writer().print("Day 03 Part 1 Result: {}\n", .{day03Result.part1});
}
