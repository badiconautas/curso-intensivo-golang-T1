20 de mai. de 2025

## Intensivão Golang\! \- Transcrição

### 00:00:00

   
**lucas Badico:** Beleza? Eh, bom, gente, o que que eu fiz aqui, tá? Eu coloquei o sumário, o sumário da aula um que o nosso amigo Google Gemini fez dentro do projeto, dentro do repositório. Coloquei a transcrição inteira também, tá? Eh, deixa eu puxar aqui. Coloquei, né, paraa primeira aula não vai ser tão importante, né, porque, pô, a gente não teve nada que vai precisar a gente revisar, mas a gente tem, vai ter aqui toda, um no OneDrive todos os arquivos. Então tem o somário com a transcrição, tem o original que o Mit gravou, vai ter o que eu gravei, inclusive eu esqueci de puxar o OBS. Eh, só um minuto. Vou puxando o OBS aqui. Você completar eh deixa eu mudar vocês o áudio de vocês pro mute. Fala aí para ver se eu tô ouvindo.  
**adolfo agnelli:** Opa. Hallå.  
**lucas Badico:** Tô ouvindo. Tô ouvindo.  
**Luiz Felipe Verissimo da silva:** Oi.  
**adolfo agnelli:** Oi,  
**lucas Badico:** Cadê? Alô, alô, alô. Falem aí para ver se o OBS tá ruim.  
   
 

### 00:01:35

   
**adolfo agnelli:** alô, alô, alô.  
**Matheus Abranches:** Oi,  
**Luiz Felipe Verissimo da silva:** Tá ouvindo?  
**lucas Badico:** Beleza,  
**Matheus Abranches:** oi.  
**Luiz Felipe Verissimo da silva:** Tá ouvindo?  
**lucas Badico:** tá perfeito. Start record aqui. Eh, então aqui no OneDrive vai ter vai ter a data, o média qualidade do ABS, o original do Meet e vai ter uma versão editada também que a gente vai depois colocar no YouTube para facilitar, tá bom? Eh, então vai ter bastante recurso aí histórico para vocês quiserem revisar e tudo mais, né? para as próximas aulas talvez seja importante, mas o sumário e a coisa vai ficar no próprio projeto, né? Eh, e aqui a gente vai manter um histórico do que a gente tá construindo, tá bom?  
**Luiz Felipe Verissimo da silva:** Você  
**lucas Badico:** Eh, beleza. O que que a gente vai falar hoje, cara? Hoje a gente vai entrar na parte de APIs, né? Cadê? Cadê aí? Cadê? Cadê? Cadê? Cadê? Só um minuto. Perdi meu abs. É isso. Ele fica na frente de tudo que eu tenho aqui.  
   
 

### 00:02:38

   
**lucas Badico:** Um minuto.  
**Luiz Felipe Verissimo da silva:** colocou o link do OneDrive em algum lugar?  
**lucas Badico:** Ah, eu vou mandar para vocês quem precisa compartilhar também. mandou certinho para vocês. Só um minutinho, gente, que perdi o meu aqui. Perfeito. Qu, OK, tô com tô com resuminho aqui. Hoje a gente vai começar a falar de duas coisas muito importantes, que é a própria API, né? Que que a gente vai tá fazendo nesse pilar? E a gente vai falar também sobre banco de dados. Deixa eu começar esse papo de banco de dados, beleza? Eh, salvar isso aqui. Eh, bom, como é que vocês definiriam para mim o que que é um servidor, né? O que que é uma um o que que é uma aplicação que Deixa eu refazer a pergunta. Como é que vocês definiam uma um back end? Qual que é a definição para vocês de back end? Угу.  
**Luiz Felipe Verissimo da silva:** Eu  
**adolfo agnelli:** que o back end aí tem aquela aquela clássica,  
**Luiz Felipe Verissimo da silva:** sei  
**adolfo agnelli:** né, que a galera usa, né, que é onde tudo acontece basicamente, né, aquele exemplo da do restaurante, né, que é o o API, né, ou seria o garçom e o frontend seria o cardápio e o back end é a cozinha, né?  
   
 

### 00:04:18

   
**adolfo agnelli:** Onde tem tudo, tem um estoque, cozinheiro,  
**lucas Badico:** Угу.  
**adolfo agnelli:** processo, tudo acontece lá.  
**lucas Badico:** Tá. Eh, tem uma tem uma parada que eu descobri essa semana. Minha esposa tava me contando  
**adolfo agnelli:** Não  
**lucas Badico:** no restaurante chamam realmente o de back front, sabia?  
**adolfo agnelli:** sabia não.  
**lucas Badico:** Eu não sabia. Eu falei sério, eu lembrei dessa analogia, mas é bem por aí mesmo. O backend é onde tudo acontece. Mas eu tô buscando aqui a resposta sobre o que que é o back end, porque essa é a abstração. Mas o que que é a versão concreta?  
**Luiz Felipe Verissimo da silva:** servidor aonde vai focar as regras de negócio, onde vai ter a conexão com o banco de dados e que vai executar a persistência dos dados.  
**lucas Badico:** Perfeito. Servidor, né? Então, a gente tá falando do computador de outra pessoa, certo?  
**adolfo agnelli:** Isso.  
**Luiz Felipe Verissimo da silva:** Não.  
**lucas Badico:** Então, no fim das contas, todo backend é só o computador de outra pessoa. Quando a gente fala de cloud, a gente não vai entrar em cloud num primeiro pilar, mas o cloud é simplesmente uma série de servidores que a gente não sabe onde tá, né?  
   
 

### 00:05:29

   
**lucas Badico:** Ele pertence a alguém que não é a gente, tá? Mas ainda é o servidor, é o computador de outra pessoa, certo?  
**adolfo agnelli:** Mhm.  
**lucas Badico:** E aí entra uma questão. Back end pode ser feito em qualquer linguagem, em qualquer tipo de aplicação, só precisa de um recurso. Que recurso que é?  
**Matheus Abranches:** B  
**adolfo agnelli:** É de dados  
**Luiz Felipe Verissimo da silva:** AP aí para se comunicar com o outro com outra parte, né? Sei  
**lucas Badico:** A  
**Luiz Felipe Verissimo da silva:** lá,  
**lucas Badico:** palavra  
**Luiz Felipe Verissimo da silva:** um serviço.  
**lucas Badico:** não é só P, é o socket, tá?  
**adolfo agnelli:** Socet  
**Luiz Felipe Verissimo da silva:** Hum.  
**lucas Badico:** Vocês entendem o que que é um socket?  
**Luiz Felipe Verissimo da silva:** uma forma de conexão.  
**lucas Badico:** É exato. Eh, eu tento, eu tento imaginar o socket, né? A gente, eu não sei muito os detalhes como é que é internamente, Eu tá? eu nessa parte me me falha, mas eu tenho que imaginar o socket como a terceira parte, né, a terceira perna de da interface, porque pera aí que a gente tem a interface aqui de clicar e visualizar, certo? Essa é um é um tipo de interface.  
   
 

### 00:06:44

   
**lucas Badico:** Aí vocês tão vendo minha tela.  
**Matheus Abranches:** Não,  
**lucas Badico:** Desculpa,  
**adolfo agnelli:** No.  
**lucas Badico:** você tá vendo a tela como tem? Mas tudo bem.  
**Matheus Abranches:** Adolfo não  
**lucas Badico:** Eu  
**Matheus Abranches:** quis  
**lucas Badico:** tava  
**Matheus Abranches:** te desculpar  
**lucas Badico:** clicando  
**Matheus Abranches:** não.  
**lucas Badico:** como se vocês tivessem vendo sempre. Ai, Deus.  
**Matheus Abranches:** Hum.  
**lucas Badico:** OK.  
**Matheus Abranches:** Угуm.  
**lucas Badico:** Eh, essa é um tipo de interface, certo? Você clica, há um uma reação e você visualiza o que você tá fazendo, certo? Ou seja, é um tipo de IO, correto?  
**Luiz Felipe Verissimo da silva:** É interface. Quando você usa essa palavra nesse contexto, é uma interface o I e, o X, né? E uma  
**lucas Badico:** Isso.  
**Luiz Felipe Verissimo da silva:** forma  
**lucas Badico:** Interface  
**Luiz Felipe Verissimo da silva:** de facilitar o usuário.  
**lucas Badico:** é uma forma de interagir, vamos chamar de uma forma de interagir, entendeu? Interface ti de uma forma de interagir, da gente interagir com a máquina, certo? O  
**Luiz Felipe Verissimo da silva:** Uhum.  
   
 

### 00:07:37

   
**lucas Badico:** terminal é outra forma de interface, certo? Eu posso vir aqui fazer git log e ele me dá um resultado. Ele vai dizer: "Olha, você não tá num você tá no lugar errado, eh, CD aqui é melhor, né?  
**Matheus Abranches:** Mhm.  
**lucas Badico:** CD  
**Matheus Abranches:** M.  
**lucas Badico:** CD. Como é que tá o nosso projeto aqui? intensiv sei.  
**Matheus Abranches:** Угуm.  
**lucas Badico:** Ele vai dizer aqui, ó, isso é uma interface, certo? Isso é uma interface. Isso é um IO. Então, isso aqui é como se fosse um IO, só que isso aqui é mais complexo, tá? Eh, fronts são sempre mais complexas, mas isso aqui é um IO input, output, certo? Um arquivo também é um IO, você lê um arquivo, né, tipo, é um IO, você tá lendo um input daquele arquivo e se você quiser, você pode escrever naquele arquivo, dar o output. Socket é uma forma de IO também. Então, pensar nas coisas mais bunda que pode que a gente pode usar para para processar o dado, como o Adolf falou, eh, pensa o Axis da Microsoft. Se a gente quisesse usar o Ax, o próprio Excel, daria para fazer um API, daria para fazer um servidor com essas duas tecnologias.  
   
 

### 00:09:15

   
**lucas Badico:** A o único requisito que aí eu não sei dizer se eles têm, eu acho que tem, é que eles precisavam disponibilizar a conexão, né, o acesso via socket, tá bom? Eh, e o que que é o socket? É como se fosse uma pastinha, tipo, imagina assim, né? Vamos abstr, você cria uma pasta e você começa a ouvir naquela pasta, tá? E aí você coloca um arquivo naquela pasta. seu sistema processa aquele arquivo e entrega o arquivo em outra pasta, entendeu? Claro que não são pastas, é tudo via network, tá bom? Mas é como se fosse, porque no fim das contas é io, tá bom? É input, output. Tô falando tudo isso só para vocês entenderem que eu vou mostrar aqui eh tudo em GO, só que o que a gente vai fazer, você consegue repetir em qualquer linguagem, tá? Principalmente quando você aprende go. Por quê? Porque o legal do go é que ele não esconde que é um io, que é um input output, entendeu? Ele não esconde. Quando você vai fazer as coisas em outras linguagens, é tudo meio que abstraído. Lembra as níveis de abstração? As outras linguagens tendem a abstrair.  
   
 

### 00:10:39

   
**lucas Badico:** Então, JavaScript adora abstrair. Node adora abstrair as paradas. Não sei se vocês já brincaram bastante com Node, mas Node adora abstrair. Você quase não vê a camada interna assim do do seu do seu web server. É só abstração, abstração em cima de abstração. Por isso, eh, e até quem quem tá mais tempo aí eu, viu já falar sobre que nasce um framework de JavaScript a cada a cada minuto praticamente, é porque estamos construindo abstração em cima de abstração. Goleng não gosta disso, tá? Goleng gosta de ser e a gente vai ver isso, tá bom? Eh, mas vamos lá. Então, a gente tá falando de APIs que rodam via network, né? Que vão socket network. Deixa eu puxar aqui. Acho que um dos softwares mais mais claros pra gente enxergar, opa, é o poxo, mano. Que estranho.  
**Luiz Felipe Verissimo da silva:** Ó, lá embaixo and fix.  
**lucas Badico:** Não, eu tô tô no browser aqui. Ele levou pro browser, mas não tá fazer o sign com com Gub, vou vou vou ficar na versão versão deslogada mesmo. É aquela parada, não faz sentido ele exigeir que a gente logue, né, para ser tudo local, mas tudo bem.  
   
 

### 00:12:51

   
**lucas Badico:** Eh, eu não sei, acho que isso aqui passar é o mais novo, né? a pessoa que tem menos experiência com papéis da do coiso. Então vamos fazer um resuminho básico. O legal do PM é que ele deixa tudo meio mastigado pra gente. O que tá aqui no postma nessa tela, ignora aqui o histórico, né, que é tudo que a gente fez. O que tá nessa tela é tudo que a gente precisa para fazer um request http, né, que é onde as nossas APs vão rodar. Você tem o verbo get, post, put, pet, delete, head, options. Normalmente a gente vai viver nesses quatro primeiros, né? E os verbos eles eles têm uma função, uma função, como é que eu posso falar essa palavra? Mecânica, mas eles também tm uma função semântica. O que eu quero dizer com isso? O get é para listar. Deixa eu pegar aqui. Eu tenho esse cara aqui. Ah, eu não tenho, não tô com um toquinho, mas eu fiz aqui um get para poder pegar o usuário do do click, certo? Fiz um get pegar o usuário duplicar, ou seja, eu tô listando.  
   
 

### 00:14:20

   
**lucas Badico:** Se eu quisesse deletar esse usuário, muito provavelmente o verbo seria delete. Então, verbo, ele tem uma função mecânica, certo? E uma função semântica. Isso é importante vocês lembrarem, porque quando a gente for modelar a PI, a gente vai estar pensando nisso, qual que é a semântica disso, certo? E aí para falar rapidinho, get é listagem, post ação. Depois, quando a gente for fazer as APIs mais avançadas vai ser interessante. Put, é substituição e pet alteração. Esses dois são os que mais se confundem, tá? E o delete é remoção, claro. Tá bom? Ah, então você tem o verbo, você tem a URL, certo? E na URL você tem os parâmetros, então você pode ter parâmetro A, aqui são queries, tá vendo? E você pode ter também o que a gente chama de Pfys, que seriam parâmetros aqui no aqui na rota. Você poderia ter colocar um ID aqui, cadê ID, por exemplo, eu tenho APS que exigem ids para você ver buscar aquele ID específico, né? Aquele item específico. Mas a gente vai ver tudo isso na hora que for montar as APS.  
   
 

### 00:15:49

   
**lucas Badico:** Mas só para um uma um resuminho aí do que que é, né? do que que é que a gente vai fazer. Além de verbo, PEF, parâmetros, né? Você tem os headers. Os headers são metainformações que você coloca no seu request e que você recebe no seu response, né? Então, normalmente a gente coloca autorização e autenticação no header. Eh, a gente coloca dados de cooks no header também, tá? E aí vem o body. Bor é normalmente um corpo ou dado que você tá mandando pra sua PI. Então, se a nossa PI for criar alguma entidade, vai ser no B que a gente vai mandar. E você pode mandar o B de várias formas, né? Você pode mandar com binário, você pode mandar com HTML, com XML, mas normalmente a gente usa Jon. Cadê a opção de Jon? Aqui, aqui a gente usa Jon, que é essa esse jeito de formatar o W, tá bom? Se vocês quiserem ver mais isso, o meu curso lá de Golengo e de APIs, na verdade é o curso de postman, né, Mateus? Eu falo bastante  
**Matheus Abranches:** É  
**lucas Badico:** disso.  
**Matheus Abranches:** sim,  
   
 

### 00:17:13

   
**lucas Badico:** É no  
**Matheus Abranches:** tem um  
**lucas Badico:** é  
**Matheus Abranches:** PMAN  
**lucas Badico:** no  
**Matheus Abranches:** que que é bem completinho. M.  
**lucas Badico:** isso. Eu falo bastante disso. Depois a gente pode deixar o link também do curso de post. Eu explico, eu caio assim, eu gasto quatro aulas para explicar tudo que dá para fazer com PMA da E entre eles eu falo de modelagem, tá bom? Eh, fechado. Passamos por isso. Vamos criar nosso servidorzinho aqui em Go, né? Vamos fazer nossa primeira nosso primeiro projetinho em Go. Deixa eu deixa eu criar aqui o a sessão de colaboração. Copiou. Deixa eu mandar aqui pro moço. Agora que entendi o que você falou, Adolfo. Você chamou o Felipe de Luía. Eu falei: "Quem? Quem? Luiz, com quem o Adolfo tá falando?" Não existe Luiz aqui. É Felipe. Ai, mas tudo  
**Luiz Felipe Verissimo da silva:** errado.  
**lucas Badico:** bem. Não é culpa, é culpa do mit, porque fica ali Luiz, aí você lê a primeira primeiro nome.  
   
 

### 00:18:30

   
**adolfo agnelli:** Sim.  
**lucas Badico:** Vamos lá, gente.  
**adolfo agnelli:** Yeah.  
**lucas Badico:** Eh, eu vou colocar aqui Hello Word, certo? Vamos lá, resumindo também, né, voltando um pouquinho pro básico de GO, toda a aplicação Go começa com Pack M, tá bom? Então, toda aplicação Go vai ter essa palavrinha packagem no seu primeiro arquivo, tá bom? Felipe tá chegando aí.  
**Luiz Felipe Verissimo da silva:** Eu  
**lucas Badico:** O Mateus já tá aqui. Mateus nem de uma permissão para você. Por será você usou o mesmo?  
**Matheus Abranches:** É,  
**lucas Badico:** Ah.  
**Matheus Abranches:** você tô no  
**lucas Badico:** Ah.  
**Matheus Abranches:** mesmo projeto, a mesma coisa.  
**lucas Badico:** O Felipe entrou  
**Luiz Felipe Verissimo da silva:** tô  
**lucas Badico:** como  
**Luiz Felipe Verissimo da silva:** anônimo.  
**lucas Badico:** anônimo. É por isso. Eh, então a gente tá com a gente precisa do do package main. Outra coisa que toda aplicação go vai ter é a função GO, a função M, que é bem simples também, certo? A função main, ela é uma função que não retorna nada, né? Depois a gente vai falar sobre anatomia de função, mas quando você tem uma função assim, ela não recebe nenhum parâmetro e ela não retorna nada.  
   
 

### 00:20:04

   
**lucas Badico:** Essa é a função que vai rodar a nossa aplicação. Paraa gente ter um hello word de verdade, o que a gente vai precisar é o log print ln. O meu o meu VSG tá trapaceando. Eu coloco o pacote automaticamente em portas poder liber né? Então a gente vai rodar aqui um hollow. Eu quero testar uma parada. Será que vocês conseguem?  
**Matheus Abranches:** digitar, escrever.  
**lucas Badico:** Não, eu queria que vocês rodassem, algum de vocês rodasse a aplicação. Tenta usar  
**Matheus Abranches:** Eh,  
**lucas Badico:** o terminal. Eu permiti como como write  
**Matheus Abranches:** eh, como é que roda o terminal aqui no caso dele do do to go?  
**lucas Badico:** eh, go run. E aí você coloca o nome do arquivo. Tem que ver se  
**Matheus Abranches:** Ah,  
**lucas Badico:** você tiver na raiz, você vai ter que colocar eh pilar um. Entra, entra na pasta que acho que é melhor aí. Boa.  
**Matheus Abranches:** é,  
**lucas Badico:** X.  
**adolfo agnelli:** ali para ver se vai  
**Matheus Abranches:** ele é o meu também pediu solicitação.  
**lucas Badico:** Tô, tô dando  
**adolfo agnelli:** tá  
**lucas Badico:** aí.  
   
 

### 00:21:24

   
**adolfo agnelli:** indo, tá indo.  
**Matheus Abranches:** Eu foi.  
**lucas Badico:** Quem que  
**adolfo agnelli:** É,  
**lucas Badico:** tá escrevendo aí? O Adolf  
**adolfo agnelli:** é, eu ten que pedir, tem que solicitar e quando ele liberar dá para escrever.  
**Matheus Abranches:** Так.  
**lucas Badico:** manda aí, Adolfo. Entra, entra no pilar um, né? Faz um cinho pro Pilar um. Não conseguiu entrar. Ah, a Sara já entrou também, né? Sarinha entrou aqui.  
**sasa cruz:** É, tá demorando um pouco para carregar aqui para mim, mas eu entrei.  
**lucas Badico:** Ah, tá.  
**adolfo agnelli:** É, é um pouquinho mais devagar, lerdinho,  
**lucas Badico:** Vai lá,  
**Matheus Abranches:** Aí  
**lucas Badico:** Dolf.  
**Matheus Abranches:** CD pilar um.  
**lucas Badico:** Isso. Se você der CDP e ele você dá um tab, ele vai completar,  
**Matheus Abranches:** Eu  
**lucas Badico:** eu acho. Alguém tá escrevendo que eu já não tô vendo mais.  
**Luiz Felipe Verissimo da silva:** Não sei. Eu também não tá vendo não.  
**Matheus Abranches:** vou escrever aqui, ó. Ó,  
**adolfo agnelli:** C  
**lucas Badico:** О.  
   
 

### 00:22:21

   
**Matheus Abranches:** no eh, CD barra  
**Luiz Felipe Verissimo da silva:** Não. É  
**Matheus Abranches:** não, né?  
**Luiz Felipe Verissimo da silva:** CD  
**Matheus Abranches:** Já é o P  
**Luiz Felipe Verissimo da silva:** espaço  
**Matheus Abranches:** direto,  
**Luiz Felipe Verissimo da silva:** P. Aí aperta  
**Matheus Abranches:** não tem  
**Luiz Felipe Verissimo da silva:** detalhe.  
**Matheus Abranches:** barra.  
**lucas Badico:** Melhor que ele já coloca  
**Matheus Abranches:** Aí pode  
**lucas Badico:** aí.  
**Matheus Abranches:** ir  
**lucas Badico:** Você  
**Matheus Abranches:** o  
**lucas Badico:** vai entrar em projetos. Nós tem que entrar em projetos depois disso. A gente tá escrevendo em projeto. Se você olhar projeto Hello Word.  
**Matheus Abranches:** projeto  
**Luiz Felipe Verissimo da silva:** É, é, ele tem a barra  
**lucas Badico:** Isso.  
**Luiz Felipe Verissimo da silva:** aí, H e aperta o tab de novo.  
**lucas Badico:** Isso. Dá o enter agora. Você tá  
**Matheus Abranches:** top.  
**lucas Badico:** agora roda o go run pra gente. Vamos ver o que que que vai sair disso aí. Com n seu  
**Matheus Abranches:** sempre  
**lucas Badico:** jogo com  
**Matheus Abranches:** erro, cara.  
**lucas Badico:** tudo bem. Go run.  
   
 

### 00:23:24

   
**lucas Badico:** Aí você precisa passar main.g.  
**Matheus Abranches:** sem  
**lucas Badico:** Coloca  
**Matheus Abranches:** espaço.  
**lucas Badico:** é run espaçma.g. Isso. Prontinho. Você rodou a sua primeira aplicação em Goha.  
**Matheus Abranches:** Top.  
**lucas Badico:** Claro que essa aplicação ela é muito boba, né? Ela não faz nada. Que que a gente vai fazer? A gente vai colocar uma API nela. Eh, pra gente fazer uma aplicação, a gente para ter eh pra gente poder eh criar uma aplicação Go completa, não é só o main. a gente precisa dar um mod init que é como se fosse o para quem é do node é você iniciar um projeto com pack init, né? Eh, com npmit. Faz aí de novo para nós o Mateus. Go mod  
**Matheus Abranches:** Tô indo o mod.  
**lucas Badico:** é sem o e é mod de módulo, sabe? Isso  
**Matheus Abranches:** Uhum.  
**lucas Badico:** com mod init. Só colocar git github.com. Eh, coloca barra Lucas Batico. Só ver como é que é meu  
**Matheus Abranches:** tem  
**lucas Badico:** slide que eu sempre esqueço.  
   
 

### 00:25:07

   
**Matheus Abranches:** eh seu usuário. Tô tentando lembrar aqui. Eu acho que é melhor olhar porque eu eu acho  
**lucas Badico:** Aí  
**Matheus Abranches:** que ele  
**lucas Badico:** eu vou  
**Matheus Abranches:** tem  
**lucas Badico:** olhar,  
**Matheus Abranches:** um acho  
**lucas Badico:** vou  
**Matheus Abranches:** que  
**lucas Badico:** olhar.  
**Matheus Abranches:** ele tem um um tracinho aí.  
**lucas Badico:** Eu acho que no BFB não tem não, mas a gente já confirma. Ah, é Lucas Badico mesmo, tudo junto.  
**Matheus Abranches:** minúsculo também, né?  
**lucas Badico:** É, pode ser menor excursão tem  
**Matheus Abranches:** Esse não tem meu case não. É só Lucas Badico.  
**lucas Badico:** barra. Coloca aí intensivo traço Hello World. Essa essa palavrinha aí que a gente tá colocando, gente, esse essa essa chave é a chave pelo qual o nosso projeto vai ser achado, tá? Então é importante, caso a gente quisesse usar isso no futuro, eh, ser o mesmo PEF que vai ficar no GitHub, entendeu? Então é por isso que a gente tá colocando G.com e o Lucas Badico. Então agora não vai ser importante, mas caso vocês quisessem fazer alguma coisa que vocês vão usar em outro projeto depois, essas palavrinhas tem que bater, entendeu?  
   
 

### 00:26:20

   
**Matheus Abranches:** Uhum.  
**lucas Badico:** Então  
**Matheus Abranches:** Aqui e aqui você você não tá puxando nada do seu GitHub, né? Você  
**lucas Badico:** não,  
**Matheus Abranches:** tá  
**lucas Badico:** só  
**Matheus Abranches:** já  
**lucas Badico:** desse nível.  
**Matheus Abranches:** só definindo o GitHub que ele vai ser vai ser o repositório dele GitHub. Basicamente  
**lucas Badico:** Isso,  
**Matheus Abranches:** isso, né? Tá,  
**lucas Badico:** isso.  
**Matheus Abranches:** tá, tá certo,  
**lucas Badico:** Tipo  
**Matheus Abranches:** né?  
**lucas Badico:** assim,  
**Matheus Abranches:** Hello?  
**lucas Badico:** não, não é nem, não é nem, Mateus, só para explicar isso, não é nem assim, a gente pode colocar um GitHub diferente, tanto que isso aqui tá em outro GitHub, nada a ver.  
**Matheus Abranches:** Uhum.  
**lucas Badico:** O ponto é, se eu quiser achar esse projeto, entendeu? Se eu quiser achar esse projeto para ele ser discoverable, eu tenho que colocar nesse PF, entendeu?  
**Matheus Abranches:** Entendi. Pode ir.  
**lucas Badico:** Sacou?  
**Luiz Felipe Verissimo da silva:** Isso vai tá salvo em algum arquivo na estrutura  
**lucas Badico:** Vai,  
**Luiz Felipe Verissimo da silva:** do GO.  
   
 

### 00:27:04

   
**lucas Badico:** vai. Você vai ver quando ele der enter, ele vai criar um, ele vai criar uma paradinha ali, ó. Ó, ele criou aquele arquivinho no G mod. Abre ele para vocês verem. Você viu? Ó, o PEF que ele colocou  
**Matheus Abranches:** al  
**lucas Badico:** entrou como goó, como modo e o e o caminho desse PEF.  
**Luiz Felipe Verissimo da silva:** Se eu altero isso aí no arquivo Gom.  
**lucas Badico:** O que que vai acontecer? você vai criar o projeto, ele não, ele não é só com Men, certo? Vocês estão ligados que a gente vai criar,  
**Luiz Felipe Verissimo da silva:** Nice.  
**lucas Badico:** né,  
**Luiz Felipe Verissimo da silva:** Vamos.  
**lucas Badico:** várias várias pastinhas. Então, digamos aqui dentro depois eu posso criar uma pastinha de internals, certo? Aí  
**Matheus Abranches:** Mhm.  
**lucas Badico:** essa pasta de internals, ele é colocar aqui P.G, tá? Só para ter arquivo, ele é um package de  
**Matheus Abranches:** Угу.  
**lucas Badico:** internos, certo? Quando eu tiver importando outras pastas do meu próprio projeto, como é que eu importo? Eu importo com, vou até copiar, barra outro package, entendeu?  
   
 

### 00:28:15

   
**lucas Badico:** Então, o impacto de mudar aquilo lá, Felipe, é que você teria de mudar os importes do seu próprio projeto, entendeu?  
**Luiz Felipe Verissimo da silva:** Uhum. Senão ele quebra.  
**lucas Badico:** É, ele ele ele é ele não vai achar as coisas porque basicamente você vai usar isso aqui e aí o Gomod vai dizer: "Ah, esse PEF é eu mesmo, né? É o meu próprio projeto, então vou procurar dentro das minhas próprias pastas, entendeu? Sacou? Então é isso, tipo, esse é o impacto, tá bom? OK. Bom, iniciamos. Qual que é a vantagem então da gente pegou mod nesse momento? A vantagem é que a gente pode instalar outros pacotes, tá? E o que que a gente vai instalar agora? A gente vai instalar o Gorila Mux, tá bom? Eh, então vamos trocar. Quem pode escrever aí a próxima coisa? Sassa quer escrever? Consegue escrever aí  
**sasa cruz:** Deixa eu ver. Calma aí.  
**lucas Badico:** só para não ficar uma pessoa só?  
**sasa cruz:** Tá aparecendo.  
**lucas Badico:** Tá,  
**sasa cruz:** Acho que  
**lucas Badico:** tá  
**sasa cruz:** eu consigo,  
   
 

### 00:29:37

   
**lucas Badico:** ótimo. Ótimo. Boa. Manda lá. S. Você vai escrever o seguinte coisa, S. Você vai instalar o pacote Gorila Mux pra gente. Você vai colocar  
**sasa cruz:** tá?  
**lucas Badico:** go install. Acho que é install. Acho que é go mod install. Desculpa. Go mod  
**sasa cruz:** M.  
**lucas Badico:** install. É go mod install. Deixa eu confirmar só só antes de o branco agora. Como instalar pacote gorila. Deu um brancão agora, desculpa, gente. Ah, não é goget, não é install, não é goget.  
**sasa cruz:** Ah, tá.  
**lucas Badico:** É goget. Mais simples. Go get espaço get. Isso. Aí você vai colocar github.com de novo. Lembra que eu falei que era importante o PEF?  
**sasa cruz:** Git hub. O quê?  
**lucas Badico:** Pom isso barra gorila com do Ls.  
**sasa cruz:** Assim  
**lucas Badico:** Isso. Barra Mux. Isso aí. Coloca um @ e coloca lei teste de último, sabe? Eh, la teste.  
   
 

### 00:31:11

   
**lucas Badico:** Isso. Perfeito. Dá o enter. Prontinho. Olha só o que que ele fez. Ele foi lá no Go, dá uma olhada, e colocou requir pro Gorila Mux. Vocês estão vendo?  
**sasa cruz:** Sim.  
**lucas Badico:** E aí ele criou um arquivinho Bsan pra gente, tá vendo? Qual que é dessa Goan? Se a gente comitar o Goan, toda vez que a gente instalar o projeto e não forçar uma reinstalação, ele vai instalar as mesmas versões daquela daquele arquivo da daquele daquela dependência. É como o package lock no Note, né? Fazendo aí uma comparação direta, tá bom? Eh, tem uma parada que é importante dizer. Go mod é algo relativamente novo no Go. A gente tem ele, se eu não me engano, há 3, 4 anos. É, acho que é até menos, eu acho, até menos que isso. Tem alguns macaco velho do Go que não usam Go Mod. Que que eles fazem? Eles trazem o código. Então, em vez da gente, em vez da gente fazer isso que a gente fez, a gente iria fazer o seguinte coisa.  
   
 

### 00:32:40

   
**lucas Badico:** Olha aqui a minha tela. Só vou citar isso. Eu nunca fiz, tá? Mas eu vou citar porque eu vi sendo feito. Eh, o que que os macaco velho fazem? Eles criam uma pastinha chamada vendor dentro do projeto, certo? Eles nem usam goms, eles usam outra parada. Eles tem uma pastinha chamada vendor e eles trazem o código para dentro dessa pasta. Eles de fato copiam o código, sabe? Eles dão git, né? Git clone e baix aquela versão do código. Por quê? Porque o go só precisa do do só precisa do do código go para rodar. Então, se a gente tivesse, em vez de ter tudo isso que a gente fez, tivesse uma pastinha aqui, né, com vendor. Vendor é porque é de terceiro, tá? A palavrinha que a gente e tivesse aqui eh Mux ou Gorila Mux, a gente poderia importar, entendeu? A gente poderia usar, mas isso é hoje em dia não é muito usado, tá? É só só o macaco velho que vai usar isso aqui. Mas é importante vocês saberem que dá para fazer. Eu nem recomendaria vocês fazerem isso, tá?  
   
 

### 00:34:14

   
**lucas Badico:** Mas é só para dizer que dá. Por muito tempo foi assim que a comunidade GO importou dependência, tá? Mas o que que a gente ganha com isso? Vamos voltar aqui pro main. E agora a gente pode usar o Gorila Max. Aqui nos importar um por um, né? Então a gente poderia importar o log e logo embaixo importa, sei lá, o FMT, que é o o formatador, né? Mas a gente não precisa fazer isso. E normalmente não é como a gente faz. A gente abre um parênteses aqui e importa tudo dentro de um único importe. Tá bom? E aí a gente vai importar agora, a gente vai importar o nosso Gorila Mux. E repara que a string que eu tô usando é a mesma string que a Sala usou na hora de importar, na hora de fazer o goget. A única diferença que eu não preciso colocar o latest, certo? Opa. É, o meu meu coiso ele ele limpa as importações. Eh, bom.  
**Matheus Abranches:** Pera aí. No caso, Lucas, eu  
**lucas Badico:** Hum.  
**Matheus Abranches:** perdi um pouquinho aí. Aí você explicou da diferença que alguns programadores aí mais macaco velho faz, né, usando  
   
 

### 00:35:40

   
**lucas Badico:** Uhum.  
**Matheus Abranches:** o o modelo de de copia e cola lá, né?  
**lucas Badico:** Uhum.  
**Matheus Abranches:** Mas essa parte do import, isso aí já é o uso, né? aquilo que vou que a gente instalou,  
**lucas Badico:** Isso  
**Matheus Abranches:** igual eu instalei um, a Sara instalou o outro e agora a gente tá fazendo o uso dele, né?  
**lucas Badico:** você não instalou, você inicializou o os  
**Matheus Abranches:** O projeto.  
**lucas Badico:** módulos. É, você  
**Matheus Abranches:** Isso.  
**lucas Badico:** inicializou o projetinho, quem instalou foi a Sara. A gente só instalou um pacote até agora  
**Matheus Abranches:** Ah, sim. Beleza,  
**lucas Badico:** e agora a gente vai usar. Exatamente. Entendeu?  
**Matheus Abranches:** beleza, beleza.  
**lucas Badico:** Então  
**Matheus Abranches:** Só  
**lucas Badico:** a  
**Matheus Abranches:** para  
**lucas Badico:** gente vê  
**Matheus Abranches:** só  
**lucas Badico:** aqui  
**Matheus Abranches:** para  
**lucas Badico:** no  
**Matheus Abranches:** ver  
**lucas Badico:** nosso  
**Matheus Abranches:** se eu  
**lucas Badico:** main.  
**Matheus Abranches:** se eu tô acompanhando bem mesmo.  
**lucas Badico:** Claro, claro.  
   
 

### 00:36:13

   
**lucas Badico:** Não, importante. Podem pausar e fazer pergunta, gente. Fiquem à vontade. Então, agora eu abri aqui o importar, a gente pode usar esse módulo. Como é que a gente usa o módulo? Não parece, mas no próprio, na própria string tem a palavra, vamos dizer assim, a o holder do módulo, que é o último termo. Então, vocês estão vendo aqui, aqui só tem um termo que é o log. Então aqui eu usei log, tá vendo? No caso do Mux, é a palavra Mux. Pode ter 10, pode ser 10 subníveis. Você tem com 10 subníveis. A última é a palavrinha, certo? Que a gente vai usar. Então, a gente aqui vai criar um router que a gente chama Muxer.  
**Matheus Abranches:** Угу.  
**lucas Badico:** E agora a gente pode então criar um web server, né? Vamos dar essa essa frasezinha aqui para crever, certo? Vocês estão vendo que o que o meu VS Code ele tá começando a reclamar de algumas coisas? Ó, ele ficou vermelhinho porque ó, você não tá usando. Caso a gente precise instanciar algo que não vai usar no GO, a gente precisa declaradamente  
   
 

### 00:37:44

   
**Luiz Felipe Verissimo da silva:** Underl  
**lucas Badico:** deixar o underline, tá? Mas no nosso caso, a gente vai usar assim. Eh, como é que é que a gente faz um um router, né? Um router ele é bem simples em com o MUX, né? Ele é bem ele é bem claro, na verdade. O router, ele tá nos ajudando aqui a definir as nossas APIs. Lembra que eu mostrei aqui no PHman? verbo e tal, verbo, pf, header e tudo mais. No router a gente tem acesso a definir esse contrato. A gente vai fazer um contrato bem simples. Agora a gente vai fazer um contrato de Hello World, certo? Eu vou fazer o primeiro e vocês depois vão me ajudar a fazer o segundo. Então, nesse primeiro contrato, eu vou dizer que eu quero uma handle funk no pato. Lembra que eu falei que tudo é um IO, né? Tudo é um input, output. Olha como o Go representa isso. E eu acho isso fantástico. Você tem um writer, opa, ele foi me ajudar. Ele ele me atrapalhou.  
**Matheus Abranches:** Auto complit às vezes faz isso.  
**lucas Badico:** É, deixa me ajudar, me atrapalhou.  
   
 

### 00:39:18

   
**lucas Badico:** Então você tem um writer e você tem um header. E aí eu preciso importar um outro pacote aqui que é o net do stand delibery http, certo?  
**Matheus Abranches:** Mhm. Esse  
**lucas Badico:** Então  
**Matheus Abranches:** esse só só um minuto, Lux.  
**lucas Badico:** pode  
**Matheus Abranches:** Esse  
**lucas Badico:** falar  
**Matheus Abranches:** newter eh ele é um uma função do próprio Mux ou é  
**lucas Badico:** isso  
**Matheus Abranches:** uma função do do é do próprio Mux, né?  
**lucas Badico:** do Mux. É o MX que tá nos nos entregando.  
**Matheus Abranches:** Tá.  
**lucas Badico:** O GO nesse pacote HTTP ele tem um um servidor, tá? Só que  
**Matheus Abranches:** Uhum.  
**lucas Badico:** o servidor que ele tem, ele, a gente até vai usar,  
**Matheus Abranches:** Угу.  
**lucas Badico:** ele não deixa tão simples você filtrar e definir os métodos e as URLs, entendeu? Então, a gente precisa de um router para ficar em cima do servidor para ele poder dizer: "Ah, eu recebi um get ou um post nessa URL, eu recebi esse PEF PM ou recebi aquele PEF PM", entendeu? Então o router é como se fosse um filtro, é como se o Houter fosse a rodovia e o o MX fosse ali o guarda de trânsito, sabe?  
   
 

### 00:40:40

   
**lucas Badico:** Para dizer: "Ah, você é caminhão, então você vai por aqui, você é carro  
**Matheus Abranches:** Uhum.  
**lucas Badico:** de passeio, você vai por aqui". Entendeu?  
**Matheus Abranches:** É como um filtro, né? Faz  
**lucas Badico:** Não é  
**Matheus Abranches:** um  
**lucas Badico:** um filtro.  
**Matheus Abranches:** direcionamento, né?  
**lucas Badico:** Exatamente. Tá bom. para deixar mais fácil as nossas funções de para lidar com as APIs, entendeu?  
**Matheus Abranches:** Entendi.  
**lucas Badico:** Que que a gente vai fazer? A gente vai a gente vai a gente vai escrever WRE Word, certo? E a gente vai dizer que isso aqui é um get, é um methods,  
**Matheus Abranches:** Угу.  
**lucas Badico:** que isso aqui é um get, tá bom? A gente vai colocar outro log. E aí a gente vai aqui colocar o http dot e esse é o nosso servidor de fato. Então é isso aqui que tá rodando, que é o servidor, mas a gente tá colocando as funções de handler, porque aquele servidor só tem o handler. A gente tá colocando as funções de handler que nascem do nosso router, entendeu? Então é de fato um filtro,  
**Matheus Abranches:** Esse  
   
 

### 00:42:26

   
**lucas Badico:** entendeu?  
**Matheus Abranches:** R é de render, né?  
**lucas Badico:** É, é não, R de  
**Matheus Abranches:** Não.  
**lucas Badico:** router aqui.  
**Matheus Abranches:** Ah, é o router que você criou,  
**lucas Badico:** Isso  
**Matheus Abranches:** tá?  
**lucas Badico:** é  
**Matheus Abranches:** Eu  
**lucas Badico:** o ro  
**Matheus Abranches:** eu já me perdi, mas  
**lucas Badico:** sem  
**Matheus Abranches:** beleza.  
**lucas Badico:** problema.  
**Matheus Abranches:** Nem achei  
**lucas Badico:** Felipe,  
**Matheus Abranches:** de novo.  
**lucas Badico:** você quer rodar aí pra gente?  
**Luiz Felipe Verissimo da silva:** Claro, claro. Eh, doutor no arquivo men, então ele vai ser Gohan. Gohan.  
**lucas Badico:** Isso,  
**Luiz Felipe Verissimo da silva:** Nem nem. Go. É isso  
**lucas Badico:** isso mesmo.  
**Luiz Felipe Verissimo da silva:** agora.G.  
**lucas Badico:** Prontinho. Eu vou  
**Luiz Felipe Verissimo da silva:** Server tá ao vivo.  
**lucas Badico:** eu vou chamar aqui pelo meu próprio post. Eu acho que dá para vocês verem minhas rotas também. A gente já testa isso, mas vamos ver aqui como é que a gente testa isso. HTTP, porque a gente tá demonstrando, a gente tá usando local host, certo?  
   
 

### 00:43:27

   
**Luiz Felipe Verissimo da silva:** Oi, Угу.  
**lucas Badico:** Eh, opa. E a gente vai chamar em hello, lembra? Vota hello. Vamos mandar para ver o que que vai vir. Ah, eu escrevi errado ou ou tá certo? Não, tá errado. Tá errado, mas tudo bem. Mas a gente provou que tá funcionando, certo? Eh, a gente provou que tá funcionando. Vamos ver aqui no se roda um post. Ó, ele não não roda. Ele diz que não tá disponível. Tá vendo? Então esse essa é anatomia básica, certo? Essa é anatomia básica de um  
**Matheus Abranches:** Tá certo, Lucas. Você rodou a porta relô e o  
**lucas Badico:** não,  
**Matheus Abranches:** print  
**lucas Badico:** mas  
**Matheus Abranches:** que  
**lucas Badico:** eu acho  
**Matheus Abranches:** ele  
**lucas Badico:** que  
**Matheus Abranches:** deu  
**lucas Badico:** palavra  
**Matheus Abranches:** foi  
**lucas Badico:** tá  
**Matheus Abranches:** Word.  
**lucas Badico:** errada.  
**Luiz Felipe Verissimo da silva:** É,  
**Matheus Abranches:** Ah, não.  
**Luiz Felipe Verissimo da silva:** Word  
**Matheus Abranches:** Mas beleza.  
**Luiz Felipe Verissimo da silva:** é mundo, né?  
   
 

### 00:44:26

   
**lucas Badico:** Tá escrito certo por alguns segundos. Eu achei que tá escrito errado, mas eu já não tenho certeza.  
**Matheus Abranches:** Tá certo.  
**lucas Badico:** Tá certo. Eh,  
**Matheus Abranches:** Vai parar.  
**lucas Badico:** tá dar uma  
**Matheus Abranches:** É.  
**lucas Badico:** paradinha  
**Matheus Abranches:** Tá.  
**lucas Badico:** aqui.  
**Matheus Abranches:** O L. O L foi a mais, hein.  
**lucas Badico:** Eu acho que foi.  
**Matheus Abranches:** V ele foi.  
**lucas Badico:** Ah, ok. Deixa eu ver só uma coisa aqui, gente. Tá, eu vou eu vou deixar um desafio para vocês aqui enquanto eu vou no banheiro. Criem uma rota aqui. cria uma rota aqui que vai responder um pong em cima de um post. Então eu vou colocar aqui no chat, no no chat. Então é um post na o na rota ping que vai responder um pong. É o clássico ping pong.  
**Luiz Felipe Verissimo da silva:** Insta,  
**lucas Badico:** Faço  
**Luiz Felipe Verissimo da silva:** você  
**lucas Badico:** aí  
**Luiz Felipe Verissimo da silva:** assume  
**lucas Badico:** enquanto  
**Luiz Felipe Verissimo da silva:** aí  
**lucas Badico:** eu vou no banheiro.  
**Luiz Felipe Verissimo da silva:** que você foi o você não escreveu ainda, né?  
   
 

### 00:46:03

   
**lucas Badico:** Eu coloquei no nosso chat aqui, ó.  
**Matheus Abranches:** É o  
**lucas Badico:** Coloquei no chat.  
**Matheus Abranches:** Adolfo, Adolfo tá digitando ali. A gente acompanha  
**lucas Badico:** Boa.  
**Matheus Abranches:** o Adolfo junto  
**lucas Badico:** Eu  
**Matheus Abranches:** com  
**lucas Badico:** vou  
**Matheus Abranches:** ele  
**lucas Badico:** no banheiro  
**Matheus Abranches:** ali.  
**lucas Badico:** aqui rapidão.  
**Matheus Abranches:** Beleza.  
**Luiz Felipe Verissimo da silva:** Aí,  
**adolfo agnelli:** Tô perdido.  
**Matheus Abranches:** Não.  
**Luiz Felipe Verissimo da silva:** vamos lá  
**Matheus Abranches:** É,  
**Luiz Felipe Verissimo da silva:** então  
**Matheus Abranches:** eu,  
**Luiz Felipe Verissimo da silva:** lá.  
**Matheus Abranches:** pelo que eu entendi, a gente vai refazer basicamente esse estilo de rota aí, só que o método vai ser post e a rota, em vez de ser hello, vai ser vai ser eh  
**Luiz Felipe Verissimo da silva:** Isso. Vamos  
**Matheus Abranches:** pink  
**Luiz Felipe Verissimo da silva:** lá. Eu eu  
**Matheus Abranches:** com  
**Luiz Felipe Verissimo da silva:** ajudo.  
**Matheus Abranches:** a resposta  
**Luiz Felipe Verissimo da silva:** Eu ajudo.  
**Matheus Abranches:** long. Eu  
   
 

### 00:46:51

   
**Luiz Felipe Verissimo da silva:** Tá  
**Matheus Abranches:** acho  
**Luiz Felipe Verissimo da silva:** tipo  
**Matheus Abranches:** que é isso,  
**Luiz Felipe Verissimo da silva:** fres  
**Matheus Abranches:** né?  
**Luiz Felipe Verissimo da silva:** isso aí.  
**Matheus Abranches:** Tá  
**Luiz Felipe Verissimo da silva:** Bora  
**Matheus Abranches:** bem  
**Luiz Felipe Verissimo da silva:** lá.  
**Matheus Abranches:** bem  
**Luiz Felipe Verissimo da silva:** Linha  
**Matheus Abranches:** parecido.  
**Luiz Felipe Verissimo da silva:** 17\. Adolfo,  
**adolfo agnelli:** Mhm.  
**Luiz Felipe Verissimo da silva:** dá um enter aí para abrir um espaço. Primeira coisa, você vai utilizar o R, né, para utilizar o router R  
**adolfo agnelli:** Sì.  
**Matheus Abranches:** É a  
**Luiz Felipe Verissimo da silva:** dot.  
**Matheus Abranches:** é a nova rota, né?  
**Luiz Felipe Verissimo da silva:** Isso é aí a gente tá chamando router. Aí você vai utilizar a assinatura do handlef lá. Abre parênteses.  
**adolfo agnelli:** Nossa,  
**Luiz Felipe Verissimo da silva:** Abre  
**adolfo agnelli:** ele tá  
**Luiz Felipe Verissimo da silva:** parênteses.  
**adolfo agnelli:** ele tá descendo para minha  
**Luiz Felipe Verissimo da silva:** É.  
**adolfo agnelli:** entrar aqui. Pronto.  
**Luiz Felipe Verissimo da silva:** Aí.  
   
 

### 00:47:28

   
**Luiz Felipe Verissimo da silva:** Aí você vai fazer o barra ping.  
**adolfo agnelli:** É meu esquema aí. Será?  
**Matheus Abranches:** É,  
**Luiz Felipe Verissimo da silva:** Ah,  
**Matheus Abranches:** mas  
**Luiz Felipe Verissimo da silva:** não,  
**Matheus Abranches:** aí  
**Luiz Felipe Verissimo da silva:** não.  
**Matheus Abranches:** você  
**Luiz Felipe Verissimo da silva:** Faltou você fechar.  
**adolfo agnelli:** Ah, tá verdade.  
**Matheus Abranches:** é  
**adolfo agnelli:** Barra  
**Matheus Abranches:** é  
**adolfo agnelli:** ping  
**Matheus Abranches:** porque  
**adolfo agnelli:** depois.  
**Matheus Abranches:** ela Ela fechou embaixo lá, ó. Tem outro.  
**adolfo agnelli:** Ah, tá. É esse auto complete que ele já  
**Matheus Abranches:** É  
**adolfo agnelli:** deixa assim.  
**Matheus Abranches:** boa.  
**adolfo agnelli:** Tá  
**Luiz Felipe Verissimo da silva:** Vai lá. Agora mesmo em cima. Isso aí. P. Abre. Sete aparências.  
**adolfo agnelli:** aqui. Respons por fou um aqui. Hum. Tá certo. Tá, né?  
**Luiz Felipe Verissimo da silva:** pass  
**adolfo agnelli:** Agora que eu vou ter o  
   
 

### 00:48:42

   
**Luiz Felipe Verissimo da silva:** Ong. Isso  
**adolfo agnelli:** o pong.  
**Matheus Abranches:** Aí o pong  
**Luiz Felipe Verissimo da silva:** aí se encadeia com o método  
**adolfo agnelli:** Tem  
**Luiz Felipe Verissimo da silva:** lá embaixo,  
**Matheus Abranches:** foi métodes.  
**Luiz Felipe Verissimo da silva:** né? Mods.  
**adolfo agnelli:** um post, né?  
**Matheus Abranches:** Isso.  
**Luiz Felipe Verissimo da silva:** Isso. Tudo maiúsculo. Beleza. Dá um espaço aí e denta isso. Carinha, entendeu? Aí o que tá acontecendo? explica para nós.  
**sasa cruz:** Entendi. Ah, nossa, mas aí se me complica. Ele deu um ele puxou o ali, deu uma função, colocou ali para chamar o ping quando a gente for dar para chamar ele. Daí ele colocou para escrever, né, que é o Writer, colocou http, colocou para ler e daí ele vai escrever o pong quando a gente no terminal lá. É isso. Ч.  
**Luiz Felipe Verissimo da silva:** Iso. Ó,  
**Matheus Abranches:** Perfeito.  
   
 

### 00:49:40

   
**Luiz Felipe Verissimo da silva:** eh,  
**adolfo agnelli:** Mas aí  
**Luiz Felipe Verissimo da silva:** tipo,  
**adolfo agnelli:** que nem  
**Luiz Felipe Verissimo da silva:** no  
**adolfo agnelli:** assim  
**Luiz Felipe Verissimo da silva:** final das  
**adolfo agnelli:** eu colocar  
**Luiz Felipe Verissimo da silva:** contas, falei  
**adolfo agnelli:** para ela coloquei. Ah, escrevi o response writer, mas o que que é? Friter  
**Luiz Felipe Verissimo da silva:** é tipo,  
**adolfo agnelli:** é  
**Luiz Felipe Verissimo da silva:** tá  
**adolfo agnelli:** um m  
**Luiz Felipe Verissimo da silva:** perguntando para ela.  
**adolfo agnelli:** f todo mundo aqui.  
**Matheus Abranches:** Não, eh,  
**Luiz Felipe Verissimo da silva:** Ah, então  
**Matheus Abranches:** eu entendo  
**Luiz Felipe Verissimo da silva:** vamos  
**Matheus Abranches:** ali  
**Luiz Felipe Verissimo da silva:** lá,  
**Matheus Abranches:** como  
**Luiz Felipe Verissimo da silva:** ó.  
**Matheus Abranches:** a  
**Luiz Felipe Verissimo da silva:** A função,  
**Matheus Abranches:** resposta a resposta da request, né? a response da request ali que a  
**Luiz Felipe Verissimo da silva:** ó,  
**Matheus Abranches:** os  
**Luiz Felipe Verissimo da silva:** vamos  
**Matheus Abranches:** dois  
**Luiz Felipe Verissimo da silva:** pensar  
**Matheus Abranches:** os  
**Luiz Felipe Verissimo da silva:** no seguinte,  
**Matheus Abranches:** dois.  
   
 

### 00:50:15

   
**Luiz Felipe Verissimo da silva:** ó. A função handle funk, que é do gurilam, ela recebe dois parâmetros, uma que é o, que é a rota, e outra é tipo pro JavaScript aqui é um callback, né? Então, que é uma função ali que vai executar alguma coisa. Essa função, por sua vez, vai receber outros dois parâmetros. Uma que é o W, que a gente nomeou de W, pode ser qualquer coisa aqui, e que ela é do tipo response writer. Isso aqui pertence a esse pacote aqui. Eh, e depois o R, que é o também pertence à aquele pacote que vai ser o request. Aí dentro da função a gente vai utilizar o a função right do W. Então esse write ele pertence a esse cara e vai devolver um arrei de bytes com essa string aqui lá. E por último a gente diz pra função qual é o tipo, qual o verbo, né, do da chamada http. conseguiram entender  
**Matheus Abranches:** Sim.  
**Luiz Felipe Verissimo da silva:** ou não?  
**lucas Badico:** Ла.  
**sasa cruz:** Só uma pergunta, mas por que em cima tá escrito get e embaixo tá escrito post? Por que que teve essa mudança?  
   
 

### 00:51:42

   
**sasa cruz:** Ah, porque ele mandou pra gente fazer um post, né? Entendi.  
**Luiz Felipe Verissimo da silva:** É que  
**lucas Badico:** Isso.  
**Luiz Felipe Verissimo da silva:** são verbos diferentes, né? Nesse caso, ele não necessita a gente fazer um post. Eh, post seria mais para, sei lá, para fazer uma atualização ou então para desenvolver uma estrutura atualização.  
**Matheus Abranches:** Você  
**lucas Badico:** Ah, normalmente é uma  
**sasa cruz:** Entendi.  
**lucas Badico:** ação, né? Eh, significa uma ação postil, né? Como ping, a gente ping teoricamente é bem clássico de de servidor, é para você verificar se aquele servidor tá online, entendeu? Você tá fazendo uma ação, me diz tá online. Aí ele manda um ponto, quer dizer que tá online, entendeu? Então aí por isso é um pst  
**sasa cruz:** Entendi.  
**Matheus Abranches:** você provoca uma ação e espera uma reação.  
**lucas Badico:** isso.  
**Matheus Abranches:** O ping ação  
**lucas Badico:** Mas tem  
**Matheus Abranches:** pong.  
**lucas Badico:** uma, é a questão  
**Matheus Abranches:** Угуm.  
**lucas Badico:** que eu falei, é além da funcionalidade, é semântica. A gente pode usar get, a gente pode usar delite.  
   
 

### 00:52:39

   
**lucas Badico:** Se quisesse delite aí podia usar, entendeu? Mas ia ser an, ia ser contrassemântica dos verbos, entendeu? tem  
**Matheus Abranches:** Угуm.  
**lucas Badico:** apsar, tipo, as pessoas não se preocupam com isso, mas teoricamente, baseado na semântica, você tem verbos melhor, eh, melhor, mais apropriados para cada para cada ação, entendeu?  
**Luiz Felipe Verissimo da silva:** Sim, ó. Você  
**lucas Badico:** Mas  
**Luiz Felipe Verissimo da silva:** sempre  
**lucas Badico:** já  
**Luiz Felipe Verissimo da silva:** pensa quando eu passar o bastão, o próximo amiguinho vai me xingar ou não? Se você quer que ele te xinga, você faz de qualquer jeito, senão você faz  
**lucas Badico:** eu  
**Luiz Felipe Verissimo da silva:** semântico.  
**lucas Badico:** diria para ele, ele vai te xingar pouco ou muito, né?  
**Luiz Felipe Verissimo da silva:** É isso aí.  
**Matheus Abranches:** Ja.  
**Luiz Felipe Verissimo da silva:** Pode ser também.  
**lucas Badico:** Ele vai te xingar pouco ou muito,  
**Matheus Abranches:** te  
**lucas Badico:** eu  
**Matheus Abranches:** xingar,  
**lucas Badico:** acho.  
**Matheus Abranches:** ele vai, não importa. Eu já aprendi  
**lucas Badico:** É.  
   
 

### 00:53:27

   
**Matheus Abranches:** que não importa o o código que eu pego, eu vou sempre xingar o amiguinho. Mas  
**lucas Badico:** Vai, vai, vai.  
**Matheus Abranches:** a diferença se vai ser muito pouco.  
**lucas Badico:** é a quantidade.  
**Matheus Abranches:** É  
**lucas Badico:** Mas é isso. Vamos ver se se tá funcionando. Eu sei se vocês rodaram de novo.  
**Matheus Abranches:** esse não.  
**Luiz Felipe Verissimo da silva:** Não rodamos.  
**adolfo agnelli:** No.  
**lucas Badico:** Salva aí, então. Salvando aqui. Ele tá dando um erro aqui, mas não tô entendendo da onde tá vindo esse erro. Não tá mostrando erro no no coiso. Mas  
**Matheus Abranches:** É porque não tá  
**lucas Badico:** quem  
**Matheus Abranches:** salvo não. Para mim não tá salvo não.  
**lucas Badico:** não aqui salvei já eu salvei?  
**Luiz Felipe Verissimo da silva:** É só na sua tela que tá vendo para  
**lucas Badico:** É.  
**Luiz Felipe Verissimo da silva:** aqui  
**lucas Badico:** Ah,  
**Luiz Felipe Verissimo da silva:** no  
**lucas Badico:** então  
**Luiz Felipe Verissimo da silva:** live.  
**lucas Badico:** tranquilo. Rodando  
**adolfo agnelli:** Okay.  
**lucas Badico:** de novo. Vamos ver se se fizeram certo.  
   
 

### 00:54:16

   
**lucas Badico:** Agora é post e é o ping. Ponto. Certinho. Topzeira. Agora vamos começar a estruturar essa aplicação como de fato.  
**Luiz Felipe Verissimo da silva:** Ô Lucas, posso fazer duas perguntas?  
**lucas Badico:** Claro. Manda.  
**Luiz Felipe Verissimo da silva:** Eh, não existe um hot reload no go no servidor  
**lucas Badico:** Vamos deixar fazer isso  
**Luiz Felipe Verissimo da silva:** só.  
**lucas Badico:** nas próximas semanas. Mas existe, existe. Eu já fiz um. Eu tenho um. Cadê?  
**Luiz Felipe Verissimo da silva:** Ah, não, tudo bem, a gente faz da própria,  
**lucas Badico:** Mas  
**Luiz Felipe Verissimo da silva:** tá tudo  
**lucas Badico:** existe,  
**Luiz Felipe Verissimo da silva:** bem. E a outra  
**lucas Badico:** existe.  
**Luiz Felipe Verissimo da silva:** questão é, eh, existe como a gente importar um arquivo para criar um router mesmo? Porque senão vai ficar ativão aí o  
**lucas Badico:** Uhum.  
**Luiz Felipe Verissimo da silva:** Mengô, né?  
**lucas Badico:** A gente vai entrar nesse ponto agora. Vou mostrar o que que a gente chama de Handler, tá bom?  
**Matheus Abranches:** Eh, deixa eu só confirmar um entendimento também aqui. Eh, ali na primeira ou na segunda função, tanto faz.  
   
 

### 00:55:19

   
**Matheus Abranches:** Depois que a gente define a nova rota, a gente vai o o handle function, aí a gente define barra reload barra relow barra ping. Aí essa funk fun funk fun funk aí ela recebe esse W. Esse W pode ser qualquer coisa, né? Você chamou de W, mas eu  
**lucas Badico:** É  
**Matheus Abranches:** posso  
**lucas Badico:** o W  
**Matheus Abranches:** chamar  
**lucas Badico:** em si.  
**Matheus Abranches:** ele.  
**lucas Badico:** Eu aqui eu posso trocar para qualquer coisa.  
**Matheus Abranches:** É.  
**lucas Badico:** O que eu não posso trocar é isso aqui, é o tipo dele.  
**Matheus Abranches:** Sim, sim. Eu  
**lucas Badico:** Eu trocar  
**Matheus Abranches:** só  
**lucas Badico:** o  
**Matheus Abranches:** queria  
**lucas Badico:** tipo dele  
**Matheus Abranches:** confirmar  
**lucas Badico:** quebra.  
**Matheus Abranches:** se isso era uma declaração, o que que era isso, entendeu?  
**lucas Badico:** É,  
**Matheus Abranches:** Eu fiquei  
**lucas Badico:** é,  
**Matheus Abranches:** na dúvida.  
**lucas Badico:** é um argumento da função, entendeu?  
**Matheus Abranches:** Um argumento mesmo que eu posso definir,  
**lucas Badico:** É,  
**Matheus Abranches:** não precisa ser específico, né?  
**lucas Badico:** é tipo assim,  
**Matheus Abranches:** Só queria confirmar  
   
 

### 00:56:01

   
**lucas Badico:** não,  
**Matheus Abranches:** essa  
**lucas Badico:** não  
**Matheus Abranches:** essa  
**lucas Badico:** precisa  
**Matheus Abranches:** dúvida,  
**lucas Badico:** ser específico no nome  
**Matheus Abranches:** mas  
**lucas Badico:** da  
**Matheus Abranches:** ele  
**lucas Badico:** coisa,  
**Matheus Abranches:** tem que ser  
**lucas Badico:** mas  
**Matheus Abranches:** tem que ser declarativo, né? Tem que tá,  
**lucas Badico:** é, mas tem que ter esses dois tipos. Ele, essa função nesse lugar tem que receber esses dois tipos, o o writer e o request, entendeu?  
**Matheus Abranches:** tá? E o padrão usado é sempre esse aí que você ensinou. Normalmente  
**lucas Badico:** Isso  
**Matheus Abranches:** se usa assim, né?  
**lucas Badico:** é porque que acontece variáveis com uma letra só são ruins, mas nesse caso a semântica o que é tá no tipo, entendeu?  
**Matheus Abranches:** Sim.  
**lucas Badico:** Então a gente a gente mantém letra só desde que mantenha próximo do tipo, entendeu?  
**Matheus Abranches:** Sim,  
**lucas Badico:** Como  
**Matheus Abranches:** entendi.  
**lucas Badico:** é uma função bem curtinha, mas vamos lá. A gente vai, a gente vai fazer agora começar pelo menos a criar os handlers.  
   
 

### 00:56:54

   
**lucas Badico:** O que que são os handlers? São as funções que vão cuidar da nossa infra, né? Deixa eu pegar o Scalidra. Vocês quiserem ver aqui o a tela do coiso que lembra que tá tem um atraso de Dralma. Depois eu vou ver se eu ficou muito explicar para mim agora. Eh, bom, vamos lá.  
**Matheus Abranches:** Угуm.  
**lucas Badico:** Que que acontece? Vamos, vamos colocar isso aqui. A gente tem o que a gente chama camadas da aplicação, certo? aplicação é como se fosse uma cebola, certo? A camada mais externa dessa cebola é o que a gente chama de eh é o que é o que seria a infra, tá? A infra, a o mundo externo, certo?  
**adolfo agnelli:** Não.  
**lucas Badico:** Como a, como o Adolfo falou, a gente tem a ideia, aquela abstração do restaurante. Eh, o você tem as camadas internas que são que tá dentro da cozinha, né? É, tipo, e você tem no back end o handler, o que tá levando e trazendo, que é no nosso caso o seria no caso o garsom, né? O Rangler ele é o que pega da camada de fora e traz pra camada de dentro, tá bom?  
   
 

### 00:58:58

   
**lucas Badico:** Então oll ele entende o contrato da infra, entende o que? Quero dizer, no nosso caso, a infra é um http, certo? A infra é um HTTP, mas a infra pode ser outras coisas. A gente pode fazer uma aplicação que é uma CELI, a gente pode fazer uma aplicação que é um GRPC, a gente pode fazer uma aplicação que é um um web socket, certo? Então, a gente tem esse conceito de que o que o que converte desses caras paraa camada da aplicação é um hanger, certo?  
**adolfo agnelli:** Sim. Hallo  
**lucas Badico:** Beleza,  
**Luiz Felipe Verissimo da silva:** E  
**lucas Badico:** vamos.  
**Luiz Felipe Verissimo da silva:** para todo mundo, tá claro o conceito do que é um handler, da palavra handler? Acho que isso ajuda a deixar, né, claro.  
**lucas Badico:** Uhum.  
**Luiz Felipe Verissimo da silva:** Угу.  
**lucas Badico:** E  
**adolfo agnelli:** claro.  
**lucas Badico:** aí o handler ele é ele é uma coisa meio boba. Vou fazer aqui uma o handler mais simples que tem, que é o quê? pegar essa nossa função, dar um nome para ela e substituir aqui. Isso é o Rangler, tá? Só que continua com o mesmo BO que o Felipe falou, velho. Se a gente colocar várias funções, isso aqui vai ficar gigantesco, correto?  
   
 

### 01:01:12

   
**lucas Badico:** Que que a gente faz? Normalmente a gente cria um pacote chamado Hangler. Aí, olha só, já muda algumas coisas. Agora o package desse cara vai ser HLY, certo? E eu vou ter aí eu vou levar essas funções para aquele package, certo? Ó, agora como é que eu acesso aquelas funções? Eu tenho que importar aquele cara. Então, vou fazer github.com/lupas/ intensivo. Como é que é intensivo aqui? Não, go mod, sempre esqueço.  
**Matheus Abranches:** É que o Hello ficou errado também.  
**lucas Badico:** P barra, lembra? Lembra barra handler. E aí, o que que eu tenho? Eu tenho ponto aquelas funções. Opa, por  
**Luiz Felipe Verissimo da silva:** sem  
**lucas Badico:** que não?  
**Luiz Felipe Verissimo da silva:** precisar exportar nada.  
**lucas Badico:** Oi,  
**Luiz Felipe Verissimo da silva:** Sem precisar  
**lucas Badico:** então  
**Luiz Felipe Verissimo da silva:** exportar nada.  
**lucas Badico:** essa que é  
**Matheus Abranches:** Eu  
**lucas Badico:** a brincadeira. Como é que você exporta as coisas em Go?  
**Luiz Felipe Verissimo da silva:** Ah, só declarando. Não, não sei porque do jeito que você fez.  
   
 

### 01:03:09

   
**Matheus Abranches:** não vejo também  
**Luiz Felipe Verissimo da silva:** Eh,  
**Matheus Abranches:** a o o export aí.  
**Luiz Felipe Verissimo da silva:** ó, mas eu percebo, por exemplo, que no arquivo handler. Go. As funções estão apagadas, o nome das funções. Isso quer dizer  
**lucas Badico:** Tem  
**Luiz Felipe Verissimo da silva:** que  
**lucas Badico:** algo errado. Tá faltando export.  
**Luiz Felipe Verissimo da silva:** é  
**lucas Badico:** Como é que exporting?  
**Luiz Felipe Verissimo da silva:** você é eu não sei dizer.  
**Matheus Abranches:** É JavaScript mesmo, export e e o nome  
**lucas Badico:** Ah,  
**Matheus Abranches:** das funções.  
**lucas Badico:** dog é genial,  
**Luiz Felipe Verissimo da silva:** Eu  
**lucas Badico:** velho. É muito estranho no primeiro momento, mas é genial.  
**Matheus Abranches:** Vou  
**adolfo agnelli:** Vamos  
**Matheus Abranches:** pesquisar.  
**adolfo agnelli:** ver.  
**Luiz Felipe Verissimo da silva:** posso resolver isso aqui na internet, mas  
**Matheus Abranches:** Então  
**lucas Badico:** É  
**Matheus Abranches:** eu  
**lucas Badico:** com  
**Matheus Abranches:** pensei nisso também.  
**lucas Badico:** letra maiúscula.  
**Matheus Abranches:** Olha,  
**lucas Badico:** Tudo que começa com uma letra maiúscula é público. Tudo que é letra minúscula é privado.  
   
 

### 01:04:04

   
**Matheus Abranches:** isso para mim é muita abstração. Sim, eu tô vendo o gol com abstrações técnicas bem interessantes, Porque  
**lucas Badico:** Por que por que que eles fizeram isso? qualquer outra linguagem, qualquer outra mesmo, se você, para você saber se aquela parada é público ou privado, você precisa olhar a definição. No go, você só precisa olhar o nome. Isso vale para função, para variável, para struct. Se for maiúsculo, é público, se for minúsculo, é privado.  
**Luiz Felipe Verissimo da silva:** PHP tinha uma galera que metia um lá embaixo no antes da  
**lucas Badico:** Ah\!  
**Luiz Felipe Verissimo da silva:** função e aí era privado. Eu já vi  
**lucas Badico:** É,  
**Luiz Felipe Verissimo da silva:** gente fazendo isso com JavaScript  
**lucas Badico:** eu faço isso  
**Luiz Felipe Verissimo da silva:** também.  
**lucas Badico:** em JavaScript também. Eu faço isso para para manter uma, tipo, é uma é uma convenção, é só convenção para dizer: "Olha, isso aqui eu não quero que seja usado lá fora, mas eh go usa isso tipo no no run time dele. Ele realmente e no compiler, né? Tornei aqui maiúsculo, ele deixou eu usar aqui, tá vendo? Se a gente rodar, vamos rodar, puxar aqui o terminal de novo.  
   
 

### 01:05:16

   
**lucas Badico:** Quem, quem pode rodar paraa gente aí para quem que não rodou ainda? Adolf,  
**Luiz Felipe Verissimo da silva:** Eu  
**lucas Badico:** Adolf já rodou.  
**Luiz Felipe Verissimo da silva:** já rodei o Rodolfo. Adolfo que Rodolfo o Adolfo escreveu a função. Não sei se ele  
**lucas Badico:** Boa.  
**Luiz Felipe Verissimo da silva:** jáou.  
**lucas Badico:** Manda lá. Dá pra gente  
**adolfo agnelli:** seria O  
**lucas Badico:** e  
**adolfo agnelli:** mas tem que entrar na base também.  
**lucas Badico:** não, você já tá, você já tá, porque a gente compartilha  
**adolfo agnelli:** Ah,  
**lucas Badico:** mesmo  
**adolfo agnelli:** tá.  
**lucas Badico:** terminal, então você já tá na pasta.  
**adolfo agnelli:** Seria o main também.  
**lucas Badico:** Isso. Main.g. Vamos ver se vai funcionar. Poder aqui a minha tela. Tã, tá funcionando. Mas um handler não é só uma função, não. Ele ele não é só isso aqui. Tá faltando coisa. Por quê? Aí a gente entra em tipos de dentro do coiso. Por quê? O legal de você ter um handler, né, de você ter um pacote Handler, é que você pode compartilhar.  
   
 

### 01:06:32

   
**lucas Badico:** Ai, desculpa, gente. Você pode compartilhar dependências e você pode instanciar essas dependências. A gente ainda não tem aqui a parte do banco, a gente vai tentar colocar ainda hoje o banco, mas digamos que a gente precisasse dentro do nosso handler colocar a dependência do banco, né? Colocar o banco ali. Eh, como que a gente faz isso? porque a gente tem só funções. Então o normal é a gente criar um tipo  
**Matheus Abranches:** Угу.  
**lucas Badico:** Hler e colocar esses essas funções que a gente tem aqui como métodos desse tipo. Eu vou mostrar como fazer aqui pra gente andar um pouquinho mais rápido. Então eu tenho aqui handler letra maiúscula, quer dizer que esse cara vai poder ser acessado lá de fora, certo? Nesse momento dentro do stct dele não tem nada, mas caso tivesse um DB, o que que ia fazer? A gente ia ter DB, tipo alguma coisa, entendeu? Mas nesse momento não temos nada. E aí eu vou substituir esses caras aqui por métodos. Como é que eu substituo por métodos? é quase igual a assinatura da da função, só que você coloca o tipo antes. Então aqui no nosso caso é o opa, é o handler aqui que eu preciso, tá vendo?  
   
 

### 01:08:19

   
**Matheus Abranches:** M.  
**lucas Badico:** E o que que a gente precisa? a gente precisa de um método para criar esse handler, que é um construtor. Em inglês, eu não tenho a palavra new como um construtor. Então você precisa criar o seu construtor. A gente usa um método público, new. Se tivesse dependência, as dependências vêm dentro desse desse método aqui. Esse essa função  
**Luiz Felipe Verissimo da silva:** terus  
**lucas Badico:** vai retornar um ponteiro de Hangler. né? Uma referência de Handler e a gente vai retornar um handlerzinho. Como que a gente usa isso? Só um minuto.  
**Luiz Felipe Verissimo da silva:** Pera aí. Ah, foi uma pergunta.  
**lucas Badico:** Vamos andar.  
**Luiz Felipe Verissimo da silva:** Ah, imaginam que quando você for chamar essas funções lá no M go, você vai fazer um  
**Matheus Abranches:** Угу.  
**Luiz Felipe Verissimo da silva:** new. Aí, R pon new ponto o nome da função. Isso ou não?  
**lucas Badico:** Tipo isso, tipo isso. É bem parecido. É bem parecido. Você tá, você tá no caminho certo. Vamos lá.  
   
 

### 01:09:44

   
**lucas Badico:** Vamos chamar de h dois pontos aqui. E eu vou chamar handler. New, certo? A partir daqui eu tenho a minha referência de handler, certo? E aí eu vou mudar isso aqui. Hhello e  
**Luiz Felipe Verissimo da silva:** distanciado  
**lucas Badico:** H  
**Luiz Felipe Verissimo da silva:** uma vez só.  
**lucas Badico:** só uma vez. Essa aqui é a mágica. Você você só vai precisar de uma conexão com o banco da sua PI, não precisa instanciar várias vezes, entendeu? Perceberam  
**Matheus Abranches:** Volta lá  
**lucas Badico:** a  
**Matheus Abranches:** na  
**lucas Badico:** mágica.  
**Matheus Abranches:** construção do Isso. Quero que volta, volte aí pra gente ver. Ah, boa, melhorou.  
**Luiz Felipe Verissimo da silva:** É tão engraçado a linguagem sobre gás posso assim dessa  
**lucas Badico:** M.  
**Luiz Felipe Verissimo da silva:** forma. Aí é engraçado a linguagem fazer o você criar o construtor. É porque eu tô acostumado com as coisas já mais abstratas, sabe?  
**lucas Badico:** Uhum.  
**Luiz Felipe Verissimo da silva:** Mais mastigadas.  
**Matheus Abranches:** É o type aí, essa parte do type eu perdi, Lucas, você pode pode me explicar ela aí novamente?  
   
 

### 01:11:05

   
**Matheus Abranches:** Eu não peguei essa  
**lucas Badico:** Esse  
**Matheus Abranches:** não.  
**lucas Badico:** aqui  
**Matheus Abranches:** É, essa parte aí vem vem do banco de dados, né? Esse DB  
**lucas Badico:** não, não, não, não vem,  
**Matheus Abranches:** eu  
**lucas Badico:** não  
**Matheus Abranches:** não  
**lucas Badico:** vem.  
**Matheus Abranches:** eu não.  
**lucas Badico:** Vamos lá. Antes eram funções soltas, certo?  
**Matheus Abranches:** Aham.  
**lucas Badico:** O que que a gente fez? A gente criou um tipo handler, certo? Para essas funções ficarem dentro desse tipo. Nesse momento não tem vantagem funcional, é só semântica,  
**Matheus Abranches:** Sim.  
**lucas Badico:** certo? Não tem não tem nenhuma vantagem específica, certo? Não tem nada aqui que seja vantajoso. Daqui a pouco vai ser, vocês vão entender.  
**Matheus Abranches:** Угу.  
**lucas Badico:** Tem nada que seja vantajoso aqui. Eh, mas é como as coisas são feitas no GO. Você define um tipo e você coloca métodos, você coloca funções naquele tipo, que é o que a gente fez aqui, entendeu? Isso aqui é como no JavaScript você definir um objeto, uma classe com funções lá dentro, com métodos, entendeu?  
   
 

### 01:12:09

   
**Matheus Abranches:** Entendi.  
**lucas Badico:** Só que aqui a gente faz com tipo, não tem classe, entendeu?  
**Matheus Abranches:** Uhum.  
**lucas Badico:** A gente faz com  
**Matheus Abranches:** Acho que  
**lucas Badico:** tipo.  
**Matheus Abranches:** eu acho que agora ficou claro. Deu, deu para entender. Tem que  
**lucas Badico:** Você vocês  
**Matheus Abranches:** depois  
**lucas Badico:** vão entender a utilidade disso aqui na na próxima  
**Matheus Abranches:** isso.  
**lucas Badico:** parte. Vocês vão entender  
**Matheus Abranches:** Beleza.  
**lucas Badico:** a utilidade disso aqui da próxima parte. Sarinha, tá acompanhando?  
**sasa cruz:** Tu  
**lucas Badico:** Boa, boa. Eh, certo. Tá tranquilo aqui,  
**Matheus Abranches:** Tranquilo.  
**lucas Badico:** tá?  
**adolfo agnelli:** Mhm.  
**lucas Badico:** Vamos fazer uma paradinha aqui. Ah, a gente tem quanto tempo? A gente tem meia hora. Eu vou acelerar uma parte aqui para eu poder mostrar o o banco de dados funcionando. Só que eu não vou explicar exatamente o que eu tô fazendo nesse momento. Eu vou explicar a parte do banco de dados, tá? Então vocês vão ter só que acreditar em mim.  
   
 

### 01:13:16

   
**lucas Badico:** Então, pos depois eu vou vou quebrar e explicar certinho isso aqui. Não. Não. M. Okay. Tô com D ligado. Que legal. Espera um momento, gente. Quem aí conhece Docker? Eu acho que é só não. Sarinha. Ô Sarinha, você já conhece o Docker?  
**sasa cruz:** Não.  
**lucas Badico:** Quanto carrega aqui, eu vou explicar. A gente vai ter que esperar mesmo. Tá carregando já. Não sei. Tá, tá, tá carregando lá. OK. Eh, é o seguinte, Sara. Eh, Docker é uma ferramenta que a gente usa para criar algumas coisas que, de outra forma, seria muito chato de criar. No, no meu caso, eu vou criar um banco de dados, tá bom? Então, a gente vai iniciar um banco de dados super rápido. Olha só, o que eu precisei fazer foi escrever esse arquivo, certo? Precisei escrever esse arquivo aqui e eu vou dar um docker compose up e ele vai criar um banco de dados para mim local. Não, não é um banco de dados de fato assim que vou usar em produção, só local pra gente desenvolver, entendeu?  
   
 

### 01:15:36

   
**lucas Badico:** Ele vai dar um tempinho aqui.  
**Luiz Felipe Verissimo da silva:** tipo de banco de dados você tá criando  
**lucas Badico:** É um C  
**Luiz Felipe Verissimo da silva:** lista  
**lucas Badico:** DB.  
**Luiz Felipe Verissimo da silva:** para mim  
**lucas Badico:** É um C DB.  
**Luiz Felipe Verissimo da silva:** relacional, não relaciono.  
**lucas Badico:** Ele é não relacional. Já vou entrar nesse ponto porque ele é o mais fácil pra gente já começar. Como a gente só tem meia hora, ele é o mais fácil pra gente já começar inserindo o dado nele, entendeu? OK. Então, eu subi o meu banco de dados. Que que eu vou precisar? Eu vou precisar criar um um banco aqui no nosso C. Ué, tá dando erro aqui. Será que deu? Deu é um put. Esse banco ele é ele é um banco bem diferentão que ele funciona via API também. Então eu vou criar um, eu vou criar uma, tô criando o que seria um uma tabela, né, que lá é uma coleção chamada notebook. Ó, ele criou, tá vendo? Tá funcionando. E ele não tem esquema, ele não tem um monte de coisa. Então, pra gente mostrar funcionando, pra gente mostrar é mais fácil.  
   
 

### 01:17:23

   
**lucas Badico:** Eh, aí se eu quiser, eu posso colocar um dado lá dentro. Notebook put. Eh, deixa eu só criar um ID aqui.  
**adolfo agnelli:** Угуm.  
**lucas Badico:** Esse banco é bem simples,  
**Luiz Felipe Verissimo da silva:** Eh  
**lucas Badico:** gente. Ele funciona do jeito que eu tô mostrando aqui para vocês, ó. Tá vendo? Ó, tá vendo? Ele gera um ID e uma revisão daquele, né? Ele aceita o ID que eu mando ele gera uma revisão daquele daquele documento. Então, só para mostrar para vocês como é que ele funciona, mas a gente vai instanciar isso no GO, tá bom? Cadê?  
**Luiz Felipe Verissimo da silva:** A revisão é uma cópia do negócio para manter uma rastreabilidade, é isso.  
**lucas Badico:** É para manter o histórico. De só ver uma paradinha aqui, porque na minha referência point que tá sessão do que só um minuto, gente, Ah, bom. O banco tá lá, a gente consegue criar documentos. Agora a gente vai trazer esse banco para cá, certo? Eh, que que a gente vai, que que a gente vai fazer? A gente vai precisar de dois pacotes, tá? Vamos voltar aqui pro pro nosso coisa.  
   
 

### 01:19:58

   
**lucas Badico:** A gente vai precisar de dois pacotes, certo? Vai precisar do Kivic V4 e do C DB V4. Tá bom. Eu vou eu vou colocar aqui pra gente ir rápido que a gente tem meia horinha. Eu quero aproveitar o máximo aqui pra gente fazer tudo que eu queria fazer hoje. Então ele tá instalando os dois pacotes pra gente, certo? a gente vai ter que vai, no caso, importar esses pacotes, certo? A gente vai precisar do contexto. Eu vou explicar contexto em outra aula, mas vocês vão ver isso aqui direto. Esse CTX, toda hora vocês vão ver o CTX, tá bom? E assim que a gente estancia o Kivic, né? Olha só. E tá aqui. OK. O o que que a gente fez aqui? Deixa eu fechar aqui o handler. Eu acho que eu preciso batermin pass local host. Deixa eu certinho. Isso mesmo. Tá certo. Vamos lá. O que que eu fiz aqui galera? Eu usei um pacote da comunidade que oferece pra gente um client do CDB, certo? Que é o Kivic.  
   
 

### 01:22:05

   
**lucas Badico:** Se vocês forem depois Kivic, Go Kivic, ele olha só. aqui, ó. Não. Então, que que é? Acho que só preciso de goft. A, esse aqui, ó, ele é um client, né? Ele é um client de de CD DB, que é a database que eu instalei aqui rapidinho, certo? Normalmente todos os bancos você vai ter esse essa esse setup que é você pedir um pint ou um pool, né? Quando a gente fala de mais ciclo ou de ciclo, a gente sempre costuma falar de pool, instanciar esse pool, né? Se der erro, você diz: "Olha, não posso iniciar o meu servidor". Por quê? Se você não tem banco de dados, você não tem servidor, né? A gente vai falar depois sobre sobre dados, mas se a gente não consegue guardar os dados em um lugar, não tem por a gente subir a nossa aplicação. Então, por isso que ao ter erro, né? E aí vocês vão ver isso aqui direto em GO. É o que o pessoal odeia em GO, eh, que é esse jeito de você verificar erro. Se o R tiverem de new, eu quero então que entre nesse if.  
   
 

### 01:23:51

   
**lucas Badico:** E nesse if aqui o caso é parar, tá bom? Então a gente tem esse client e com esse client a gente pode então criar objetos, né? Criar criar coisas. E aí é a mágica. Lembra que eu falei para você que a centia ter deixado aberto. A gente vai colocar o o client aqui dentro do do Hangler. Eh, então vai ter C de client Kivic D client. E aqui você vai ter eh C BB, que é um Kivic client. E eu acho que esse cara é uma referência, eu acho. A gente vai a gente vai chutar aqui. Então, a gente vai count DB igual a C, certo? Ah, a gente não tá usando isso, esse cara aqui. Acho que não. OK. Tão acompanhando o que eu tô fazendo?  
**Luiz Felipe Verissimo da silva:** M.  
**lucas Badico:** Ah, não, pera aí, faltou uma parada, desculpa. Ah, faltar uma paradinha aqui que eu preciso criar. Eu preciso dizer qual DB eu vou usar. OK. O nosso caso é o notebook safe aqui em cima. OK. Notebook. E ele devolve o que esse cara?  
   
 

### 01:26:20

   
**lucas Badico:** E aí isso aqui eu coloco aqui, só que aí vai quality type. Deixa eu ver. Pera aí, gente, que aí a gente entra na Ah, mas tá certo, ó.  
**Luiz Felipe Verissimo da silva:** O  
**lucas Badico:** Não, ele não tem não. Tá certo. Ah, ver aqui deu certo.  
**Luiz Felipe Verissimo da silva:** O  
**lucas Badico:** É, tá certo. Não tô vendo nada errado aqui não, mas tudo bem. Mas o meu ponto é qual que é o tipo do DB cl DB? Bom, posso ver? Ah, onde é que tá o cent? Aqui. KQDB aqui. DBG. Ah, é um é um DP. Se rendler eu tenho que mudar para KC DB. Pronto. OK. Eh, então o que que a gente fez? A gente tá que tá reclamando aqui. Ah, é porque é V4 aqui. OK, perfeito. Eh, canote. Sempre acontece isso, gente, só quando a gente tá com pressa. Ah, é porque ele ele não tem contexto, mas não precisa mais desse contexto. OK. Então é só assim.  
   
 

### 01:28:37

   
**lucas Badico:** Acho que agora vai rodar. Vamos ver se vai rodar esse bichinho. Sério? Não vai funcionar. Acho que não vai. É, não tá funcionando. É que eu precisoar não. Ele que tá, a gente não vai funcionar hoje. A gente vai ter que continuar a parte do banco na quinta-feira. Eu  
**Matheus Abranches:** Tranquilo.  
**lucas Badico:** vou vou ver corrigir esse erro e a gente como segue na quinta. Mas pelo menos eu cheguei nesse ponto aqui que vocês conseguem ver que agora, né, uma vez que funciona, a gente tem  
**Luiz Felipe Verissimo da silva:** Угуm.  
**lucas Badico:** o nosso banco aqui para poder nesses métodos interagir com o banco, correto? começa  
**Matheus Abranches:** Uhum.  
**lucas Badico:** a fazer sentido.  
**Matheus Abranches:** Sim.  
**adolfo agnelli:** M.  
**lucas Badico:** Boa, boa. Eh, refizeram então um pouquinho, a gente entendeu um pouquinho de APIs. A gente fez aí uma aplicaçãozinha básica, uma API super básica, começamos a definir a anatomia de como é que é uma aplicação GO, né? E vimos essa parte de tanto tipagem, né, como criar tipos e métodos, quanto o que que é exportável ou não, correto?  
   
 

### 01:30:27

   
**lucas Badico:** E vale vale lembrar uma paradinha aqui, ó. Eh, esse valor aqui dentro do handler, ele é privado porque ele começa com letra minúscula, percebe? Então eu só consigo manipular ele dentro desse pacote. Se eu quisesse que esse valor fosse exposto, eu teria de criar ele com letra maiúscula. Então, lembrem-se desse detalhe, tá bom?  
**Luiz Felipe Verissimo da silva:** Até mesmo  
**lucas Badico:** Eh,  
**Luiz Felipe Verissimo da silva:** dentro de um instruct.  
**lucas Badico:** isso, stretante, variáveis que você cria. Digamos que eu criei uma variável aqui, criei uma variável aqui, certo? Sei lá, eh, full. Quando é assim, eu preciso dizer qual que é o tipo string. Não tá lembro. Acho que é string. Criei full. Nesse momento ela é privada, ela não pode ser usada do lado de fora. Se eu colocar com letra mágica, ela pode ser usado lá de fora, entendeu? Então, tudo tudo tudo que é minúsculo é privado, tudo que é maiúsculo é é público. E aí métodos, por exemplo, se a gente tá aqui, a gente vai criar um método de criar de sei lá, precisa criar um método que a gente quer que seja privado, só se iniciar ele como minúsculo, porque ele só pode ser chamado aqui dentro, entendeu?  
   
 

### 01:32:02

   
**lucas Badico:** Tudo no Goa funciona assim, tá bom?  
**adolfo agnelli:** Sim.  
**lucas Badico:** Então, a gente viu essa parte, a gente falou sobre verbos, né? Verbos de HTTP. Eh, eu rodei aqui super rápido um servidorzinho, um servidorzinho, um banco de CD DB pra gente, eu queria que a gente criasse um documento, mas não vai rolar. Eh, e se e na quinta-feira a gente começa de fato a fazer o ciclo de uma aplicação. O que que é o ciclo de aplicação? Você vai instanciar seu banco, você vai, né, ter seu handler que vai usar o seu banco e com isso você vai interagir com esse banco, criar, deletar, certo? Editar. A gente vai começar a fazer isso, tá bom? Eh, certinhos? Curtiram a aula de hoje? Vamos parar um pouquinho aí, um pouquinho antes da das, mas é, teve tem problemas e é o mundo real. Então, quando eu descobrir o que que é o erro, eu eu mostro para vocês na aula que vem, tá bom?  
**adolfo agnelli:** Tranquilo.  
**lucas Badico:** Curtiram essa aula?  
**adolfo agnelli:** Tem ainda. Tá. Acho que dar uma revisada, entendeu?  
   
 

### 01:33:18

   
**adolfo agnelli:** alguns pontinhos ali, mas tá tranquilo ainda.  
**lucas Badico:** Boa  
**Luiz Felipe Verissimo da silva:** aparecendo um caminho legal. Eu eu já tava vendo ali que mesmo no handler de  
**adolfo agnelli:** Ah.  
**Luiz Felipe Verissimo da silva:** camada tem algumas coisas que você coloca como ponteiro para poder utilizar como referência. Eh, eu não  
**lucas Badico:** se  
**Luiz Felipe Verissimo da silva:** sei  
**lucas Badico:** a  
**Luiz Felipe Verissimo da silva:** se  
**lucas Badico:** gente  
**Luiz Felipe Verissimo da silva:** eu  
**lucas Badico:** vai  
**Luiz Felipe Verissimo da silva:** tô falando uma uma besteira, mas eu  
**lucas Badico:** Uhum.  
**Luiz Felipe Verissimo da silva:** tenho a impressão de que quando o dado é manipulável, se você quer manipular ele, você acaba sempre precisando utilizar ele como ponteiro dentro de G.  
**lucas Badico:** Não é essa razão, exatamente, tá? Eh, a gente depois vai entrar explicar certinho essa parte, essa parte de de custo, tá? Isso aí tem a ver com custo. Eh, mas a gente entra nisso em uma outra aula, tá bom?  
**Luiz Felipe Verissimo da silva:** Não tá  
**lucas Badico:** Deixa  
**Luiz Felipe Verissimo da silva:** beleza.  
**lucas Badico:** a gente a gente falar de modelagem de dado das itens mesmo. Aí a gente entra falar porque você quer as coisas como referência e as coisas como valor, tá bom? Mas aí a gente cai nisso, tá bom? Mas no início é é me estrelinha, pô. Porque algumas coisas tem o tem a estrelinha e todas as coisas que não tem estrelinha, entendeu? Eh, ficar a gente fica meio perdido no início, é normal. porque que você vai querer muito porque você vai querer outro, tá bom?  
**Matheus Abranches:** Beleza,  
**lucas Badico:** Gente,  
**adolfo agnelli:** Feou.  
**lucas Badico:** brigadão curti bastante essa segunda aula e a gente se vê na quinta-feira, beleza? tchau  
**adolfo agnelli:** Tranquilo,  
**Matheus Abranches:** show.  
**adolfo agnelli:** valeu.  
**Luiz Felipe Verissimo da silva:** Mas gente tchau.  
**lucas Badico:** tchau  
**Matheus Abranches:** tchau  
**sasa cruz:** Até quinta.  
   
 

### A transcrição foi encerrada após 01:35:01

*Esta transcrição editável foi gerada por computador e pode conter erros. As pessoas também podem alterar o texto depois que ele for criado.*