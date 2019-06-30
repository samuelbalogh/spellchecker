



function checkPhrase() {
  let text = document.getElementById("textToCheck").value
  let oReq = new XMLHttpRequest();
  oReq.addEventListener("load", reqListener);
  let encodedText = encodeURIComponent(text)
  let query = "text=" + encodedText
  oReq.open("POST", "http://localhost:8001/check");
  oReq.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");

  oReq.addEventListener('load', function(event) {
    console.log('Yeah! Data sent and response loaded.');
  });

  oReq.addEventListener('error', function(event) {
    console.log('Oops! Something goes wrong.');
  });

  oReq.send(query);
  return false;
}


function reqListener() {
  let suggestion = "Did you mean: <br> <br> " 
  console.log(this)
  var data = JSON.parse(this.responseText);
  data.forEach(function(element) {
      suggestion = suggestion + " " + element
  })
  suggestion = suggestion + "?"
  console.log(suggestion)
  document.getElementById("correction").innerHTML = suggestion;
}
