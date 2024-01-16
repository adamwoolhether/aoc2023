const std = @import("std");

pub fn isSymbol(char: u8) bool {
    return char != '.' and !std.ascii.isDigit(char);
}
