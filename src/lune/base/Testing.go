// This code is transcompiled by LuneScript.
package lnsc
import . "github.com/ifritJP/LuneScript/lune/base/runtime_go"
var init_Testing bool
var Testing__mod__ string
var Testing_testModuleMap *LnsMap
type Testing_TestcaseFunc func (arg1 *Testing_Ctrl)
// 164: decl @lune.@base.@Testing.registerTestcase
func Testing_registerTestcase(modName string,caseName string,testcase Testing_TestcaseFunc) {
    var info *Testing_TestModuleInfo
    
    {
        _info := Testing_testModuleMap.Items[modName]
        if _info == nil{
            info = NewTesting_TestModuleInfo(modName)
            
            Testing_testModuleMap.Set(modName,info)
        } else {
            info = _info.(*Testing_TestModuleInfo)
        }
    }
    var result *Testing_Result
    result = NewTesting_Result(caseName, 0, 0)
    info.FP.AddCase(caseName, NewTesting_TestCase(testcase, result))
}

// 173: decl @lune.@base.@Testing.run
func Testing_run(modPath string) {
    {
        __collection866 := Testing_testModuleMap
        __sorted866 := __collection866.CreateKeyListStr()
        __sorted866.Sort( LnsItemKindStr, nil )
        for _, _name := range( __sorted866.Items ) {
            info := __collection866.Items[ _name ].(Testing_TestModuleInfoDownCast).ToTesting_TestModuleInfo()
            name := _name.(string)
            if name == modPath{
                info.FP.Run()
            }
        }
    }
}

// 180: decl @lune.@base.@Testing.runAll
func Testing_runAll() {
    {
        __collection879 := Testing_testModuleMap
        __sorted879 := __collection879.CreateKeyListStr()
        __sorted879.Sort( LnsItemKindStr, nil )
        for _, ___key879 := range( __sorted879.Items ) {
            info := __collection879.Items[ ___key879 ].(Testing_TestModuleInfoDownCast).ToTesting_TestModuleInfo()
            info.FP.Run()
        }
    }
}

// 186: decl @lune.@base.@Testing.outputAllResult
func Testing_outputAllResult(stream Lns_oStream) {
    {
        __collection897 := Testing_testModuleMap
        __sorted897 := __collection897.CreateKeyListStr()
        __sorted897.Sort( LnsItemKindStr, nil )
        for _, ___key897 := range( __sorted897.Items ) {
            info := __collection897.Items[ ___key897 ].(Testing_TestModuleInfoDownCast).ToTesting_TestModuleInfo()
            info.FP.OutputResult(stream)
        }
    }
}

// declaration Class -- Result
type Testing_ResultMtd interface {
    CheckEq(arg1 LnsAny, arg2 LnsAny, arg3 string, arg4 string, arg5 LnsAny, arg6 string, arg7 LnsInt) bool
    CheckNotEq(arg1 LnsAny, arg2 LnsAny, arg3 string, arg4 string, arg5 LnsAny, arg6 string, arg7 LnsInt) bool
    Err(arg1 string, arg2 string, arg3 LnsInt)
    IsNil(arg1 LnsAny, arg2 string, arg3 LnsAny, arg4 string, arg5 LnsInt) bool
    IsNotNil(arg1 LnsAny, arg2 string, arg3 LnsAny, arg4 string, arg5 LnsInt) bool
    IsNotTrue(arg1 LnsAny, arg2 string, arg3 LnsAny, arg4 string, arg5 LnsInt) bool
    IsTrue(arg1 LnsAny, arg2 string, arg3 LnsAny, arg4 string, arg5 LnsInt) bool
    OutputResult(arg1 Lns_oStream)
}
type Testing_Result struct {
    name string
    okNum LnsInt
    ngNum LnsInt
    FP Testing_ResultMtd
}
func Testing_Result2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Testing_Result).FP
}
type Testing_ResultDownCast interface {
    ToTesting_Result() *Testing_Result
}
func Testing_ResultDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Testing_ResultDownCast)
    if ok { return work.ToTesting_Result() }
    return nil
}
func (obj *Testing_Result) ToTesting_Result() *Testing_Result {
    return obj
}
func NewTesting_Result(arg1 string, arg2 LnsInt, arg3 LnsInt) *Testing_Result {
    obj := &Testing_Result{}
    obj.FP = obj
    obj.InitTesting_Result(arg1, arg2, arg3)
    return obj
}
func (self *Testing_Result) InitTesting_Result(arg1 string, arg2 LnsInt, arg3 LnsInt) {
    self.name = arg1
    self.okNum = arg2
    self.ngNum = arg3
}
// 30: decl @lune.@base.@Testing.Result.outputResult
func (self *Testing_Result) OutputResult(stream Lns_oStream) {
    stream.Write(Lns_getVM().String_format("test total: %s %d (OK:%d, NG:%d)\n", []LnsAny{self.name, self.okNum + self.ngNum, self.okNum, self.ngNum}))
}

// 36: decl @lune.@base.@Testing.Result.err
func (self *Testing_Result) Err(mess string,mod string,lineNo LnsInt) {
    self.ngNum = self.ngNum + 1
    
    Lns_io_stderr.Write(Lns_getVM().String_format("error: %s:%d: %s\n", []LnsAny{mod, lineNo, mess}))
}

// 41: decl @lune.@base.@Testing.Result.isTrue
func (self *Testing_Result) IsTrue(val1 LnsAny,val1txt string,msg LnsAny,mod string,lineNo LnsInt) bool {
    if val1 == true{
        self.okNum = self.okNum + 1
        
        return true
    }
    self.FP.Err(Lns_getVM().String_format("not true -- %s:%s:[%s]\n", []LnsAny{Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( msg) ||
        Lns_GetEnv().SetStackVal( "") ).(string), val1txt, val1}), mod, lineNo)
    return false
}

// 51: decl @lune.@base.@Testing.Result.isNotTrue
func (self *Testing_Result) IsNotTrue(val1 LnsAny,val1txt string,msg LnsAny,mod string,lineNo LnsInt) bool {
    if Lns_op_not(val1){
        self.okNum = self.okNum + 1
        
        return true
    }
    self.FP.Err(Lns_getVM().String_format("is true -- %s:%s:[%s]\n", []LnsAny{Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( msg) ||
        Lns_GetEnv().SetStackVal( "") ).(string), val1txt, val1}), mod, lineNo)
    return false
}

// 61: decl @lune.@base.@Testing.Result.isNil
func (self *Testing_Result) IsNil(val1 LnsAny,val1txt string,msg LnsAny,mod string,lineNo LnsInt) bool {
    if val1 == nil{
        self.okNum = self.okNum + 1
        
        return true
    }
    self.FP.Err(Lns_getVM().String_format("is not nil -- %s:%s:[%s]\n", []LnsAny{Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( msg) ||
        Lns_GetEnv().SetStackVal( "") ).(string), val1txt, val1}), mod, lineNo)
    return false
}

// 71: decl @lune.@base.@Testing.Result.isNotNil
func (self *Testing_Result) IsNotNil(val1 LnsAny,val1txt string,msg LnsAny,mod string,lineNo LnsInt) bool {
    if val1 != nil{
        self.okNum = self.okNum + 1
        
        return true
    }
    self.FP.Err(Lns_getVM().String_format("is nil -- %s:%s:[%s]\n", []LnsAny{Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( msg) ||
        Lns_GetEnv().SetStackVal( "") ).(string), val1txt, val1}), mod, lineNo)
    return false
}

// 81: decl @lune.@base.@Testing.Result.checkEq
func (self *Testing_Result) CheckEq(val1 LnsAny,val2 LnsAny,val1txt string,val2txt string,msg LnsAny,mod string,lineNo LnsInt) bool {
    if val1 == val2{
        self.okNum = self.okNum + 1
        
        return true
    }
    self.FP.Err(Lns_getVM().String_format("not equal -- %s:%s:[%s] != %s:[%s]\n", []LnsAny{Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( msg) ||
        Lns_GetEnv().SetStackVal( "") ).(string), val1txt, val1, val2txt, val2}), mod, lineNo)
    return false
}

// 94: decl @lune.@base.@Testing.Result.checkNotEq
func (self *Testing_Result) CheckNotEq(val1 LnsAny,val2 LnsAny,val1txt string,val2txt string,msg LnsAny,mod string,lineNo LnsInt) bool {
    if val1 != val2{
        self.okNum = self.okNum + 1
        
        return true
    }
    self.FP.Err(Lns_getVM().String_format("equal -- %s:%s:[%s] == %s:[%s]\n", []LnsAny{Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( msg) ||
        Lns_GetEnv().SetStackVal( "") ).(string), val1txt, val1, val2txt, val2}), mod, lineNo)
    return false
}


// declaration Class -- Ctrl
type Testing_CtrlMtd interface {
    CheckEq(arg1 LnsAny, arg2 LnsAny, arg3 string, arg4 string, arg5 LnsAny, arg6 string, arg7 LnsInt) bool
    CheckNotEq(arg1 LnsAny, arg2 LnsAny, arg3 string, arg4 string, arg5 LnsAny, arg6 string, arg7 LnsInt) bool
    Err(arg1 string, arg2 string, arg3 LnsInt)
    Get_result() *Testing_Result
    IsNil(arg1 LnsAny, arg2 string, arg3 LnsAny, arg4 string, arg5 LnsInt) bool
    IsNotNil(arg1 LnsAny, arg2 string, arg3 LnsAny, arg4 string, arg5 LnsInt) bool
    IsNotTrue(arg1 LnsAny, arg2 string, arg3 LnsAny, arg4 string, arg5 LnsInt) bool
    IsTrue(arg1 LnsAny, arg2 string, arg3 LnsAny, arg4 string, arg5 LnsInt) bool
    OutputResult(arg1 Lns_oStream)
}
type Testing_Ctrl struct {
    result *Testing_Result
    FP Testing_CtrlMtd
}
func Testing_Ctrl2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Testing_Ctrl).FP
}
type Testing_CtrlDownCast interface {
    ToTesting_Ctrl() *Testing_Ctrl
}
func Testing_CtrlDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Testing_CtrlDownCast)
    if ok { return work.ToTesting_Ctrl() }
    return nil
}
func (obj *Testing_Ctrl) ToTesting_Ctrl() *Testing_Ctrl {
    return obj
}
func NewTesting_Ctrl(arg1 *Testing_Result) *Testing_Ctrl {
    obj := &Testing_Ctrl{}
    obj.FP = obj
    obj.InitTesting_Ctrl(arg1)
    return obj
}
func (self *Testing_Ctrl) InitTesting_Ctrl(arg1 *Testing_Result) {
    self.result = arg1
}
func (self *Testing_Ctrl) Get_result() *Testing_Result{ return self.result }
func (self *Testing_Ctrl) CheckEq(arg1 LnsAny,arg2 LnsAny,arg3 string,arg4 string,arg5 LnsAny,arg6 string,arg7 LnsInt) bool {
    return self.result. FP.CheckEq( arg1,arg2,arg3,arg4,arg5,arg6,arg7)
}
func (self *Testing_Ctrl) CheckNotEq(arg1 LnsAny,arg2 LnsAny,arg3 string,arg4 string,arg5 LnsAny,arg6 string,arg7 LnsInt) bool {
    return self.result. FP.CheckNotEq( arg1,arg2,arg3,arg4,arg5,arg6,arg7)
}
func (self *Testing_Ctrl) Err(arg1 string,arg2 string,arg3 LnsInt) {
self.result. FP.Err( arg1,arg2,arg3)
}
func (self *Testing_Ctrl) IsNil(arg1 LnsAny,arg2 string,arg3 LnsAny,arg4 string,arg5 LnsInt) bool {
    return self.result. FP.IsNil( arg1,arg2,arg3,arg4,arg5)
}
func (self *Testing_Ctrl) IsNotNil(arg1 LnsAny,arg2 string,arg3 LnsAny,arg4 string,arg5 LnsInt) bool {
    return self.result. FP.IsNotNil( arg1,arg2,arg3,arg4,arg5)
}
func (self *Testing_Ctrl) IsNotTrue(arg1 LnsAny,arg2 string,arg3 LnsAny,arg4 string,arg5 LnsInt) bool {
    return self.result. FP.IsNotTrue( arg1,arg2,arg3,arg4,arg5)
}
func (self *Testing_Ctrl) IsTrue(arg1 LnsAny,arg2 string,arg3 LnsAny,arg4 string,arg5 LnsInt) bool {
    return self.result. FP.IsTrue( arg1,arg2,arg3,arg4,arg5)
}
func (self *Testing_Ctrl) OutputResult(arg1 Lns_oStream) {
self.result. FP.OutputResult( arg1)
}

// declaration Class -- TestCase
type Testing_TestCaseMtd interface {
    Get_func() Testing_TestcaseFunc
    Get_result() *Testing_Result
}
type Testing_TestCase struct {
    _func Testing_TestcaseFunc
    result *Testing_Result
    FP Testing_TestCaseMtd
}
func Testing_TestCase2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Testing_TestCase).FP
}
type Testing_TestCaseDownCast interface {
    ToTesting_TestCase() *Testing_TestCase
}
func Testing_TestCaseDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Testing_TestCaseDownCast)
    if ok { return work.ToTesting_TestCase() }
    return nil
}
func (obj *Testing_TestCase) ToTesting_TestCase() *Testing_TestCase {
    return obj
}
func NewTesting_TestCase(arg1 Testing_TestcaseFunc, arg2 *Testing_Result) *Testing_TestCase {
    obj := &Testing_TestCase{}
    obj.FP = obj
    obj.InitTesting_TestCase(arg1, arg2)
    return obj
}
func (self *Testing_TestCase) InitTesting_TestCase(arg1 Testing_TestcaseFunc, arg2 *Testing_Result) {
    self._func = arg1
    self.result = arg2
}
func (self *Testing_TestCase) Get_func() Testing_TestcaseFunc{ return self._func }
func (self *Testing_TestCase) Get_result() *Testing_Result{ return self.result }

// declaration Class -- TestModuleInfo
type Testing_TestModuleInfoMtd interface {
    AddCase(arg1 string, arg2 *Testing_TestCase)
    Get_name() string
    Get_runned() bool
    Get_testcaseMap() *LnsMap
    OutputResult(arg1 Lns_oStream)
    Run()
}
type Testing_TestModuleInfo struct {
    runned bool
    name string
    testcaseMap *LnsMap
    FP Testing_TestModuleInfoMtd
}
func Testing_TestModuleInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Testing_TestModuleInfo).FP
}
type Testing_TestModuleInfoDownCast interface {
    ToTesting_TestModuleInfo() *Testing_TestModuleInfo
}
func Testing_TestModuleInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Testing_TestModuleInfoDownCast)
    if ok { return work.ToTesting_TestModuleInfo() }
    return nil
}
func (obj *Testing_TestModuleInfo) ToTesting_TestModuleInfo() *Testing_TestModuleInfo {
    return obj
}
func NewTesting_TestModuleInfo(arg1 string) *Testing_TestModuleInfo {
    obj := &Testing_TestModuleInfo{}
    obj.FP = obj
    obj.InitTesting_TestModuleInfo(arg1)
    return obj
}
func (self *Testing_TestModuleInfo) Get_runned() bool{ return self.runned }
func (self *Testing_TestModuleInfo) Get_name() string{ return self.name }
func (self *Testing_TestModuleInfo) Get_testcaseMap() *LnsMap{ return self.testcaseMap }
// 133: DeclConstr
func (self *Testing_TestModuleInfo) InitTesting_TestModuleInfo(name string) {
    self.runned = false
    
    self.name = name
    
    self.testcaseMap = NewLnsMap( map[LnsAny]LnsAny{})
    
}

// 139: decl @lune.@base.@Testing.TestModuleInfo.addCase
func (self *Testing_TestModuleInfo) AddCase(name string,testCase *Testing_TestCase) {
    self.testcaseMap.Set(name,testCase)
}

// 143: decl @lune.@base.@Testing.TestModuleInfo.run
func (self *Testing_TestModuleInfo) Run() {
    self.runned = true
    
    Lns_print([]LnsAny{Lns_getVM().String_format("module: %s %s", []LnsAny{self.name, Lns_getVM().String_rep("=", 30)})})
    {
        __collection702 := self.testcaseMap
        __sorted702 := __collection702.CreateKeyListStr()
        __sorted702.Sort( LnsItemKindStr, nil )
        for _, _name := range( __sorted702.Items ) {
            testcase := __collection702.Items[ _name ].(Testing_TestCaseDownCast).ToTesting_TestCase()
            name := _name.(string)
            Lns_print([]LnsAny{Lns_getVM().String_format("%s: %s", []LnsAny{name, Lns_getVM().String_rep("-", 15)})})
            testcase.FP.Get_func()(NewTesting_Ctrl(testcase.FP.Get_result()))
        }
    }
}

// 152: decl @lune.@base.@Testing.TestModuleInfo.outputResult
func (self *Testing_TestModuleInfo) OutputResult(stream Lns_oStream) {
    if Lns_op_not(self.runned){
        return 
    }
    Lns_print([]LnsAny{Lns_getVM().String_format("module: %s %s", []LnsAny{self.name, Lns_getVM().String_rep("=", 30)})})
    {
        __collection756 := self.testcaseMap
        __sorted756 := __collection756.CreateKeyListStr()
        __sorted756.Sort( LnsItemKindStr, nil )
        for _, ___key756 := range( __sorted756.Items ) {
            testcase := __collection756.Items[ ___key756 ].(Testing_TestCaseDownCast).ToTesting_TestCase()
            testcase.FP.Get_result().FP.OutputResult(stream)
        }
    }
}


func Lns_Testing_init() {
    if init_Testing { return }
    init_Testing = true
    Testing__mod__ = "@lune.@base.@Testing"
    Lns_InitMod()
    Testing_testModuleMap = NewLnsMap( map[LnsAny]LnsAny{})
}
func init() {
    init_Testing = false
}
