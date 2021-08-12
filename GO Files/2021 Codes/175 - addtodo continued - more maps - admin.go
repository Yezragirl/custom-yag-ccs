{{/* Adds tasks to todolist for all maps as separate entries continued */}}
{{$elysian := sendMessageRetID 573860427721605120 (joinStr "" .ExecData.msg " :Elysian")}} 
{{addMessageReactions 573860427721605120 $elysian "✔️" "❌"}}
{{$silvestria := sendMessageRetID 573860427721605120 (joinStr "" .ExecData.msg " :Silvestria")}} 
{{addMessageReactions 573860427721605120 $silvestria "✔️" "❌"}}
{{$grayson := sendMessageRetID 573860427721605120 (joinStr "" .ExecData.msg " :Grayson")}} 
{{addMessageReactions 573860427721605120 $grayson "✔️" "❌"}}
{{$venus := sendMessageRetID 573860427721605120 (joinStr "" .ExecData.msg " :Venus")}} 
{{addMessageReactions 573860427721605120 $venus "✔️" "❌"}}
{{$sphinx := sendMessageRetID 573860427721605120 (joinStr "" .ExecData.msg " :Sphinx")}} 
{{addMessageReactions 573860427721605120 $sphinx "✔️" "❌"}}