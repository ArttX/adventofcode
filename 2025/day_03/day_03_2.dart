import "dart:io";

void main() async {
  final file = File("day_03/input.txt");
  final input = await file.readAsLines();

  final joltages =
      input
          .map((row) => row.split("").map((c) => int.parse(c)).toList())
          .toList();

  var sum = 0;
  for (var row in joltages) {
    var full = "";
    for (final (i, _) in List.filled(12, "").indexed) {
      (int, int) maxInRow = (-1, 0);
      for (var item in row.sublist(0, row.length - (11 - i)).indexed) {
        if (item.$2 > maxInRow.$2) maxInRow = item;
      }
      full += maxInRow.$2.toString();
      row = row.sublist(maxInRow.$1 + 1);
    }
    sum += int.parse(full);
  }

  print("Total Joltage: $sum");
}
