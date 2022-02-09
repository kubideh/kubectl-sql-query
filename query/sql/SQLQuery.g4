// SQLQuery.g4
grammar SQLQuery;

// Rules
query: selectStatement EOF ;

selectStatement: selectClause fromClause (whereClause)? ;

selectClause: SELECT columns ;

fromClause: FROM table ;

columns: ALL_COLUMNS | column (',' column)* ;

column: IDENTIFIER ;

table: IDENTIFIER ;

whereClause: WHERE predicate ;

predicate: comparison | predicate AND predicate ;

comparison: lhs EQ rhs ;

lhs: IDENTIFIER ;

rhs: STRING ;

// Tokens
STRING:         '\'' (~['])* '\'' ;
ALL_COLUMNS:    '*' ;
EQ:             '=' ;
FROM:           F R O M ;
SELECT:         S E L E C T ;
WHERE:          W H E R E ;
AND:            A N D ;
IDENTIFIER:     [a-zA-Z0-9.] [a-zA-Z0-9-.]* ;
WHITESPACE:     [ \u000B\t\r\n]+ -> skip ;

fragment A: [aA];
fragment B: [bB];
fragment C: [cC];
fragment D: [dD];
fragment E: [eE];
fragment F: [fF];
fragment G: [gG];
fragment H: [hH];
fragment I: [iI];
fragment J: [jJ];
fragment K: [kK];
fragment L: [lL];
fragment M: [mM];
fragment N: [nN];
fragment O: [oO];
fragment P: [pP];
fragment Q: [qQ];
fragment R: [rR];
fragment S: [sS];
fragment T: [tT];
fragment U: [uU];
fragment V: [vV];
fragment W: [wW];
fragment X: [xX];
fragment Y: [yY];
fragment Z: [zZ];