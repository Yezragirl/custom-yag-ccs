{{$upsidelist := sdict "a" "ɐ" "b" "q" "c" "ɔ" "d" "p" "e" "ǝ" "f" "ɟ" "g" "ƃ" "h" "ɥ" "i" "ᴉ" "j" "ɾ" "k" "ʞ" "l" "l" "m" "ɯ" "n" "u" "o" "o" "p" "d" "q" "b" "r" "ɹ" "s" "s" "t" "ʇ" "u" "n" "v" "ʌ" "w" "ʍ" "x" "x" "y" "ʎ" "z" "z" "A" "∀" "B" "q" "C" "Ɔ" "D" "p" "E" "Ǝ" "F" "Ⅎ" "G" "פ" "H" "H" "I" "I" "J" "ſ" "K" "ʞ" "L" "˥" "M" "W" "N" "N" "O" "O" "P" "Ԁ" "Q" "Q" "R" "ɹ" "S" "S" "T" "┴" "U" "∩" "V" "Λ" "W" "M" "X" "X" "Y" "⅄" "Z" "Z" "?" "¿" "." "˙" "'" "," "," "'" "!" "¡" "(" ")" ")" "(" "{" "}" "}" "{" "[" "]" "]" "[" ">" "<" "<" ">" "^" "v"}}
{{$output := ""}}
{{range (split .StrippedMsg "")}}
{{if eq (randInt 0 3) 0}} 
          {{- $output = (joinStr "" $output (index $upsidelist .)) -}}
{{else if eq (randInt 0 3) 1}}
		  {{$output = joinStr "" $output (lower .)}}
{{else if eq (randInt 0 3) 2}}
		  {{$output = joinStr "" $output (upper .)}}
     {{- end -}}
{{end}}
{{$output}}





{{$output := ""}}
{{range (split .StrippedMsg "")}}
{{if eq (randInt 0 2) 0}}
{{$output = joinStr "" $output (lower .)}}
{{else}}
{{$output = joinStr "" $output (upper .)}}
{{end}}
{{end}}
{{$output}}
{{deleteTrigger 1}}