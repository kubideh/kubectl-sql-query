// SQLQuery.g4
grammar SQLQuery;

// Rules
query: statement EOF ;

statement: SELECT fieldList FROM tableName (whereClause)? ;

fieldList: field (',' field)* ;

field: ALL | IDENTIFIER ;

tableName: IDENTIFIER ;

whereClause: WHERE expr ;

expr: evaluation | expr AND expr ;

evaluation: key EQ value ;

key: IDENTIFIER ;

value: IDENTIFIER ;

// Tokens
ALL: '*' ;
EQ: '=' ;
FROM: 'FROM' ;
SELECT : 'SELECT' ;
WHERE: 'WHERE' ;
AND: 'AND' ;
IDENTIFIER: [a-zA-Z0-9] [a-zA-Z0-9-.]* ;
WHITESPACE : [ \u000B\t\r\n]+ -> skip ;
