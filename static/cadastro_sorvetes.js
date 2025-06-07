document.addEventListener('DOMContentLoaded', () => {
    const urlParams = new URLSearchParams(window.location.search);
    const sorveteId = urlParams.get('id');

    const formSorvete = document.getElementById('formSorvete');
    if (!formSorvete) {
        console.error("Formulário 'formSorvete' não encontrado.");
        return;
    }

    if (sorveteId) {
        carregarSorveteParaEdicao(sorveteId);
    }

    formSorvete.addEventListener('submit', (event) => {
        event.preventDefault(); // Impede o envio padrão do formulário
        salvarSorvete(sorveteId); // Passa o ID para a função de salvar
    });
});

function carregarSorveteParaEdicao(id) {
    fetch(`/api/sorvetes/${id}`) // URL relativa
        .then(response => {
            if (!response.ok) {
                // Tenta ler o texto do erro para dar mais detalhes
                return response.text().then(text => {
                    throw new Error(`Erro ao buscar sorvete: ${response.status} - ${text}`);
                });
            }
            return response.json(); // Espera JSON do backend
        })
        .then(sorveteData => {
            // IMPORTANTE: Acessar as propriedades com minúscula (camelCase)
            document.getElementById("sabor").value = sorveteData.sabor;
            document.getElementById("preco").value = sorveteData.preco;
            document.getElementById("tipo").value = sorveteData.tipo;
            document.getElementById("disponivel").checked = sorveteData.disponivel;
            document.getElementById("descricao").value = sorveteData.descricao;
            document.getElementById("id").value = sorveteData.id; // Preenche o campo hidden ID
        })
        .catch(error => {
            console.error("Erro ao carregar sorvete para edição:", error);
            alert(`Erro ao carregar sorvete para edição: ${error.message}.`);
        });
}

function salvarSorvete(sorveteId) {
    const sabor = document.getElementById("sabor").value.trim();
    const preco = parseFloat(document.getElementById("preco").value);
    const tipo = document.getElementById("tipo").value.trim();
    const disponivel = document.getElementById("disponivel").checked;
    const descricao = document.getElementById("descricao").value.trim();
    // O ID do campo hidden pode ser usado para determinar se é uma edição
    const idInput = document.getElementById("id").value.trim();

    if (!sabor || isNaN(preco) || !tipo) {
        alert("Preencha os campos obrigatórios: Sabor, Preço e Tipo!");
        return;
    }

    // Cria o objeto sorvete com os nomes das propriedades em camelCase
    const sorvete = {
        id: idInput ? parseInt(idInput) : 0, // Envia 0 ou um valor que o backend ignore para criar
        sabor: sabor,
        preco: preco,
        tipo: tipo,
        disponivel: disponivel,
        descricao: descricao
    };

    let url = `/api/sorvetes`; // URL base para criar
    let method = "POST";

    if (sorveteId) { // Se há um ID na URL, é uma atualização (PUT)
        url = `/api/sorvetes/${sorveteId}`; // URL para atualizar com ID
        method = "PUT";
    }

    fetch(url, {
        method: method,
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(sorvete)
    })
    .then(response => {
        if (!response.ok) {
            return response.text().then(text => { throw new Error(`Erro ao salvar sorvete: ${response.status} - ${text}`); });
        }
        return response.json(); // Espera JSON do backend
    })
    .then(() => {
        alert("Sorvete salvo com sucesso!");
        window.location.href = "/"; // Redireciona para a página principal (lista)
    })
    .catch(error => {
        console.error("Erro ao salvar sorvete:", error);
        alert(`Erro ao salvar sorvete: ${error.message}`);
    });
}