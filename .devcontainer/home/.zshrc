export ZSH=$HOME/.oh-my-zsh

ZSH_THEME="powerlevel10k/powerlevel10k"

plugins=(git fzf zsh-autosuggestions)

source $ZSH/oh-my-zsh.sh

DISABLE_AUTO_UPDATE=true
DISABLE_UPDATE_PROMPT=true

[[ ! -f ~/.p10k.zsh ]] || source ~/.p10k.zsh;

hash -d w=/workspaces;
export EDITOR='code';

export PATH=$PATH:/usr/local/go/bin;