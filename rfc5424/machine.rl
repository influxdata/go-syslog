%%{
machine rfc5424;

# unsigned alphabet
alphtype uint8;

action add_char {
    cr.Add(fc)
}

action set_prival {
    prival = NewPrival(*cr.ReduceToInt(chars.UTF8DecimalCodePointsToInt))
}

action set_version {
    version = NewVersion(*cr.ReduceToInt(chars.UTF8DecimalCodePointsToInt))
}

action set_hostname {
    hostname = *cr.ReduceToString(chars.UTF8DecimalCodePointsToString)
}

action set_appname {
    appname = *cr.ReduceToString(chars.UTF8DecimalCodePointsToString)
}

action set_procid {
    fmt.Println("set")
    //procid = *cr.ReduceToString(chars.UTF8DecimalCodePointsToString)
}

action set_msgid {
    msgid = *cr.ReduceToString(chars.UTF8DecimalCodePointsToString)
}

action tag_rfc3339 {
    poss["timestamp:ini"] = fpc
}

action set_rfc3339 {
    if t, e := time.Parse(time.RFC3339Nano, data[poss["timestamp:ini"]:p]); e != nil {
        err = fmt.Errorf("error %s [col %d:%d]", e, poss["timestamp:ini"], p);
        fhold; fgoto line;
    } else {
        timestamp = &t
    }
}

action err_nilvalue {
    err = fmt.Errorf("error parsing <nilvalue>");
}

nilvalue = '-' @lerr(err_nilvalue);

sp = ' ';

nonzerodigit = '1'..'9';

# 0..59
sexagesimal = '0'..'5' . '0'..'9';

printusascii = '!'..'~';

# 1..191 or 0
prival = (('1' ( '9' ( '0'..'1' ){,1} | '0'..'8' ( '0'..'9' ){,1} ){,1}) | ( '2'..'9' ('0'..'9'){,1} )) | '0';

pri = '<' prival $add_char %set_prival '>';

version = (nonzerodigit digit{0,2}) $add_char %set_version;

datemday = ('0' . '1'..'9' | '1'..'2' . '0'..'9' | '3' . '0'..'1');

datemonth = ('0' . '1'..'9' | '1' . '0'..'2');

datefullyear = digit{4};

fulldate = datefullyear  '-' datemonth  '-' datemday;

timehour = ('0'..'1' . '0'..'9' | '2' . '0'..'3');

timeminute = sexagesimal;

timesecond = sexagesimal;

timesecfrac = '.' digit{1,6};

timenumoffset = ('+' | '-') timehour ':' timeminute;

timeoffset = 'Z' | timenumoffset;

partialtime = timehour ':' timeminute ':' timesecond . timesecfrac?;

fulltime = partialtime . timeoffset;

timestamp = nilvalue | (fulldate >tag_rfc3339 . 'T' . fulltime %set_rfc3339); 

hostname = nilvalue | (printusascii{1,255}  $add_char %set_hostname); # !nilvalue

appname = nilvalue | (printusascii{1,48} $add_char %set_appname);

procid = nilvalue | (printusascii{1,128} $add_char %set_procid);

msgid = nilvalue | (printusascii{1,32} @add_char %set_msgid);

header = pri version sp timestamp sp hostname sp appname sp procid sp msgid;

utf8string = any*; # (todo) > complete

sdname = printusascii{1,32} -- ('=' | sp | ']' | '"');

paramvalue = utf8string; # (todo) > escape '"', '\' and ']'

sdparam = paramname:sdname '=' '"' paramvalue '"';

sdelement = '['  sdid:sdname  (sp sdparam)*  ']';

structureddata = nilvalue | sdelement+;

bom = 0xEF 0xBB 0xBF;

msgutf8 = bom utf8string;

msgany = any*;

msg = msgany | msgutf8;

line := (any - [\n\r])* @{ fgoto main; }; # [^\n]* '\n' @{ fgoto main; }; 

main := header sp structureddata (sp msg)?;

}%%

// (todo) > segnarsi posizioni (non piu chars) e alla fine fare substring + eventuali conversioni/controlli
// (todo) > escludere il nilvalue manualmente