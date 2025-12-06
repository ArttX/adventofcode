import "dart:io";

void main() async {
  final file = File("day_06/input.txt");
  final input = await file.readAsLines();

  List<List<String>> values = [];
  for (final row in input) {
    final parts = row.trim().replaceAll(RegExp(r"\s+"), " ").split(" ");
    values.add(parts);
  }
  final actions = values.removeLast();

  // Do actions
  var result = 0;
  for (var i = 0; i < values[0].length; i++) {
    var current = actions[i] == "+" ? 0 : 1;
    for (var j = 0; j < values.length; j++) {
      if (actions[i] == "+") current += int.parse(values[j][i]);
      if (actions[i] == "*") current *= int.parse(values[j][i]);
    }
    result += current;
  }

  print("Result: $result");
}
