package syntax

import (
        "fmt"
)

%%{
    machine md_parser;
    write data;
    access lex.;
    variable p lex.p;
    variable pe lex.pe;

    whitespace = [ \t]+;
    # newline = [\n\r]+ | whitespace [\n\r]+;
    newline = [\n\r];
    # textword = [a-zA-Z0-9]+;
    code_block_token = [~]{3}|[`]{3};
    textword = [^ \t\n\r]+ -- code_block_token;

}%%

type lexer struct {
    // It must be an array containting the data to process.
    data []byte

    // Data end pointer. This should be initialized to p plus the data length on every run of the machine. In Java and Ruby code this should be initialized to the data length.
    pe int

    // Data pointer. In C/D code this variable is expected to be a pointer to the character data to process. It should be initialized to the beginning of the data block on every run of the machine. In Java and Ruby it is used as an offset to data and must be an integer. In this case it should be initialized to zero on every run of the machine.
    p int


    // This must be a pointer to character data. In Java and Ruby code this must be an integer. See Section 6.3 for more information.
    ts int

    // Also a pointer to character data.
    te int


    // This must be an integer value. It is a variable sometimes used by scanner code to keep track of the most recent successful pattern match.
    act int

    // Current state. This must be an integer and it should persist across invocations of the machine when the data is broken into blocks that are processed independently. This variable may be modified from outside the execution loop, but not from within.
    cs int

    // This must be an integer value and will be used as an offset to stack, giving the next available spot on the top of the stack.
    top int

    result Blocks
}

func newLexer(data []byte) *lexer {
    lex := &lexer{
        data: data,
        pe: len(data),
        result: Blocks{},
    }
    %% write init;
    return lex
}

func (lex *lexer) Lex(out *mdSymType) int {
    eof := lex.pe
    tok := 0
    %%{
        main := |*
            textword => {
                if mdDebug == 1 {
                    fmt.Printf("%q", string(lex.data[lex.ts:lex.te]))
                }
                out.token = wordToken{rawToken(string(lex.data[lex.ts:lex.te]))}
                tok = TOKEN;
                fbreak;
            };
            whitespace => {
                if mdDebug == 1 {
                    fmt.Print("+")
                }
                out.token = whitespaceToken{rawToken(string(lex.data[lex.ts:lex.te]))}
                tok = WHITESPACE;
                fbreak;
            };
            code_block_token => {
                if mdDebug == 1 {
                    fmt.Print("T_CODE_BLOCK")
                }
                out.token = fencedCodeblockToken{rawToken(string(lex.data[lex.ts:lex.te]))}
                tok = T_CODE_BLOCK;
                fbreak;
            };
            newline => {
                if mdDebug == 1 {
                    fmt.Print("EOL")
                }

                out.token = eolToken{rawToken(string(lex.data[lex.ts:lex.te]))}
                tok = NEWLINE;
                fbreak;
            };
        *|;

         write exec;
    }%%

    return tok;
}

func (lex *lexer) Error(e string) {
    fmt.Println("lexer error:", e)
}
