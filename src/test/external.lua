local obj = {}

function obj.func1()
   print( "func1" )
end

obj.val = 0

function obj:func2( val )
   self.val = self.val + val
   print( "func2", self.val )
end

return obj
