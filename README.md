# ws_wordwrap - MySQL UDF for wordwrap.

### about

This is MySQL User Defined Function written by cgo.  
Word wrap from string.  
Like php `wordwrap` function.  

(php wordwrap)  
https://www.php.net/manual/en/function.wordwrap.php

|arg|explain|
|---|-------|
|arg1|`input` input string|
|arg2|`limit` limit length|
|arg3|`split` split string|

### how to install

    $ ./build.sh

(notice)  

* require root privilege

### example

(simple1)  

    MariaDB [(none)]> select ws_wordwrap('aaabbbccc', 3, '<br>');
    +-------------------------------------+
    | ws_wordwrap('aaabbbccc', 3, '<br>') |
    +-------------------------------------+
    | aaa<br>bbb<br>ccc                   |
    +-------------------------------------+

(simple2)  

    MariaDB [(none)]> select ws_wordwrap('aaabbbccc', 4, '\n');
    +-----------------------------------+
    | ws_wordwrap('aaabbbccc', 4, '\n') |
    +-----------------------------------+
    | aaab
    bbcc
    c                       |
    +-----------------------------------+


(multibyte)

    MariaDB [(none)]> select ws_wordwrap('𠮷野家で𠮷野がご飯をたべる', 2, '\n');
    +-------------------------------------------------------------+
    | ws_wordwrap('?野家で?野がご飯をたべる', 2, '\n')            |
    +-------------------------------------------------------------+
    | 𠮷野
    家で
    𠮷野
    がご
    飯を
    たべ
    る                            |
    +-------------------------------------------------------------+

(compare php)

mysql  

    MariaDB [(none)]> select ws_wordwrap('aa<br>abbbccc<br>dddee<br>e', 3, '<br>');
    +-------------------------------------------------------+
    | ws_wordwrap('aa<br>abbbccc<br>dddee<br>e', 3, '<br>') |
    +-------------------------------------------------------+
    | aa<br>abb<br>bcc<br>c<br>ddd<br>ee<br>e               |
    +-------------------------------------------------------+

php  

    php > $input = 'aa<br>abbbccc<br>dddee<br>e';
    php > echo wordwrap($input, 3, '<br>', true);
    aa<br>abb<br>bcc<br>c<br>ddd<br>ee<br>e

(safe duplicate)

    MariaDB [(none)]> select ws_wordwrap(ws_wordwrap('𠮷野家で𠮷野がご飯をたべる', 2, '\n'), 2, '\n');
    +-----------------------------------------------------------------------------------+
    | ws_wordwrap(ws_wordwrap('?野家で?野がご飯をたべる', 2, '\n'), 2, '\n')            |
    +-----------------------------------------------------------------------------------+
    | 𠮷野
    家で
    𠮷野
    がご
    飯を
    たべ
    る                                                  |
    +-----------------------------------------------------------------------------------+

