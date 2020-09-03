# POC UI Onboarding

+ Crie aplicativos de Cross-Platform para desktop com HTML, CSS e script!

+ Criar telas de desktop é fácil utilizando o [sciter](https://sciter.com/).

+ Veja mais no [GitHub](https://github.com/sciter-sdk/go-sciter).

+ Talvez esses [vídeos](https://www.youtube.com/playlist?list=PLub5C2vM5SjKvkbFfposhyg1V2gpXnviM) ajudem você a entender um pouco sobre a biblioteca.

+ Alguns [Tutoriais](https://sciter.com/tutorials/).

+ Entenda sobre a [Arquitetura do Sciter](https://sciter.com/developers/engine-architecture/).

+ [Documentação](https://sciter.com/developers/sciter-docs/).

+ Sobre [Licenciamento e preços](https://sciter.com/prices/).

+ Assistente Sciter para [Visual Studio Code](https://sciter.com/sciter-assistant-for-visual-studio-code/).

+ [Ferramentas de desenvolvimento](https://sciter.com/developers/development-tools/).

## Minhas impressões: 

+ Facilidade de uso.
+ Com poucos KB, fazemos muito, ao contrário do Electron. Que também produz aplicativos de plataforma cruzada para desktop. 
+ Não encontrei a documentação e itens deste de forma fácil, porém no [site](https://sciter.com/) você pode pesquisar problemas de forma eficiente. 
+ Embedable, pode ser usado com qualquer linguagem de programação. 
+ Não é compatível com todos os recursos do CSS3 😥. 
+ Não usamos JavaScript, mas TIScript é uma versão estendida do ECMAScript. O que não é necessariamente uma coisa ruim, mas no início produz um certo desconforto, veja mais [aqui](https://sciter.com/developers/for-web-programmers/tiscript-vs-javascript/)

## Como testar local: 

+ Configurar o Sciter de acordo com o SO

  - No momento, apenas Go 1.10 ou superior é compatível.
  - Baixe o [sciter-sdk](https://sciter.com/download/).
  - Extraia a biblioteca de tempo de execução do sciter de sciter-sdk para o PATH do sistema.
  - As bibliotecas de tempo de execução vivem bin bin.lnx bin.osx com sufixo como dll sooudylib.

  - Windows: basta copiar `bin\64\sciter.dll` para `c:\windows\system32`
  - Linux:
  
    ```shell
        user@zup:~$ cd sciter-sdk/bin.lnx/x64
        user@zup:~$ export LIBRARY_PATH=$PWD
        user@zup:~$ echo $PWD >> libsciter.conf
        user@zup:~$ sudo cp libsciter.conf /etc/ld.so.conf.d/
        user@zup:~$ sudo ldconfig
        user@zup:~$ ldconfig -p | grep sciter
    ```

  - OSX:
  
    ```shell
        cd sciter-sdk/bin.osx/
        export DYLD_LIBRARY_PATH=$PWD
    ```

  - Configure o ambiente GCC para CGO.
  - [mingw64-gcc](https://sourceforge.net/projects/mingw-w64/) (5.2.0 e 7.2.0 são testados) é recomendado para usuários do Windows.
  - No Linux, gcc (4.8 ou superior) e gtk + -3.0 são necessários.
  - go get -x github.com/sciter-sdk/go-sciter

Execute o exemplo e divirta-se :)

## Veja o funcionamento atual da POC:

![UI Onboarding](./img/record-v3.gif)
