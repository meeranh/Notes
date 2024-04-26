# Binary Shifts
+ If you have a number, say `2` for example, its binary value will be `10`.
+ The purpose of binary shifts are to multiply and divide a number by 2 in binary.

# Left Binary Shift
+ If you perform a binary shift to the left, that means that *the bits of a number will shift one position to the right*, which is basically *a multiplication of 2*.
+ For example, if we right shift `10`, it will result in `100`. An extra zero will be added to the beginning
+ We can verify this: `10` is `2` in decimal, and `100` is `4` in decimal and since `2 * 2 = 4`, we can conclude that *a left binary shift is equivalent to multiplying a number by 2*.

# Right Binary Shift
+ Similarly, a right binary shift will move bits of a binary digit to the right, and it acts as a division operation.
+ If we have `100` as a binary number, a right binary shift will result in `10`, the rightmost bit (which was a `0` in our case) will be discarded.
+ We can verify this: `100` is `4` in decimal, and `10` is `2` in decimal and since `4 / 2 = 2`, we can conclude that *a right binary shift is equivalent to dividing a number by 2*.
