{{$slice_db := (dbGet 0 .StrippedMsg).Value}}
{{$slice := (cslice).AppendSlice $slice_db}}
{{$len := len $slice}}
{{sendMessage nil (joinStr "" "Length of Slice: " $len)}}
{{$dotlength := 0}}
{{$printslice := cslice}}
{{$slicelength := 0}}
{{- range $slice -}}
    {{- $dotlength = len . -}}
    {{- $slicelength = len (joinStr "|" $printslice.StringSlice) -}}
    {{- if gt (add $slicelength $dotlength) 1500 -}}
	{{- sendMessage nil (joinStr ", " $printslice.StringSlice) -}}
        {{- $printslice = cslice -}}
	{{- $printslice = $printslice.AppendSlice (cslice .) -}}
    {{- else -}}
        {{- $printslice = $printslice.AppendSlice (cslice .) -}}
    {{- end -}}
{{- end -}}
{{sendMessage nil (joinStr ", " $printslice.StringSlice)}}
Show Slice Complete




{{$slice_db := (dbGet 0 .StrippedMsg).Value}}
{{$slice := (cslice).AppendSlice $slice_db}}
{{$len := len $slice}}
{{sendMessage nil (joinStr "" "Length of Slice: " $len)}}
{{- range $slice -}}
{{.}}
{{- end -}}
Show Slice Complete





{{$slice_db := (dbGet 0 .StrippedMsg).Value}}
{{$slice := (cslice).AppendSlice $slice_db}}
{{$len := len $slice}}
{{sendMessage nil (joinStr "" "Length of Slice: " $len)}}
{{$dotlength := 0}}
{{$printslice := cslice}}
{{$slicelength := 0}}
{{- range $slice -}}
    {{- $dotlength = len . -}}
    {{- $slicelength = len (joinStr "|" $printslice.StringSlice) -}}
    {{- if gt (add $slicelength $dotlength) 1500 -}}
	{{- sendMessage nil (joinStr ", " $printslice.StringSlice) -}}
        {{- $printslice = cslice -}}
	{{- $printslice = $printslice.AppendSlice (cslice .) -}}
    {{- else -}}
        {{- $printslice = $printslice.AppendSlice (cslice .) -}}
    {{- end -}}
{{- end -}}
{{sendMessage nil (joinStr ", " $printslice.StringSlice)}}
Show Slice Complete



{{$find := .StrippedMsg}}
{{$min := 0}}
{{$max := 75}}
{{if .ExecData}}
{{$min = sub .ExecData.count 1}}
{{$max = add $min 75}}
{{$find = .ExecData.find}}
{{end}}
{{$slice_db := (dbGet 0 $find).Value}}
{{$slice := (cslice).AppendSlice $slice_db}}
{{$len := len $slice}}
{{sendMessage nil (joinStr "" "Length of Slice: " $len)}}

{{$count := 0}}
{{$dotlength := 0}}
{{$printslice := cslice}}
{{$slicelength := 0}}
{{- range $slice -}}
    {{- $count = add $count 1 -}}
    {{- if and (gt $count $min) (le $count $max) -}}
    {{- sendMessage nil . -}}
{{- end -}} {{- end -}} 
{{if gt $max $len}}
    {{- else -}}
    {{- execCC 276 nil 10 (sdict "count" $count "find" $find) -}}
{{- end -}}