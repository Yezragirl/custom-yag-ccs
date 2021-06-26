{{$lim := 10 -}}
{{if not .ExecData -}}
  {{$tribes :=(cslice).AppendSlice (dbGet 0 "tribes").Value -}}
  {{execCC .CCID nil 2 (sdict "tribes" $tribes) -}}
{{else -}}
  {{$ub := $lim -}}
  {{if lt (len .ExecData.tribes) $lim -}}
    {{$ub = len .ExecData.tribes -}}
  {{end -}}
  {{range (slice .ExecData.tribes 0 $ub) -}}
  {{dbSet 0 (joinStr "" "tribe_" (lower .)) (sdict "name" . "members" cslice)-}}
  {{dbDel 0 (joinStr "" "tribe_" .)-}}
  DB Entry Deleted for {{.}}.
  DB Entry Created for {{(lower .)}}.
  {{end -}}
  {{if gt (len .ExecData.tribes) $lim -}}
    {{execCC .CCID nil 8 (sdict "tribes" (slice .ExecData.tribes $lim)) -}}
  {{end -}}
{{end -}}