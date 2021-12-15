grammar Calc;

expr
    : left=expr op=('*'|'/') right=expr                                 # OpExpr
    | left=expr op=('+'|'-') right=expr                                 # OpExpr
    | identifier                                                        # IdExpr
    | literal                                                           # LitExpr
    | '(' expr ')'                                                      # ParenExpr
    ;

identifier
    : Id
    ;

literal
    : DecimalLiteral
    ;

MUL: '*';
DIV: '/';
ADD: '+';
SUB: '-';
Id  : [a-zA-Z][a-zA-Z0-9_]* ;

DecimalLiteral:                 DecimalIntegerLiteral '.' [0-9] [0-9_]*
              |                 '.' [0-9] [0-9_]*
              |                 DecimalIntegerLiteral
              ;

fragment DecimalIntegerLiteral
    : '0'
    | [1-9] [0-9_]*
    ;

WS : [ \t\r\n]+ -> skip ;    // toss out whitespace

// antlr -Dlanguage=Go -o parser -visitor Calc.g4
