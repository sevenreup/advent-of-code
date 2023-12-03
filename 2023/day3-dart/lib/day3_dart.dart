import 'dart:io';

import 'package:day3_dart/visual.dart';

void solve(List<String> args) async {
  var input = await File('input.txt').readAsString();
  List<List<String>> grid = [];

  for (var element in input.split("\n")) {
    grid.add(element.runes.map(String.fromCharCode).toList());
  }

  List<int> numbers = [];
  for (var y = 0; y < grid.length; y++) {
    var row = grid[y];
    var numAccum = "";
    bool hasHit = false;
    for (var x = 0; x < row.length; x++) {
      var char = row[x];
      if (int.tryParse(char) != null) {
        numAccum += char;
        var hit = scan(grid, x, y);
        if (hit && !hasHit) {
          hasHit = true;
        }
      } else {
        if (numAccum != "") {
          if (hasHit) {
            numbers.add(int.parse(numAccum));
            grid.printBoundingBox(x, y, numAccum.length);
            print("");
          }
          hasHit = false;
        }
        numAccum = "";
      }
    }
  }

  print(numbers.reduce((previousValue, element) => previousValue + element));
}

bool scan(List<List<String>> grid, int x, int y) {
  if (y + 1 < grid.length) {
    var isBelow = hasHit(grid, x, y + 1);
    if (isBelow) {
      return true;
    }
  }
  if (x - 1 > 0) {
    var isLeft = hasHit(grid, x - 1, y);
    if (isLeft) {
      return true;
    }
  }
  if (x + 1 < grid[y].length) {
    var isRight = hasHit(grid, x + 1, y);
    if (isRight) {
      return true;
    }
  }

  if (y - 1 >= 0) {
    var isAbove = hasHit(grid, x, y - 1);
    if (isAbove) {
      return true;
    }
  }
  if (y + 1 < grid.length && x + 1 < grid[y].length) {
    var isBelowRight = hasHit(grid, x + 1, y + 1);
    if (isBelowRight) {
      return true;
    }
  }
  if (y + 1 < grid.length && x - 1 >= 0) {
    var isBelowLeft = hasHit(grid, x - 1, y + 1);
    if (isBelowLeft) {
      return true;
    }
  }
  if (y - 1 >= 0 && x + 1 < grid[y].length) {
    var isAboveRight = hasHit(grid, x + 1, y - 1);
    if (isAboveRight) {
      return true;
    }
  }
  if (y - 1 >= 0 && x - 1 >= 0) {
    var isAboveLeft = hasHit(grid, x - 1, y - 1);
    if (isAboveLeft) {
      return true;
    }
  }
  return false;
}

bool hasHit(List<List<String>> grid, int x, int y) {
  var char = grid[y][x];
  if (char.trim().isEmpty) return false;
  return !"0123456789.".contains(char);
}
