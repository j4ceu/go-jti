const token = getCookie("token")
fetch("http://localhost:8000/v1/phone-number/odd", {
  method: "GET",
  headers: {
    "Content-type": "application/json; charset=UTF-8",
    "Authorization": "Bearer " + token
  }
}).then(
  res => {
    res.json().then(
      data => {
        console.log(data.data);
        if (data.data.length > 0) {
          var temp = "";
          data.data.forEach((itemData, key) => {
            counter = "modalsOdd" + key
            temp += "<tr>";
            temp += "<td>" + itemData.phone_number + "</td>";
            temp += `
              <td class="gap-2 d-flex justify-content-center">
                <button type="button" class="btn btn-primary justify-content-center" data-mdb-ripple-init data-mdb-modal-init data-mdb-target="#${counter}" style="background-color: #dd4b39;">
                  Edit
                </button>
                <button type="button" class="btn btn-secondary justify-content-center" data-id="${counter}" id="delete-${counter}"data-mdb-ripple-init>
                  Delete
                </button>
              </td></tr>
            `
            temp += `
              <!-- Modal -->
              <div class="modal fade" id="${counter}" tabindex="-1" aria-labelledby="exampleModalLabel" aria-hidden="true">
                <div class="modal-dialog">
                  <div class="modal-content">
                    <div class="modal-header">
                      <h5 class="modal-title" id="exampleModalLabel">Handphone Data</h5>
                      <button type="button" class="btn-close" data-mdb-ripple-init data-mdb-dismiss="modal" aria-label="Close"></button>
                    </div>
                    <div class="modal-body">
                    <div class="form-outline mb-4" data-mdb-input-init>
                      <input type="text" id="number-${counter}" class="form-control active" value="${itemData.phone_number}" />
                      <label class="form-label" for="number-${counter}">No Handphone</label>
                    </div>
                    <div class="form-outline mb-44" data-mdb-input-init>
                      <input type="text" id="provider-${counter}" class="form-control active" value="${itemData.provider}" disabled />
                      <label class="form-label" for="provider-${counter}">Provider</label>
                    </div>    
                    </div>
                    <div class="modal-footer">
                      <button type="button" class="btn btn-secondary" data-mdb-ripple-init data-mdb-dismiss="modal">Close</button>
                      <button type="button" class="btn btn-primary" id="submit-${counter}"  data-id="${counter}" style="background-color: #dd4b39;" data-mdb-ripple-init>Save changes</button>
                    </div>
                  </div>
                </div>
              </div>
            `
          });
          document.getElementById('odd-data').innerHTML = temp;

          data.data.forEach((itemData, key) => {
            counter = "modalsOdd" + key

            let submitBtn = document.getElementById(`submit-${counter}`);

            submitBtn.addEventListener("click", (e) => {

              counter = submitBtn.dataset.id
              const token = getCookie("token")
              console.log("counter " + counter)
              e.preventDefault();

              let handphone = document.getElementById(`number-${counter}`);
              console.log(handphone)

              if (handphone.value == "") {
                alert("Ensure you input a value in phone number fields!");
              } else {
                // perform operation with form input
                // TODO: update
                updateDataWithEcnrypt(handphone,itemData.id)
              }
            })

            //Delete
            let deleteBtn = document.getElementById(`delete-${counter}`);
            deleteBtn.addEventListener("click", (e) => {
              counter = deleteBtn.dataset.id
              const token = getCookie("token")
              e.preventDefault();

              // TODO: delete
              fetch(`http://localhost:8000/v1/phone-number/${itemData.id}`, {
                method: 'DELETE',
                headers: {
                  "Content-type": "application/json; charset=UTF-8",
                  "Authorization": "Bearer " + token
                },
              }).then((response) => {
                if (!response.ok) {
                  alert("Failed to delete the data!");
                } else {
                  if (alert('Data successfully deleted!')) { }
                  else window.location.reload();
                }
              })

            })
          })
        }
      }
    )
  }
)

fetch("http://localhost:8000/v1/phone-number/even", {
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
            counter = "modalsEven" + key

            temp += "<tr>";
            temp += "<td>" + itemData.phone_number + "</td>";
            temp += `
              <td class="gap-2 d-flex justify-content-center">
                <button type="button" class="btn btn-primary justify-content-center" data-mdb-ripple-init data-mdb-modal-init data-mdb-target="#${counter}" style="background-color: #dd4b39;">
                  Edit
                </button>
                <button type="button" class="btn btn-secondary justify-content-center" data-mdb-ripple-init id="delete-${counter}" data-id="${counter}">
                  Delete
                </button>
              </td></tr>
            `
            temp += `
            <!-- Modal -->
              <div class="modal fade" id="${counter}" tabindex="-1" aria-labelledby="exampleModalLabel" aria-hidden="true">
                <div class="modal-dialog">
                  <div class="modal-content">
                    <div class="modal-header">
                      <h5 class="modal-title" id="exampleModalLabel">Handphone Data</h5>
                      <button type="button" class="btn-close" data-mdb-ripple-init data-mdb-dismiss="modal" aria-label="Close"></button>
                    </div>
                    <div class="modal-body">
                    <div class="form-outline mb-4" data-mdb-input-init>
                      <input type="text" id="number-${counter}" class="form-control active" value="${itemData.phone_number}" />
                      <label class="form-label" for="number-${counter}">No Handphone</label>
                    </div>
                    <div class="form-outline mb-44" data-mdb-input-init>
                      <input type="text" id="provider-${counter}" class="form-control active" value="${itemData.provider}" disabled />
                      <label class="form-label" for="provider-${counter}">Provider</label>
                    </div>    
                    </div>
                    <div class="modal-footer">
                      <button type="button" class="btn btn-secondary" data-mdb-ripple-init data-mdb-dismiss="modal">Close</button>
                      <button type="button" class="btn btn-primary" data-mdb-ripple-init id="submit-${counter}" data-id="${counter}" style="background-color: #dd4b39;">Save changes</button>
                    </div>
                  </div>
                </div>
              </div>
              `

          });

          document.getElementById('even-data').innerHTML = temp;

          data.data.forEach((itemData, key) => {
            counter = "modalsEven" + key

            let submitBtn = document.getElementById(`submit-${counter}`);
            //Update
            submitBtn.addEventListener("click", (e) => {
              counter = submitBtn.dataset.id
              const token = getCookie("token")
              console.log("counter " + counter)
              e.preventDefault();

              let handphone = document.getElementById(`number-${counter}`);
              console.log(handphone)

              if (handphone.value == "") {
                alert("Ensure you input a value in phone number fields!");
              } else {
                // perform operation with form input
                // TODO: update
                updateDataWithEcnrypt(handphone,itemData.id)
              }
            })

            //Delete
            let deleteBtn = document.getElementById(`delete-${counter}`);
            deleteBtn.addEventListener("click", (e) => {
              counter = deleteBtn.dataset.id
              const token = getCookie("token")
              e.preventDefault();

              // TODO: delete
              fetch(`http://localhost:8000/v1/phone-number/${itemData.id}`, {
                method: 'DELETE',
                headers: {
                  "Content-type": "application/json; charset=UTF-8",
                  "Authorization": "Bearer " + token
                },
              }).then((response) => {
                if (!response.ok) {
                  alert("Failed to delete the data!");
                } else {
                  if (alert('Data successfully deleted!')) { }
                  else window.location.reload();
                }
              })

            })
          })
        }
      })
  }
)



function getCookie(name) {
  const value = `; ${document.cookie}`;
  const parts = value.split(`; ${name}=`);
  if (parts.length === 2) return parts.pop().split(';').shift();
}

async function updateDataWithEcnrypt(handphone, id) {
  const token = getCookie("token")

  // Encrypt both data items concurrently
  const handphone_data = await Promise.all([
    encryptAES(handphone.value)
  ]);

  fetch(`http://localhost:8000/v1/phone-number/${id}`, {
    method: 'PUT',
    headers: {
      "Content-type": "application/json; charset=UTF-8",
      "Authorization": "Bearer " + token
    },
    body: JSON.stringify({
      phone_number: handphone_data[0],
    })
  }).then((response) => {
    if (!response.ok) {
      alert("Failed to update the data!");
    } else {
      if (alert('Data successfully updated!')) { }
      else window.location.reload();
    }
  })


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
