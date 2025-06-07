document.addEventListener('DOMContentLoaded', () => {
    const urlParams = new URLSearchParams(window.location.search);
    const restauranteId = urlParams.get('id'); 

    const formRestaurante = document.getElementById('formRestaurante'); 
    if (!formRestaurante) {
        console.error("Formulário 'formRestaurante' não encontrado.");
        return;
    }

    if (restauranteId) { 
        carregarRestauranteParaEdicao(restauranteId); 
    }

    formRestaurante.addEventListener('submit', (event) => { 
        event.preventDefault();
        salvarRestaurante(restauranteId); 
    });
});

function carregarRestauranteParaEdicao(id) {
    fetch(`/api/restaurantes/${id}`) 
        .then(response => {
            if (!response.ok) {
                return response.text().then(text => {
                    throw new Error(`Erro ao buscar restaurante: ${response.status} - ${text}`);
                });
            }
            return response.json();
        })
        .then(restauranteData => { 
            document.getElementById("nome").value = restauranteData.nome; 
            document.getElementById("telefone").value = restauranteData.telefone; 
            document.getElementById("endereco").value = restauranteData.endereco; 
            document.getElementById("tipoCozinha").value = restauranteData.tipoCozinha; 
            document.getElementById("id").value = restauranteData.id;
        })
        .catch(error => {
            console.error("Erro ao carregar restaurante para edição:", error);
            alert(`Erro ao carregar restaurante para edição: ${error.message}.`);
        });
}

function salvarRestaurante(restauranteId) { 
    const nome = document.getElementById("nome").value.trim(); 
    const telefone = document.getElementById("telefone").value.trim(); 
    const endereco = document.getElementById("endereco").value.trim(); 
    const tipoCozinha = document.getElementById("tipoCozinha").value.trim(); 
    const idInput = document.getElementById("id").value.trim();

    if (!nome || !telefone || !endereco || !tipoCozinha) { 
        alert("Preencha todos os campos obrigatórios: Nome, Telefone, Endereço e Tipo de Cozinha!"); 
        return;
    }

    const restaurante = { 
        id: idInput ? parseInt(idInput) : 0,
        nome: nome, 
        telefone: telefone, 
        endereco: endereco, 
        tipoCozinha: tipoCozinha 
    };

    let url = `/api/restaurantes`; 
    let method = "POST";

    if (restauranteId) { 
        url = `/api/restaurantes/${restauranteId}`; 
        method = "PUT";
    }

    fetch(url, {
        method: method,
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(restaurante) 
    })
    .then(response => {
        if (!response.ok) {
            return response.text().then(text => { throw new Error(`Erro ao salvar restaurante: ${response.status} - ${text}`); }); 
        }
        return response.json();
    })
    .then(() => {
        alert("Restaurante salvo com sucesso!"); 
        window.location.href = "/";
    })
    .catch(error => {
        console.error("Erro ao salvar restaurante:", error); 
        alert(`Erro ao salvar restaurante: ${error.message}`); 
    });
}