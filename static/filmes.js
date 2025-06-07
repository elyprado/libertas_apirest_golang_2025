


const modalcadastro = new bootstrap.Modal(document.getElementById('modalcadastro'))

var idfilmeatual = 0;

function alterar(idfilme) {
    //implemente o método fetch, buscando os dados com idusuario
    //preencha o resultados nos 3 inputs e abra o modal para edição
    idfilmeatual = idfilme;
    fetch("http://127.0.0.1:8080/filmes/"+idfilme)
         .then(resp => resp.json())
         .then(dados => {
            document.getElementById("nome").value = dados.nome;
            document.getElementById("classificacao").value = dados.classificacao;
            document.getElementById("genero").value = dados.genero;
            document.getElementById("ano").value = dados.ano;
            document.getElementById("autor").value = dados.autor;
            modalcadastro.show();
         });
}
function excluir(idfilme) {
    fetch("http://127.0.0.1:8080/filmes/"+idfilme,
        {
            method: "DELETE"
        } 
    ).then(function () {
        //recarrega a lista
        listar();
    });
}

function salvar() {
    let vnome = document.getElementById("nome").value;
    let vclassificacao = document.getElementById("classificacao").value;
    let vgenero = document.getElementById("genero").value;
    let vautor = document.getElementById("autor").value;
    let vano = document.getElementById("ano").value;

    let filme = {
        nome: vnome, classificacao: vclassificacao, genero: vgenero, autor: vautor, ano: vano
    }

    let url;
    let metodo;
    if (idfilmeatual>0) {
        //alterar
        url = "http://127.0.0.1:8080/filmes/"+idfilmeatual;
        metodo = "PUT";
    } else {
        //inserir
        url = "http://127.0.0.1:8080/filmes";
        metodo = "POST";
    }

    fetch(url,
        {
            method: metodo,
            body: JSON.stringify(filme),
            headers: {
                "Content-Type" : "application/json"
            }
        }
    ).then(function () {
        //recarrega a lista
        listar();
        //esconde o modal
        modalcadastro.hide();
    })

}

function novo() {
    idfilmeatual = 0;
    document.getElementById("nome").value = "";
    document.getElementById("classificacao").value = "";
    document.getElementById("genero").value = "";
    document.getElementById("ano").value = "";
    document.getElementById("autor").value = "";
    modalcadastro.show();
}

function listar() {
    const lista = document.getElementById("lista");
    lista.innerHTML = "<tr><td colspan='5'>Carregando...</td></tr>";
    
    fetch("http://127.0.0.1:8080/filmes")
         .then(resp => resp.json())
         .then(dados => mostrar(dados));
}
function mostrar(dados) {
    const lista = document.getElementById("lista");
    lista.innerHTML = "";
    for (let i in dados) {
        lista.innerHTML += "<tr>" 
                        + "<td>" + dados[i].id + "</td>"
                        + "<td>" + dados[i].nome + "</td>"
                        + "<td>" + dados[i].classificacao + "</td>"
                        + "<td>" + dados[i].genero + "</td>"
                        + "<td>" + dados[i].ano + "</td>"
                        + "<td>" + dados[i].autor + "</td>"
                        + "<td>"
+ "<button type='button' class='btn btn-primary' onclick='alterar("+dados[i].id+")'>A</button>"
+ "<button type='button' class='btn btn-danger' onclick='excluir("+dados[i].id+")'>X</button>"
                        + "</td>"
                        + "</tr>";
    }
}
