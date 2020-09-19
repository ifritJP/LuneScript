package main

import "fmt"
import "time"
import "math"

type LnsInt = int
type LnsReal = float64
type LnsAny = interface{}

var LnsNone interface{} = nil

type LnsAlgeVal interface {
    getTxt() string
}


type LnsEnv struct {
    valStack []LnsAny
    stackPos int
}



var cur_LnsEnv = LnsEnv{ []LnsAny{}, -1 }

type LnsList struct {
    Items []LnsAny
}

func NewLnsList( list []LnsAny ) *LnsList {
    return &LnsList{ list }
}
func (lnsList *LnsList) Insert( val LnsAny ) {
    if val != nil {
        lnsList.Items = append( lnsList.Items, val )
    }
}
func (lnsList *LnsList) Remove( index LnsAny ) LnsAny {
    if index == nil {
        ret := lnsList.Items[ len(lnsList.Items) - 1 ]
        lnsList.Items = lnsList.Items[ : len(lnsList.Items) - 1 ]
        return ret
    } else {
        work := index.(LnsInt)
        ret := lnsList.Items[ work ]
        lnsList.Items =
            append( lnsList.Items[ : work ], lnsList.Items[ work+1: ]... )
        return ret
    }
}
func (lnsList *LnsList) GetAt( index int ) LnsAny {
    return lnsList.Items[ index - 1 ]
}


func Lns_isCondTrue( stem LnsAny ) bool {
    if stem == nil {
        return false;
    }
    switch stem.(type) {
    case bool:
        return stem.(bool)
    default:
        return true
    }
}

func Lns_op_not( stem LnsAny ) bool {
    return !Lns_isCondTrue( stem );
}

/** 多値返却の先頭を返す
*/
func Lns_car( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 {
        return nil
    }
    if multi[0] == nil {
        return nil
    }
    return multi[0]
}

func Lns_2DDD( multi ...LnsAny ) []LnsAny {
    return multi
}

func Lns_print( multi []LnsAny ) {
    for index, val := range( multi ) {
        if index != 0 {
            fmt.Print( "\t" )
        }
        fmt.Print( Lns_ToString( val ) )
    }
    fmt.Print( "\n" )
}


func Lns_ToString( val LnsAny ) string {
    if val == nil {
        return "nil"
    }
    switch val.(type) {
    case LnsInt:
        return fmt.Sprintf( "%d", val )
    case LnsReal:
        if digit, frac := math.Modf( val.(LnsReal) ); frac == 0 {
            return fmt.Sprintf( "%g.0", digit )
        }
        return fmt.Sprintf( "%g", val )
    case bool:
        return fmt.Sprintf( "%v", val )
    case string:
        return val.(string);
    default:
        return fmt.Sprintf( "table:%T", val )
    }
}

/**
 * スタックを一段上げる
 */
func Lns_incStack() bool {
    cur_LnsEnv.valStack = append( cur_LnsEnv.valStack, nil )
    cur_LnsEnv.stackPos++;
    return false;
}

/**
 * 値 pVal をスタックの top にセットし、値 pVal の条件判定結果を返す
 *
 * スタックに lns_ddd_t は詰めない。
 * 呼び出し側で lns_ddd_t の先頭要素を指定すること。
 *
 * @param pVal スタックに詰む値
 * @return pVal の条件判定結果。 lns_isCondTrue()。
 */
func Lns_setStackVal( val LnsAny ) bool {
    cur_LnsEnv.valStack[ cur_LnsEnv.stackPos ] = val
    return Lns_isCondTrue( val )
}

/**
 * スタックから値を pop する。
 *
 * @return pop した値。
 */
func Lns_popVal( dummy bool ) LnsAny {
    val := cur_LnsEnv.valStack[ cur_LnsEnv.stackPos ]
    cur_LnsEnv.stackPos--
    cur_LnsEnv.valStack = cur_LnsEnv.valStack[:cur_LnsEnv.stackPos+1]
    return val;
}

func test() {

    fmt.Println( Lns_isCondTrue( 1 ), Lns_isCondTrue( nil ),
        Lns_isCondTrue( true ), Lns_isCondTrue( false ) )


    {
        prev := time.Now()
        lst := [] int { 0, 1, 2, 3, 4, 5, 6, 7, 8, 9 }
        var lst2 [] int
        for count:=1; count< 1000000; count++ {
            lst2 = append( lst, lst... )
        }
        fmt.Println( lst2 )

        lnsList := LnsList{ []LnsAny{} }
        for count:=0; count < 10; count++ {
            lnsList.Insert( count )
        }
        fmt.Println( lnsList )
        lnsList.Remove(nil)
        fmt.Println( lnsList )
        lnsList.Remove(2)
        fmt.Println( lnsList )
        lnsList.Remove(4)
        fmt.Println( lnsList )
        lnsList.Insert( 100 )
        fmt.Println( lnsList )

        for _, val := range( lnsList.Items ) {
            fmt.Println( val )
        }

        hoge := func () [] LnsAny {
            return []LnsAny{ 1 }
        }
        foo := []LnsAny{1, 2, hoge() }
        fmt.Println( foo )

         
        now := time.Now()
        fmt.Println( (now.Sub( prev )).Milliseconds() )
    }
}

