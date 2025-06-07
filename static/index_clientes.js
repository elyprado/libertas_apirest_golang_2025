function deletar(id) {
    if (confirm("Deseja realmente excluir este cliente?")) {
        fetch(`/api/clientes/${id}`, { method: "DELETE" }) 
            .then(response => {
                if (!response.ok) {
                    return response.text().then(text => {
                        throw new Error(`HTTP error! status: ${response.status} - ${text}`);
                    });
                }
                return response.text();
            })
            .then(() => {
                alert("Cliente excluído com sucesso!");
                listarClientes(); 
            })
            .catch(error => {
                console.error("Erro ao deletar cliente:", error); 
                alert(`Erro ao deletar o cliente: ${error.message}. Verifique o console.`); 
            });
    }
}

function listarClientes() { 
    const tabela = document.getElementById("tabela-clientes"); 
    if (!tabela) {
        console.error("Elemento 'tabela-clientes' não encontrado no DOM."); 
        return;
    }
    tabela.innerHTML = "<tr><td colspan='6' class='text-center'>Carregando clientes...</td></tr>";

    fetch("/api/clientes") 
        .then(response => {
            if (!response.ok) {
                return response.text().then(text => { throw new Error(`Erro ao buscar clientes: ${text}`); }); 
            }
            return response.json();
        })
        .then(dados => mostrarClientes(dados)) 
        .catch(error => {
            tabela.innerHTML = `<tr><td colspan='6' class='text-danger text-center'>Erro ao carregar clientes</td></tr>`; 
            console.error("Erro ao listar clientes:", error); 
            alert(`Erro ao listar clientes: ${error.message}`); 
        });
}

function mostrarClientes(dados) { 
    const tabela = document.getElementById("tabela-clientes"); 
    if (!tabela) {
        console.error("Elemento 'tabela-clientes' não encontrado no DOM."); 
        return;
    }
    tabela.innerHTML = "";

    if (dados.length === 0) {
        tabela.innerHTML = "<tr><td colspan='6' class='text-center'>Nenhum cliente cadastrado.</td></tr>"; 
        return;
    }

    for (let i = 0; i < dados.length; i++) {
        const tr = document.createElement("tr");
        const cliente = dados[i]; 

        tr.innerHTML = `
            <td>${cliente.id}</td>
            <td>${cliente.nome}</td>
            <td>${cliente.email || 'N/A'}</td>
            <td>${cliente.telefone || 'N/A'}</td>
            <td>${cliente.endereco || 'N/A'}</td>
            <td>
                <a href="/static/cadastro_clientes.html?id=${cliente.id}" class="btn btn-primary btn-sm me-2">Editar</a> <button class="btn btn-danger btn-sm" onclick='deletar(${cliente.id})'>Excluir</button>
            </td>
        `;
        tabela.appendChild(tr);
    }
}

document.addEventListener('DOMContentLoaded', listarClientes); 