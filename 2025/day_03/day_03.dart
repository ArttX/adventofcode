import "dart:io";

void main() async {
  final file = File("day_03/input.txt");
  final input = await file.readAsLines();

  final joltages =
      input
          .map((row) => row.split("").map((c) => int.parse(c)).toList())
          .toList();

  var sum = 0;
  for (final row in joltages) {
    var maxJolt = 0;
    for (var i = 0; i < row.length - 1; i++) {
      for (var y = i + 1; y < row.length; y++) {
        if (i == y) continue; // Skip on same number combo
        final num = int.parse("${row[i]}${row[y]}");
        if (num > maxJolt) maxJolt = num;
      }
    }
    sum += maxJolt;
  }

  print("Total Joltage: $sum");
}
