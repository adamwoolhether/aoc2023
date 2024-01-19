const std = @import("std");
const allocator = std.heap.page_allocator;

const cardSum = struct {
    score: usize,
    matchAmt: usize,
};

pub fn card_score(winningNums: []const u8, secondHalf: []const u8) anyerror!cardSum {
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
        score = std.math.pow(usize, 2, matches - 1);
    }

    const ret = cardSum{
        .score = score,
        .matchAmt = matches,
    };

    return ret;
}
