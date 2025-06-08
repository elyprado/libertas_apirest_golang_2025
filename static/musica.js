const apiURL = 'http://localhost:8080/musicas'; // ajuste a URL da sua API

const modal = new bootstrap.Modal(document.getElementById('modalcadastro'));
const form = document.getElementById('form-item');

function listar() {
  fetch(apiURL)
    .then(res => res.json())
    .then(data => {
      const lista = document.getElementById('lista');
      lista.innerHTML = '';
      data.forEach(musica => {
        lista.innerHTML += `
          <tr>
            <td>${musica.id}</td>
            <td>${musica.titulo}</td>
            <td>${musica.artista}</td>
            <td>${musica.album}</td>
            <td>${musica.ano}</td>
            <td>${musica.genero}</td>
            <td>
              <button class="btn btn-sm btn-primary" onclick="editar(${musica.id})">Editar</button>
              <button class="btn btn-sm btn-danger" onclick="deletar(${musica.id})">Excluir</button>
            </td>
          </tr>
        `;
      });
    })
    .catch(err => alert('Erro ao carregar músicas: ' + err));
}

function novo() {
  form.reset();
  document.getElementById('id').value = '';
  modal.show();
}

function salvar() {
  const id = document.getElementById('id').value;
  const musica = {
    titulo: document.getElementById('titulo').value,
    artista: document.getElementById('artista').value,
    album: document.getElementById('album').value,
    ano: parseInt(document.getElementById('ano').value),
    genero: document.getElementById('genero').value
  };

  if (!musica.titulo || !musica.artista || !musica.album || !musica.ano || !musica.genero) {
    alert('Por favor, preencha todos os campos.');
    return;
  }

  let method = 'POST';
  let url = apiURL;

  if (id) {
    method = 'PUT';
    url = `${apiURL}/${id}`;
  }

  fetch(url, {
    method: method,
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(musica)
  })
    .then(res => {
      if (!res.ok) throw new Error('Erro na requisição');
      return res.json();
    })
    .then(() => {
      modal.hide();
      listar();
    })
    .catch(err => alert('Erro ao salvar música: ' + err));
}

function editar(id) {
  fetch(`${apiURL}/${id}`)
    .then(res => {
      if (!res.ok) throw new Error('Música não encontrada');
      return res.json();
    })
    .then(musica => {
      document.getElementById('id').value = musica.id;
      document.getElementById('titulo').value = musica.titulo;
      document.getElementById('artista').value = musica.artista;
      document.getElementById('album').value = musica.album;
      document.getElementById('ano').value = musica.ano;
      document.getElementById('genero').value = musica.genero;
      modal.show();
    })
    .catch(err => alert('Erro ao buscar música: ' + err));
}

function deletar(id) {
  if (!confirm('Deseja realmente excluir esta música?')) return;

  fetch(`${apiURL}/${id}`, { method: 'DELETE' })
    .then(res => {
      if (!res.ok) throw new Error('Erro ao deletar');
      listar();
    })
    .catch(err => alert('Erro ao excluir música: ' + err));
}
