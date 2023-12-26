const std = @import("std");
const day01 = @import("days/day01/day01.zig");

pub fn main() anyerror!void {
    var day01Result = try day01.day01();

    try std.io.getStdOut().writer().print("Day01 Result: {}\n", .{day01Result});
}
