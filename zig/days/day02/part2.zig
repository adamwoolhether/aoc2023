const std = @import("std");

const Game = struct {
    red: u64,
    green: u64,
    blue: u64,
};

pub fn maxPowerGames(line: []u8) anyerror!u64 {
    var game = std.mem.splitAny(u8, line, ":");
    _ = game.next();

    var results = std.mem.splitAny(u8, game.next().?, ";");

    var highestCnt = Game{
        .red = 0,
        .green = 0,
        .blue = 0,
    };

    while (results.next()) |result| {
        var picks = std.mem.splitAny(u8, result, ",");

        while (picks.next()) |pick| {
            const trimmed = std.mem.trim(u8, pick, " ");
            var cube = std.mem.splitAny(u8, trimmed, " ");

            const amt = try std.fmt.parseInt(u8, cube.next().?, 10);
            const color = cube.next().?[0];

            switch (color) {
                'r' => if (amt > highestCnt.red) {
                    highestCnt.red = amt;
                },
                'b' => {
                    if (amt > highestCnt.blue) {
                        highestCnt.blue = amt;
                    }
                },
                'g' => if (amt > highestCnt.green) {
                    highestCnt.green = amt;
                },
                else => unreachable,
            }
        }
    }

    // std.debug.print("red: {}, green: {}, blue: {}\n", .{ highestCnt.red, highestCnt.green, highestCnt.blue });
    // std.debug.print("total: {}\n", .{highestCnt.red * highestCnt.green * highestCnt.blue});

    const result: u64 = highestCnt.red * highestCnt.green * highestCnt.blue;

    return result;
}
