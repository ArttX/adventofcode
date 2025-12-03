import 'dart:io';

void main() async {
  final file = File("day_01/input.txt");
  final instructions = await file.readAsLines();

  int current = 50;
  int count = 0;
  for (final inst in instructions) {
    final num = int.parse(inst.substring(1));
    // Make rotation
    if (inst.startsWith("R")) {
      for (int i = num; i > 0; i--) {
        current += 1;
        if (current > 99) current = 0;
      }
    }
    if (inst.startsWith("L")) {
      for (int i = num; i > 0; i--) {
        current -= 1;
        if (current < 0) current = 99;
      }
    }
    // Check for rotation
    if (current == 0) {
      count++;
    }
  }
  print("Count: $count");
}
