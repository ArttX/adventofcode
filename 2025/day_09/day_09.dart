import "dart:io";

int calculateArea((int, int) p1, (int, int) p2) {
  final length = (p2.$1 - p1.$1).abs() + 1;
  final width = (p2.$2 - p1.$2).abs() + 1;
  return length * width;
}

void main() async {
  final file = File("day_09/input.txt");
  final input = await file.readAsLines();

  final cords =
      input.map((row) {
        final parts = row.split(",");
        return (int.parse(parts[0]), int.parse(parts[1]));
      }).toList();

  var biggestArea = 0;
  for (var i = 0; i < cords.length - 1; i++) {
    for (var j = i + 1; j < cords.length; j++) {
      final area = calculateArea(cords[i], cords[j]);
      if (area > biggestArea) biggestArea = area;
    }
  }
  print("Biggest area: $biggestArea");
}
