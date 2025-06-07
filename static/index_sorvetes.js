function deletar(id) {
    if (confirm("Deseja realmente excluir este sorvete?")) {
        fetch(`/api/sorvetes/${id}`, { method: "DELETE" })
            .then(response => {
                if (!response.ok) {
                    // Tenta ler o texto do erro para dar mais detalhes
                    return response.text().then(text => {
                        throw new Error(`HTTP error! status: ${response.status} - ${text}`);
                    });
                }
                // Se tudo OK, não tenta parsear JSON se o backend não retornar nada no DELETE
                // O backend pode retornar uma mensagem de sucesso em JSON, então podemos tentar json()
                // mas text() é mais seguro se for apenas um status 200 OK sem corpo.
                // Vou deixar como .text() para compatibilidade, já que o Go retorna map[string]string.
                return response.text();
            })
            .then(() => {
                alert("Sorvete excluído com sucesso!");
                listarSorvetes(); // Recarrega a lista
            })
            .catch(error => {
                console.error("Erro ao deletar sorvete:", error);
                alert(`Erro ao deletar o sorvete: ${error.message}. Verifique o console.`);
            });
    }
}

function listarSorvetes() {
    const tabela = document.getElementById("tabela-sorvetes");
    if (!tabela) {
        console.error("Elemento 'tabela-sorvetes' não encontrado no DOM.");
        return;
    }
    tabela.innerHTML = "<tr><td colspan='7' class='text-center'>Carregando sorvetes...</td></tr>";

    fetch("/api/sorvetes") // URL relativa ao host (localhost:8080)
        .then(response => {
            if (!response.ok) {
                return response.text().then(text => { throw new Error(`Erro ao buscar sorvetes: ${text}`); });
            }
            return response.json(); // Espera JSON do backend
        })
        .then(dados => mostrarSorvetes(dados))
        .catch(error => {
            tabela.innerHTML = `<tr><td colspan='7' class='text-danger text-center'>Erro ao carregar sorvetes</td></tr>`;
            console.error("Erro ao listar sorvetes:", error);
            alert(`Erro ao listar sorvetes: ${error.message}`);
        });
}

function mostrarSorvetes(dados) {
    const tabela = document.getElementById("tabela-sorvetes");
    if (!tabela) {
        console.error("Elemento 'tabela-sorvetes' não encontrado no DOM.");
        return;
    }
    tabela.innerHTML = ""; // Limpa a tabela antes de popular

    if (dados.length === 0) {
        tabela.innerHTML = "<tr><td colspan='7' class='text-center'>Nenhum sorvete cadastrado.</td></tr>";
        return;
    }

    for (let i = 0; i < dados.length; i++) {
        const tr = document.createElement("tr");
        const sorvete = dados[i];
        // IMPORTANTE: Acessar as propriedades com minúscula (camelCase)
        // porque o JSON do Go está usando os nomes definidos nos tags `json:"..."`
        const statusDisponivel = sorvete.disponivel ? '<span class="badge bg-success">Sim</span>' : '<span class="badge bg-danger">Não</span>';

        tr.innerHTML = `
            <td>${sorvete.id}</td>
            <td>${sorvete.sabor}</td>
            <td>R$ ${sorvete.preco ? sorvete.preco.toFixed(2) : '0.00'}</td>
            <td>${sorvete.tipo || 'N/A'}</td>
            <td>${statusDisponivel}</td>
            <td>${sorvete.descricao || 'N/A'}</td>
            <td>
                <a href="/static/cadastro_sorvetes.html?id=${sorvete.id}" class="btn btn-primary btn-sm me-2">Editar</a>
                <button class="btn btn-danger btn-sm" onclick='deletar(${sorvete.id})'>Excluir</button>
            </td>
        `;
        tabela.appendChild(tr);
    }
}

document.addEventListener('DOMContentLoaded', listarSorvetes);