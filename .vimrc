filetype off
filetype plugin indent on
syntax on

set nocompatible
set number
set incsearch
set paste

"""  Insert 2 spaces instead TAB
set shiftwidth=2
set expandtab
set tabstop=2

""" Display status line always
set laststatus=2
set statusline+=%F\ %l\:%c
set statusline+=\ %p%%
set statusline+=%=
set statusline+=%#CursorColumn#
set statusline+=\ %y
set statusline+=\ [%{&fileformat}\]
set statusline+=\ %{&fileencoding?&fileencoding:&encoding}

""" Display line ending
" set list
""" Convert file to UNIX type (:set ff)
" set fileformat=unix

""" GO: gofmt 
au BufWritePost *.go !gofmt -w %
