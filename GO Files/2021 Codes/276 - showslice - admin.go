{{/*showslice shows a slice in a legible way*/}}
{{$find := or .ExecData.find .StrippedMsg}}
{{$min := .ExecData.max|toInt}}
{{$max := add $min 75}}
{{$slice := cslice.AppendSlice (dbGet 0 $find).Value}}
{{$len := len $slice}}
{{if $min}}
	{{else}}
    {{sendMessage nil (print "Length of Slice: " $len)}}
{{end}}
{{range $count, $v := $slice}}
    {{- $count = add $count 1}}
    {{- if and (gt $count $min) (le $count $max)}}
        {{- sendMessage nil .}}
    {{- end -}}
{{end}} 
{{if gt $max $len}}
    Show Slice Complete
{{else}}
    {{execCC 276 nil 10 (sdict "max" $max "find" $find)}}
{{end}}