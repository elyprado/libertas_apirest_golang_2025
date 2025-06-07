function deletar(id) {
    if (confirm("Deseja realmente excluir este restaurante?")) {
        fetch(`/api/restaurantes/${id}`, { method: "DELETE" })
            .then(response => {
                if (!response.ok) {
                    return response.text().then(text => {
                        throw new Error(`HTTP error! status: ${response.status} - ${text}`);
                    });
                }
                return response.text();
            })
            .then(() => {
                alert("Restaurante excluído com sucesso!");
                listarRestaurantes();
            })
            .catch(error => {
                console.error("Erro ao deletar restaurante:", error);
                alert(`Erro ao deletar o restaurante: ${error.message}. Verifique o console.`);
            });
    }
}

function listarRestaurantes() {
    const tabela = document.getElementById("tabela-restaurantes");
    if (!tabela) {
        console.error("Elemento 'tabela-restaurantes' não encontrado no DOM.");
        return;
    }
    tabela.innerHTML = "<tr><td colspan='6' class='text-center'>Carregando restaurantes...</td></tr>";

    fetch("/api/restaurantes") // <--- MUDANÇA AQUI
        .then(response => {
            if (!response.ok) {
                return response.text().then(text => { throw new Error(`Erro ao buscar restaurantes: ${text}`); });
            }
            return response.json();
        })
        .then(dados => mostrarRestaurantes(dados))
        .catch(error => {
            tabela.innerHTML = `<tr><td colspan='6' class='text-danger text-center'>Erro ao carregar restaurantes</td></tr>`;
            console.error("Erro ao listar restaurantes:", error);
            alert(`Erro ao listar restaurantes: ${error.message}`);
        });
}

function mostrarRestaurantes(dados) {
    const tabela = document.getElementById("tabela-restaurantes");
    if (!tabela) {
        console.error("Elemento 'tabela-restaurantes' não encontrado no DOM.");
        return;
    }
    tabela.innerHTML = "";

    if (dados.length === 0) {
        tabela.innerHTML = "<tr><td colspan='6' class='text-center'>Nenhum restaurante cadastrado.</td></tr>";
        return;
    }

    for (let i = 0; i < dados.length; i++) {
        const tr = document.createElement("tr");
        const restaurante = dados[i];

        tr.innerHTML = `
            <td>${restaurante.id}</td>
            <td>${restaurante.nome}</td>
            <td>${restaurante.telefone || 'N/A'}</td>
            <td>${restaurante.endereco || 'N/A'}</td>
            <td>${restaurante.tipoCozinha || 'N/A'}</td> <td>
                <a href="/static/cadastro_restaurantes.html?id=${restaurante.id}" class="btn btn-primary btn-sm me-2">Editar</a> <button class="btn btn-danger btn-sm" onclick='deletar(${restaurante.id})'>Excluir</button>
            </td>
        `;
        tabela.appendChild(tr);
    }
}

document.addEventListener('DOMContentLoaded', listarRestaurantes);