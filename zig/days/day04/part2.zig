const std = @import("std");

pub fn winAmt(winningAmt: []usize) anyerror!usize {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    const alloc = gpa.allocator();
    defer {
        if (gpa.deinit() == .leak) {
            @panic("dealloc failed");
        }
    }

    var totalWins = std.ArrayList(usize).init(alloc);
    defer totalWins.deinit();

    var idx: usize = 0;
    while (idx < winningAmt.len) : (idx += 1) {
        try totalWins.append(1);
    }

    for (winningAmt, 0..) |amt, i| {
        var j: usize = 1;
        while (j <= amt and i + j < winningAmt.len) : (j += 1) {
            totalWins.items[i + j] += totalWins.items[i];
        }
    }

    var total: usize = 0;

    for (totalWins.items) |count| {
        total += count;
    }

    return total;
}
