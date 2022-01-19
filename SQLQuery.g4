// SQLQuery.g4
grammar SQLQuery;

// Rules
query: statement EOF ;

statement: SELECT '*' FROM tableName (whereClause)? ;

tableName: IDENTIFIER ;

whereClause: WHERE expr ;

expr: evaluation | expr AND expr ;

evaluation: field EQ value ;

field: IDENTIFIER ;

value: IDENTIFIER ;

// Tokens
EQ: '=' ;
FROM: 'FROM' ;
SELECT : 'SELECT' ;
WHERE: 'WHERE' ;
AND: 'AND' ;
IDENTIFIER: [a-zA-Z0-9] [a-zA-Z0-9-.]* ;
WHITESPACE : [ \u000B\t\r\n]+ -> skip ;
