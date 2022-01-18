console.log(window.runtime)
// Get input + focus
let nameElement = document.getElementById("name");
nameElement.focus();

// Setup the greet function
window.greet = function () {

  // Get name
  let name = nameElement.value;

  // Call App.Greet(name)
  window.go.main.App.Greet(name).then((result) => {
    // Update result with data back from App.Greet()
    document.getElementById("result").innerText = result;
  });
};

nameElement.onkeydown = function (e) {
  console.log(e)
  if (e.keyCode == 13) {
    window.greet()
  }
}

const toBase64 = file => new Promise((resolve, reject) => {
  const reader = new FileReader();
  reader.readAsDataURL(file);
  reader.onload = () => resolve(reader.result);
  reader.onerror = error => reject(error);
});

const mergeFileToPDFButton = document.getElementById('merge-files-to-pdf');
mergeFileToPDFButton.addEventListener('click', async () => {
  window.go.main.App.OpenDirectoryDialog().then(() => {

  });
  return;
  //get files
  var files = document.getElementById("files").files;
  console.log(files);
  let base64Strings = Array.from(files).map(file => {
    return toBase64(file);
  });
  const images = await Promise.all(base64Strings);

  window.go.main.App.ConvertBase64ToImage(images).then((result) => {
    console.log(result);
  });
});
