{{$md := 1}}

{{if (dbGet .Reaction.UserID "key")}}{{else}}{{dbSet .Reaction.UserID "key" 1}}{{end}}

{{if eq .Reaction.Emoji.ID 656672093244620800}}{{$md = 2}}{{else if eq .Reaction.Emoji.ID 656672093274112020}}{{$md = 3}}{{else if eq .Reaction.Emoji.ID 656672093320249354}}{{$md = 4}}{{else if eq .Reaction.Emoji.ID 656672092934111282}}{{$md = 5}}{{else if eq .Reaction.Emoji.ID 656672093395484702}}{{$md = 7}}{{else if eq .Reaction.Emoji.ID 656672093110272055}}{{$md = 9}}{{else if eq .Reaction.Emoji.ID 656672093143957504}}{{$md = 11}}{{else if eq .Reaction.Emoji.ID 656672093622108171}}{{$md = 13}}{{else if eq .Reaction.Emoji.ID 656672092816670732}}{{$md = 16}}{{else if eq .Reaction.Emoji.ID 656672093014065183}}{{$md = 17}}{{else if eq .Reaction.Emoji.ID 656672093202808845}}{{$md = 19}}{{else if eq .Reaction.Emoji.ID 656672093655662592}}{{$md = 23}}{{else if eq .Reaction.Emoji.ID 656672093647274005}}{{$md = 25}}{{else if eq .Reaction.Emoji.ID 656672093941006346}}{{$md = 29}}{{else if eq .Reaction.Emoji.ID 656672094335270934}}{{$md = 31}}{{else if eq .Reaction.Emoji.ID 656672093764845619}}{{$md = 37}}{{else if eq .Reaction.Emoji.ID 656672094075092992}}{{$md = 41}}{{else if eq .Reaction.Emoji.ID 656672094599512074}}{{$md = 43}}{{else if eq .Reaction.Emoji.ID 656672094267899961}}{{$md = 47}}{{else if eq .Reaction.Emoji.ID 656672094553112616}}{{$md = 49}}{{else if eq .Reaction.Emoji.ID 656672094863622164}}{{$md = 53}}{{else if eq .Reaction.Emoji.ID 656672094930599936}}{{$md = 59}}{{else if eq .Reaction.Emoji.ID 656672094590861312}}{{$md = 61}}{{else if eq .Reaction.Emoji.ID 656672094360436752}}{{$md = 67}}{{else if eq .Reaction.Emoji.ID 656672095765266434}}{{$md = 71}}{{else if eq .Reaction.Emoji.ID 656672094486003722}}{{$md = 73}}{{else if eq .Reaction.Emoji.ID 656672094502780928}}{{$md = 79}}{{else if eq .Reaction.Emoji.ID 656672094771347496}}{{$md = 81}}{{else if eq .Reaction.Emoji.ID 656672094704369716}}{{$md = 83}}{{else if eq .Reaction.Emoji.ID 656672094570020864}}{{$md = 89}}{{else if eq .Reaction.Emoji.ID 656672094863622167}}{{$md = 97}}{{else if eq .Reaction.Emoji.ID 656672094825742356}}{{$md = 101}}{{else if eq .Reaction.Emoji.ID 656672094687330334}}{{$md = 103}}{{else if eq .Reaction.Emoji.ID 656672094880268288}}{{$md = 107}}{{else if eq .Reaction.Emoji.ID 656672094762958868}}{{$md = 109}}{{else if eq .Reaction.Emoji.ID 656672095153029120}}{{$md = 113}}{{else if eq .Reaction.Emoji.ID 656672094758633482}}{{$md = 121}}{{else if eq .Reaction.Emoji.ID 656672095136382977}}{{$md = 127}}{{else if eq .Reaction.Emoji.ID 656672094335139873}}{{$md = 131}}{{else if eq .Reaction.Emoji.ID 656672094863753263}}{{$md = 137}}{{else if eq .Reaction.Emoji.ID 656672094855233536}}{{$md = 139}}{{else if eq .Reaction.Emoji.ID 656672095190646817}}{{$md = 149}}{{else if eq .Reaction.Emoji.ID 656672094595317771}}{{$md = 151}}{{else if eq .Reaction.Emoji.ID 656672094733598730}}{{$md = 157}}{{else if eq .Reaction.Emoji.ID 656672094909759489}}{{$md = 163}}{{else if eq .Reaction.Emoji.ID 656672095077400586}}{{$md = 167}}{{else if eq .Reaction.Emoji.ID 656672094922342400}}{{$md = 169}}{{else if eq .Reaction.Emoji.ID 656672094972542977}}{{$md = 173}}{{else if eq .Reaction.Emoji.ID 656672094624415791}}{{$md = 179}}{{else if eq .Reaction.Emoji.ID 656672094574346291}}{{$md = 181}}{{end}}

{{$db := (toInt (dbGet .Reaction.UserID "key").Value)}}

{{if .ReactionAdded}}{{$db = mult $db $md}}{{else}}{{$db = div $db $md}}{{end}}
{{dbSet .Reaction.UserID "key" (toString $db)}}





{{if eq (toInt (mod $db 16)) 0}}
{{if eq (toInt (mod $db 64)) 0}}
{{if eq (toInt (mod $db 128)) 0}}{{$A := 1}}{{$As := 1}}{{$Cs := 1}}
{{else}}{{$A := 0}}{{$As := 1}}{{$Cs := 1}}
{{end}}
{{else}}
{{if eq (toInt (mod $db 32)) 0}}
{{$A := 1}}{{$As := 0}}{{$Cs := 1}}
{{else}}
{{$A := 0}}{{$As := 0}}{{$Cs := 1}}
{{end}}
{{end}}
{{else}}
{{if eq (toInt (mod $db 4)) 0}}
{{if eq (toInt (mod $db 8)) 0}}
{{$A := 1}}{{$As := 1}}{{$Cs := 0}}
{{else}}
{{$A := 0}}{{$As := 1}}{{$Cs := 0}}
{{end}}
{{else}}
{{if eq (toInt (mod $db 2)) 0}}
{{$A := 1}}{{$As := 0}}{{$Cs := 0}}
{{else}}
{{$A := 0}}{{$As := 0}}{{$Cs := 0}}
{{end}}
{{end}}
{{end}}


{{if eq (toInt (mod $db 81)) 0}}
{{if eq (toInt (mod $db 729)) 0}}
{{if eq (toInt (mod $db 2187)) 0}}
{{$Am := 1}}{{$B := 1}}{{$L := 1}}
{{else}}
{{$Am := 0}}{{$B := 1}}{{$L := 1}}
{{end}}
{{else}}
{{if eq (toInt (mod $db 243)) 0}}
{{$Am := 1}}{{$B := 0}}{{$L := 1}}
{{else}}
{{$Am := 0}}{{$B := 0}}{{$L := 1}}
{{end}}
{{end}}
{{else}}
{{if eq (toInt (mod $db 9)) 0}}
{{if eq (toInt (mod $db 27)) 0}}
{{$Am := 1}}{{$B := 1}}{{$L := 0}}
{{else}}
{{$Am := 0}}{{$B := 1}}{{$L := 0}}
{{end}}
{{else}}
{{if eq (toInt (mod $db 3)) 0}}
{{$Am := 1}}{{$B := 0}}{{$L := 0}}
{{else}}
{{$Am := 0}}{{$B := 0}}{{$L := 0}}
{{end}}
{{end}}
{{end}}


{{if eq (toInt (mod $db 25)) 0}}
{{if eq (toInt (mod $db 125)) 0}}
{{$Ba := 1}}{{$Ec := 1}}
{{else}}
{{$Ba := 0}}{{$Ec := 1}}
{{end}}
{{else}}
{{if eq (toInt (mod $db 5)) 0}}
{{$Ba := 1}}{{$Ec := 0}}
{{else}}
{{$Ba := 0}}{{$Ec := 0}}
{{end}}
{{end}}

{{if eq (toInt (mod $db 49)) 0}}
{{if eq (toInt (mod $db 343)) 0}}
{{$Bl := 1}}{{$G := 1}}
{{else}}
{{$Bl := 0}}{{$G := 0}}
{{end}}
{{else}}
{{if eq (toInt (mod $db 7)) 0}}
{{$Bl := 1}}{{$G := 0}}
{{else}}
{{$Bl := 0}}{{$G := 0}}
{{end}}
{{end}}


{{if eq (toInt (mod $db 121)) 0}}
{{if eq (toInt (mod $db 1331)) 0}}
{{$Bu := 1}}{{$N := 1}}
{{else}}
{{$Bu := 0}}{{$N := 1}}
{{end}}
{{else}}
{{if eq (toInt (mod $db 11)) 0}}
{{$Bu := 1}}{{$N := 0}}
{{else}}
{{$Bu := 0}}{{$N := 0}}
{{end}}
{{end}}

{{if eq (toInt (mod $db 169)) 0}}
{{if eq (toInt (mod $db 2197)) 0}}
{{$Cp := 1}}{{$W := 1}}
{{else}}
{{$Cp := 0}}{{$W := 1}}
{{end}}
{{else}}
{{if eq (toInt (mod $db 13)) 0}}
{{$Cp := 1}}{{$W := 0}}
{{else}}
{{$Cp := 0}}{{$W := 0}}
{{end}}
{{end}}


{{if eq (toInt (mod $db 17)) 0}}
{{$C := 1}}}}
{{else}}
{{$C := 0}}
{{end}}
{{if eq (toInt (mod $db 19)) 0}}
{{$D := 1}}
{{else}}
{{$D := 0}}
{{end}}
{{if eq (toInt (mod $db 23)) 0}}
{{$Dk := 1}}
{{else}}
{{$Dk := 0}}
{{end}}
{{if eq (toInt (mod $db 29)) 0}}
{{$E := 1}}
{{else}}
{{$E := 0}}
{{end}}
{{if eq (toInt (mod $db 31)) 0}}
{{$Fi := 1}}
{{else}}
{{$Fi := 0}}
{{end}}
{{if eq (toInt (mod $db 37)) 0}}
{{$F := 1}}
{{else}}
{{$F := 0}}
{{end}}
{{if eq (toInt (mod $db 41)) 0}}
{{$Fu := 1}}
{{else}}
{{$Fu := 0}}
{{end}}
{{if eq (toInt (mod $db 43)) 0}}
{{$Gz := 1}}
{{else}}
{{$Gz := 0}}
{{end}}
{{if eq (toInt (mod $db 47)) 0}}
{{$Go := 1}}
{{else}}
{{$Go := 0}}
{{end}}
{{if eq (toInt (mod $db 53)) 0}}
{{$H := 1}}
{{else}}
{{$H := 0}}
{{end}}
{{if eq (toInt (mod $db 59)) 0}}
{{$I := 1}}
{{else}}
{{$I := 0}}
{{end}}
{{if eq (toInt (mod $db 61)) 0}}
{{$Ja := 1}}
{{else}}{{$Ja := 0}}
{{end}}
{{if eq (toInt (mod $db 67)) 0}}
{{$J := 1}}
{{else}}{{$J := 0}}
{{end}}
{{if eq (toInt (mod $db 71)) 0}}{{$Kd := 1}}
{{else}}{{$Kd := 0}}
{{end}}
{{if eq (toInt (mod $db 73)) 0}}
{{$K := 1}}
{{else}}
{{$K := 0}}
{{end}}
{{if eq (toInt (mod $db 79)) 0}}
{{$Kp := 1}}
{{else}}
{{$Kp := 0}}
{{end}}
{{if eq (toInt (mod $db 83)) 0}}
{{$Li := 1}}
{{else}}
{{$Li := 0}}
{{end}}
{{if eq (toInt (mod $db 89)) 0}}
{{$M := 1}}
{{else}}
{{$M := 0}}
{{end}}
{{if eq (toInt (mod $db 97)) 0}}
{{$Mv := 1}}
{{else}}
{{$Mv := 0}}
{{end}}
{{if eq (toInt (mod $db 101)) 0}}
{{$Mi := 1}}
{{else}}
{{$Mi := 0}}
{{end}}
{{if eq (toInt (mod $db 103)) 0}}
{{$Mn := 1}}
{{else}}
{{$Mn := 0}}
{{end}}
{{if eq (toInt (mod $db 107)) 0}}
{{$Mz := 1}}
{{else}}
{{$Mz := 0}}
{{end}}
{{if eq (toInt (mod $db 109)) 0}}
{{$Mu := 1}}
{{else}}
{{$Mu := 0}}
{{end}}
{{if eq (toInt (mod $db 113)) 0}}
{{$Nk := 1}}
{{else}}
{{$Nk := 0}}
{{end}}
{{if eq (toInt (mod $db 127)) 0}}
{{$P := 1}}
{{else}}
{{$P := 0}}
{{end}}
{{if eq (toInt (mod $db 131)) 0}}
{{$R := 1}}
{{else}}
{{$R := 0}}
{{end}}
{{if eq (toInt (mod $db 137)) 0}}
{{$Sg := 1}}
{{else}}
{{$Sg := 0}}
{{end}}
{{if eq (toInt (mod $db 139)) 0}}
{{$S := 1}}
{{else}}
{{$S := 0}}
{{end}}
{{if eq (toInt (mod $db 149)) 0}}
{{$T := 1}}
{{else}}
{{$T := 0}}
{{end}}
{{if eq (toInt (mod $db 151)) 0}}
{{$Th := 1}}
{{else}}
{{$Th := 0}}
{{end}}
{{if eq (toInt (mod $db 157)) 0}}
{{$Tw := 1}}
{{else}}
{{$Tw := 0}}
{{end}}
{{if eq (toInt (mod $db 163)) 0}}
{{$Vk := 1}}
{{else}}
{{$Vk := 0}}
{{end}}
{{if eq (toInt (mod $db 167)) 0}}
{{$V := 1}}
{{else}}
{{$V := 0}}
{{end}}
{{if eq (toInt (mod $db 173)) 0}}
{{$Wr := 1}}
{{else}}
{{$Wr := 0}}
{{end}}
{{if eq (toInt (mod $db 179)) 0}}
{{$Y := 1}}
{{else}}
{{$Y := 0}}
{{end}}
{{if eq (toInt (mod $db 181)) 0}}
{{$Z := 1}}
{{else}}
{{$Z := 0}}
{{end}}



{{if or $Th $H $Mv}}{{addRoleID 656600212760035330}}{{else}}{{removeRoleID 656600212760035330}}{{end}}
{{if or $Z $Sg $As $Bu}}{{addRoleID 656600261095063602}}{{else}}{{removeRoleID 656600261095063602}}{{end}}
{{if or $Mu $Kd $Ba $M}}{{addRoleID 656600483712073761}}{{else}}{{removeRoleID 656600483712073761}}{{end}}
{{if or $S $Go $Cs $F}}{{addRoleID 656681611470831657}}{{else}}{{removeRoleID 656681611470831657}}{{end}}
{{if or $Tw $T $Mv $Z $K}}{{addRoleID 656600333899792416}}{{else}}{{removeRoleID 656600333899792416}}{{end}}
{{if or $J $W}}{{addRoleID 656601756029157408}}{{else}}{{removeRoleID 656601756029157408}}{{end}}
{{if or $Li $Ja $I $G $Dk $N}}{{addRoleID 656683108162076682}}{{else}}{{removeRoleID 656683108162076682}}{{end}}
{{if or $Vk $P $Mz $Mi $M $L $Ec $C $A $Wr}}{{addRoleID 656602904660148243}}{{else}}{{removeRoleID 656602904660148243}}{{end}}
{{if or $Y $T $Nk $I}}{{addRoleID 656684052727857192}}{{else}}{{removeRoleID 656684052727857192}}{{end}}
{{if or $V $Mu $Mz $C}}{{addRoleID 656684324695048203}}{{else}}{{removeRoleID 656684324695048203}}{{end}}
{{if or $L $Kp $F}}{{addRoleID 656685631765872642}}{{else}}{{removeRoleID 656685631765872642}}{{end}}
{{if or $Tw $N $Mn $Li $G $Gz $K $Fu $Fi $Dk $Cp}}{{addRoleID 656687072161169428}}{{else}}{{removeRoleID 656687072161169428}}{{end}}
{{if or $S $Mi $M $Ec $D $A $R $Go}}{{addRoleID 656687242110173185}}{{else}}{{removeRoleID 656687242110173185}}{{end}}
{{if or $S $Mi $M $Ec $Wr $P $E}}{{addRoleID 656689004992790533}}{{else}}{{removeRoleID 656689004992790533}}{{end}}