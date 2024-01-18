const std = @import("std");
const allocator = std.heap.page_allocator;

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

    while (try reader.readUntilDelimiterOrEof(buffer, '\n')) |line| {
        // std.debug.print("{s}\n", .{line});

        var l = std.mem.splitAny(u8, line, "|");
        const firstHalf = l.next() orelse continue;
        const secondHalf = l.next() orelse continue;

        var firstHalfIter = std.mem.splitAny(u8, firstHalf, ":");
        _ = firstHalfIter.next(); // Discard the unwanted text

        const winningNums = firstHalfIter.next() orelse continue;

        const scoreRes = try card_score(winningNums, secondHalf);
        part1Res += scoreRes;
    }

    const result = day04_result{
        .part1 = part1Res,
        .part2 = part2Res,
    };

    return result;
}

fn card_score(winningNums: []const u8, secondHalf: []const u8) anyerror!usize {
    var winMap = std.AutoHashMap(usize, void).init(allocator);
    defer winMap.deinit();

    var winNumSplit = std.mem.tokenizeAny(u8, winningNums, " ");
    while (winNumSplit.next()) |winNum| {
        const numResult = try std.fmt.parseInt(usize, winNum, 10);
        try winMap.put(numResult, {});
    }

    var myNums = std.mem.tokenizeAny(u8, secondHalf, " ");

    var matches: usize = 0;
    while (myNums.next()) |myNum| {
        const num = try std.fmt.parseInt(usize, myNum, 10);
        if (winMap.contains(num)) {
            matches += 1;
        }
    }

    var score: usize = 0;
    if (matches > 0) {
        // Probably a better way to do this, but since we're dealing with
        // a usize, compiler needs to ensure there is no overflow.
        matches -= 1;
    }

    score = std.math.pow(usize, 2, matches);

    return score;
}
