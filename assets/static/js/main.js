const elementToc = document.getElementById("generator__output__text");

// Add class to copied element to denote success
function flashElement(element) {
  element.classList.add("flash");
  document.addEventListener("transitionend", function () {
    setTimeout(function () {
      element.classList.remove("flash");
    }, 1000);
  });
}

//add onclike event to copy button
const copyButton = document.getElementById("copy-button");
const flashMessage = document.getElementById("flash-message");

copyButton.addEventListener("click", () => {
  navigator.clipboard.writeText(elementToc.innerText);
});

function fetchRandom() {
  fetch("/random")
    .then((response) => response.json())
    .then((data) => {
      elementToc.innerText = data;
      flashElement(flashMessage);
    });
}
