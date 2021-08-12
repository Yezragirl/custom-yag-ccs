{{/*addslice adds an item to a specified slice*/}}
{{$find := (index .CmdArgs 1)}}
{{$slice_db := (dbGet 0 (index .CmdArgs 0)).Value}}
{{$slice := (cslice).AppendSlice $slice_db}}
{{$slice = $slice.Append $find}}
Added {{$find}} to Slice {{(index .CmdArgs 0)}}.
{{dbSet 0 (index .CmdArgs 0) $slice}}
{{deleteTrigger 10}}
{{deleteResponse 10}}