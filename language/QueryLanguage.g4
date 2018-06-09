grammar QueryLanguage;

query
    : matchClause
    ;


matchClause
    : Match WS+ node WS* ((relationship WS* node)+)? WS+ Return WS+ returnValue (WS* ',' WS* returnValue)? EOF
    ;


node
    : '(' aliasIdentity (WS* '{' WS* filter (WS* ',' WS* filter)? WS* '}')? ')'
    ;

filter
    : field WS* ':' WS* fieldValue
    ;

relationship
    : '<-[' aliasIdentity (WS* '{' WS* filter (WS* ',' WS* filter)? WS* '}')? ']-'
    | '-[' aliasIdentity (WS* '{' WS* filter (WS* ',' WS* filter)? WS* '}')? ']->'
    ;

returnValue
    : alias
    | alias '.' field
    ;

aliasIdentity
    : alias ':' nodeName
    | ':' nodeName
    ;

Match
    : ('M')('A')('T')('C')('H')
    ;

Return
    : ('R')('E')('T')('U')('R')('N')
    ;

alias
    : AnyText
    ;

field
    : AnyText
    ;

fieldValue
    : '"' WS* AnyText (WS+ AnyText)* '"'
    | NumberText
    ;

nodeName
    : AnyText
    ;


fragment NUMBER_CLASSES
    : [0-9]
    ;


fragment UNICODE_SYMBOL_CLASSES
    : [a-zA-Z0-9_]
    ;

fragment UNICODE_SYMBOL_CLASSES_WITH_SPACES
    : [a-zA-Z0-9_ ]
    ;

WS
    : [\t\r\n ]+
    ;

NumberText
    : NUMBER_CLASSES+
    ;

AnyText
    : UNICODE_SYMBOL_CLASSES+
    ;
