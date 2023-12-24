const elementToc = document.getElementById("generator__output__text");
//add onclike event to copy button
const copyButton = document.getElementById("copy-button");
const flashMessage = document.getElementById("flash-message");
const copyIcon = document.getElementById("copy-icon");

// Add class to copied element to denote success
function flashElement(element) {
  element.classList.add("flash");
  copyIcon.classList.add("hidden");
  setTimeout(() => {
    element.classList.remove("flash");
    copyIcon.classList.remove("hidden");
  }, 800);
}

copyButton.addEventListener("click", () => {
  navigator.clipboard.writeText(elementToc.innerText);
  flashElement(flashMessage);
});
