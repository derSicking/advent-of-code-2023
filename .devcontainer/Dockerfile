FROM ubuntu:jammy-20231004

RUN apt update;
RUN yes | unminimize 2>&1;

RUN DEBIAN_FRONTEND=noninteractive TZ=Etc/UTC apt-get -y install --no-install-recommends tzdata locales ca-certificates man-db sudo ssh vim zsh git curl wget;
RUN apt-get clean -y;
RUN rm -rf /var/lib/apt/lists/*;

RUN update-ca-certificates;

RUN wget "https://dl.google.com/go/$(curl https://go.dev/dl/?mode=json | grep -o 'go.*.linux-amd64.tar.gz' | head -n 1 | tr -d '\r\n' )";
RUN tar -C /usr/local -xzf *.tar.gz;
RUN rm *.tar.gz;

RUN locale-gen en_US.UTF-8;

RUN adduser --shell /bin/zsh --disabled-password --gecos "" vscode;
RUN chown -R vscode:vscode /home/vscode;
RUN echo vscode ALL=\(root\) NOPASSWD:ALL > /etc/sudoers.d/vscode;
RUN chmod 0440 /etc/sudoers.d/vscode;

RUN runuser -l vscode -c 'sh -c "$(wget https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh -O -)"';
RUN runuser -l vscode -c 'git clone --depth 1 https://github.com/junegunn/fzf.git ~/.fzf';
RUN runuser -l vscode -c '~/.fzf/install';

RUN runuser -l vscode -c 'git clone --depth=1 https://github.com/romkatv/powerlevel10k.git ${ZSH_CUSTOM:-$HOME/.oh-my-zsh/custom}/themes/powerlevel10k';
RUN runuser -l vscode -c 'git clone https://github.com/zsh-users/zsh-autosuggestions ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-autosuggestions';
COPY home/* /home/vscode/

RUN runuser -l vscode -c '/usr/local/go/bin/go install -v golang.org/x/tools/gopls@latest';
RUN runuser -l vscode -c '/usr/local/go/bin/go install -v github.com/cweill/gotests/gotests@v1.6.0';
RUN runuser -l vscode -c '/usr/local/go/bin/go install -v github.com/fatih/gomodifytags@v1.16.0';
RUN runuser -l vscode -c '/usr/local/go/bin/go install -v github.com/josharian/impl@v1.1.0';
RUN runuser -l vscode -c '/usr/local/go/bin/go install -v github.com/haya14busa/goplay/cmd/goplay@v1.0.0';
RUN runuser -l vscode -c '/usr/local/go/bin/go install -v github.com/go-delve/delve/cmd/dlv@latest';
RUN runuser -l vscode -c '/usr/local/go/bin/go install -v honnef.co/go/tools/cmd/staticcheck@latest';