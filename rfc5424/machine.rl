%%{
machine rfc5424;

action add_char {
    cr.Add(fc)
}

action set_prival {
    prival = NewPrival(*cr.Reduce())
}

action set_version {
    version = NewVersion(*cr.Reduce())
}

action tag_fulldate {
    poss["fulldate:ini"] = fpc
}

action set_fulldate {
    if _, e := time.Parse("2006-01-02", data[poss["fulldate:ini"]:p]); e != nil {
        err = fmt.Errorf("error %s [col %d:%d]", e, poss["fulldate:ini"], p);
        fhold; fgoto line;
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

version = (nonzerodigit . digit{0,2}) @add_char %set_version;

datemday = ('0' . '1'..'9' | '1'..'2' . '0'..'9' | '3' . '0'..'1') $add_char;

datemonth = ('0' . '1'..'9' | '1' . '0'..'2') $add_char;

datefullyear = digit{4} $add_char;

fulldate = datefullyear >tag_fulldate '-' datemonth  '-' datemday %set_fulldate;

timehour = ('0'..'1' . '0'..'9' | '2' . '0'..'3');

timeminute = sexagesimal;

timesecond = sexagesimal;

timesecfrac = '.' . digit{1,6};

timenumoffset = ('+' | '-') timehour ':' timeminute;

timeoffset = 'Z' . timenumoffset;

partialtime = timehour ':'  timeminute ':' timesecond . timesecfrac?;

fulltime = partialtime . timeoffset;

timestamp = nilvalue | fulldate . 'T' . fulltime; 

hostname = nilvalue | printusascii{1,255};

# 1..191 or 0
prival = (('1' ( '9' ( '0'..'1' ){,1} | '0'..'8' ( '0'..'9' ){,1} ){,1}) | ( '2'..'9' ('0'..'9'){,1} )) | '0';

pri = '<' . prival @add_char %set_prival . '>';

header = pri . version . sp . timestamp;

line := (any - [\n\r])* @{ fgoto main; }; # [^\n]* '\n' @{ fgoto main; }; 

main := header;

}%%
