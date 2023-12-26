const std = @import("std");
const day01 = @import("days/day01/day01.zig");
const day02 = @import("days/day02/day02.zig");

pub fn main() anyerror!void {
    var day01Result = try day01.day01();
    var day02Result = try day02.day02();

    try std.io.getStdOut().writer().print("Day01 Result: {}\n", .{day01Result});
    try std.io.getStdOut().writer().print("Day02 Result: {}\n", .{day02Result});
}
