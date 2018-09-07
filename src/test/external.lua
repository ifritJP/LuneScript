local obj = {}

function obj.func1( val )
   print( "func1", val )
end

obj.val = 0

function obj:func2( val )
   self.val = self.val + val
   print( "func2", self.val )
end

return obj
