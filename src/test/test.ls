-- test
--[[
   'test'"
]]
let val1 : int, val2 : real = ?a, 1.0;
let val3 : real, val4, val5 : real = 0x1, 1.0e-2, 0.1e+3;
let val6 : int[] = [1, 2, 3];
let val7 : int[@] = [@1,2,3];
let val8 : string = "1234";
let val9, val10 = '1234""', [[":a23]];
fn plus1( val1: int, val2: int, val3: int[] ) : int, int {
  val1 = -10@int + val2 + val3[1];
  return val1 + 1, val2 + 1;
}
