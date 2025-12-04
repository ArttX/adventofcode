import "dart:io";

void main() async {
  final file = File("day_04/input.txt");
  final input = await file.readAsLines();

  final grid = input.map((row) => row.split("")).toList();

  List<(int, int)> directions = [
    (-1, -1),
    (0, -1),
    (1, -1),
    (1, 0),
    (1, 1),
    (0, 1),
    (-1, 1),
    (-1, 0),
  ];
  final max_row = grid.length;
  final max_col = grid[0].length;

  var accessibleRollCount = 0;

  while (true) {
    for (var row = 0; row < max_row; row++) {
      for (var col = 0; col < max_col; col++) {
        if (grid[row][col] != "@") continue;
        var rollCount = 0;
        for (final dir in directions) {
          if (row + dir.$2 < 0 || row + dir.$2 > max_col - 1) continue;
          if (col + dir.$1 < 0 || col + dir.$1 > max_row - 1) continue;
          if (grid[row + dir.$2][col + dir.$1] == "@" ||
              grid[row + dir.$2][col + dir.$1] == "x")
            rollCount++;
        }
        if (rollCount < 4) {
          accessibleRollCount++;
          grid[row][col] = "x"; // Mark for removal
        }
      }
    }
    // Remove marked
    var removed = 0;
    for (var row = 0; row < max_row; row++) {
      for (var col = 0; col < max_col; col++) {
        if (grid[row][col] == "x") {
          grid[row][col] = ".";
          removed++;
        }
      }
    }
    if (removed == 0) break;
  }
  print("Roll count: $accessibleRollCount");
}
