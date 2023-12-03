extension Visual on List<List<String>> {
  void printBoundingBox(int endX, int startY, int count) {
    var startX = endX - count;
    List<List<String>> acc = [[], [], []];
    for (var i = startX; i < endX; i++) {
      if (i == startX) {
        acc[0] = [_getChar(i, startY, aboveLeft: true)];
        acc[1] = [_getChar(i, startY, left: true)];
        acc[2] = [_getChar(i, startY, belowLeft: true)];
        acc[1] = [...acc[1], this[startY][i]];
      } else if (i == endX - 1) {
        acc[1] = [...acc[1], this[startY][i]];
        acc[1] = [...acc[1], _getChar(i, startY, right: true)];
      }
      else {
        acc[1] = [...acc[1], this[startY][i]];
      }
      acc[0] = [...acc[0], _getChar(i, startY, above: true)];
      acc[2] = [...acc[2], _getChar(i, startY, below: true)];
      if (i == endX - 1) {
        acc[0] = [...acc[0], _getChar(i, startY, aboveRight: true)];
        acc[2] = [...acc[2], _getChar(i, startY, belowRight: true)];
      }
    }

    // print acc list

    for (var element in acc) {
      print(element.join(""));
    }
  }

  String _getChar(
    int x,
    int y, {
    bool left = false,
    bool right = false,
    bool above = false,
    bool below = false,
    bool aboveLeft = false,
    bool aboveRight = false,
    bool belowLeft = false,
    bool belowRight = false,
  }) {
    if (below) {
      if (y + 1 < length) {
        var isBelow = this[y + 1][x];
        return isBelow;
      }
    }
    if (left) {
      if (x - 1 >= 0) {
        var isLeft = this[y][x - 1];
        return isLeft;
      }
    }

    if (right) {
      if (x + 1 < this[y].length) {
        var isRight = this[y][x + 1];
        return isRight;
      }
    }
    if (above) {
      if (y - 1 >= 0) {
        var isAbove = this[y - 1][x];
        return isAbove;
      }
    }
    if (belowRight) {
      if (y + 1 < length && x + 1 < this[y].length) {
        var isBelowRight = this[y + 1][x + 1];
        return isBelowRight;
      }
    }
    if (belowLeft) {
      if (y + 1 < length && x - 1 >= 0) {
        var isBelowLeft = this[y + 1][x - 1];
        return isBelowLeft;
      }
    }
    if (aboveRight) {
      if (y - 1 >= 0 && x + 1 < this[y].length) {
        var isAboveRight = this[y - 1][x + 1];
        return isAboveRight;
      }
    }
    if (aboveLeft) {
      if (y - 1 >= 0 && x - 1 >= 0) {
        var isAboveLeft = this[y - 1][x - 1];
        return isAboveLeft;
      }
    }
    return ".";
  }
}



/**
 * 
 * .......
 * ..1234..
 * .......
 * 
 */