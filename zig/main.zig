const std = @import("std");
const day01 = @import("days/day01/day01.zig");
const day02 = @import("days/day02/day02.zig");

pub fn main() anyerror!void {
    var day01Result = try day01.run();
    var day02Result = try day02.run();

    try std.io.getStdOut().writer().print("Day01 Part1 Result: {}\n", .{day01Result.part1});
    try std.io.getStdOut().writer().print("Day01 Part2 Result: {}\n", .{day01Result.part2});
    try std.io.getStdOut().writer().print("Day02 Part1 Result: {}\n", .{day02Result.part1});
    try std.io.getStdOut().writer().print("Day02 Part2 Result: {}\n", .{day02Result.part2});
}
