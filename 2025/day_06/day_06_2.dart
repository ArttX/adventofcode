import "dart:io";

void main() async {
  final file = File("day_06/input.txt");
  final input = await file.readAsLines();
  final actionRow =
      input.removeLast().trim().split(RegExp(r"\s+")).reversed.toList();

  var maxRowLen = 0;
  for (final row in input) {
    if (row.length > maxRowLen) maxRowLen = row.length;
  }

  final numStrings = [];
  for (var row in input) {
    row = row.padRight(maxRowLen);
    final numRow = [];
    for (var i = row.length - 1; i >= 0; i--) {
      numRow.add(row[i]);
    }
    numStrings.add(numRow);
  }

  // Build numbers
  List<List<int>> groups = [];
  List<int> nums = [];
  for (var i = 0; i <= input[0].length; i++) {
    var str = "";
    for (final row in numStrings) str += row[i];
    if (str.trim() == "") {
      groups.add(nums);
      nums = [];
      continue;
    }
    nums.add(int.parse(str));
    if (i == input[0].length) groups.add(nums);
  }

  var sum = 0;
  for (final (i, group) in groups.indexed) {
    final mark = actionRow[i];
    var result = mark == "+" ? 0 : 1;
    for (final it in group) {
      if (mark == "+") result += it;
      if (mark == "*") result *= it;
    }
    sum += result;
  }

  print("Sum: $sum");
}
