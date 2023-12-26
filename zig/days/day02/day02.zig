const std = @import("std");
const filepath = "/days/day02/games.txt";

pub fn day02() anyerror!u32 {
    const file = try std.fs.io.cwd().openFile(filepath, .{});
    defer file.close()
}