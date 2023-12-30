let form = document.getElementById("form");
form.addEventListener("submit", (e) => {
  e.preventDefault();

  let handphone = document.getElementById("Handphone").value
  let provider = document.getElementById("Provider").value

  if (handphone == "" || provider == "") {
    alert("Ensure you input a value in both fields!");
  } else {
    // perform operation with form input
    postDataWithEcnrypt(handphone.toString(), provider)
    
  }
});


async function postDataWithEcnrypt(handphone, provider) {
  const token = getCookie("token")

     // Encrypt both data items concurrently
     const [handphone_data, provider_data] = await Promise.all([
      encryptAES(handphone),
      encryptAES(provider)
    ]);

    console.log("handphone enkrip " + handphone_data)
   

    fetch("http://localhost:8000/v1/phone-number", {
      method: "POST",
      body: JSON.stringify({
        phone_number: handphone_data,
        provider: provider_data,
      }),
      headers: {
        "Content-type": "application/json; charset=UTF-8",
        "Authorization": "Bearer " + token
      }
    })
      .then((response) => {
        if (!response.ok) {
          alert("Failed to submit the form!");
        } else {
          alert("This form has been successfully submitted!");
        }
      })

}

let auto = document.getElementById("generate-number");
auto.addEventListener("click", (e) => {
  e.preventDefault();
  // perform operation with form input
  const token = getCookie("token")
  fetch("http://localhost:8000/v1/phone-number/generate-number", {
    method: "GET",
    headers: {
      "Content-type": "application/json; charset=UTF-8",
      "Authorization": "Bearer " + token
    }
  }).then(
    res => {
      res.json().then(
        data => {
          if (data.data.length > 0) {

            var temp = "";
            data.data.forEach((itemData, key) => {
              var counter = "number-" + key
              temp += "<tr>";
              temp += `<td> ${key + 1}  <input class="form-check-input" type="radio" name="flexRadioDefault" id="${counter}" data-number="${itemData.phone_number}" data-provider="${itemData.provider}"> </td>`
              temp += "<td>" + itemData.phone_number + "</td>";
              temp += "<td>" + itemData.provider + "</td></tr>";

            });
            document.getElementById('generate-number-table').innerHTML = temp;

          }
        }
      )
    }
  )


});

let choose = document.getElementById("choose");
choose.addEventListener("click", (e) => {
  var radioChecked = atLeastOneRadio()
  if (radioChecked === undefined) {
    alert("Please pick your number first")
  } else {
    let handphone = document.getElementById("Handphone")

    $("#Provider").val(radioChecked.dataset.provider)
    $("#Provider").trigger("change.select2")
    handphone.value = radioChecked.dataset.number
    handphone.classList.add("active")
    $("#exampleModal").modal("hide");
  }


});



function atLeastOneRadio() {
  var radios = document.getElementsByTagName('input');
  var value;
  for (var i = 0; i < radios.length; i++) {
    if (radios[i].type === 'radio' && radios[i].checked) {
      // get value, set checked flag or do whatever you need to
      value = radios[i];
    }
  }
  return value
}

function getCookie(name) {
  const value = `; ${document.cookie}`;
  const parts = value.split(`; ${name}=`);
  if (parts.length === 2) return parts.pop().split(';').shift();
}


async function encryptAES(message) {

  const ENCRYPT_KEY = "my32digitkey12345678901234567890"; // 256-bit key (32 bytes)
  const ENCRYPT_IV = "my16digitIvKey12"; // 128-bit IV (16 bytes)

  const key = new TextEncoder().encode(ENCRYPT_KEY);
  const iv = new TextEncoder().encode(ENCRYPT_IV);

  // 2. Define the algorithm configuration
  const algorithm = {
    name: "AES-CBC",
    iv: iv
  };

  // 3. Import the key
  const keyObject = await crypto.subtle.importKey(
    "raw",
    key,
    algorithm,
    false,
    ["encrypt"]
  );

  // 4. Encode the message
  const data = new TextEncoder().encode(message);

  // 5. Encrypt the data
  const encrypted = await crypto.subtle.encrypt(algorithm, keyObject, data);

  const encryptedArray = new Uint8Array(encrypted);
  const encryptedString = btoa(String.fromCharCode(...encryptedArray));

  return encryptedString;
}

