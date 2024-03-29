// This code is transcompiled by LuneScript.
package lnsc
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_Testing bool
var Testing__mod__ string
var Testing_testModuleMap *LnsMap2_[string,*Testing_TestModuleInfo]
type Testing_TestcaseFunc func (_env *LnsEnv, arg1 *Testing_Ctrl)
// 169: decl @lune.@base.@Testing.registerTestcase
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

// 178: decl @lune.@base.@Testing.run
func Testing_run(_env *LnsEnv, modPath string) {
    {
        __forsortCollection0 := Testing_testModuleMap
        __forsortSorted0 := __forsortCollection0.CreateKeyListStr()
        __forsortSorted0.Sort( _env, LnsItemKindStr, nil )
        for _, _name := range( __forsortSorted0.Items ) {
            info := __forsortCollection0.Items[ _name ]
            name := _name
            if name == modPath{
                info.FP.Run(_env)
            }
        }
    }
}

// 185: decl @lune.@base.@Testing.runAll
func Testing_runAll(_env *LnsEnv) {
    {
        __forsortCollection0 := Testing_testModuleMap
        __forsortSorted0 := __forsortCollection0.CreateKeyListStr()
        __forsortSorted0.Sort( _env, LnsItemKindStr, nil )
        for _, ___forsortKey0 := range( __forsortSorted0.Items ) {
            info := __forsortCollection0.Items[ ___forsortKey0 ]
            info.FP.Run(_env)
        }
    }
}

// 191: decl @lune.@base.@Testing.outputAllResult
func Testing_outputAllResult(_env *LnsEnv, stream Lns_oStream) {
    {
        __forsortCollection0 := Testing_testModuleMap
        __forsortSorted0 := __forsortCollection0.CreateKeyListStr()
        __forsortSorted0.Sort( _env, LnsItemKindStr, nil )
        for _, ___forsortKey0 := range( __forsortSorted0.Items ) {
            info := __forsortCollection0.Items[ ___forsortKey0 ]
            info.FP.OutputResult(_env, stream)
        }
    }
    stream.Flush(_env)
}

// 35: decl @lune.@base.@Testing.Result.outputResult
func (self *Testing_Result) OutputResult(_env *LnsEnv, stream Lns_oStream) {
    stream.Write(_env, _env.GetVM().String_format("test total: %s %d (OK:%d, NG:%d)\n", Lns_2DDD(self.name, self.okNum + self.ngNum, self.okNum, self.ngNum)))
}
// 41: decl @lune.@base.@Testing.Result.err
func (self *Testing_Result) Err(_env *LnsEnv, mess string,mod string,lineNo LnsInt) {
    self.ngNum = self.ngNum + 1
    Lns_io_stderr.Write(_env, _env.GetVM().String_format("error: %s:%d: %s\n", Lns_2DDD(mod, lineNo, mess)))
}
// 46: decl @lune.@base.@Testing.Result.isTrue
func (self *Testing_Result) IsTrue(_env *LnsEnv, val1 LnsAny,val1txt string,msg LnsAny,mod string,lineNo LnsInt) bool {
    if val1 == true{
        self.okNum = self.okNum + 1
        return true
    }
    self.FP.Err(_env, _env.GetVM().String_format("not true -- %s:%s:[%s]\n", Lns_2DDD(_env.PopVal( _env.IncStack() ||
        _env.SetStackVal( msg) ||
        _env.SetStackVal( "") ).(string), val1txt, val1)), mod, lineNo)
    return false
}
// 56: decl @lune.@base.@Testing.Result.isNotTrue
func (self *Testing_Result) IsNotTrue(_env *LnsEnv, val1 LnsAny,val1txt string,msg LnsAny,mod string,lineNo LnsInt) bool {
    if Lns_op_not(val1){
        self.okNum = self.okNum + 1
        return true
    }
    self.FP.Err(_env, _env.GetVM().String_format("is true -- %s:%s:[%s]\n", Lns_2DDD(_env.PopVal( _env.IncStack() ||
        _env.SetStackVal( msg) ||
        _env.SetStackVal( "") ).(string), val1txt, val1)), mod, lineNo)
    return false
}
// 66: decl @lune.@base.@Testing.Result.isNil
func (self *Testing_Result) IsNil(_env *LnsEnv, val1 LnsAny,val1txt string,msg LnsAny,mod string,lineNo LnsInt) bool {
    if val1 == nil{
        self.okNum = self.okNum + 1
        return true
    }
    self.FP.Err(_env, _env.GetVM().String_format("is not nil -- %s:%s:[%s]\n", Lns_2DDD(_env.PopVal( _env.IncStack() ||
        _env.SetStackVal( msg) ||
        _env.SetStackVal( "") ).(string), val1txt, val1)), mod, lineNo)
    return false
}
// 76: decl @lune.@base.@Testing.Result.isNotNil
func (self *Testing_Result) IsNotNil(_env *LnsEnv, val1 LnsAny,val1txt string,msg LnsAny,mod string,lineNo LnsInt) bool {
    if val1 != nil{
        self.okNum = self.okNum + 1
        return true
    }
    self.FP.Err(_env, _env.GetVM().String_format("is nil -- %s:%s:[%s]\n", Lns_2DDD(_env.PopVal( _env.IncStack() ||
        _env.SetStackVal( msg) ||
        _env.SetStackVal( "") ).(string), val1txt, val1)), mod, lineNo)
    return false
}
// 86: decl @lune.@base.@Testing.Result.checkEq
func (self *Testing_Result) CheckEq(_env *LnsEnv, val1 LnsAny,val2 LnsAny,val1txt string,val2txt string,msg LnsAny,mod string,lineNo LnsInt) bool {
    if val1 == val2{
        self.okNum = self.okNum + 1
        return true
    }
    self.FP.Err(_env, _env.GetVM().String_format("not equal -- %s:%s:[%s] != %s:[%s]\n", Lns_2DDD(_env.PopVal( _env.IncStack() ||
        _env.SetStackVal( msg) ||
        _env.SetStackVal( "") ).(string), val1txt, val1, val2txt, val2)), mod, lineNo)
    return false
}
// 99: decl @lune.@base.@Testing.Result.checkNotEq
func (self *Testing_Result) CheckNotEq(_env *LnsEnv, val1 LnsAny,val2 LnsAny,val1txt string,val2txt string,msg LnsAny,mod string,lineNo LnsInt) bool {
    if val1 != val2{
        self.okNum = self.okNum + 1
        return true
    }
    self.FP.Err(_env, _env.GetVM().String_format("equal -- %s:%s:[%s] == %s:[%s]\n", Lns_2DDD(_env.PopVal( _env.IncStack() ||
        _env.SetStackVal( msg) ||
        _env.SetStackVal( "") ).(string), val1txt, val1, val2txt, val2)), mod, lineNo)
    return false
}
// 144: decl @lune.@base.@Testing.TestModuleInfo.addCase
func (self *Testing_TestModuleInfo) AddCase(_env *LnsEnv, name string,testCase *Testing_TestCase) {
    self.testcaseMap.Set(name,testCase)
}
// 148: decl @lune.@base.@Testing.TestModuleInfo.run
func (self *Testing_TestModuleInfo) Run(_env *LnsEnv) {
    self.runned = true
    Lns_print(Lns_2DDD(_env.GetVM().String_format("module: %s %s", Lns_2DDD(self.name, _env.GetVM().String_rep("=", 30)))))
    {
        __forsortCollection0 := self.testcaseMap
        __forsortSorted0 := __forsortCollection0.CreateKeyListStr()
        __forsortSorted0.Sort( _env, LnsItemKindStr, nil )
        for _, _name := range( __forsortSorted0.Items ) {
            testcase := __forsortCollection0.Items[ _name ]
            name := _name
            Lns_print(Lns_2DDD(_env.GetVM().String_format("%s: %s", Lns_2DDD(name, _env.GetVM().String_rep("-", 15)))))
            testcase.FP.Get_func(_env)(_env, NewTesting_Ctrl(_env, testcase.FP.Get_result(_env)))
        }
    }
}
// 157: decl @lune.@base.@Testing.TestModuleInfo.outputResult
func (self *Testing_TestModuleInfo) OutputResult(_env *LnsEnv, stream Lns_oStream) {
    if Lns_op_not(self.runned){
        return 
    }
    Lns_print(Lns_2DDD(_env.GetVM().String_format("module: %s %s", Lns_2DDD(self.name, _env.GetVM().String_rep("=", 30)))))
    {
        __forsortCollection0 := self.testcaseMap
        __forsortSorted0 := __forsortCollection0.CreateKeyListStr()
        __forsortSorted0.Sort( _env, LnsItemKindStr, nil )
        for _, ___forsortKey0 := range( __forsortSorted0.Items ) {
            testcase := __forsortCollection0.Items[ ___forsortKey0 ]
            testcase.FP.Get_result(_env).FP.OutputResult(_env, stream)
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
func Testing_Result_toSlice(slice []LnsAny) []*Testing_Result {
    ret := make([]*Testing_Result, len(slice))
    for index, val := range slice {
        ret[index] = val.(Testing_ResultDownCast).ToTesting_Result()
    }
    return ret
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
func Testing_Ctrl_toSlice(slice []LnsAny) []*Testing_Ctrl {
    ret := make([]*Testing_Ctrl, len(slice))
    for index, val := range slice {
        ret[index] = val.(Testing_CtrlDownCast).ToTesting_Ctrl()
    }
    return ret
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
// advertise -- 114
func (self *Testing_Ctrl) CheckEq(_env *LnsEnv, arg1 LnsAny,arg2 LnsAny,arg3 string,arg4 string,arg5 LnsAny,arg6 string,arg7 LnsInt) bool {
    return self.result. FP.CheckEq( _env, arg1,arg2,arg3,arg4,arg5,arg6,arg7)
}
// advertise -- 114
func (self *Testing_Ctrl) CheckNotEq(_env *LnsEnv, arg1 LnsAny,arg2 LnsAny,arg3 string,arg4 string,arg5 LnsAny,arg6 string,arg7 LnsInt) bool {
    return self.result. FP.CheckNotEq( _env, arg1,arg2,arg3,arg4,arg5,arg6,arg7)
}
// advertise -- 114
func (self *Testing_Ctrl) Err(_env *LnsEnv, arg1 string,arg2 string,arg3 LnsInt) {
self.result. FP.Err( _env, arg1,arg2,arg3)
}
// advertise -- 114
func (self *Testing_Ctrl) IsNil(_env *LnsEnv, arg1 LnsAny,arg2 string,arg3 LnsAny,arg4 string,arg5 LnsInt) bool {
    return self.result. FP.IsNil( _env, arg1,arg2,arg3,arg4,arg5)
}
// advertise -- 114
func (self *Testing_Ctrl) IsNotNil(_env *LnsEnv, arg1 LnsAny,arg2 string,arg3 LnsAny,arg4 string,arg5 LnsInt) bool {
    return self.result. FP.IsNotNil( _env, arg1,arg2,arg3,arg4,arg5)
}
// advertise -- 114
func (self *Testing_Ctrl) IsNotTrue(_env *LnsEnv, arg1 LnsAny,arg2 string,arg3 LnsAny,arg4 string,arg5 LnsInt) bool {
    return self.result. FP.IsNotTrue( _env, arg1,arg2,arg3,arg4,arg5)
}
// advertise -- 114
func (self *Testing_Ctrl) IsTrue(_env *LnsEnv, arg1 LnsAny,arg2 string,arg3 LnsAny,arg4 string,arg5 LnsInt) bool {
    return self.result. FP.IsTrue( _env, arg1,arg2,arg3,arg4,arg5)
}
// advertise -- 114
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
func Testing_TestCase_toSlice(slice []LnsAny) []*Testing_TestCase {
    ret := make([]*Testing_TestCase, len(slice))
    for index, val := range slice {
        ret[index] = val.(Testing_TestCaseDownCast).ToTesting_TestCase()
    }
    return ret
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
    Get_testcaseMap(_env *LnsEnv) *LnsMap2_[string,*Testing_TestCase]
    OutputResult(_env *LnsEnv, arg1 Lns_oStream)
    Run(_env *LnsEnv)
}
type Testing_TestModuleInfo struct {
    runned bool
    name string
    testcaseMap *LnsMap2_[string,*Testing_TestCase]
    FP Testing_TestModuleInfoMtd
}
func Testing_TestModuleInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Testing_TestModuleInfo).FP
}
func Testing_TestModuleInfo_toSlice(slice []LnsAny) []*Testing_TestModuleInfo {
    ret := make([]*Testing_TestModuleInfo, len(slice))
    for index, val := range slice {
        ret[index] = val.(Testing_TestModuleInfoDownCast).ToTesting_TestModuleInfo()
    }
    return ret
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
func (self *Testing_TestModuleInfo) Get_testcaseMap(_env *LnsEnv) *LnsMap2_[string,*Testing_TestCase]{ return self.testcaseMap }
// 138: DeclConstr
func (self *Testing_TestModuleInfo) InitTesting_TestModuleInfo(_env *LnsEnv, name string) {
    self.runned = false
    self.name = name
    self.testcaseMap = NewLnsMap2_[string,*Testing_TestCase]( map[string]*Testing_TestCase{})
}


func Lns_Testing_init(_env *LnsEnv) {
    if init_Testing { return }
    init_Testing = true
    Testing__mod__ = "@lune.@base.@Testing"
    Lns_InitMod()
    Testing_testModuleMap = NewLnsMap2_[string,*Testing_TestModuleInfo]( map[string]*Testing_TestModuleInfo{})
}
func init() {
    init_Testing = false
}
