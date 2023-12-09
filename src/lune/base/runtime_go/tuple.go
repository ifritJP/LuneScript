/*
MIT License

Copyright (c) 2018,2020 ifritJP

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package runtimelns

type LnsTuple1[T1 any] struct {
	Val1 T1
}
func Lns_CreateLnsTuple1[T1 any]( arg1 T1 ) LnsTuple1[T1] {
    return LnsTuple1[T1]{ arg1 }
}

type LnsTuple2[T1, T2 any] struct {
	Val1 T1
	Val2 T2
}
func Lns_CreateLnsTuple2[T1, T2 any]( arg1 T1, arg2 T2 ) LnsTuple2[T1, T2] {
    return LnsTuple2[T1, T2]{ arg1, arg2 }
}

type LnsTuple3[T1, T2, T3 any] struct {
	Val1 T1
	Val2 T2
	Val3 T3
}
func Lns_CreateLnsTuple3[T1, T2, T3 any]( arg1 T1, arg2 T2, arg3 T3 ) LnsTuple3[T1, T2, T3] {
    return LnsTuple3[T1, T2, T3]{ arg1, arg2, arg3 }
}

type LnsTuple4[T1, T2, T3, T4 any] struct {
	Val1 T1
	Val2 T2
	Val3 T3
	Val4 T4
}

func Lns_CreateLnsTuple4[T1, T2, T3, T4 any]( arg1 T1, arg2 T2, arg3 T3, arg4 T4 ) LnsTuple4[T1, T2, T3, T4] {
    return LnsTuple4[T1, T2, T3, T4]{ arg1, arg2, arg3, arg4 }
}

type LnsTuple5[T1, T2, T3, T4, T5 any] struct {
	Val1 T1
	Val2 T2
	Val3 T3
	Val4 T4
	Val5 T5
}

func Lns_CreateLnsTuple5[T1, T2, T3, T4, T5 any](
    arg1 T1, arg2 T2, arg3 T3, arg4 T4, arg5 T5 ) LnsTuple5[T1, T2, T3, T4, T5] {
    return LnsTuple5[T1, T2, T3, T4, T5]{ arg1, arg2, arg3, arg4, arg5 }
}

type LnsTuple6[T1, T2, T3, T4, T5, T6 any] struct {
	Val1 T1
	Val2 T2
	Val3 T3
	Val4 T4
	Val5 T5
	Val6 T6
}

func Lns_CreateLnsTuple6[T1, T2, T3, T4, T5, T6 any](
    arg1 T1, arg2 T2, arg3 T3, arg4 T4, arg5 T5, arg6 T6 ) LnsTuple6[T1, T2, T3, T4, T5, T6] {
    return LnsTuple6[T1, T2, T3, T4, T5, T6]{ arg1, arg2, arg3, arg4, arg5, arg6 }
}

type LnsTuple7[T1, T2, T3, T4, T5, T6, T7 any] struct {
	Val1 T1
	Val2 T2
	Val3 T3
	Val4 T4
	Val5 T5
	Val6 T6
	Val7 T7
}

func Lns_CreateLnsTuple7[T1, T2, T3, T4, T5, T6, T7 any](
    arg1 T1, arg2 T2, arg3 T3, arg4 T4, arg5 T5, arg6 T6, arg7 T7 ) LnsTuple7[T1, T2, T3, T4, T5, T6, T7] {
    return LnsTuple7[T1, T2, T3, T4, T5, T6, T7]{arg1, arg2, arg3, arg4, arg5, arg6, arg7 }
}

type LnsTuple8[T1, T2, T3, T4, T5, T6, T7, T8 any] struct {
	Val1 T1
	Val2 T2
	Val3 T3
	Val4 T4
	Val5 T5
	Val6 T6
	Val7 T7
	Val8 T8
}

func Lns_CreateLnsTuple8[T1, T2, T3, T4, T5, T6, T7, T8 any](
    arg1 T1, arg2 T2, arg3 T3, arg4 T4, arg5 T5, arg6 T6, arg7 T7, arg8 T8 ) LnsTuple8[T1, T2, T3, T4, T5, T6, T7, T8] {
    return LnsTuple8[T1, T2, T3, T4, T5, T6, T7, T8]{
        arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8 }
}

type LnsTuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9 any] struct {
	Val1 T1
	Val2 T2
	Val3 T3
	Val4 T4
	Val5 T5
	Val6 T6
	Val7 T7
	Val8 T8
	Val9 T9
}

func Lns_CreateLnsTuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9 any](
    arg1 T1, arg2 T2, arg3 T3, arg4 T4, arg5 T5, arg6 T6, arg7 T7, arg8 T8, arg9 T9 ) LnsTuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9] {
    return LnsTuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]{
        arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9 }
}

type LnsTuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10 any] struct {
	Val1  T1
	Val2  T2
	Val3  T3
	Val4  T4
	Val5  T5
	Val6  T6
	Val7  T7
	Val8  T8
	Val9  T9
	Val10 T10
}

func Lns_CreateLnsTuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10 any](
    arg1 T1, arg2 T2, arg3 T3, arg4 T4, arg5 T5, arg6 T6, arg7 T7, arg8 T8, arg9 T9, arg10 T10 ) LnsTuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10] {
    return LnsTuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]{
        arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10 }
}

