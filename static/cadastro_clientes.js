document.addEventListener('DOMContentLoaded', () => {
    const urlParams = new URLSearchParams(window.location.search);
    const clienteId = urlParams.get('id'); 

    const formCliente = document.getElementById('formCliente'); 
    if (!formCliente) {
        console.error("Formulário 'formCliente' não encontrado."); 
        return;
    }

    if (clienteId) { 
        carregarClienteParaEdicao(clienteId); 
    }

    formCliente.addEventListener('submit', (event) => { 
        event.preventDefault();
        salvarCliente(clienteId); 
    });
});

function carregarClienteParaEdicao(id) {
    fetch(`/api/clientes/${id}`) 
        .then(response => {
            if (!response.ok) {
                return response.text().then(text => {
                    throw new Error(`Erro ao buscar cliente: ${response.status} - ${text}`); 
                });
            }
            return response.json();
        })
        .then(clienteData => { 
            document.getElementById("nome").value = clienteData.nome;
            document.getElementById("email").value = clienteData.email; 
            document.getElementById("telefone").value = clienteData.telefone;
            document.getElementById("endereco").value = clienteData.endereco;
            document.getElementById("id").value = clienteData.id;
        })
        .catch(error => {
            console.error("Erro ao carregar cliente para edição:", error); 
            alert(`Erro ao carregar cliente para edição: ${error.message}.`); 
        });
}

function salvarCliente(clienteId) { 
    const nome = document.getElementById("nome").value.trim();
    const email = document.getElementById("email").value.trim();     // <--- NOVO CAMPO
    const telefone = document.getElementById("telefone").value.trim();
    const endereco = document.getElementById("endereco").value.trim();
    const idInput = document.getElementById("id").value.trim();

    if (!nome || !email || !telefone || !endereco) { 
        alert("Preencha todos os campos obrigatórios: Nome, Email, Telefone e Endereço!"); 
        return;
    }

    const cliente = { 
        id: idInput ? parseInt(idInput) : 0,
        nome: nome,
        email: email, 
        telefone: telefone,
        endereco: endereco,
    };

    let url = `/api/clientes`; 
    let method = "POST";

    if (clienteId) { 
        url = `/api/clientes/${clienteId}`; 
        method = "PUT";
    }

    fetch(url, {
        method: method,
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(cliente) 
    })
    .then(response => {
        if (!response.ok) {
            return response.text().then(text => { throw new Error(`Erro ao salvar cliente: ${response.status} - ${text}`); }); 
        }
        return response.json();
    })
    .then(() => {
        alert("Cliente salvo com sucesso!"); 
        window.location.href = "/";
    })
    .catch(error => {
        console.error("Erro ao salvar cliente:", error); 
        alert(`Erro ao salvar cliente: ${error.message}`); 
    });
}