import "dart:io";

void main() async {
  final file = File("day_07/input.txt");
  final input = await file.readAsLines();

  final grid = input.map((row) => row.split("")).toList();

  Map<int, int> beams = {};

  var sIdx = grid[0].indexOf("S");
  beams[sIdx] = 1;

  var splits = 1;
  for (final row in grid) {
    if (!row.contains("^")) continue;
    for (var i = 0; i < row.length; i++) {
      if (row[i] == "^" && beams.containsKey(i)) {
        final count = beams[i] ?? 0;
        beams[i + 1] = count + (beams[i + 1] ?? 0);
        beams[i - 1] = count + (beams[i - 1] ?? 0);
        beams[i] = 0;

        splits += count;
      }
    }
  }

  print("Answer: $splits");
}
