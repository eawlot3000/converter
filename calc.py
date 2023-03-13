#!/usr/bin/env python3
import sys

# TODO: Add support for negative numbers
# binary remainder for division [10101011100011 / 1100101011 = 1101 R: 110110100]
def add(a, b, base):
  base_convert = {16: hex, 8: oct, 2: bin}
  if base in base_convert:
    a, b = int(a, base), int(b, base)
    result = a + b
    return base_convert[base](result)[2:]
  else:
    return "Invalid base"

def subtract(a, b, base):
  base_convert = {16: hex, 8: oct, 2: bin}
  if base in base_convert:
    a, b = int(a, base), int(b, base)
    result = a - b
    return base_convert[base](result)[2:]
  else:
    return "Invalid base"

def multiply(a, b, base):
  base_convert = {16: hex, 8: oct, 2: bin}
  if base in base_convert:
    a, b = int(a, base), int(b, base)
    result = a * b
    return base_convert[base](result)[2:]
  else:
    return "Invalid base"

def divide(a, b, base):
  base_convert = {16: hex, 8: oct, 2: bin}
  if base in base_convert:
    a, b = int(a, base), int(b, base)
    result = a // b
    return base_convert[base](result)[2:]
  else:
    return "Invalid base"


if __name__ == '__main__':
  try:
    arg1, arg2, arg3, arg4 = sys.argv[1], sys.argv[2], sys.argv[3], sys.argv[4]
  except:
    print("Usage: python3 calc.py [OPERATOR] [BASE] [DATA1] [DATA2] \ne.g. python3 calc.py add 2 100 110011: add 2 binary numbers 100 and 110011")
    sys.exit(0)

  if arg2 not in ["2", "8", "16"]:
    print('heyy ur almost there! please use an integer(2 for binary, 8 for octal, 16 for hex)')
    sys.exit(0)
  elif arg1 == "add":
    print(add(arg3, arg4, int(arg2)))
  elif arg1 == "subtract":
    print(subtract(arg3, arg4, int(arg2)))
  elif arg1 == "multiply":
    print(multiply(arg3, arg4, int(arg2)))
  elif arg1 == "divide":
    print(divide(arg3, arg4, int(arg2)))
  else:
    print("Invalid operator")
