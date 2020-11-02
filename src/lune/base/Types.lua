--lune/base/Types.lns
local _moduleObj = {}
local __mod__ = '@lune.@base.@Types'
local _lune = {}
if _lune2 then
   _lune = _lune2
end
if not _lune2 then
   _lune2 = _lune
end
local Conv = {}
_moduleObj.Conv = Conv
Conv._val2NameMap = {}
function Conv:_getTxt( val )
   local name = self._val2NameMap[ val ]
   if name then
      return string.format( "Conv.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end
function Conv._from( val )
   if Conv._val2NameMap[ val ] then
      return val
   end
   return nil
end
    
Conv.__allList = {}
function Conv.get__allList()
   return Conv.__allList
end

Conv.C = 0
Conv._val2NameMap[0] = 'C'
Conv.__allList[1] = Conv.C
Conv.Go = 1
Conv._val2NameMap[1] = 'Go'
Conv.__allList[2] = Conv.Go

local CheckingUptodateMode = {}
_moduleObj.CheckingUptodateMode = CheckingUptodateMode
CheckingUptodateMode._val2NameMap = {}
function CheckingUptodateMode:_getTxt( val )
   local name = self._val2NameMap[ val ]
   if name then
      return string.format( "CheckingUptodateMode.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end
function CheckingUptodateMode._from( val )
   if CheckingUptodateMode._val2NameMap[ val ] then
      return val
   end
   return nil
end
    
CheckingUptodateMode.__allList = {}
function CheckingUptodateMode.get__allList()
   return CheckingUptodateMode.__allList
end

CheckingUptodateMode.Force = 'force'
CheckingUptodateMode._val2NameMap['force'] = 'Force'
CheckingUptodateMode.__allList[1] = CheckingUptodateMode.Force
CheckingUptodateMode.Normal = 'none'
CheckingUptodateMode._val2NameMap['none'] = 'Normal'
CheckingUptodateMode.__allList[2] = CheckingUptodateMode.Normal
CheckingUptodateMode.Touch = 'touch'
CheckingUptodateMode._val2NameMap['touch'] = 'Touch'
CheckingUptodateMode.__allList[3] = CheckingUptodateMode.Touch


local TransCtrlInfo = {}
_moduleObj.TransCtrlInfo = TransCtrlInfo
function TransCtrlInfo.setmeta( obj )
  setmetatable( obj, { __index = TransCtrlInfo  } )
end
function TransCtrlInfo.new( warningShadowing, compatComment, checkingDefineAbbr, stopByWarning, uptodateMode, validLuaval )
   local obj = {}
   TransCtrlInfo.setmeta( obj )
   if obj.__init then
      obj:__init( warningShadowing, compatComment, checkingDefineAbbr, stopByWarning, uptodateMode, validLuaval )
   end
   return obj
end
function TransCtrlInfo:__init( warningShadowing, compatComment, checkingDefineAbbr, stopByWarning, uptodateMode, validLuaval )

   self.warningShadowing = warningShadowing
   self.compatComment = compatComment
   self.checkingDefineAbbr = checkingDefineAbbr
   self.stopByWarning = stopByWarning
   self.uptodateMode = uptodateMode
   self.validLuaval = validLuaval
end


function TransCtrlInfo.create_normal(  )

   return TransCtrlInfo.new(false, false, true, false, CheckingUptodateMode.Touch, false)
end


return _moduleObj
