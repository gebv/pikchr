%{
package	syntax

import "fmt"

%}

%union	{
  token Token
  block Block
  blocks Blocks
}

%token <token> TOKEN NEWLINE T_CODE_BLOCK WHITESPACE


%type <token> word
%type <block> wblock codeblock words
%type <blocks> doc_blocks  wblocks

%%

doc: doc_blocks {
  mdlex.(*lexer).result.AddBlock($1...)
}|;

doc_blocks:
  wblock {
    if mdDebug == 1 {
      fmt.Println(" ->", $1)
    }

    $$ = Blocks{$1}
  }|
  doc_blocks wblock {
    if mdDebug == 1 {
      fmt.Println(" ->", $2)
    }
    $$ = append($1, $2)
  }|
  codeblock {
    if mdDebug == 1 {
      fmt.Println(" ->", $1)
    }
    $$ = Blocks{$1}
  }|
  doc_blocks codeblock {
    if mdDebug == 1 {
      fmt.Println(" ->", $2)
    }
    $$ = append($1, $2)
  }
;

wblocks: wblock {
  $$ = Blocks{$1}
} | wblocks wblock {
  $1.AddBlock($2)
  $$ = $1
};

wblock:
  NEWLINE {
    block := &LineBlock{}
    block.AddToken($1)
    $$ = block
  }|
  words NEWLINE {
    $1.(*LineBlock).AddToken($2)
    $$ = $1
  }
;


codeblock:
  T_CODE_BLOCK wblocks T_CODE_BLOCK NEWLINE{
    $$ = NewCodeBlock($1, $3, $2...)
  }
;

words:
  word {
    block := &LineBlock{}
    block.AddToken($1)
    $$ = block
  }|
  words word{
    $1.(*LineBlock).AddToken($2)
    $$ = $1
  }
;

word: TOKEN | WHITESPACE;
