#!/usr/bin/env python3
# this code is not using any built-in function i.e. bin(), oct(), hex() to
# convert for better understanding of the algorithm conversion processes itself
import re
import sys

# ==== FROM BINARY ====
# TODO: binary decimal part? "100110011.101"
def binary_decimal(b):
  b = str(b)[::-1]
  reverse = ''
  output = 0
  # only integer part
  for i in range(len(b)):
    reverse += b[i]
    output += int(reverse[i]) * 2 ** i
  return output

def binary_octal(b):
  dic = {'001':1, '010':2, '011':3, '100':4, '101':5, '110':6, '111':7}
  out = ''
  pr = re.sub(r'(\d)(?=(\d{3})+$)',r'\1\t',str(b)).split('\t')
  if len(pr[0]) != 3:
    pr[0] = '0' * (3 - len(pr[0])) + pr[0]
  for i in range(len(pr)):
    pr[i] = dic[pr[i]]
  return ''.join(map(str, pr))

def binary_hexadecimal(b):
  dic = {'0000':0, '0001':1, '0010':2, '0011':3, '0100':4, '0101':5, '0110':6, '0111':7, '1000':8, '1001':9, '1010':'A', '1011':'B', '1100':'C', '1101':'D', '1110':'E', '1111':'F'}
  out = ''
  pr = re.sub(r'(\d)(?=(\d{4})+$)',r'\1\t',str(b)).split('\t')
  if len(pr[0]) != 4:
    pr[0] = '0' * (4 - len(pr[0])) + pr[0]
  for i in range(len(pr)):
    pr[i] = dic[pr[i]]
  return ''.join(map(str, pr))

# ==== FROM DECIMAL ====
def decimal_binary(d):
  out = ''
  while d > 0:
    out += str(d % 2)
    d //= 2
  return out[::-1]

def decimal_octal(d):
  out = ''
  while d > 0:
    out += str(d % 8)
    d //= 8
  return out[::-1]

def decimal_hexadecimal(d):
  dic = {0:'0', 1:'1', 2:'2', 3:'3', 4:'4', 5:'5', 6:'6', 7:'7', 8:'8', 9:'9', 10:'A', 11:'B', 12:'C', 13:'D', 14:'E', 15:'F'}
  out = ''
  while d > 0:
    out += dic[d % 16]
    d //= 16
  return out[::-1]




if __name__ == '__main__':
  try:
    arg1, arg2, arg3 = sys.argv[1], sys.argv[2], sys.argv[3]
  except:
    print("Usage: python3 convert.py [FROM] [TO] [DATA]\ne.g. python3 convert.py b d 100110011")
    sys.exit()
  # >>> binary >>>
  if arg1 == 'b':
    if arg2 == 'd':
      print(binary_decimal(arg3))
    elif arg2 == 'o':
      print(binary_octal(arg3))
    elif arg2 == 'h':
      print(binary_hexadecimal(arg3))

  # >>> decimal >>>
  elif arg1 == 'd':
    arg3 = int(arg3)
    if arg2 == 'b':
      print(decimal_binary(arg3))
    elif arg2 == 'o':
      print(decimal_octal(arg3))
    elif arg2 == 'h':
      print(decimal_hexadecimal(arg3))
