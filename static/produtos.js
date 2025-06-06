const modalcadastro = new bootstrap.Modal(document.getElementById('modalcadastro'));

var idProdutoAtual;

function alterar(id) {
    idProdutoAtual = id;
    fetch("http://127.0.0.1:8080/produtos/" + id)
        .then(resp => resp.json())
        .then(dados => {
            document.getElementById("nome").value = dados.descricao;
            document.getElementById("marca").value = dados.fornecedor;
            document.getElementById("quantidade").value = dados.estoque;
            document.getElementById("preco").value = dados.valor;
            document.getElementById("observacao").value = dados.detalhes;
            modalcadastro.show();
        });
}

function excluir(id) {
    fetch("http://127.0.0.1:8080/produtos/" + id, {
        method: "DELETE"
    }).then(() => {
        listar();
    });
}

function salvar() {
    let produto = {
        descricao: document.getElementById("nome").value,
        fornecedor: document.getElementById("marca").value,
        estoque: parseInt(document.getElementById("quantidade").value),
        valor: parseFloat(document.getElementById("preco").value),
        detalhes: document.getElementById("observacao").value
    };

    let url, metodo;
    if (idProdutoAtual > 0) {
        url = "http://127.0.0.1:8080/produtos/" + idProdutoAtual;
        metodo = "PUT";
    } else {
        url = "http://127.0.0.1:8080/produtos";
        metodo = "POST";
    }

    fetch(url, {
        method: metodo,
        body: JSON.stringify(produto),
        headers: {
            "Content-Type": "application/json"
        }
    }).then(() => {
        listar();
        modalcadastro.hide();
    });
}

function novo() {
    idProdutoAtual = 0;
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

    fetch("http://127.0.0.1:8080/produtos")
        .then(resp => resp.json())
        .then(dados => mostrar(dados));
}

function mostrar(dados) {
    const lista = document.getElementById("lista");
    lista.innerHTML = "";
    for (let i in dados) {
        lista.innerHTML += `
            <tr>
                <td>${dados[i].id}</td>
                <td>${dados[i].descricao}</td>
                <td>${dados[i].fornecedor}</td>
                <td>${dados[i].estoque}</td>
                <td>${dados[i].valor.toFixed(2)}</td>
                <td>${dados[i].detalhes}</td>
                <td>
                    <button type='button' class='btn btn-primary' onclick='alterar(${dados[i].id})'>✏️</button>
                    <button type='button' class='btn btn-danger' onclick='excluir(${dados[i].id})'>🗑️</button>
                </td>
            </tr>
        `;
    }
}
