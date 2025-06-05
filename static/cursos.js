const modalcadastro = new bootstrap.Modal(
  document.getElementById("modalcadastro")
);

var idcursoatual;

function alterar(idcurso) {
  idcursoatual = idcurso;
  fetch("http://localhost:8080/cursos/" + idcurso, {
    method: "GET",
    mode: "cors",
  })
    .then((resp) => resp.json())
    .then((dados) => {
      document.getElementById("nome").value = dados.nome;
      document.getElementById("cargaHoraria").value = dados.cargaHoraria;
      document.getElementById("descricao").value = dados.descricao;
      document.getElementById("valor").value = dados.valor;
      modalcadastro.show();
    })
    .catch((err) => console.error("Erro ao buscar curso:", err));
}

function excluir(idcurso) {
  fetch("http://localhost:8080/cursos/" + idcurso, {
    method: "DELETE",
    mode: "cors",
  })
    .then(() => {
      listar();
    })
    .catch((err) => console.error("Erro ao excluir curso:", err));
}

function salvar() {
  let vnome = document.getElementById("nome").value;
  let vcargaHoraria = document.getElementById("cargaHoraria").value;
  let vdescricao = document.getElementById("descricao").value;
  let vvalor = parseFloat(document.getElementById("valor").value) || 0;

  let curso = {
    nome: vnome,
    cargaHoraria: vcargaHoraria,
    descricao: vdescricao,
    valor: vvalor,
  };

  let url, metodo;
  if (idcursoatual > 0) {
    url = "http://localhost:8080/cursos/" + idcursoatual;
    metodo = "PUT";
  } else {
    url = "http://localhost:8080/cursos";
    metodo = "POST";
  }

  fetch(url, {
    method: metodo,
    mode: "cors",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(curso),
  })
    .then(() => {
      listar();
      modalcadastro.hide();
    })
    .catch((err) => console.error("Erro ao salvar curso:", err));
}

function novo() {
  idcursoatual = 0;
  document.getElementById("nome").value = "";
  document.getElementById("cargaHoraria").value = "";
  document.getElementById("descricao").value = "";
  document.getElementById("valor").value = "";
  modalcadastro.show();
}

function listar() {
  const lista = document.getElementById("lista");
  lista.innerHTML = "<tr><td colspan='5'>Carregando...</td></tr>";

  fetch("http://localhost:8080/cursos", {
    method: "GET",
    mode: "cors",
  })
    .then((resp) => resp.json())
    .then((dados) => mostrar(dados))
    .catch((err) => {
      console.error("Erro ao listar cursos:", err);
      lista.innerHTML = "<tr><td colspan='5'>Erro ao carregar.</td></tr>";
    });
}

function mostrar(dados) {
  const lista = document.getElementById("lista");
  lista.innerHTML = "";
  for (let i in dados) {
    lista.innerHTML +=
      "<tr>" +
      "<td>" +
      dados[i].idcurso +
      "</td>" +
      "<td>" +
      dados[i].nome +
      "</td>" +
      "<td>" +
      dados[i].cargaHoraria +
      "</td>" +
      "<td>" +
      dados[i].descricao +
      "</td>" +
      "<td>" +
      dados[i].valor +
      "</td>" +
      "<td>" +
      "<button type='button' class='btn btn-primary' onclick='alterar(" +
      dados[i].idcurso +
      ")'>A</button> " +
      "<button type='button' class='btn btn-danger' onclick='excluir(" +
      dados[i].idcurso +
      ")'>X</button>" +
      "</td>" +
      "</tr>";
  }
}
