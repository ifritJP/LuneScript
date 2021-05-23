// This code is transcompiled by LuneScript.
package lnsc
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_Testing bool
var Testing__mod__ string
var Testing_testModuleMap *LnsMap
type Testing_TestcaseFunc func (_env *LnsEnv, arg1 *Testing_Ctrl)
// 168: decl @lune.@base.@Testing.registerTestcase
func Testing_registerTestcase(_env *LnsEnv, modName string,caseName string,testcase Testing_TestcaseFunc) {
    var info *Testing_TestModuleInfo
    
    {
        _info := Testing_testModuleMap.Get(modName)
        if _info == nil{
            info = NewTesting_TestModuleInfo(_env, modName)
            
            Testing_testModuleMap.Set(modName,info)
        } else {
            info = _info.(*Testing_TestModuleInfo)
        }
    }
    var result *Testing_Result
    result = NewTesting_Result(_env, caseName, 0, 0)
    info.FP.AddCase(_env, caseName, NewTesting_TestCase(_env, testcase, result))
}

// 177: decl @lune.@base.@Testing.run
func Testing_run(_env *LnsEnv, modPath string) {
    {
        __collection866 := Testing_testModuleMap
        __sorted866 := __collection866.CreateKeyListStr()
        __sorted866.Sort( LnsItemKindStr, nil )
        for _, _name := range( __sorted866.Items ) {
            info := __collection866.Items[ _name ].(Testing_TestModuleInfoDownCast).ToTesting_TestModuleInfo()
            name := _name.(string)
            if name == modPath{
                info.FP.Run(_env)
            }
        }
    }
}

// 184: decl @lune.@base.@Testing.runAll
func Testing_runAll(_env *LnsEnv) {
    {
        __collection879 := Testing_testModuleMap
        __sorted879 := __collection879.CreateKeyListStr()
        __sorted879.Sort( LnsItemKindStr, nil )
        for _, ___key879 := range( __sorted879.Items ) {
            info := __collection879.Items[ ___key879 ].(Testing_TestModuleInfoDownCast).ToTesting_TestModuleInfo()
            info.FP.Run(_env)
        }
    }
}

// 190: decl @lune.@base.@Testing.outputAllResult
func Testing_outputAllResult(_env *LnsEnv, stream Lns_oStream) {
    {
        __collection897 := Testing_testModuleMap
        __sorted897 := __collection897.CreateKeyListStr()
        __sorted897.Sort( LnsItemKindStr, nil )
        for _, ___key897 := range( __sorted897.Items ) {
            info := __collection897.Items[ ___key897 ].(Testing_TestModuleInfoDownCast).ToTesting_TestModuleInfo()
            info.FP.OutputResult(_env, stream)
        }
    }
}

// declaration Class -- Result
type Testing_ResultMtd interface {
    CheckEq(_env *LnsEnv, arg1 LnsAny, arg2 LnsAny, arg3 string, arg4 string, arg5 LnsAny, arg6 string, arg7 LnsInt) bool
    CheckNotEq(_env *LnsEnv, arg1 LnsAny, arg2 LnsAny, arg3 string, arg4 string, arg5 LnsAny, arg6 string, arg7 LnsInt) bool
    Err(_env *LnsEnv, arg1 string, arg2 string, arg3 LnsInt)
    IsNil(_env *LnsEnv, arg1 LnsAny, arg2 string, arg3 LnsAny, arg4 string, arg5 LnsInt) bool
    IsNotNil(_env *LnsEnv, arg1 LnsAny, arg2 string, arg3 LnsAny, arg4 string, arg5 LnsInt) bool
    IsNotTrue(_env *LnsEnv, arg1 LnsAny, arg2 string, arg3 LnsAny, arg4 string, arg5 LnsInt) bool
    IsTrue(_env *LnsEnv, arg1 LnsAny, arg2 string, arg3 LnsAny, arg4 string, arg5 LnsInt) bool
    OutputResult(_env *LnsEnv, arg1 Lns_oStream)
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
func NewTesting_Result(_env *LnsEnv, arg1 string, arg2 LnsInt, arg3 LnsInt) *Testing_Result {
    obj := &Testing_Result{}
    obj.FP = obj
    obj.InitTesting_Result(_env, arg1, arg2, arg3)
    return obj
}
func (self *Testing_Result) InitTesting_Result(_env *LnsEnv, arg1 string, arg2 LnsInt, arg3 LnsInt) {
    self.name = arg1
    self.okNum = arg2
    self.ngNum = arg3
}
// 34: decl @lune.@base.@Testing.Result.outputResult
func (self *Testing_Result) OutputResult(_env *LnsEnv, stream Lns_oStream) {
    stream.Write(_env, _env.LuaVM.String_format("test total: %s %d (OK:%d, NG:%d)\n", []LnsAny{self.name, self.okNum + self.ngNum, self.okNum, self.ngNum}))
}

// 40: decl @lune.@base.@Testing.Result.err
func (self *Testing_Result) Err(_env *LnsEnv, mess string,mod string,lineNo LnsInt) {
    self.ngNum = self.ngNum + 1
    
    Lns_io_stderr.Write(_env, _env.LuaVM.String_format("error: %s:%d: %s\n", []LnsAny{mod, lineNo, mess}))
}

// 45: decl @lune.@base.@Testing.Result.isTrue
func (self *Testing_Result) IsTrue(_env *LnsEnv, val1 LnsAny,val1txt string,msg LnsAny,mod string,lineNo LnsInt) bool {
    if val1 == true{
        self.okNum = self.okNum + 1
        
        return true
    }
    self.FP.Err(_env, _env.LuaVM.String_format("not true -- %s:%s:[%s]\n", []LnsAny{_env.PopVal( _env.IncStack() ||
        _env.SetStackVal( msg) ||
        _env.SetStackVal( "") ).(string), val1txt, val1}), mod, lineNo)
    return false
}

// 55: decl @lune.@base.@Testing.Result.isNotTrue
func (self *Testing_Result) IsNotTrue(_env *LnsEnv, val1 LnsAny,val1txt string,msg LnsAny,mod string,lineNo LnsInt) bool {
    if Lns_op_not(val1){
        self.okNum = self.okNum + 1
        
        return true
    }
    self.FP.Err(_env, _env.LuaVM.String_format("is true -- %s:%s:[%s]\n", []LnsAny{_env.PopVal( _env.IncStack() ||
        _env.SetStackVal( msg) ||
        _env.SetStackVal( "") ).(string), val1txt, val1}), mod, lineNo)
    return false
}

// 65: decl @lune.@base.@Testing.Result.isNil
func (self *Testing_Result) IsNil(_env *LnsEnv, val1 LnsAny,val1txt string,msg LnsAny,mod string,lineNo LnsInt) bool {
    if val1 == nil{
        self.okNum = self.okNum + 1
        
        return true
    }
    self.FP.Err(_env, _env.LuaVM.String_format("is not nil -- %s:%s:[%s]\n", []LnsAny{_env.PopVal( _env.IncStack() ||
        _env.SetStackVal( msg) ||
        _env.SetStackVal( "") ).(string), val1txt, val1}), mod, lineNo)
    return false
}

// 75: decl @lune.@base.@Testing.Result.isNotNil
func (self *Testing_Result) IsNotNil(_env *LnsEnv, val1 LnsAny,val1txt string,msg LnsAny,mod string,lineNo LnsInt) bool {
    if val1 != nil{
        self.okNum = self.okNum + 1
        
        return true
    }
    self.FP.Err(_env, _env.LuaVM.String_format("is nil -- %s:%s:[%s]\n", []LnsAny{_env.PopVal( _env.IncStack() ||
        _env.SetStackVal( msg) ||
        _env.SetStackVal( "") ).(string), val1txt, val1}), mod, lineNo)
    return false
}

// 85: decl @lune.@base.@Testing.Result.checkEq
func (self *Testing_Result) CheckEq(_env *LnsEnv, val1 LnsAny,val2 LnsAny,val1txt string,val2txt string,msg LnsAny,mod string,lineNo LnsInt) bool {
    if val1 == val2{
        self.okNum = self.okNum + 1
        
        return true
    }
    self.FP.Err(_env, _env.LuaVM.String_format("not equal -- %s:%s:[%s] != %s:[%s]\n", []LnsAny{_env.PopVal( _env.IncStack() ||
        _env.SetStackVal( msg) ||
        _env.SetStackVal( "") ).(string), val1txt, val1, val2txt, val2}), mod, lineNo)
    return false
}

// 98: decl @lune.@base.@Testing.Result.checkNotEq
func (self *Testing_Result) CheckNotEq(_env *LnsEnv, val1 LnsAny,val2 LnsAny,val1txt string,val2txt string,msg LnsAny,mod string,lineNo LnsInt) bool {
    if val1 != val2{
        self.okNum = self.okNum + 1
        
        return true
    }
    self.FP.Err(_env, _env.LuaVM.String_format("equal -- %s:%s:[%s] == %s:[%s]\n", []LnsAny{_env.PopVal( _env.IncStack() ||
        _env.SetStackVal( msg) ||
        _env.SetStackVal( "") ).(string), val1txt, val1, val2txt, val2}), mod, lineNo)
    return false
}


// declaration Class -- Ctrl
type Testing_CtrlMtd interface {
    CheckEq(_env *LnsEnv, arg1 LnsAny, arg2 LnsAny, arg3 string, arg4 string, arg5 LnsAny, arg6 string, arg7 LnsInt) bool
    CheckNotEq(_env *LnsEnv, arg1 LnsAny, arg2 LnsAny, arg3 string, arg4 string, arg5 LnsAny, arg6 string, arg7 LnsInt) bool
    Err(_env *LnsEnv, arg1 string, arg2 string, arg3 LnsInt)
    Get_result(_env *LnsEnv) *Testing_Result
    IsNil(_env *LnsEnv, arg1 LnsAny, arg2 string, arg3 LnsAny, arg4 string, arg5 LnsInt) bool
    IsNotNil(_env *LnsEnv, arg1 LnsAny, arg2 string, arg3 LnsAny, arg4 string, arg5 LnsInt) bool
    IsNotTrue(_env *LnsEnv, arg1 LnsAny, arg2 string, arg3 LnsAny, arg4 string, arg5 LnsInt) bool
    IsTrue(_env *LnsEnv, arg1 LnsAny, arg2 string, arg3 LnsAny, arg4 string, arg5 LnsInt) bool
    OutputResult(_env *LnsEnv, arg1 Lns_oStream)
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
func NewTesting_Ctrl(_env *LnsEnv, arg1 *Testing_Result) *Testing_Ctrl {
    obj := &Testing_Ctrl{}
    obj.FP = obj
    obj.InitTesting_Ctrl(_env, arg1)
    return obj
}
func (self *Testing_Ctrl) InitTesting_Ctrl(_env *LnsEnv, arg1 *Testing_Result) {
    self.result = arg1
}
func (self *Testing_Ctrl) Get_result(_env *LnsEnv) *Testing_Result{ return self.result }
// advertise -- 113
func (self *Testing_Ctrl) CheckEq(_env *LnsEnv, arg1 LnsAny,arg2 LnsAny,arg3 string,arg4 string,arg5 LnsAny,arg6 string,arg7 LnsInt) bool {
    return self.result. FP.CheckEq( _env, arg1,arg2,arg3,arg4,arg5,arg6,arg7)
}
// advertise -- 113
func (self *Testing_Ctrl) CheckNotEq(_env *LnsEnv, arg1 LnsAny,arg2 LnsAny,arg3 string,arg4 string,arg5 LnsAny,arg6 string,arg7 LnsInt) bool {
    return self.result. FP.CheckNotEq( _env, arg1,arg2,arg3,arg4,arg5,arg6,arg7)
}
// advertise -- 113
func (self *Testing_Ctrl) Err(_env *LnsEnv, arg1 string,arg2 string,arg3 LnsInt) {
self.result. FP.Err( _env, arg1,arg2,arg3)
}
// advertise -- 113
func (self *Testing_Ctrl) IsNil(_env *LnsEnv, arg1 LnsAny,arg2 string,arg3 LnsAny,arg4 string,arg5 LnsInt) bool {
    return self.result. FP.IsNil( _env, arg1,arg2,arg3,arg4,arg5)
}
// advertise -- 113
func (self *Testing_Ctrl) IsNotNil(_env *LnsEnv, arg1 LnsAny,arg2 string,arg3 LnsAny,arg4 string,arg5 LnsInt) bool {
    return self.result. FP.IsNotNil( _env, arg1,arg2,arg3,arg4,arg5)
}
// advertise -- 113
func (self *Testing_Ctrl) IsNotTrue(_env *LnsEnv, arg1 LnsAny,arg2 string,arg3 LnsAny,arg4 string,arg5 LnsInt) bool {
    return self.result. FP.IsNotTrue( _env, arg1,arg2,arg3,arg4,arg5)
}
// advertise -- 113
func (self *Testing_Ctrl) IsTrue(_env *LnsEnv, arg1 LnsAny,arg2 string,arg3 LnsAny,arg4 string,arg5 LnsInt) bool {
    return self.result. FP.IsTrue( _env, arg1,arg2,arg3,arg4,arg5)
}
// advertise -- 113
func (self *Testing_Ctrl) OutputResult(_env *LnsEnv, arg1 Lns_oStream) {
self.result. FP.OutputResult( _env, arg1)
}

// declaration Class -- TestCase
type Testing_TestCaseMtd interface {
    Get_func(_env *LnsEnv) Testing_TestcaseFunc
    Get_result(_env *LnsEnv) *Testing_Result
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
func NewTesting_TestCase(_env *LnsEnv, arg1 Testing_TestcaseFunc, arg2 *Testing_Result) *Testing_TestCase {
    obj := &Testing_TestCase{}
    obj.FP = obj
    obj.InitTesting_TestCase(_env, arg1, arg2)
    return obj
}
func (self *Testing_TestCase) InitTesting_TestCase(_env *LnsEnv, arg1 Testing_TestcaseFunc, arg2 *Testing_Result) {
    self._func = arg1
    self.result = arg2
}
func (self *Testing_TestCase) Get_func(_env *LnsEnv) Testing_TestcaseFunc{ return self._func }
func (self *Testing_TestCase) Get_result(_env *LnsEnv) *Testing_Result{ return self.result }

// declaration Class -- TestModuleInfo
type Testing_TestModuleInfoMtd interface {
    AddCase(_env *LnsEnv, arg1 string, arg2 *Testing_TestCase)
    Get_name(_env *LnsEnv) string
    Get_runned(_env *LnsEnv) bool
    Get_testcaseMap(_env *LnsEnv) *LnsMap
    OutputResult(_env *LnsEnv, arg1 Lns_oStream)
    Run(_env *LnsEnv)
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
func NewTesting_TestModuleInfo(_env *LnsEnv, arg1 string) *Testing_TestModuleInfo {
    obj := &Testing_TestModuleInfo{}
    obj.FP = obj
    obj.InitTesting_TestModuleInfo(_env, arg1)
    return obj
}
func (self *Testing_TestModuleInfo) Get_runned(_env *LnsEnv) bool{ return self.runned }
func (self *Testing_TestModuleInfo) Get_name(_env *LnsEnv) string{ return self.name }
func (self *Testing_TestModuleInfo) Get_testcaseMap(_env *LnsEnv) *LnsMap{ return self.testcaseMap }
// 137: DeclConstr
func (self *Testing_TestModuleInfo) InitTesting_TestModuleInfo(_env *LnsEnv, name string) {
    self.runned = false
    
    self.name = name
    
    self.testcaseMap = NewLnsMap( map[LnsAny]LnsAny{})
    
}

// 143: decl @lune.@base.@Testing.TestModuleInfo.addCase
func (self *Testing_TestModuleInfo) AddCase(_env *LnsEnv, name string,testCase *Testing_TestCase) {
    self.testcaseMap.Set(name,testCase)
}

// 147: decl @lune.@base.@Testing.TestModuleInfo.run
func (self *Testing_TestModuleInfo) Run(_env *LnsEnv) {
    self.runned = true
    
    Lns_print([]LnsAny{_env.LuaVM.String_format("module: %s %s", []LnsAny{self.name, _env.LuaVM.String_rep("=", 30)})})
    {
        __collection702 := self.testcaseMap
        __sorted702 := __collection702.CreateKeyListStr()
        __sorted702.Sort( LnsItemKindStr, nil )
        for _, _name := range( __sorted702.Items ) {
            testcase := __collection702.Items[ _name ].(Testing_TestCaseDownCast).ToTesting_TestCase()
            name := _name.(string)
            Lns_print([]LnsAny{_env.LuaVM.String_format("%s: %s", []LnsAny{name, _env.LuaVM.String_rep("-", 15)})})
            testcase.FP.Get_func(_env)(_env, NewTesting_Ctrl(_env, testcase.FP.Get_result(_env)))
        }
    }
}

// 156: decl @lune.@base.@Testing.TestModuleInfo.outputResult
func (self *Testing_TestModuleInfo) OutputResult(_env *LnsEnv, stream Lns_oStream) {
    if Lns_op_not(self.runned){
        return 
    }
    Lns_print([]LnsAny{_env.LuaVM.String_format("module: %s %s", []LnsAny{self.name, _env.LuaVM.String_rep("=", 30)})})
    {
        __collection756 := self.testcaseMap
        __sorted756 := __collection756.CreateKeyListStr()
        __sorted756.Sort( LnsItemKindStr, nil )
        for _, ___key756 := range( __sorted756.Items ) {
            testcase := __collection756.Items[ ___key756 ].(Testing_TestCaseDownCast).ToTesting_TestCase()
            testcase.FP.Get_result(_env).FP.OutputResult(_env, stream)
        }
    }
}


func Lns_Testing_init(_env *LnsEnv) {
    if init_Testing { return }
    init_Testing = true
    Testing__mod__ = "@lune.@base.@Testing"
    Lns_InitMod()
    Testing_testModuleMap = NewLnsMap( map[LnsAny]LnsAny{})
}
func init() {
    init_Testing = false
}
