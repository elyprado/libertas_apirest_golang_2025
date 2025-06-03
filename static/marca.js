

const modalcadastro = new bootstrap.Modal(document.getElementById('modalcadastro'))

var idmarcaatual;

function alterar(idmarca) {
    //implemente o método fetch, buscando os dados com idusuario
    //preencha o resultados nos 3 inputs e abra o modal para edição
    idmarcaatual = idmarca;
    fetch("http://127.0.0.1:8080/marcas/"+idmarca)
         .then(resp => resp.json())
         .then(dados => {
            document.getElementById("nome").value = dados.nome;
            document.getElementById("nicho").value = dados.nicho;
            document.getElementById("cnpj").value = dados.cnpj;
            document.getElementById("site").value = dados.site;
            modalcadastro.show();
         });
}
function excluir(idmarca) {
    fetch("http://127.0.0.1:8080/marcas/"+idmarca,
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
    let vnicho = document.getElementById("nicho").value;
    let vcnpj = document.getElementById("cnpj").value;
    let vsite = document.getElementById("site").value;

    let marca = {
        nome: vnome, nicho: vnicho, cnpj: vcnpj, site: vsite
    }

    let url;
    let metodo;
    if (idmarcaatual>0) {
        //alterar
        url = "http://127.0.0.1:8080/marcas/"+idmarcaatual;
        metodo = "PUT";
    } else {
        //inserir
        url = "http://127.0.0.1:8080/marcas";
        metodo = "POST";
    }

    fetch(url,
        {
            method: metodo,
            body: JSON.stringify(marca),
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
    idmarcaatual = 0;
    document.getElementById("nome").value = "";
    document.getElementById("nicho").value = "";
    document.getElementById("cnpj").value = "";
    document.getElementById("site").value = "";
    modalcadastro.show();
}

function listar() {
    const lista = document.getElementById("lista");
    lista.innerHTML = "<tr><td colspan='5'>Carregando...</td></tr>";

    fetch("http://127.0.0.1:8080/marcas")
         .then(resp => resp.json())
         .then(dados => mostrar(dados));
}
function mostrar(dados) {
    const lista = document.getElementById("lista");
    lista.innerHTML = "";
    for (let i in dados) {
        lista.innerHTML += "<tr>" 
                        + "<td>" + dados[i].idmarca + "</td>"
                        + "<td>" + dados[i].nome + "</td>"
                        + "<td>" + dados[i].nicho + "</td>"
                        + "<td>" + dados[i].cnpj + "</td>"
                        + "<td>" + dados[i].site + "</td>"
                        + "<td>"
+ "<button type='button' class='btn btn-primary' onclick='alterar("+dados[i].idmarca+")'>A</button>"
+ "<button type='button' class='btn btn-danger' onclick='excluir("+dados[i].idmarca+")'>X</button>"
                        + "</td>"
                        + "</tr>";
    }
}