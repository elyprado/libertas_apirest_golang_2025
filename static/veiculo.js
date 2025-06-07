

const modalcadastro = new bootstrap.Modal(document.getElementById('modalcadastro'))

var idveiculoatual;

function alterar(idveiculo) {
    //implemente o método fetch, buscando os dados com idusuario
    //preencha o resultados nos inputs e abra o modal para edição
    idveiculoatual = idveiculo;
    fetch("http://127.0.0.1:8080/veiculos/"+idveiculo)
         .then(resp => resp.json())
         .then(dados => {
            document.getElementById("modelo").value = dados.modelo;
            document.getElementById("marca").value = dados.marca;
            document.getElementById("ano").value = dados.ano;
            document.getElementById("cor").value = dados.cor;
            document.getElementById("preco").value = dados.preco;
            modalcadastro.show();
         });
}
function excluir(idveiculo) {
    fetch("http://127.0.0.1:8080/veiculos/"+idveiculo,
        {
            method: "DELETE"
        } 
    ).then(function () {
        //recarrega a lista
        listar();
    });
}

function salvar() {
    let vmodelo = document.getElementById("modelo").value;
    let vmarca = document.getElementById("marca").value;
    let vano = document.getElementById("ano").value;
    let vcor = document.getElementById("cor").value;
    let vpreco = document.getElementById("preco").value;

    let marca = {
        modelo: vmodelo, marca: vmarca, ano: vano, cor: vcor, preco: vpreco
    }

    let url;
    let metodo;
    if (idveiculoatual>0) {
        //alterar
        url = "http://127.0.0.1:8080/veiculos/"+idveiculoatual;
        metodo = "PUT";
    } else {
        //inserir
        url = "http://127.0.0.1:8080/veiculos";
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
    idveiculoatual = 0;
    document.getElementById("modelo").value = "";
    document.getElementById("marca").value = "";
    document.getElementById("ano").value = "";
    document.getElementById("cor").value = "";
    document.getElementById("preco").value = "";
    modalcadastro.show();
}

function listar() {
    const lista = document.getElementById("lista");
    lista.innerHTML = "<tr><td colspan='5'>Carregando...</td></tr>";

    fetch("http://127.0.0.1:8080/veiculos")
         .then(resp => resp.json())
         .then(dados => mostrar(dados));
}
function mostrar(dados) {
    const lista = document.getElementById("lista");
    lista.innerHTML = "";
    for (let i in dados) {
        lista.innerHTML += "<tr>" 
                        + "<td>" + dados[i].idveiculo + "</td>"
                        + "<td>" + dados[i].modelo + "</td>"
                        + "<td>" + dados[i].marca + "</td>"
                        + "<td>" + dados[i].ano + "</td>"
                        + "<td>" + dados[i].cor + "</td>"
                        + "<td>" + dados[i].preco + "</td>"
                        + "<td>"
+ "<button type='button' class='btn btn-primary' onclick='alterar("+dados[i].idveiculo+")'>A</button>"
+ "<button type='button' class='btn btn-danger' onclick='excluir("+dados[i].idveiculo+")'>X</button>"
                        + "</td>"
                        + "</tr>";
    }
}