const std = @import("std");

const Game = struct {
    red: u8,
    green: u8,
    blue: u8,
};

const maxAmt = Game{
    .red = 12,
    .green = 13,
    .blue = 14,
};

pub fn possibleGames(line: []u8) anyerror!u8 {
    var game = std.mem.splitAny(u8, line, ":");

    var idParts = std.mem.splitAny(u8, game.next().?, " ");
    _ = idParts.next();

    const gameID = try std.fmt.parseInt(u8, idParts.next().?, 10);
    var results = std.mem.splitAny(u8, game.next().?, ";");

    while (results.next()) |result| {
        var picks = std.mem.splitAny(u8, result, ",");

        while (picks.next()) |pick| {
            const trimmed = std.mem.trim(u8, pick, " ");
            var cube = std.mem.splitAny(u8, trimmed, " ");

            const amt = try std.fmt.parseInt(u8, cube.next().?, 10);
            const color = cube.next().?[0];

            switch (color) {
                'r' => if (amt > maxAmt.red) {
                    return 0;
                },
                'b' => {
                    if (amt > maxAmt.blue) {
                        return 0;
                    }
                },
                'g' => if (amt > maxAmt.green) {
                    return 0;
                },
                else => unreachable,
            }
        }
    }

    return gameID;
}
