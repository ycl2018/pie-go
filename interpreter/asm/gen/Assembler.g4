/** A generic bytecode assembler whose instructions take 0..3 operands.
 *  Instruction set is dictated externally with a String[].  Implement
 *  specifics by subclassing and defining gen() methods. Comments start
 *  with ';' and all instructions end with '\n'.  Handles both register (rN)
 *  and stack-based assembly instructions.  Labels are "ID:".  "main:" label
 *  is where we start execution.  Use .globals and .def for global data
 *  and function definitions, respectively.
 */
grammar Assembler;

@parser::header {
type AsmGenerator interface {
    Gen(instrs ...antlr.Token)
    CheckForUnresolvedReferences()
    DefineFunction(id antlr.Token, nargs, nlocals int)
    DefineDataSize(n int)
    DefineLabel(id antlr.Token)
}
}

// START: members
@parser::structmembers {AsmGenerator}
// END: members

program
    :   globals?
        ( functionDeclaration | instr | label | NEWLINE )+
        {p.CheckForUnresolvedReferences()}
    ;
   
// how much data space
// START: data
globals : NEWLINE* '.globals' INT NEWLINE {p.DefineDataSize($INT.int)} ;
// END: data

//  .def fact: args=1, locals=0
// START: func
functionDeclaration
    : '.def' name=ID ':' 'args' '=' a=INT ',' 'locals' '=' n=INT NEWLINE
      {p.DefineFunction($name, $a.int, $n.int)}
    ;
// END: func

// START: instr
instr
    :   ID NEWLINE                         {p.Gen($ID)}
    |   ID operand NEWLINE                 {p.Gen($ID,$operand.start)}
    |   ID a=operand ',' b=operand NEWLINE {p.Gen($ID,$a.start,$b.start)}
    |   ID a=operand ',' b=operand ',' c=operand NEWLINE
        {p.Gen($ID,$a.start,$b.start,$c.start)}
    ;
// END: instr

// START: operand
operand
    :   ID   // basic code label; E.g., "loop"
    |   REG  // register name; E.g., "r0"
    |   FUNC // function label; E.g., "f()"
    |   INT
// ...
// END: operand
    |   CHAR
    |   STRING
    |   FLOAT
    ;

label
    :   ID ':' {p.DefineLabel($ID)}
    ;

REG :   'r' INT ;

ID  :   LETTER (LETTER | '0'..'9')* ;

FUNC:   ID '()' ;

fragment
LETTER
    :   ('a'..'z' | 'A'..'Z')
    ;
    
INT :   '-'? '0'..'9'+ ;

CHAR:   '\'' . '\'' ;

STRING: '"' STR_CHARS '"' ;

fragment STR_CHARS : ~'"'* ;

FLOAT
    :   INT '.' INT*
    |   '.' INT+
    ;

WS  :   (' '|'\t')+ -> skip ;

NEWLINE
    :   (';' (.)*?)? '\r'? '\n'  // optional comment followed by newline
    ;