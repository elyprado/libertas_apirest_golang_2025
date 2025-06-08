const modalcadastro = new bootstrap.Modal(document.getElementById('modalcadastro'))

var matriculaAtual;

function alterar(matricula) {
    matriculaAtual = matricula;
    fetch("http://127.0.0.1:8080/alunos/" + matricula)
        .then(resp => resp.json())
        .then(dados => {
            document.getElementById("nome").value = dados.nome;
            document.getElementById("idade").value = dados.idade;
            document.getElementById("curso").value = dados.curso;
            document.getElementById("matricula").value = dados.matricula;
            document.getElementById("email").value = dados.email;
            modalcadastro.show();
        });
}

function excluir(matricula) {
    fetch("http://127.0.0.1:8080/alunos/" + matricula, {
        method: "DELETE"
    }).then(function () {
        listar();
    });
}

function salvar() {
    let vnome = document.getElementById("nome").value;
    let vidade = document.getElementById("idade").value;
    let vcurso = document.getElementById("curso").value;
    let vmatricula = document.getElementById("matricula").value;
    let vemail = document.getElementById("email").value;

    let aluno = {
        nome: vnome,
        idade: parseInt(vidade),
        curso: vcurso,
        matricula: vmatricula,
        email: vemail
    }

    let url;
    let metodo;
    if (matriculaAtual) {
        url = "http://127.0.0.1:8080/alunos/" + matriculaAtual;
        metodo = "PUT";
    } else {
        url = "http://127.0.0.1:8080/alunos";
        metodo = "POST";
    }

    fetch(url, {
        method: metodo,
        body: JSON.stringify(aluno),
        headers: {
            "Content-Type": "application/json"
        }
    }).then(function () {
        listar();
        modalcadastro.hide();
    })
}

function novo() {
    matriculaAtual = null;
    document.getElementById("nome").value = "";
    document.getElementById("idade").value = "";
    document.getElementById("curso").value = "";
    document.getElementById("matricula").value = "";
    document.getElementById("email").value = "";
    modalcadastro.show();
}

function listar() {
    const lista = document.getElementById("lista");
    lista.innerHTML = "<tr><td colspan='6'>Carregando...</td></tr>";

    fetch("http://127.0.0.1:8080/alunos")
        .then(resp => resp.json())
        .then(dados => mostrar(dados));
}

function mostrar(dados) {
    const lista = document.getElementById("lista");
    lista.innerHTML = "";
    for (let i in dados) {
        lista.innerHTML += "<tr>"
            + "<td>" + dados[i].nome + "</td>"
            + "<td>" + dados[i].idade + "</td>"
            + "<td>" + dados[i].curso + "</td>"
            + "<td>" + dados[i].matricula + "</td>"
            + "<td>" + dados[i].email + "</td>"
            + "<td>"
            + "<button type='button' class='btn btn-primary' onclick='alterar(\"" + dados[i].matricula + "\")'>A</button> "
            + "<button type='button' class='btn btn-danger' onclick='excluir(\"" + dados[i].matricula + "\")'>X</button>"
            + "</td>"
            + "</tr>";
    }
}
