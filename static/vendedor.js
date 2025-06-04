

const modalcadastro = new bootstrap.Modal(document.getElementById('modalcadastro'))

var idvendedoratual;

function alterar(idvendedor) {
    //implemente o método fetch, buscando os dados com idvendedor
    //preencha o resultados nos 3 inputs e abra o modal para edição
    idvendedoratual = idvendedor;
    fetch("http://127.0.0.1:8080/vendedores/"+idvendedor)
         .then(resp => resp.json())
         .then(dados => {
            document.getElementById("nome").value = dados.nome;
            document.getElementById("email").value = dados.email;
            document.getElementById("cpf").value = dados.cpf;
            document.getElementById("telefone").value = dados.telefone;
            document.getElementById("estado").value = dados.estado;
            modalcadastro.show();
         });
}
function excluir(idvendedor) {
    fetch("http://127.0.0.1:8080/vendedores/"+idvendedor,
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
    let vemail = document.getElementById("email").value;
    let vcpf = document.getElementById("cpf").value;
    let vtelefone = document.getElementById("telefone").value;
    let vestado = document.getElementById("estado").value;

    let vendedor = {
        nome: vnome, email: vemail, cpf:vcpf, telefone: vtelefone, estado: vestado, senha: ''
    }

    let url;
    let metodo;
    if (idvendedoratual>0) {
        //alterar
        url = "http://127.0.0.1:8080/vendedores/"+idvendedoratual;
        metodo = "PUT";
    } else {
        //inserir
        url = "http://127.0.0.1:8080/vendedores";
        metodo = "POST";
    }

    fetch(url,
        {
            method: metodo,
            body: JSON.stringify(vendedor),
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
    idvendedoratual = 0;
    document.getElementById("nome").value = "";
    document.getElementById("email").value = "";
    document.getElementById("cpf").value = "";
    document.getElementById("telefone").value = "";
    document.getElementById("estado").value = "";
    modalcadastro.show();
}

function listar() {
    const lista = document.getElementById("lista");
    lista.innerHTML = "<tr><td colspan='5'>Carregando...</td></tr>";

    fetch("http://127.0.0.1:8080/vendedores")
         .then(resp => resp.json())
         .then(dados => mostrar(dados));
}
function mostrar(dados) {
    const lista = document.getElementById("lista");
    lista.innerHTML = "";
    for (let i in dados) {
        lista.innerHTML += "<tr>" 
                        + "<td>" + dados[i].idvendedor + "</td>"
                        + "<td>" + dados[i].nome + "</td>"
                        + "<td>" + dados[i].email + "</td>"
                        + "<td>" + dados[i].cpf + "</td>"
                        + "<td>" + dados[i].telefone + "</td>"
                        + "<td>" + dados[i].estado + "</td>"
                        + "<td>"
+ "<button type='button' class='btn btn-primary' onclick='alterar("+dados[i].idvendedor+")'>A</button>"
+ "<button type='button' class='btn btn-danger' onclick='excluir("+dados[i].idvendedor+")'>X</button>"
                        + "</td>"
                        + "</tr>";
    }
}