const modalcadastro = new bootstrap.Modal(document.getElementById('modalcadastro')); 

var idusuarioatual;

function alterar(id) {
  fetch("http://127.0.0.1:8080/imoveis/" + id)
    .then(resp => resp.json())
    .then(dados => {
      idimovelatual = id;  
      document.getElementById('endereco').value = dados.endereco;
      document.getElementById('cep').value = dados.cep;
      document.getElementById('valor').value = dados.valor;
      document.getElementById('contato').value = dados.contato;
      document.getElementById('STATUS').value = dados.STATUS;
      modalcadastro.show();
    })
    .catch(err => console.error("Erro ao buscar Imovel:", err));
}

function excluir(id) {
  if (window.confirm("Deseja realmente excluir este imovel?")) {
    fetch("http://127.0.0.1:8080/imoveis/" + id, {  
      method: "DELETE",
  })
  .then(() => listar())
  .catch(err => console.error("Erro ao excluir Imovel:", err));
  }
}

function salvar() {
  let vendereco = document.getElementById("endereco").value;
  let vcep = document.getElementById("cep").value;
  let vvalor = document.getElementById("valor").value;
  let vcontato = document.getElementById("contato").value;
  let vstatus = document.getElementById("STATUS").value;

  let imovel = {  
    endereco: vendereco,
    cep: vcep,
    valor: vvalor,
    contato: vcontato,
    STATUS: vstatus
  };

  let url, metodo;
  if (idusuarioatual > 0) {
    url = "http://127.0.0.1:8080/imoveis/" + idusuarioatual; // Enviar ID para PUT
    metodo = "PUT";
  } else {
    url = "http://127.0.0.1:8080/imoveis/";  // Para POST (novo cadastro)
    metodo = "POST";
  }

  fetch(url, {
    method: metodo,
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify(imovel)
  })
  .then(() => {
    listar();
    modalcadastro.hide();
  })
  .catch(err => console.error("Erro ao salvar imovel:", err));
}

function novo() {
  idusuarioatual = 0; // Limpa o idusuarioatual para um novo cadastro
  document.getElementById("endereco").value = "";
  document.getElementById("cep").value = "";
  document.getElementById("valor").value = "";
  document.getElementById("contato").value = "";
  document.getElementById("STATUS").value = "";
  modalcadastro.show();
}

function listar() {
  const listar = document.getElementById("lista");
  listar.innerHTML = "<tr><td colspan='5'>Carregando...</td></tr>";  

  fetch("http://127.0.0.1:8080/imoveis/")
    .then(resp => resp.json())
    .then(dados => mostrar(dados))
    .catch(err => {
      console.error("Erro ao listar imoveis:", err);
      listar.innerHTML = "<tr><td colspan='5'>Erro ao carregar dados.</td></tr>";  
    });
}

function mostrar(dados) {
  const lista = document.getElementById("lista");
  lista.innerHTML = "";
  for (let i in dados) {
    lista.innerHTML += "<tr>"
      + "<td>" + dados[i].id + "</td>"
      + "<td>" + dados[i].endereco + "</td>"
      + "<td>" + dados[i].cep + "</td>"
      + "<td>" + dados[i].valor + "</td>"
      + "<td>" + dados[i].contato + "</td>"
      + "<td>" + dados[i].STATUS + "</td>"
      + "<td>" 
      + "<button type='button' class='btn btn-primary btn-sm' onclick='alterar(" + dados[i].id + ")'>Alterar</button> "
      + "<button type='button' class='btn btn-secondary btn-sm' onclick='excluir(" + dados[i].id + ")'>Excluir</button>"
      + "</td>"
      + "</tr>";
  }
}