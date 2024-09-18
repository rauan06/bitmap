def bitwise_and(x, y):
  return x & y


def bitwise_or(x, y):
  return x | y


def bitwise_xor(x, y):
    return x ^ y


def bitwise_not(x, y):
    return ~-x

def bitwise_rightShift(x, y):
    return x >> y

def bitwise_lefShift(x, y):
    return x << y

def bitwise_logicalRightShift(x, y):
    return (x >> y) & ((1 << (x.bit_length() - y)) - 1)


print(bitwise_and(6, 13))

# 0000 0110
# 6 >> 1 : 0000 0011

# 0000 0110
# 6 << 1 : 0000 1100

# 0000 0110
# 6 >>> 1 : 0000 0011