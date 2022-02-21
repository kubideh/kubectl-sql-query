/*
 * The MIT License (MIT)
 *
 * Copyright (c) 2020 by Martin Mirchev
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of this software and
 * associated documentation files (the "Software"), to deal in the Software without restriction,
 * including without limitation the rights to use, copy, modify, merge, publish, distribute,
 * sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all copies or
 * substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT
 * NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
 * NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
 * DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 *
 * Project : sqlite-parser; an ANTLR4 grammar for SQLite https://github.com/bkiers/sqlite-parser
 * Developed by : Bart Kiers, bart@big-o.nl
 */

// $antlr-format alignTrailingComments on, columnLimit 150, maxEmptyLinesToKeep 1, reflowComments off, useTab off
// $antlr-format allowShortRulesOnASingleLine on, alignSemicolons ownLine

lexer grammar SQLiteLexer;

SCOL:      ';';
DOT:       '.';
OPEN_PAR:  '(';
CLOSE_PAR: ')';
COMMA:     ',';
ASSIGN:    '=';
STAR:      '*';
PLUS:      '+';
MINUS:     '-';
TILDE:     '~';
PIPE2:     '||';
DIV:       '/';
MOD:       '%';
LT2:       '<<';
GT2:       '>>';
AMP:       '&';
PIPE:      '|';
LT:        '<';
LT_EQ:     '<=';
GT:        '>';
GT_EQ:     '>=';
EQ:        '==';
NOT_EQ1:   '!=';
NOT_EQ2:   '<>';

// http://www.sqlite.org/lang_keywords.html
ABORT_:             A B O R T;
AFTER_:             A F T E R;
ALL_:               A L L;
AND_:               A N D;
AS_:                A S;
ASC_:               A S C;
BETWEEN_:           B E T W E E N;
BY_:                B Y;
CASE_:              C A S E;
CAST_:              C A S T;
COLLATE_:           C O L L A T E;
COLUMN_:            C O L U M N;
CROSS_:             C R O S S;
CURRENT_DATE_:      C U R R E N T '_' D A T E;
CURRENT_TIME_:      C U R R E N T '_' T I M E;
CURRENT_TIMESTAMP_: C U R R E N T '_' T I M E S T A M P;
DEFAULT_:           D E F A U L T;
DESC_:              D E S C;
DETACH_:            D E T A C H;
DISTINCT_:          D I S T I N C T;
ELSE_:              E L S E;
END_:               E N D;
ESCAPE_:            E S C A P E;
EXCEPT_:            E X C E P T;
EXISTS_:            E X I S T S;
FAIL_:              F A I L;
FROM_:              F R O M;
GLOB_:              G L O B;
GROUP_:             G R O U P;
HAVING_:            H A V I N G;
IF_:                I F;
IGNORE_:            I G N O R E;
IN_:                I N;
INDEXED_:           I N D E X E D;
INNER_:             I N N E R;
INSTEAD_:           I N S T E A D;
INTERSECT_:         I N T E R S E C T;
INTO_:              I N T O;
IS_:                I S;
ISNULL_:            I S N U L L;
JOIN_:              J O I N;
KEY_:               K E Y;
LEFT_:              L E F T;
LIKE_:              L I K E;
LIMIT_:             L I M I T;
MATCH_:             M A T C H;
NATURAL_:           N A T U R A L;
NO_:                N O;
NOT_:               N O T;
NOTNULL_:           N O T N U L L;
NULL_:              N U L L;
OF_:                O F;
OFFSET_:            O F F S E T;
ON_:                O N;
OR_:                O R;
ORDER_:             O R D E R;
OUTER_:             O U T E R;
RAISE_:             R A I S E;
RECURSIVE_:         R E C U R S I V E;
REGEXP_:            R E G E X P;
RIGHT_:             R I G H T;
ROLLBACK_:          R O L L B A C K;
ROW_:               R O W;
ROWS_:              R O W S;
SELECT_:            S E L E C T;
TABLE_:             T A B L E;
TEMP_:              T E M P;
TEMPORARY_:         T E M P O R A R Y;
THEN_:              T H E N;
TO_:                T O;
UNION_:             U N I O N;
UPDATE_:            U P D A T E;
USING_:             U S I N G;
VALUES_:            V A L U E S;
WHEN_:              W H E N;
WHERE_:             W H E R E;
WITH_:              W I T H;
FIRST_VALUE_:       F I R S T '_' V A L U E;
OVER_:              O V E R;
PARTITION_:         P A R T I T I O N;
RANGE_:             R A N G E;
PRECEDING_:         P R E C E D I N G;
UNBOUNDED_:         U N B O U N D E D;
CURRENT_:           C U R R E N T;
FOLLOWING_:         F O L L O W I N G;
CUME_DIST_:         C U M E '_' D I S T;
DENSE_RANK_:        D E N S E '_' R A N K;
LAG_:               L A G;
LAST_VALUE_:        L A S T '_' V A L U E;
LEAD_:              L E A D;
NTH_VALUE_:         N T H '_' V A L U E;
NTILE_:             N T I L E;
PERCENT_RANK_:      P E R C E N T '_' R A N K;
RANK_:              R A N K;
ROW_NUMBER_:        R O W '_' N U M B E R;
TRUE_:              T R U E;
FALSE_:             F A L S E;
WINDOW_:            W I N D O W;
NULLS_:             N U L L S;
FIRST_:             F I R S T;
LAST_:              L A S T;
FILTER_:            F I L T E R;
GROUPS_:            G R O U P S;
EXCLUDE_:           E X C L U D E;
TIES_:              T I E S;
OTHERS_:            O T H E R S;
DO_:                D O;
NOTHING_:           N O T H I N G;

IDENTIFIER:
    '"' (~'"' | '""')* '"'
    | '`' (~'`' | '``')* '`'
    | '[' ~']'* ']'
    | [a-zA-Z_.] [a-zA-Z_0-9.]*
; // TODO check: needs more chars in set

NUMERIC_LITERAL: ((DIGIT+ ('.' DIGIT*)?) | ('.' DIGIT+)) (E [-+]? DIGIT+)? | '0x' HEX_DIGIT+;

BIND_PARAMETER: '?' DIGIT* | [:@$] IDENTIFIER;

STRING_LITERAL: '\'' ( ~'\'' | '\'\'')* '\'';

BLOB_LITERAL: X STRING_LITERAL;

SINGLE_LINE_COMMENT: '--' ~[\r\n]* (('\r'? '\n') | EOF) -> channel(HIDDEN);

MULTILINE_COMMENT: '/*' .*? '*/' -> channel(HIDDEN);

SPACES: [ \u000B\t\r\n] -> channel(HIDDEN);

UNEXPECTED_CHAR: .;

fragment HEX_DIGIT: [0-9a-fA-F];
fragment DIGIT:     [0-9];

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
