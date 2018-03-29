%%{
machine rfc5424;

action acc_prival {
    privalChars = append(privalChars, fc)
}

action int_prival {
    prival = NewPrival(utf8ToNum(privalChars))
}

action acc_version {
    versionChars = append(versionChars, fc)
}

action int_version {
    version = NewVersion(utf8ToNum(versionChars))
}

action get_fulldate_year {
    fmt.Printf("year: %#v\n", fc)
}

action get_fulldate_month {
    fmt.Printf("month: %#v\n", fc)
}

action get_fulldate_mday {

}

action predicate_fulldate {

}

nilvalue = '-';

sp = ' ';

nonzerodigit = '1'..'9';

printusascii = '!'..'~';

version = (nonzerodigit . digit{0,2}) @acc_version %int_version;

datemday = digit{0,2};

datemonth = digit{0,2};

datefullyear = digit{4};

fulldate = datefullyear @get_fulldate_year '-' datemonth @get_fulldate_month '-' datemday @get_fulldate_mday %predicate_fulldate;

# fulltime = ;

timestamp = nilvalue | fulldate ; #. 'T' . fulltime;

hostname = nilvalue | printusascii{1,255};

# 1..191 or 0
prival = (('1' ( '9' ( '0'..'1' ){,1} | '0'..'8' ( '0'..'9' ){,1} ){,1}) | ( '2'..'9' ('0'..'9'){,1} )) | '0';

pri = '<' . prival @acc_prival %int_prival . '>';

header = pri . version . sp;

}%%
