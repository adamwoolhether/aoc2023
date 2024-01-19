const std = @import("std");
const allocator = std.heap.page_allocator;
const part1 = @import("./part1.zig");
const part2 = @import("./part2.zig");

const filepath = "days/day04/scratchcards.txt";

const day04_result = struct {
    part1: usize,
    part2: usize,
};

pub fn run() anyerror!day04_result {
    var part1Res: usize = 0;
    var part2Res: usize = 0;

    const file = try std.fs.cwd().openFile(filepath, .{});
    defer file.close();

    var buf_reader = std.io.bufferedReader(file.reader());
    const reader = buf_reader.reader();

    var buffer = try allocator.alloc(u8, 1024);
    defer allocator.free(buffer);

    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    const alloc = gpa.allocator();
    defer {
        if (gpa.deinit() == .leak) {
            @panic("dealloc faile");
        }
    }

    var winCardsAmt = std.ArrayList(usize).init(alloc);
    defer winCardsAmt.deinit();

    while (try reader.readUntilDelimiterOrEof(buffer, '\n')) |line| {
        // std.debug.print("{s}\n", .{line});

        var l = std.mem.splitAny(u8, line, "|");
        const firstHalf = l.next() orelse continue;
        const secondHalf = l.next() orelse continue;

        var firstHalfIter = std.mem.splitAny(u8, firstHalf, ":");
        _ = firstHalfIter.next(); // Discard the unwanted text

        const winningNums = firstHalfIter.next() orelse continue;

        const scoreRes = try part1.card_score(winningNums, secondHalf);
        part1Res += scoreRes.score;

        try winCardsAmt.append(scoreRes.matchAmt);
    }

    part2Res = try part2.winAmt(winCardsAmt.items);

    const result = day04_result{
        .part1 = part1Res,
        .part2 = part2Res,
    };

    return result;
}
