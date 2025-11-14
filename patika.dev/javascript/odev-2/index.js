let myname = prompt("İsminizi girin.");
document.querySelector("#myName").innerHTML = myname;
if (myname == "") alert("İsminizi girmeyi unuttunuz.");

let ulDom = document.querySelector("#list");
let input = document.querySelector("#task");
let ekleBtn = document.querySelector("#liveToastBtn");
let closeBtn = document.querySelector("#removeItemBtn");

let list = localStorage.getItem("list")
  ? JSON.parse(localStorage.getItem("list"))
  : [];

list.forEach((element) => {
  let itemElement = document.createElement("li");
  itemElement.classList.add("list-item");
  itemElement.addEventListener("click", mission);
  itemElement.innerHTML = `${element}<span onclick="removeItem(event)" class="close">&times;</span>
  `;
  ulDom.append(itemElement);
});

function addItem(item) {
  let task = First(item);
  let itemElement = document.createElement("li");
  itemElement.addEventListener("click", mission);
  itemElement.innerHTML = `${task}<span onclick="removeItem(event)" id="removeItemBtn" class="close">&times;</span>
  `;
  ulDom.append(itemElement);
  list.push(task);
  saveToLocalStorage(list);
  showToast("Başarılı bir şekilde eklendi");
}

function removeItem(e) {
  list = list.filter((x) => x !== e.target.previousSibling.textContent);
  e.target.parentElement.remove();
  saveToLocalStorage(list);
  showToast("Başarılı bir şekilde silindi");
}

function checkItems(e) {
  if (input.value == "") {
    showToast("Boş değer giremezsiniz");
  } else if (isItemsSame(input.value)) {
    showToast("Aynı değerleri giremezsiniz");
  } else {
    addItem(input.value);
  }
}

function mission(event) {
  event.target.classList.toggle("checked");
  console.log(event.target.classList);
}

function showToast(message) {
  let option = {
    animation: true,
    delay: 3000,
  };
  document.querySelector(".toast-body").innerHTML = message;
  const toast = document.querySelector("#liveToast");
  const toastElement = new bootstrap.Toast(toast, option);
  toastElement.show();
}

function isItemsSame(item) {
  let ok = list.filter((x) => x == item);
  if (ok.length == 0) {
    return 0;
  } else {
    return 1;
  }
}

function First(item) {
  return item;
}

function saveToLocalStorage(list) {
  localStorage.setItem("list", JSON.stringify(list));
}
