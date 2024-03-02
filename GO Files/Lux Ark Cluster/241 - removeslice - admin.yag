{{/*removeslice name of slice, item to remove - removes an item from a slice*/}}
{{$find := (index .CmdArgs 1)}}
{{$slice_db := (dbGet 0 (index .CmdArgs 0)).Value}}
{{$slice := (cslice).AppendSlice $slice_db}}
{{$slice_new := cslice}}
{{- range $slice -}}
{{- if ne . $find -}}
{{- $slice_new = $slice_new.Append . -}}
{{- end -}}{{- end -}}
Removed {{$find}} from Slice {{(index .CmdArgs 0)}}.
{{dbSet 0 (index .CmdArgs 0) $slice_new}}
{{deleteTrigger 10}}
{{deleteResponse 10}}