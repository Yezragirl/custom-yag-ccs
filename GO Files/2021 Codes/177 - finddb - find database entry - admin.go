{{/* finddb  */}}
{{$db:= (dbGet (toInt64 (index .Args 1)) (index .Args 2))}}
{{json $db.Value}}