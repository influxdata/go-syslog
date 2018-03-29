%%{
machine rfc5424;

# 1..191 or 0
prival = (('1' ( '9' ( '0'..'1' ){,1} | '0'..'8' ( '0'..'9' ){,1} ){,1}) | ( '2'..'9' ('0'..'9'){,1} )) | '0';

action acc_prival { privalChars = append(privalChars, fc) }
action int_prival { prival = NewPrival(utf8ToNum(privalChars)) }

pri = '<' . prival @acc_prival %int_prival . '>';
}%%
