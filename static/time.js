const modalCadastro = new bootstrap.Modal(document.getElementById('modalcadastro'));

function mostrarCarregamento() {
    const spinner = document.getElementById("loadingSpinner");
    spinner.classList.remove("d-none");
}

function esconderCarregamento() {
    const spinner = document.getElementById("loadingSpinner");
    spinner.classList.add("d-none");
}

function novo() {
    document.getElementById("nome").value = "";
    document.getElementById("cidade").value = ""; 
    document.getElementById("estado").value = "";
    document.getElementById("fundacao").value = "";
    document.getElementById("estadio").value = "";
    document.getElementById("idtime").value = "";
    modalCadastro.show();
}

function salvar() {
    const nome = document.getElementById("nome").value.trim();
    const cidade = document.getElementById("cidade").value.trim();
    const estado = document.getElementById("estado").value.trim();
    const fundacao = document.getElementById("fundacao").value.trim();
    const estadio = document.getElementById("estadio").value.trim();
    const idTime = document.getElementById("idtime").value.trim();

    if (!nome || !cidade || !estado || !fundacao || !estadio) {
        alert("Preencha todos os campos!");
        return;
    }

    const time = {
        idtime: (idTime == "" ? null : parseInt(fundacao)),
        nome: (nome == "" ? null : nome),
        cidade: (cidade == "" ? null : cidade),
        estado: (estado == "" ? null : estado),
        fundacao: (fundacao == "" || isNaN(parseInt(fundacao)) ? null : parseInt(fundacao)),
        estadio: (estadio == "" ? null : estadio)
    };

    let url = "http://127.0.0.1:8080/time";
    let metodo = "POST";

    if (idTime) {
        console.log("idTime", idTime);
        url += "/" + idTime;
        metodo = "PUT";
    }
    console.log(idTime, url, metodo);
    mostrarCarregamento();
    fetch(url, {
        method: metodo,
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(time)
    })
    .then(resposta => {
        if (!resposta.ok) {
            throw new Error("Erro ao salvar time");
        }
        return resposta.text();  
    })
    .then(() => {
        listar();
        const modal = bootstrap.Modal.getInstance(document.getElementById('modalcadastro'));
        modal.hide();  
    })
    .catch(erro => {
        console.error("Erro ao salvar time:", erro);
    });
}

function editar(id) {    
    mostrarCarregamento();
    fetch("http://127.0.0.1:8080/time/" + id, {
        method: "GET"})
    .then(resposta => {
        console.log(resposta);
        if (!resposta.ok) {
            throw new Error("Erro ao buscar time");
        }
        return resposta.json();
    })
    .then(dados => {
        document.getElementById("nome").value = dados.nome;
        document.getElementById("cidade").value = dados.cidade;
        document.getElementById("estado").value = dados.estado;
        document.getElementById("fundacao").value = dados.fundacao;
        document.getElementById("estadio").value = dados.estadio;
        document.getElementById("idtime").value = dados.idtime;
        modalCadastro.show();
    })
    .catch(erro => {
        console.error("Erro ao buscar time:", erro);
    })
    .finally(() => {
        esconderCarregamento();
    });
}

function listar(pesquisar) {
    mostrarCarregamento();
    const lista = document.getElementById("lista");
    let rota = "http://127.0.0.1:8080/time";
    const params = new URLSearchParams();

    if (pesquisar) {
        const pesquisa = document.getElementById("pesquisa").value.trim();
        if (pesquisa != null
            && pesquisa != undefined
            && pesquisa != ""
        ) {
            params.append("nome", pesquisa); 
            params.append("cidade", pesquisa);
            params.append("estado", pesquisa);
            params.append("fundacao", pesquisa);
            params.append("estadio", pesquisa);
        }
    }

    if ([...params].length > 0) {
        rota += "?" + params.toString();
    }
    console.log("Rota:", rota);

    fetch(rota) 
        .then(resposta => {
            if (!resposta.ok) {
                throw new Error("Erro ao buscar time");
            }
            return resposta.json();
        })
        .then(dados => mostrar(dados))
        .catch(erro => {
            lista.innerHTML = `<tr><td colspan='12' class='text-danger text-center'>Nenhum time encontrado</td></tr>`;
            console.error("Erro ao listar times:", erro);
            esconderCarregamento();
        });
}


function excluir(id) {
    mostrarCarregamento();
    fetch("http://127.0.0.1:8080/time/" + id, {
        method: "DELETE"
    })
    .then(resposta => {
        if (!resposta.ok) {
            throw new Error("Erro ao excluir time");
        }
        return resposta.text();
    })
    .then(() => {
        listar();
    })
    .catch(erro => {
        console.error("Erro ao excluir time:", erro);
    });
}

function mostrar(dados) {
    const lista = document.getElementById("lista");
    lista.innerHTML = "";
    if(dados == null || dados.length == 0) {
        lista.innerHTML = `<tr><td colspan='12' class='text-danger text-center'>Nenhum time encontrado</td></tr>`;
        esconderCarregamento();
        return;
    }
    for (let i = 0; i < dados.length; i++) {
        const tr = document.createElement("tr");
        tr.innerHTML = `
            <td>${dados[i].idtime}</td>
            <td>${dados[i].nome}</td>
            <td>${dados[i].cidade}</td>
            <td>${dados[i].estado}</td>
            <td>${dados[i].fundacao}</td>
            <td>${dados[i].estadio}</td>
            <td>
                <button onclick='editar(${dados[i].idtime})'><img src='imgs/edit.svg'></button>
                <button onclick='excluir(${dados[i].idtime})'><img src='imgs/x-square.svg'></button>
            </td>`;
        lista.appendChild(tr);
    }
    esconderCarregamento();
}

listar(false);