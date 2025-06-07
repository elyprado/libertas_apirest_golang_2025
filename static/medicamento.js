var modalCadastro = new bootstrap.Modal(document.getElementById('modalCadastro'));
var modalExcluir = new bootstrap.Modal(document.getElementById('excluirCadastro'));
var modalToast = new bootstrap.Toast(document.getElementById('toast'));
var loading = document.querySelector('#loading')

function init() {
    document.querySelector('#cancelarCadastro').onclick = function () {
        modalCadastro.hide()
    }

    document.querySelector('#naoExcluir').onclick = function () {
        modalExcluir.hide()
    }

    document.querySelector('#imgPesquisa').onclick = function () {
        pesquisar()
    };

    document.querySelector('#inputPesquisar').addEventListener('keydown', function (event) {
        if (event.key === 'Enter') {
            document.querySelector('#inputPesquisar').blur()

            if (document.querySelector('#imgPesquisa').classList.contains('ph-x')) {
                getCadastros(document.querySelector('#inputPesquisar').value.trim())
            } else {
                pesquisar();

            }
        }
    });

    getCadastros()
}

// ################### UTILS ###################

function abrirLoading() {
    loading.style.display = "flex"
}

function fecharLoading() {
    loading.style.display = "none"
}

function toast(mensagem, success) {
    //document.querySelector('.toast-body').innerHTML = mensagem
    //document.querySelector('.toast').style.backgroundColor = (success ? 'var(--primaria-700)' : 'var(--error)');
    window.alert(mensagem);
    //modalToast.show()
}

function validarInputs(...inputs) {
    let temErro = false;

    inputs.forEach(input => {
        if (input.value.trim() === "") {
            input.classList.add('inputError');
            temErro = true;
        } else {
            input.classList.remove('inputError');
        }
    });

    return !temErro;
}



function getCadastros(nome) {
    abrirLoading()
    var parametro = (nome == undefined ? "" : "?nome=" + nome)

    fetch('http://127.0.0.1:8080/medicamentos' + parametro,
        {
            method: "GET",
            headers: { 'Content-Type': 'application/json' },
        })
        .then(resp => resp.json())
        .then(dados => listarCadastros(dados))
        .catch(err => console.error("Erro ao buscar dados:", err))
        .finally(function () {
            fecharLoading()
        });
}

function getCadastro(id) {
    abrirLoading()
    fetch("http://127.0.0.1:8080/medicamentos/" + id,
        {
            method: "GET",
            headers: { 'Content-Type': 'application/json' },
        })
        .then(resp => resp.json())
        .then(dados => editarCadastro(dados))
        .catch(err => console.error("Erro ao buscar dados:", err))
        .finally(function () {
            fecharLoading()
        });
}

function deleteCadastro(id) {
    abrirLoading()
    fetch("http://127.0.0.1:8080/medicamentos/" + id,
        {
            method: "DELETE",
            headers: { 'Content-Type': 'application/json' },
        })
        .then(resp => resp.json())
        .then(function (retorno) {
            toast(retorno.message, retorno.success);
        }).finally(function () {
            getCadastros();
        });
}

function postPutCadastro(obj, id) {
    abrirLoading()
    let body = JSON.stringify(obj);
    let endereco = "http://127.0.0.1:8080/medicamentos";
    let metodo = "POST";
    if (id) {
        endereco = "http://127.0.0.1:8080/medicamentos/" + id;
        metodo = "PUT";
    }
    fetch(endereco,
        {
            method: metodo,
            body: body,
            headers: { 'Content-Type': 'application/json' },
        })
        .then(resp => resp.json())
        .then(function (retorno) {
            toast(retorno.message, retorno.success);
        })
        .catch(err => console.error("Erro ao buscar dados:", err))
        .finally(function () {
            modalCadastro.hide()
            getCadastros();
        });
}


// ################### MODAL ###################

function novoCadastro() {
    // alterar titulo
    document.querySelector('#tituloModal').innerHTML = "Novo medicamento";
    // limpar campos
    document.querySelector('#inputnome').value = '';
    document.querySelector('#inputquantidade').value = '';
    document.querySelector('#inputtipo').value = '';
    document.querySelector('#inputfabricante').value = '';
    // remover erros
    document.querySelector('#inputnome').classList.remove('inputError');
    document.querySelector('#inputquantidade').classList.remove('inputError');
    document.querySelector('#inputtipo').classList.remove('inputError');
    document.querySelector('#inputfabricante').classList.remove('inputError');
    // botao salvar
    document.querySelector('#salvarCadastro').onclick = function () {
        salvar()
    }
    // abrir modal
    modalCadastro.show();
}

function editarCadastro(dados) {
    //alterar titulo
    document.querySelector('#tituloModal').innerHTML = "Alterar medicamento";
    // preencher campos
    document.querySelector('#inputnome').value = dados.nome;
    document.querySelector('#inputquantidade').value = (dados.quantidade);
    document.querySelector('#inputtipo').value = dados.tipo;
    document.querySelector('#inputfabricante').value = dados.fabricante;
    // remover erros
    document.querySelector('#inputnome').classList.remove('inputError');
    document.querySelector('#inputquantidade').classList.remove('inputError');
    document.querySelector('#inputtipo').classList.remove('inputError');
    document.querySelector('#inputfabricante').classList.remove('inputError');
    // botao salvar
    document.querySelector('#salvarCadastro').onclick = function () {
        salvar(dados.idmedicamento)
    }
    // abrir modal
    modalCadastro.show();
}

function excluir(id) {
    modalExcluir.show()
    document.querySelector('#simExcluir').onclick = function () {
        deleteCadastro(id)
        modalExcluir.hide()
    }
}

function salvar(id) {
    console.log(id)
    const inputnome = document.querySelector('#inputnome');
    const inputquantidade = document.querySelector('#inputquantidade');
    const inputtipo = document.querySelector('#inputtipo');
    const inputfabricante = document.querySelector('#inputfabricante');

    let cadastro = {};
    cadastro.nome = inputnome.value.trim();
    cadastro.quantidade = inputquantidade.value.trim();
    cadastro.tipo = inputtipo.value.trim();
    cadastro.fabricante = inputfabricante.value.trim();

    if (validarInputs(inputnome, inputquantidade, inputtipo, inputfabricante)) {
        postPutCadastro(cadastro, id)
    } else {
        toast('Preencha todos os campos!', false)
    }
}
// ################### TABELA ###################

function listarCadastros(dados) {
    if (dados == null) {
        document.querySelector('.seminfo').style.display = "flex"
        document.querySelector('.cardTabela').style.display = "none"
    } else {
        document.querySelector('.seminfo').style.display = "none"
        document.querySelector('.cardTabela').style.display = ""

        var tab = '';
        for (let i in dados) {
            tab += "<tr>"
                + "<td class='colunaCentro'>" + dados[i].idmedicamento + "</td>"
                + "<td>" + (dados[i].nome == null ? "-" : dados[i].nome) + "</td>"
                + "<td class='colunaCentro'>" + (dados[i].quantidade == null ? '-' : dados[i].quantidade) + "</td>"
                + "<td>" + (dados[i].tipo == null ? "-" : dados[i].tipo) + "</td>"
                + "<td>" + (dados[i].fabricante == null ? "-" : dados[i].fabricante) + "</td>"
                + `<td class="colunaIcone"><div id='btnEdit' title="Alterar" onclick='getCadastro(${dados[i].idmedicamento})'><i class="ph ph-pencil-simple"></i></div></td>`
            + `<td class="colunaIcone"><div id='btnDelete' title="Excluir" onclick='excluir(${dados[i].idmedicamento})'><i class="ph ph-trash"></i></div></td>`
                + "</tr>";
        }

        document.querySelector("#lista").innerHTML = tab;
    }

}

function pesquisar() {
    var icone = document.querySelector('#imgPesquisa')
    var nome = document.querySelector('#inputPesquisar');
    if (icone.classList.contains('ph-magnifying-glass')) {
        icone.classList.remove('ph-magnifying-glass');
        icone.classList.add('ph-x');
        getCadastros(nome.value)
    } else if (icone.classList.contains('ph-x')) {
        icone.classList.remove('ph-x');
        icone.classList.add('ph-magnifying-glass');
    
        nome.value = ''
        getCadastros(null)
    }
}

document.addEventListener('DOMContentLoaded', init);