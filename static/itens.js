const modalcadastro = new bootstrap.Modal(document.getElementById('modalcadastro'));

var iditematual;

function alterar(iditem) {
    
    iditematual = iditem;
    fetch("http://127.0.0.1:8080/itens/" + iditem)
        .then(resp => resp.json())
        .then(dados => {
            document.getElementById("nome").value = dados.nome;
            document.getElementById("marca").value = dados.marca;
            document.getElementById("quantidade").value = dados.quantidade;
            document.getElementById("preco").value = dados.preco;
            document.getElementById("observacao").value = dados.observacao;
            modalcadastro.show();
        });
}

function excluir(iditem) {
    fetch("http://127.0.0.1:8080/itens/" + iditem, {
        method: "DELETE"
    }).then(function () {
        listar();
    });
}

function salvar() {
    let vnome = document.getElementById("nome").value;
    let vmarca = document.getElementById("marca").value;
    let vquantidade = parseInt(document.getElementById("quantidade").value);
    let vpreco = parseFloat(document.getElementById("preco").value);
    let vobservacao = document.getElementById("observacao").value;

    let item = {
        nome: vnome,
        marca: vmarca,
        quantidade: vquantidade,
        preco: vpreco,
        observacao: vobservacao
    };

    let url;
    let metodo;
    if (iditematual > 0) {
        // Alterar
        url = "http://127.0.0.1:8080/itens/" + iditematual;
        metodo = "PUT";
    } else {
        // Inserir
        url = "http://127.0.0.1:8080/itens";
        metodo = "POST";
    }

    fetch(url, {
        method: metodo,
        body: JSON.stringify(item),
        headers: {
            "Content-Type": "application/json"
        }
    }).then(function () {
        listar();
        modalcadastro.hide();
    });
}

function novo() {
    iditematual = 0;
    document.getElementById("nome").value = "";
    document.getElementById("marca").value = "";
    document.getElementById("quantidade").value = "";
    document.getElementById("preco").value = "";
    document.getElementById("observacao").value = "";
    modalcadastro.show();
}

function listar() {
    const lista = document.getElementById("lista");
    lista.innerHTML = "<tr><td colspan='7'>Carregando...</td></tr>";

    fetch("http://127.0.0.1:8080/itens")
        .then(resp => resp.json())
        .then(dados => mostrar(dados));
}

function mostrar(dados) {
    const lista = document.getElementById("lista");
    lista.innerHTML = "";
    for (let i in dados) {
        lista.innerHTML += "<tr>" +
            "<td>" + dados[i].id + "</td>" +
            "<td>" + dados[i].nome + "</td>" +
            "<td>" + dados[i].marca + "</td>" +
            "<td>" + dados[i].quantidade + "</td>" +
            "<td>" + dados[i].preco.toFixed(2) + "</td>" +
            "<td>" + dados[i].observacao + "</td>" +
            "<td>" +
            "<button type='button' class='btn btn-primary' onclick='alterar(" + dados[i].id + ")'>A</button>" +
            "<button type='button' class='btn btn-danger' onclick='excluir(" + dados[i].id + ")'>X</button>" +
            "</td>" +
            "</tr>";
    }
}
